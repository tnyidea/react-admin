package httpserver

type DataProviderDataResponse struct {
	Status  *int        `json:"-"`
	Message *string     `json:"-"`
	Data    interface{} `json:"data,omitempty"`
	Total   *int        `json:"total,omitempty"`
}

type DataProviderErrorResponse struct {
	Status  *int    `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
