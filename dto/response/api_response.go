package response

type DataResponse struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
}
