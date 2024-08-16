package utils

type HttpError struct {
	Status int // Http ステータスコード
	Message string // エラーメッセージ
	Success bool  //成功したか
}

func NewHttpError(StatusCode int,Message string) *HttpError {
	return &HttpError{
		Status: StatusCode,
		Message: Message,
	}
}

func (err *HttpError) Error() string {
	return err.Message
}