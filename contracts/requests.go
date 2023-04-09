package contracts

type AppAddRequest struct {
	AppName  *string `json:"App_Name"`
	AppImage *string `json:"App_Image"`
}

type NodeAddRequest struct {
	NodeIp *string `json:"address"`
	NodeId *string `json:"node_id"`
}

type NodeDeleteRequest struct {
	NodeIp *string `json:"address"`
	NodeId *string `json:"node_id"`
}

type NodeContainerRequest struct {
	AppImage      *string `json:"App_Image"`
	NewContainers *string `json:"Number_Containers"`
}

type NodeContainerResponse struct {
	Port *string `json:"Port_number"`
}
