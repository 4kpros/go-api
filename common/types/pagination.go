package types

type PaginationRequest struct {
	Page  int `json:"page" query:"page" doc:"Current page" example:"1"`
	Limit int `json:"limit" query:"limit" doc:"Max items per page" example:"10"`
}

type Pagination struct {
	CurrentPage  int   `json:"currentPage"`
	NextPage     int   `json:"nextPage"`
	PreviousPage int   `json:"previousPage"`
	TotalPages   int64 `json:"totalPages"`
	Count        int64 `json:"count"`
	Limit        int   `json:"limit"`
	Offset       int   `json:"offset"`
}

func (p *Pagination) UpdateFields(count *int64) {
	p.Count = *count                             // Update count
	DivUp(count, &p.Limit, &p.TotalPages)        // Update total pages
	NextPage(&p.NextPage, &p.TotalPages)         // Update next page
	PreviousPage(&p.PreviousPage, &p.TotalPages) // Update previous page
	// Update current page
	if int64(p.CurrentPage) > p.TotalPages {
		p.CurrentPage = int(p.TotalPages)
	} else if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}
	// Update offset
	if p.CurrentPage > 1 {
		var tempOffset = (p.CurrentPage - 1)
		p.Offset = tempOffset * p.Limit
	} else {
		p.Offset = 0
	}
}

// This function round up(ceil) A/B but extremely faster
func DivUp(numerator *int64, denominator *int, result *int64) {
	*result = 1 + (*numerator-1)/int64(*denominator)
}

func NextPage(page *int, totalPages *int64) {
	if int64(*page) <= 0 {
		*page = 1
	}
	if int64(*page) < *totalPages {
		*page++
		return
	}
	*page = int(*totalPages)
}

func PreviousPage(page *int, totalPages *int64) {
	if int64(*page) > *totalPages {
		*page = int(*totalPages)
	}
	if int64(*page) > 1 {
		*page--
		return
	}
	*page = 1
}
