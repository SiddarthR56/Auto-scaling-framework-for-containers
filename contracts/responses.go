package contracts

type BaseResponse struct {
	Error   *string `json:"error"`
	Message *string `json:"message"`
}

type NodeContainerResponse struct {
	Port *string `json:"Port_number"`
}
