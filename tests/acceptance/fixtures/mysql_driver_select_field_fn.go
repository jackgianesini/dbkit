package fixtures

import (
	"context"
	"github.com/lab210-dev/dbkit"
	"github.com/lab210-dev/dbkit/connector/drivers"
	"github.com/lab210-dev/dbkit/specs"
	"github.com/lab210-dev/dbkit/tests/models"
)

func (fixture *Fixture) MysqlDriverSelectWithFieldFn(ctx context.Context) (err error) {
	fields := []specs.DriverField{
		drivers.NewField().
			SetName("Id").
			SetCustom("SELECT COUNT(id) FROM posts WHERE posts.user_id = %Id%", []specs.DriverField{
				drivers.NewField().SetColumn("id").SetIndex(0).SetName("Id"),
			}),
	}

	selectPayload := dbkit.NewPayload[*models.UsersModel]()
	selectPayload.SetFields(fields)

	err = fixture.Connector().Select(ctx, selectPayload)
	if err != nil {
		return err
	}

	fixture.Assert().Len(selectPayload.Result(), 3)
	fixture.Assert().EqualValues(2, selectPayload.Result()[0].Id)
	fixture.Assert().EqualValues(1, selectPayload.Result()[1].Id)
	fixture.Assert().EqualValues(1, selectPayload.Result()[2].Id)

	return
}