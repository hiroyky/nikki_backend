package gqlmodel

import (
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/lib"
)

func GenArticle(article *dbmodel.Article) *Article {
	return &Article{
		ID:             lib.EncodeGraphQLID("Article", article.ArticleID),
		Title:          article.Title,
		Body:           article.Body,
		Description:    article.Description,
		PublishStatus:  article.PublishStatus,
		ThumbnailImage: article.ThumbnailImage,
		PostedAt:       article.PostedAt,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
	}
}
