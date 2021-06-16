package routes

type errorJSON struct {
	Msg        string `json:"msg"`
	ErrMsg     error  `json:"errMsg"`
	StatusCode int    `json:"statusCode"`
}
