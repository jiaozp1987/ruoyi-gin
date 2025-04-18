package ajaxResult

import "net/http"

var ( // 多变量分组声明
	CODE_TAG = "code"
	MSG_TAG  = "msg"
	DATA_TAG = "data"
)

type AjaxResutlt struct {
	Result map[string]interface{}
}

func NewAjaxResutlt() *AjaxResutlt {
	return NewAjaxResutltFactoryWithData(0, "", nil)
}

func NewAjaxResutltDefault(code int, msg string) *AjaxResutlt {
	return NewAjaxResutltFactoryWithData(code, msg, nil)
}

func NewAjaxResutltFactoryWithData(code int, msg string, data interface{}) *AjaxResutlt {
	result := make(map[string]interface{})
	result[CODE_TAG] = code
	result[MSG_TAG] = msg
	result[DATA_TAG] = data
	return &AjaxResutlt{Result: result}
}

func (ar *AjaxResutlt) SuccessFullArgs(msg string, data interface{}) {
	ar.Result[CODE_TAG] = http.StatusOK
	ar.Result[MSG_TAG] = msg
	ar.Result[DATA_TAG] = data
}

func (ar *AjaxResutlt) SuccessWithMsg(msg string) {
	ar.SuccessFullArgs(msg, nil)
}

func (ar *AjaxResutlt) SuccessWithData(data interface{}) {
	ar.SuccessFullArgs("操作成功", data)
}

func (ar *AjaxResutlt) SuccessDefault() {
	ar.SuccessFullArgs("操作成功", nil)
}

func (ar *AjaxResutlt) WarnFullArgs(msg string, data interface{}) {
	ar.Result[CODE_TAG] = http.StatusInternalServerError
	ar.Result[MSG_TAG] = msg
	ar.Result[DATA_TAG] = data
}

func (ar *AjaxResutlt) WarnDefault(msg string) {
	ar.WarnFullArgs(msg, nil)
}

func (ar *AjaxResutlt) ErrorFullArgs(msg string, data interface{}) {
	ar.Result[CODE_TAG] = http.StatusInternalServerError
	ar.Result[MSG_TAG] = msg
	ar.Result[DATA_TAG] = data
}

func (ar *AjaxResutlt) ErrorWithMsg(msg string) {
	ar.ErrorFullArgs(msg, nil)
}

func (ar *AjaxResutlt) ErrorWithData(data interface{}) {
	ar.ErrorFullArgs("操作失败", data)
}

func (ar *AjaxResutlt) ErrorDefault() {
	ar.SuccessFullArgs("操作失败", nil)
}

func (ar *AjaxResutlt) IsSuccess() bool {
	return http.StatusOK == ar.Result[CODE_TAG]
}

func (ar *AjaxResutlt) IsError() bool {
	return http.StatusInternalServerError == ar.Result[CODE_TAG]
}
