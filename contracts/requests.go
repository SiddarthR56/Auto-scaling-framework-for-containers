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

type ContainerRestartRequest struct {
	ContainerId *string `json:"container_id"`
}

type ContainerModifyRequest struct {
	AppName      *string `json:"app_name"`
	ContainerNum *int    `json:"container_num"`
}
