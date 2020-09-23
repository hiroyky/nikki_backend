package service

import (
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateAdminPagination(t *testing.T) {
	limit := 20
	offset := 5
	page := &adminmodel.Pagination{Limit: &limit, Offset: &offset}
	actualLimit, actualOffset := ValidateAdminPagination(page, 10, 100)
	expectedLimit := 20
	expectedOffset := 5

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestValidateAdminPagination_nil(t *testing.T) {
	actualLimit, actualOffset := ValidateAdminPagination(nil, 10, 100)
	expectedLimit := 10
	expectedOffset := 0

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestValidateAdminPagination_未指定(t *testing.T) {
	page := &adminmodel.Pagination{}
	actualLimit, actualOffset := ValidateAdminPagination(page, 10, 100)
	expectedLimit := 10
	expectedOffset := 0

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestValidateAdminPagination_offsetが負の値(t *testing.T) {
	offset := -1
	page := &adminmodel.Pagination{Offset: &offset}
	actualLimit, actualOffset := ValidateAdminPagination(page, 10, 100)
	expectedLimit := 10
	expectedOffset := 0

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestValidateAdminPagination_limitが負の値(t *testing.T) {
	limit := -1
	page := &adminmodel.Pagination{Limit: &limit}
	actualLimit, actualOffset := ValidateAdminPagination(page, 10, 100)
	expectedLimit := 10
	expectedOffset := 0

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestValidateAdminPagination_limitが大きすぎ(t *testing.T) {
	limit := 101
	page := &adminmodel.Pagination{Limit: &limit}
	actualLimit, actualOffset := ValidateAdminPagination(page, 10, 100)
	expectedLimit := 100
	expectedOffset := 0

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestValidateAdminPagination_境界値(t *testing.T) {
	limit := 100
	offset := 0
	page := &adminmodel.Pagination{Limit: &limit, Offset: &offset}
	actualLimit, actualOffset := ValidateAdminPagination(page, 10, 100)
	expectedLimit := 100
	expectedOffset := 0

	assert.Equal(t, expectedLimit, actualLimit)
	assert.Equal(t, expectedOffset, actualOffset)
}
