package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/lib"

	"github.com/hiroyky/nikki_backend/domain/gqlmodel"
)

func (r *queryResolver) Article(ctx context.Context, id string) (*gqlmodel.Article, error) {
	dbID, err := lib.DecodeGraphQLID(id)
	if err != nil {
		return nil, err
	}

	article, err := dbmodel.FindArticleG(ctx, dbID)
	if err != nil {
		return nil, err
	}

	return gqlmodel.GenArticle(article), nil
}

func (r *queryResolver) Articles(ctx context.Context) (*gqlmodel.ArticleConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
