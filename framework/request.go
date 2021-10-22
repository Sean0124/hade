package framework

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

// 代表请求包含的方法
type IRequest interface {
	// 请求地址 url 中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// 路由匹配中带的参数
	// 形如 /book/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	// form 表单中带的参数
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// 获取路由参数
func (ctx *Context) Param(key string) interface{} {
	if ctx.params != nil {
		if val, ok := ctx.params[key]; ok {
			return val
		}
	}
	return nil
}

// 路由匹配中带的参数
// 形如 /book/:id
func (ctx *Context) ParmInt(key string, def int) (int, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// 基础信息
func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) Host() string {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIp() string {
	r := ctx.request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// header
func (ctx *Context) Headers() map[string][]string {
	return map[string][]string(ctx.request.Header)
}

func (ctx *Context) Header(key string) (string, bool) {
	vals := ctx.request.Header.Values(key)
	if vals == nil || len(vals) <= 0 {
		return "", false
	}
	return vals[0], true
}

// cookie
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

func (ctx *Context) Cookie(key string) (string, bool) {
	cookies := ctx.Cookies()
	if val, ok := cookies[key]; ok {
		return val, true
	}
	return "", false
}
