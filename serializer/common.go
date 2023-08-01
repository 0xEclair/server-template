package serializer

type Response struct {
	Code int32       `json:"len"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}
