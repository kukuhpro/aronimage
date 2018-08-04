package core

type response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Token   map[string]interface{} `json:"token"`
	Data    map[string]interface{} `json:"data"`
}

type ApiResponse struct {
	status  int
	message string
	data    map[string]interface{}
	token   map[string]interface{}
}

func (api *ApiResponse) SetStatus(status int) {
	api.status = status
}

func (api *ApiResponse) SetMessage(message string) {
	api.message = message
}

func (api *ApiResponse) SetData(data map[string]interface{}) {
	api.data = data
}

func (api *ApiResponse) SetToken(tokendata map[string]interface{}) {
	api.token = tokendata
}

func (api *ApiResponse) ToResponse() response {
	var res response

	res.Status = api.status
	res.Message = api.message
	res.Data = api.data
	res.Token = api.token

	api.data = nil

	return res
}
