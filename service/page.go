package service

func ValidateLimit(limit *int, def int) int {
	if limit == nil {
		return def
	}
	if *limit < 0 || *limit > def {
		return def
	}
	return *limit
}

func ValidateOffset(offset *int) int {
	if offset == nil {
		return 0
	}
	if *offset < 0 {
		return 0
	}
	return *offset
}
