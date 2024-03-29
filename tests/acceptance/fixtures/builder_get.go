package fixtures

import (
	"context"
	"github.com/kitstack/dbkit"
	"github.com/kitstack/dbkit/tests/models"
)

func (fixture *Fixture) BuilderGet(ctx context.Context) (err error) {

	user, err := dbkit.Use[*models.UsersModel](ctx, fixture.Connector()).SetFields("Id").Get(1)

	fixture.Assert().NoError(err)
	fixture.Assert().EqualValues(1, user.Id)

	return
}

func (fixture *Fixture) BuilderGetWithJoin(ctx context.Context) (err error) {

	comment, err := dbkit.Use[*models.CommentsModel](ctx, fixture.Connector()).SetFields("Id", "Post.Id", "Post.Comments.Id", "Post.Comments.Content").Get(1)

	fixture.Assert().NoError(err)
	fixture.Assert().EqualValues(1, comment.Id)
	fixture.Assert().EqualValues(1, comment.Post.Id)
	fixture.Assert().Len(comment.Post.Comments, 2)
	fixture.Assert().EqualValues(1, comment.Post.Comments[0].Id)
	fixture.Assert().EqualValues(2, comment.Post.Comments[1].Id)

	return
}
