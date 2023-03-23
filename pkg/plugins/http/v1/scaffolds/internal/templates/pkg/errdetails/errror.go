package errdetails

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &ErrorDetail{}

// ErrorDetail scaffolds a file that defines the config package
type ErrorDetail struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *ErrorDetail) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "errdetails", "error.go")
	}

	f.TemplateBody = errdetailsTemplate

	return nil
}

var errdetailsTemplate = `{{ .Boilerplate }}

package errdetails

import "net/http"

// BizError
// 业务自定义 error
type BizError struct {
	// Code 是错误码，大部分情况下等于 HTTP 状态码，也可以定义自己的服务错误码
	Code int ` + "`" + `json:"code,omitempty" example:"400"` + "`" + `
	// Reason 是具体的错误原因
	Reason string ` + "`" + `json:"reason,omitempty" example:"Bad Request"` + "`" + `
	// Metadata 是错误携带的元数据，在错误中可以填入一些自定义字段来保存出现错误时的上下文信息
	Metadata map[string]string ` + "`" + `json:"metadata,omitempty" example:"user_id:foo"` + "`" + `
}

func (e BizError) Error() string {
	return e.Reason
}

func New(code int, err error, metadata map[string]string) BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return BizError{
		Code:     code,
		Reason:   errMsg,
		Metadata: metadata,
	}
}

func BadRequest(err error, metadata map[string]string) BizError {
	return New(http.StatusBadRequest, err, metadata)
}

func InternalError(err error, metadata map[string]string) BizError {
	return New(http.StatusInternalServerError, err, metadata)
}

func Conflict(err error, metadata map[string]string) BizError {
	return New(http.StatusConflict, err, metadata)
}

func Unauthorized(err error, metadata map[string]string) BizError {
	return New(http.StatusUnauthorized, err, metadata)
}

func Forbidden(err error, metadata map[string]string) BizError {
	return New(http.StatusForbidden, err, metadata)
}

func NotFound(err error, metadata map[string]string) BizError {
	return New(http.StatusNotFound, err, metadata)
}

func TooManyRequest(err error, metadata map[string]string) BizError {
	return New(http.StatusTooManyRequests, err, metadata)
}

`
