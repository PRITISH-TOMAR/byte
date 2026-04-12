package response

type APIResponse[T any] struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    *T          `json:"data,omitempty"`
    Error   *ErrorBody  `json:"error,omitempty"`
}

type ErrorBody struct {
    Code    string `json:"code"`
    Details string `json:"details,omitempty"`
}