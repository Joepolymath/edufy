package common

type APIResponse struct {
	Message any         `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
	Detail  any         `json:"detail,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
