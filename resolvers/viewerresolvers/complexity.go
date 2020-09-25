package viewerresolvers

import (
	"github.com/hiroyky/nikki_backend/domain/gql"
	"github.com/hiroyky/nikki_backend/domain/gql/viewermodel"
	"github.com/hiroyky/nikki_backend/service"
)

var ViewerComplexRoot *ComplexityRoot = nil

func init() {
	root := &ComplexityRoot{}

	root.Article.PreviousArticle = func(childComplexity int) int {
		return childComplexity + 5
	}
	root.Article.NextArticle = func(childComplexity int) int {
		return childComplexity + 5
	}

	root.Query.Articles = func(childComplexity int, sort []*gql.SortOrder, page *viewermodel.Pagination) int {
		limit, _ := service.ValidateViewerPagination(page, 10, 100)
		return limit * childComplexity
	}
	root.Query.Article = func(childComplexity int, id string) int {
		return childComplexity + 1
	}

	ViewerComplexRoot = root
}
