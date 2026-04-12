package response

func Success[T any](message string, data *T) APIResponse[T] {
    return APIResponse[T]{
        Success: true,
        Message: message,
        Data:    data,
    }
}

func Error(code string, message string, details string) APIResponse[any] {
    return APIResponse[any]{
        Success: false,
        Message: message,
        Error: &ErrorBody{
            Code:    code,
            Details: details,
        },
    }
}