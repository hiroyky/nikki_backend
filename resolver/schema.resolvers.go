package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/domain/gqlmodel"
	"github.com/hiroyky/nikki_backend/lib"
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
	articles, err := dbmodel.Articles().AllG(ctx)
	if err != nil {
		return nil, err
	}

	return gqlmodel.GenArticleConnection(articles), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
