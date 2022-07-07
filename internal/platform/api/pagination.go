package api

// Pagination represents pagination
type Pagination struct {
	Page       int `form:"page"`
	PerPage    int `form:"per_page"`
	NextPage   int `form:"-"`
	TotalPages int `form:"-"`
	TotalItems int `form:"-"`
}

// Sanitize sanitizes pagination
func (p *Pagination) Sanitize() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.PerPage < 1 {
		p.PerPage = 50
	}

	if p.PerPage > 100 {
		p.PerPage = 100
	}
}

// Offset return computed db offset
func (p Pagination) Offset() int {
	return p.PerPage * (p.Page - 1)
}
