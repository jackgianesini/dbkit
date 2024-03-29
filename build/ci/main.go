package main

import (
	"context"
	"dagger.io/dagger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Acceptance tests")

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// vendor := client.CacheVolume("vendor")

	dir := client.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"/vendor"},
	})

	db := client.Container().
		From("mysql:8").
		WithEnvVariable("MYSQL_USER", os.Getenv("MYSQL_USER")).
		WithEnvVariable("MYSQL_PASSWORD", os.Getenv("MYSQL_PASSWORD")).
		WithEnvVariable("MYSQL_DATABASE", os.Getenv("MYSQL_DATABASE")).
		WithEnvVariable("MYSQL_ROOT_PASSWORD", os.Getenv("MYSQL_ROOT_PASSWORD")).
		WithMountedFile("/docker-entrypoint-initdb.d/dump.sql", dir.File("/tests/acceptance/data/acceptance.sql")).
		WithExposedPort(3306).
		WithExec(nil)

	client.
		Container().
		From("golang:latest").
		WithServiceBinding("db", db).
		WithEnvVariable("MYSQL_HOST", os.Getenv("MYSQL_HOST")).
		WithEnvVariable("MYSQL_USER", os.Getenv("MYSQL_USER")).
		WithEnvVariable("MYSQL_PASSWORD", os.Getenv("MYSQL_PASSWORD")).
		WithEnvVariable("MYSQL_DATABASE", os.Getenv("MYSQL_DATABASE")).
		WithEnvVariable("MYSQL_ROOT_PASSWORD", os.Getenv("MYSQL_ROOT_PASSWORD")).
		WithEnvVariable("DEBUG", os.Getenv("DEBUG")).
		WithMountedDirectory("/src", dir).
		//WithMountedCache("/src/vendor", vendor).
		WithWorkdir("/src").
		WithExec([]string{"go", "run", "./tests/acceptance"}).Stdout(ctx)

	/*	_, err = golang.WithExec([]string{"go", "run", "./tests/acceptance"}).Stdout(ctx)

		if err != nil {
			return err
		}*/

	return nil
}
