package pkg_success

type ClientSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ClientSuccessWithPagination struct {
	Message     string `json:"message"`
	Data        any    `json:"data"`
	CurrentPage int    `json:"current_page"`
	TotalPage   int    `json:"total_page"`
	TotalData   int    `json:"total_data"`
	PerPage     int    `json:"per_page"`
}

func SuccessMessage(message string) *ClientSuccess {
	return &ClientSuccess{
		Message: message,
	}
}

func SuccessWithData(data any) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}

func SuccessWithPagination(data any, currentPage int, totalPage int, perPage int, totalData int) *ClientSuccessWithPagination {
	return &ClientSuccessWithPagination{
		Message:     "Success",
		Data:        data,
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		PerPage:     perPage,
		TotalData:   totalData,
	}
}
