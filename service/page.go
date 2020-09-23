package service

import (
	"github.com/hiroyky/nikki_backend/domain/gql/adminmodel"
	"github.com/hiroyky/nikki_backend/domain/gql/viewermodel"
)

// @return limit, offset
func ValidateViewerPagination(page *viewermodel.Pagination, defaultLimit, maxLimit int) (int, int) {
	if page == nil {
		return defaultLimit, 0
	}

	return validateLimit(page.Limit, defaultLimit, maxLimit), validateOffset(page.Offset)
}

// @return limit, offset
func ValidateAdminPagination(page *adminmodel.Pagination, defaultLimit, maxLimit int) (int, int) {
	if page == nil {
		return defaultLimit, 0
	}

	return validateLimit(page.Limit, defaultLimit, maxLimit), validateOffset(page.Offset)
}

func validateLimit(limit *int, def, max int) int {
	if limit == nil {
		return def
	}
	if *limit < 0 {
		return def
	}
	if *limit > max {
		return max
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
