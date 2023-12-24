package request

type UserDownload struct {
	ResourceId   int32 `json:"resource_id"`
	ResourceType int32 `json:"resource_type"`
}
