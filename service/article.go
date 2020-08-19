package service

import (
	"context"
	"fmt"
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

func NewArticle(ctx context.Context, article dbmodel.Article) (*dbmodel.Article, error) {
	if err := article.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}
	return &article, nil
}

func UpdateArticle(ctx context.Context, id int, article dbmodel.Article) (*dbmodel.Article, error) {
	found, err := GetArticle(ctx, id)
	if err != nil {
		return nil, err
	}

	article.ArticleID = found.ArticleID
	if _, err := article.UpdateG(ctx, boil.Infer()); err != nil {
		return nil, err
	}
	return &article, nil
}

func GetArticle(ctx context.Context, id int) (*dbmodel.Article, error) {
	return dbmodel.FindArticleG(ctx, id)
}

func FindArticles(ctx context.Context, limit, offset int) ([]*dbmodel.Article, error) {
	return dbmodel.Articles(
		qm.Limit(limit),
		qm.Offset(offset),
	).AllG(ctx)
}

func CountArticles(ctx context.Context) (int64, error) {
	return dbmodel.Articles().CountG(ctx)
}

func GetPreviousArticle(ctx context.Context, baseArticleID int, basePostedAt time.Time) (*dbmodel.Article, error) {
	res, err := dbmodel.Articles(
		qm.Where(fmt.Sprintf("%s < ?", dbmodel.ArticleColumns.PostedAt), basePostedAt),
		qm.Where(fmt.Sprintf("%s != ?", dbmodel.ArticleColumns.ArticleID), baseArticleID),
		qm.OrderBy(fmt.Sprintf("%s DESC", dbmodel.ArticleColumns.PostedAt)),
		qm.OrderBy(fmt.Sprintf("%s DESC", dbmodel.ArticleColumns.ArticleID)),
		qm.Limit(1),
	).AllG(ctx)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}

func GetNextArticle(ctx context.Context, baseArticleID int, basePostedAt time.Time) (*dbmodel.Article, error) {
	res, err := dbmodel.Articles(
		qm.Where(fmt.Sprintf("%s > ?", dbmodel.ArticleColumns.PostedAt), basePostedAt),
		qm.Where(fmt.Sprintf("%s != ?", dbmodel.ArticleColumns.ArticleID), baseArticleID),
		qm.OrderBy(fmt.Sprintf("%s ASC", dbmodel.ArticleColumns.PostedAt)),
		qm.OrderBy(fmt.Sprintf("%s ASC", dbmodel.ArticleColumns.ArticleID)),
		qm.Limit(1),
	).AllG(ctx)

	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}
