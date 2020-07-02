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

func ToGQLAdminArticleConnectionFromDBArticles(inputs []*dbmodel.Article) *adminmodel.ArticleConnection {
	nodes := make([]*adminmodel.Article, len(inputs))
	edges := make([]*adminmodel.ArticleEdge, len(inputs))
	for i, v := range inputs {
		node := ToGQLAdminArticleFromDBArticle(v)
		nodes[i] = node
		edges[i] = &adminmodel.ArticleEdge{
			Cursor: node.ID,
			Node:   node,
		}
	}

	connection := &adminmodel.ArticleConnection{
		PageInfo: nil,
		Edges:    edges,
		Nodes:    nodes,
	}

	return connection
}
