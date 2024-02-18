package logiares

type Pagination struct {
	Count       int
	CurrentPage any
	PerPage     int
	TotalPage   int
}

func (p *Pagination) ParsePagination() interface{} {
	return map[string]interface{}{
		"count":       p.Count,
		"currentPage": p.CurrentPage,
		"perPage":     p.PerPage,
		"totalPage":   p.TotalPage,
	}
}
