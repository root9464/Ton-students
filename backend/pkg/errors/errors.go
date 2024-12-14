package errors

import "fmt"

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Error представляет собой структуру ошибки с кодом и сообщением.
type Error struct {
	Code    int     // Код ошибки (например, HTTP статус)
	Message Message // Описание ошибки
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Новый способ создания ошибки с кодом и сообщением
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: Message{Status: "fail", Message: message},
	}
}

func GetErroField(err error) (int, Message) {
	if errWrap, ok := err.(*Error); ok {
		return errWrap.Code, errWrap.Message
	}
	return 500, Message{Status: "fail", Message: "Internal server error"}
}
