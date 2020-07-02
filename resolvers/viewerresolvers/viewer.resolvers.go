package viewerresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hiroyky/nikki_backend/domain/gql/viewermodel"
	"github.com/hiroyky/nikki_backend/lib"
	"github.com/hiroyky/nikki_backend/presenter"
	"github.com/hiroyky/nikki_backend/service"
)

func (r *queryResolver) Article(ctx context.Context, id string) (*viewermodel.Article, error) {
	dbID, err := lib.DecodeGraphQLID(id)
	if err != nil {
		return nil, err
	}

	dst, err := service.GetArticle(ctx, dbID)
	if err != nil {
		return nil, err
	}

	return presenter.ToGQLViewerArticleFromDBArticle(dst), nil
}

func (r *queryResolver) Articles(ctx context.Context, page *viewermodel.Pagination) (*viewermodel.ArticleConnection, error) {
	limit, offset := service.ValidateViewerPagination(page, 100)
	articles, err := service.FindArticles(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	totalCount, err := service.CountArticles(ctx)
	if err != nil {
		return nil, err
	}
	return presenter.ToGQLViewerArticleConnectionFromDBArticles(articles, int(totalCount), limit, offset), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
