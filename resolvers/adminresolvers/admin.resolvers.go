package adminresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/lib"
	"github.com/hiroyky/nikki_backend/presenter"
	"github.com/hiroyky/nikki_backend/service"
)

func (r *mutationResolver) NewArticle(ctx context.Context, input adminmodel.ArticleInput) (*adminmodel.Article, error) {
	article := presenter.ToDBArticleFromAdminArticleInput(&input)
	dst, err := service.NewArticle(ctx, article)
	if err != nil {
		return nil, err
	}

	return presenter.ToGQLAdminArticleFromDBArticle(dst), nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, id string, input adminmodel.ArticleInput) (*adminmodel.Article, error) {
	dbID, err := lib.DecodeGraphQLID(id)
	if err != nil {
		return nil, err
	}

	dst, err := service.UpdateArticle(ctx, dbID, presenter.ToDBArticleFromAdminArticleInput(&input))
	if err != nil {
		return nil, err
	}

	return presenter.ToGQLAdminArticleFromDBArticle(dst), nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*adminmodel.Article, error) {
	dbID, err := lib.DecodeGraphQLID(id)
	if err != nil {
		return nil, err
	}

	dst, err := service.GetArticle(ctx, dbID)
	if err != nil {
		return nil, err
	}

	return presenter.ToGQLAdminArticleFromDBArticle(dst), nil
}

func (r *queryResolver) Articles(ctx context.Context) (*adminmodel.ArticleConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
