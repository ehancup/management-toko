package br

type BaseSuccessResponse[T interface{}] struct {
	Message string `json:"message"`
	// Success bool   `json:"success" binding:"default:true" example:"true"`
	Data    T      `json:"data"`
}

type BaseSuccessResponsePagination[T any] struct {
	Message    string             `json:"message"`
	// Success    bool               `json:"success" binding:"default:true" example:"true"`
	Data       T                  `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
