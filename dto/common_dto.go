package dto

// ======================================
// - common ID
type CommonIDDto struct {
	ID uint `json:"id" param:"id" validate:"required"`
}

// ======================================
// - Paging Dto
type PagingDto struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (m *PagingDto) GetPage() int {
	if m.Page <= 0 {
		m.Page = 1
	}
	return m.Page
}

func (m *PagingDto) GetLimit() int {
	if m.Limit <= 0 {
		m.Limit = 1
	}
	return m.Limit
}
