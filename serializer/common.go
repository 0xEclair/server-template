package serializer

type Response struct {
	Code int32       `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}
