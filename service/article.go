package service

import (
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/lib"
)

func ToDBArticleFromAdminArticleInput(input *adminmodel.ArticleInput) *dbmodel.Article {
	article := &dbmodel.Article{
		Title:          input.Title,
		Body:           input.Body,
		Description:    input.Description,
		PublishStatus:  input.PublishStatus,
		ThumbnailImage: input.ThumbnailImage,
		PostedAt:       input.PostedAt,
	}
	return article
}

func UpdateDBArticleFromAdminArticleInput(article *dbmodel.Article, input *adminmodel.ArticleInput) *dbmodel.Article {
	dest := ToDBArticleFromAdminArticleInput(input)
	dest.ArticleID = article.ArticleID
	return dest
}

func ToGQLAdminArticleFromDBArticle(input *dbmodel.Article) *adminmodel.Article {
	article := &adminmodel.Article{
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
	return article
}
