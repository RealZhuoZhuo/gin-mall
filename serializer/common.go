package serializer

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
	Error  string `json:"error"`
}
type DataList struct {
	Items any  `json:"items"`
	Total uint `json:"total"`
}
type TokenData struct {
	User  any    `json:"user"`
	Token string `json:"token"`
}
