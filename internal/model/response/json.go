package response

type (
	Json struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
