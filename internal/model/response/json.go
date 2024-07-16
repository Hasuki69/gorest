package response

type (
	//Json[T any] struct {
	//	Message string `json:"message"`
	//	Data    *T     `json:"data"`
	//}

	Json struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
