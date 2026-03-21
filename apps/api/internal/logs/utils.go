package logs

func BuildLimitOffset(page int, pageSize int) (int, int) {
	limit := pageSize * page
	offset := (page - 1) * pageSize
	return limit, offset
}
