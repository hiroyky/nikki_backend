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

func GenArticleConnection(articles []*dbmodel.Article) *ArticleConnection {
	return &ArticleConnection{
		PageInfo: &PageInfo{},
		Edges:    genArticleEdges(articles),
		Nodes:    genArticleNodes(articles),
	}
}

func genArticleNodes(articles []*dbmodel.Article) []*Article {
	result := make([]*Article, len(articles))
	for i, v := range articles {
		result[i] = GenArticle(v)
	}
	return result
}

func genArticleEdges(articles []*dbmodel.Article) []*ArticleEdge {
	result := make([]*ArticleEdge, len(articles))
	for i, v := range articles {
		dest := GenArticle(v)
		result[i] = &ArticleEdge{
			Cursor: dest.ID,
			Node:   dest,
		}
	}
	return result
}
