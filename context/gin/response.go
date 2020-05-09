package gin

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro-plat/hydra/context"
)

var _ context.IResponse = &response{}

type response struct {
	*gin.Context
	specials []string
}

//Header 设置头信息到response里
func (c *response) SetHeader(k string, v string) {
	c.Context.Request.Response.Header.Set(k, v)
}

//Abort 根据错误码终止应用
func (c *response) Abort(s int) {
	c.Context.AbortWithStatus(s)
}

//AbortWithError 根据错误码与错误消息终止应用
func (c *response) AbortWithError(s int, err error) {
	c.Context.AbortWithError(s, err)
}

//GetStatusCode 获取response状态码
func (c *response) GetStatusCode() int {
	return c.Context.Writer.Status()
}

//SetStatusCode 设置response状态码
func (c *response) SetStatusCode(s int) {
	c.Context.Writer.WriteHeader(s)
}

//Written 响应是否已写入
func (c *response) Written() bool {
	return c.Context.Writer.Written()
}

//File 输入文件
func (c *response) File(path string) {
	if c.Context.Writer.Written() {
		panic(fmt.Sprint("不能重复写入到响应流::", path))
	}
	c.Context.File(path)
}

//WriteAny 将结果写入响应流，自动检查内容处理状态码
func (c *response) WriteAny(content interface{}) error {
	status := 200
	content := ""
	switch v := content.(type) {
	case IResult:
		status = v.GetCode()
		content = v.GetResult()
		return
	case IError:
		content = v.GetError()
		status = v.GetCode()
	case error:
		status = 400
		content = content.(error)
	default:
		tp := reflect.TypeOf(content).Kind()
		value := reflect.ValueOf(content)
		if tp == reflect.Ptr {
			value = value.Elem()
		}
		switch tp {
		case reflect.Chan, reflect.Map, reflect.Slice:
			if value.Len() == 0 {

			}
		}
	}
	c.Write(status, content)
	return
}
func (c *response) swap(content interface{}) interface{} {

}

//Write 将状态码、内容写入到响应流中
func (c *response) Write(s int, v string) error {
	if c.Request.Response.Header.Get("Content-Type") == "" {
		c.Request.Response.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if c.Context.Writer.Written() {
		panic(fmt.Sprint("不能重复写入到响应流:status:", s, v))
	}
	c.Context.Writer.WriteHeader(s)
	_, err := c.Context.Writer.WriteString(v)
	return err
}

//Redirect 转跳g刚才gc
func (c *response) Redirect(code int, url string) {
	c.Context.Redirect(code, url)
}

//AddSpecial 添加响应的特殊字符
func (c *response) AddSpecial(t string) {
	if c.specials == nil {
		c.specials = make([]string, 0, 1)
	}
	c.specials = append(c.specials, t)
}

//GetSpecials 获取多个响应特殊字符
func (c *response) GetSpecials() string {
	return strings.Join(c.specials, " ")
}

//GetTrace 获取trace信息
func (c *response) GetTrace() string {
	return ""
}
