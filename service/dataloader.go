package service

import (
	"context"
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

const defaultWaitTime = 2 * time.Millisecond
const defaultMaxBatch = 100

var articleByIDLoader *dbmodel.ArticleLoader = nil

func ArticleByIDLoader(ctx context.Context) *dbmodel.ArticleLoader {
	if articleByIDLoader == nil {
		articleByIDLoader = dbmodel.NewArticleLoader(dbmodel.ArticleLoaderConfig{
			Fetch:    func(keys []int) ([]*dbmodel.Article, []error) { return fetchArticleByID(ctx, keys) },
			Wait:     defaultWaitTime,
			MaxBatch: defaultMaxBatch,
		})
	}
	return articleByIDLoader
}

func fetchArticleByID(ctx context.Context, articleIDs []int) ([]*dbmodel.Article, []error) {
	articles := make([]*dbmodel.Article, len(articleIDs))
	errors := make([]error, len(articleIDs))

	iArticleIDs := make([]interface{}, len(articleIDs))
	for i, v := range articleIDs {
		iArticleIDs[i] = v
	}

	rows, err := dbmodel.Articles(qm.WhereIn(dbmodel.ArticleColumns.ArticleID+" IN ?", iArticleIDs...)).AllG(ctx)
	if err != nil {
		for i, _ := range articleIDs {
			errors[i] = err
		}
	}

	for i, id := range articleIDs {
		for _, article := range rows {
			if article.ArticleID == id {
				articles[i] = article
			}
		}
	}

	return articles, errors
}
