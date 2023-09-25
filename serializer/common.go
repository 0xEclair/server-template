package serializer

type Response struct {
	Code int32       `json:"code"`
	Data interface{} `json:"data"`
	Err  string      `json:"err,omitempty"`
}
