package common

import "strings"

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`

	//Support cursor with UID
	FakeCursor string `json:"fake_cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Fullfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}

	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}
