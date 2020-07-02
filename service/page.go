package service

import (
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/viewermodel"
)

// @return limit, offset
func ValidateViewerPagination(page *viewermodel.Pagination, defaultLimit int) (int, int) {
	if page == nil {
		return defaultLimit, 0
	}

	return validateLimit(page.Limit, defaultLimit), validateOffset(page.Offset)
}

// @return limit, offset
func ValidateAdminPagination(page *adminmodel.Pagination, defaultLimit int) (int, int) {
	if page == nil {
		return defaultLimit, 0
	}

	return validateLimit(page.Limit, defaultLimit), validateOffset(page.Offset)
}

func validateLimit(limit *int, def int) int {
	if limit == nil {
		return def
	}
	if *limit < 0 || *limit > def {
		return def
	}
	return *limit
}

func validateOffset(offset *int) int {
	if offset == nil {
		return 0
	}
	if *offset < 0 {
		return 0
	}
	return *offset
}
