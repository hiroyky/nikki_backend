package presenter

import (
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/viewermodel"
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

func ToGQLAdminArticleConnectionFromDBArticles(inputs []*dbmodel.Article, totalCount, limit, offset int) *adminmodel.ArticleConnection {
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
		PageInfo: genAdminPageInfo(len(inputs), totalCount, limit, offset),
		Edges:    edges,
		Nodes:    nodes,
	}

	return connection
}

func ToGQLViewerArticleFromDBArticle(input *dbmodel.Article) *viewermodel.Article {
	return &viewermodel.Article{
		ID:             lib.EncodeGraphQLID(dbmodel.TableNames.Articles, input.ArticleID),
		Title:          input.Title,
		Body:           input.Body,
		Description:    input.Description,
		ThumbnailImage: input.ThumbnailImage,
		PostedAt:       input.PostedAt,
		UpdatedAt:      input.UpdatedAt,
	}
}

func ToGQLViewerArticleConnectionFromDBArticles(inputs []*dbmodel.Article, totalCount, limit, offset int) *viewermodel.ArticleConnection {
	nodes := make([]*viewermodel.Article, len(inputs))
	edges := make([]*viewermodel.ArticleEdge, len(inputs))
	for i, v := range inputs {
		node := ToGQLViewerArticleFromDBArticle(v)
		nodes[i] = node
		edges[i] = &viewermodel.ArticleEdge{
			Cursor: node.ID,
			Node:   node,
		}
	}

	connection := &viewermodel.ArticleConnection{
		PageInfo: genViewerPageInfo(len(inputs), totalCount, limit, offset),
		Edges:    edges,
		Nodes:    nodes,
	}

	return connection
}

func genAdminPageInfo(count, totalCount, limit, offset int) *adminmodel.PageInfo {
	return &adminmodel.PageInfo{
		Page:             getPage(limit, offset),
		PaginationLength: getPaginationLength(limit, totalCount),
		HasNextPage:      offset+limit < totalCount,
		HasPreviousPage:  offset > 0,
		Count:            count,
		TotalCount:       totalCount,
		Limit:            limit,
		Offset:           offset,
	}
}

func genViewerPageInfo(count, totalCount, limit, offset int) *viewermodel.PageInfo {
	return &viewermodel.PageInfo{
		Page:             getPage(limit, offset),
		PaginationLength: getPaginationLength(limit, totalCount),
		HasNextPage:      offset+limit < totalCount,
		HasPreviousPage:  offset > 0,
		Count:            count,
		TotalCount:       totalCount,
		Limit:            limit,
		Offset:           offset,
	}
}

func getPage(limit, offset int) int {
	if limit == 0 {
		return 0
	}
	return offset/limit + 1
}

func getPaginationLength(limit, totalCount int) int {
	if limit == 0 {
		return 0
	}
	return totalCount/limit + 1
}
