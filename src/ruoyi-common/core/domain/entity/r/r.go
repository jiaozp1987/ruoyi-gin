package r

import "net/http"

var (
	SUCCESS = http.StatusOK
	FAIL    = http.StatusInternalServerError
)

type R[T any] struct {
	Code int
	Msg  string
	Data T
}

func NewR[T any](data T) *R[T] {
	return &R[T]{
		Code: 0,
		Msg:  "",
		Data: data,
	}
}

func (r *R[T]) Ok() {
	r.Msg = "操作成功"
	r.Code = http.StatusOK
}
func (r *R[T]) OkDefault(data T) {
	r.Ok()
	r.Data = data
}
func (r *R[T]) OkFullArgss(data T, msg string) {
	r.OkDefault(data)
	r.Msg = msg
}

func (r *R[T]) Fail(msg string) {
	r.Msg = msg
}

func (r *R[T]) RestResult(data T, code int, msg string) *R[T] {
	newR := NewR(data)
	newR.Msg = msg
	newR.Code = code
	return newR
}

func (r *R[T]) IsSuccess() bool {
	return SUCCESS == r.Code
}
