package entity

// RequestMeta hold request information
type RequestMeta struct {
	Limit  int
	Offset int
}

// ResponseMeta hold response information
type ResponseMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
