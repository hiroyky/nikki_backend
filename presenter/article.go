package presenter

import (
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/lib"
)

func ToDBArticleFromAdminArticleInput(input *adminmodel.ArticleMutationInput) dbmodel.Article {
	return dbmodel.Article{
		Title:          input.Title,
		Body:           input.Body,
		Description:    input.Description,
		PublishStatus:  input.PublishStatus,
		ThumbnailImage: input.ThumbnailImage,
		PostedAt:       input.PostedAt,
	}
}

func ToGQLAdminArticleFromDBArticle(input *dbmodel.Article) *adminmodel.Article {
	return &adminmodel.Article{
		ID:             lib.EncodeGraphQLID(dbmodel.TableNames.Articles, input.ArticleID),
		Title:          input.Title,
		Body:           input.Body,
		Description:    input.Description,
		PublishStatus:  input.PublishStatus,
		ThumbnailImage: input.ThumbnailImage,
		PostedAt:       input.PostedAt,
		CreatedAt:      input.CreatedAt,
		UpdatedAt:      input.UpdatedAt,
	}
}
