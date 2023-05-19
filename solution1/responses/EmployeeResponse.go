package responses

type SuccessResponse struct {

	Status int         `json:"status"`
	Message string     `json:"message"`
	Data map[string]interface{} `json:"data"`
}

type ErrorResponse struct {
	Status int         `json:"status"`
	Message string     `json:"message"`
}
