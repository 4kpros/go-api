package types

type Pagination struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	TotalPages   int64
	Count        int64
	Limit        int
	Offset       int
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

func DivUp(n *int64, d *int, r *int64) {
	*r = 1 + (*n-1)/int64(*d)
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