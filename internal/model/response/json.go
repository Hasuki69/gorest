package response

type (
	Json struct {
		Message string      `json:"message"`
		Count   int         `json:"count"`
		Data    interface{} `json:"data"`
	}
)
