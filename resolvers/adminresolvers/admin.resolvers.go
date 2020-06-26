package adminresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/lib"
	"github.com/hiroyky/nikki_backend/service"
	"github.com/volatiletech/sqlboiler/boil"
)

func (r *mutationResolver) NewArticle(ctx context.Context, input adminmodel.ArticleInput) (*adminmodel.Article, error) {
	article := service.ToDBArticleFromAdminArticleInput(&input)
	if err := article.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}
	return service.ToGQLAdminArticleFromDBArticle(article), nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, id string, input adminmodel.ArticleInput) (*adminmodel.Article, error) {
	dbID, err := lib.DecodeGraphQLID(id)
	if err != nil {
		return nil, err
	}

	article, err := dbmodel.FindArticleG(ctx, dbID)
	if err != nil {
		return nil, err
	}

	dst := service.UpdateDBArticleFromAdminArticleInput(article, &input)
	if err != nil {
		return nil, err
	}
	if _, err := dst.UpdateG(ctx, boil.Infer()); err != nil {
		return nil, err
	}

	return service.ToGQLAdminArticleFromDBArticle(dst), nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*adminmodel.Article, error) {
	panic(fmt.Errorf("not implemented"))
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
