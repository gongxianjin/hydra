package ctx

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/conf/server"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/encoding"
	"github.com/micro-plat/lib4go/types"
)

type request struct {
	ctx        context.IInnerContext
	serverConf server.IServerConf
	*body
	path *rpath
}

//newRequest 构建请求的Request
//自动对请求进行解码，响应结果进行编码。
//当指定为gbk,gb2312后,请求方式为application/x-www-form-urlencoded或application/xml、application/json时内容必须编码为指定的格式，否则会解码失败
func newRequest(c context.IInnerContext, s server.IServerConf) *request {
	rpath := newRpath(c, s)
	return &request{
		ctx:  c,
		body: newBody(c, rpath),
		path: rpath,
	}
}

//Path 获取请求路径信息
func (r *request) Path() context.IPath {
	return r.path
}

//Path 获取请求路径信息
func (r *request) Param(key string) string {
	return r.ctx.Param(key)
}

//Bind 根据输入参数绑定对象
func (r *request) Bind(obj interface{}) error {
	if err := r.ctx.ShouldBind(&obj); err != nil {
		return err
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	if _, err := govalidator.ValidateStruct(obj); err != nil {
		err = fmt.Errorf("输入参数有误 %v", err)
		return err
	}
	return nil
}

//Check 检查输入参数和配置参数是否为空
func (r *request) Check(field ...string) error {
	data, _ := r.body.GetBodyMap()
	for _, key := range field {
		if _, ok := r.ctx.GetFormValue(key); ok {
			continue
		}
		if _, ok := r.ctx.GetQuery(key); ok {
			continue
		}
		if v, ok := data[key]; !ok || fmt.Sprint(v) == "" {
			return fmt.Errorf("输入参数:%s值不能为空", key)
		}
	}
	return nil
}

//GetKeys 获取字段名称
func (r *request) GetKeys() []string {
	keys := make([]string, 0, len(r.ctx.GetForm()))
	for k := range r.ctx.GetForm() {
		keys = append(keys, k)
	}
	data, _ := r.body.GetBodyMap()
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

//GetData 获取请求的参数信息
func (r *request) GetData() (map[string]interface{}, error) {
	forms := r.ctx.GetForm()
	body, err := r.body.GetBodyMap()
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	for k, v := range forms {
		data[k] = v[0]
	}
	for k, v := range body {
		data[k] = v
	}

	return data, nil

}

//Get 获取字段的值
func (r *request) Get(name string) (result string, ok bool) {
	var fromBody bool
	defer func() {
		if ok && !fromBody { //只对url,form中的参数进行解码,body中的参数已经解码了无需再解码
			u, err := url.QueryUnescape(result)
			if err != nil {
				panic(fmt.Errorf("url.unescape出错:%w", err))
			}
			rx, err := encoding.Decode(u, r.path.GetRouter().GetEncoding())
			if err != nil {
				result = u
				return
			}
			result = string(rx)
		}
	}()

	if result, ok = r.ctx.GetFormValue(name); ok {
		return
	}
	fromBody = true
	m, err := r.body.GetBodyMap()
	if err != nil {
		return "", false
	}
	v, b := m[name]

	return fmt.Sprint(v), b
}

//GetString 获取字符串
func (r *request) GetString(name string, def ...string) string {
	if v, ok := r.Get(name); ok {
		return v
	}
	return types.GetStringByIndex(def, 0, "")
}

func (r *request) GetInt(name string, def ...int) int {
	v, _ := r.Get(name)
	return types.GetInt(v, def...)
}

func (r *request) GetMax(name string, o ...int) int {
	v := r.GetInt(name, o...)
	return types.GetMax(v, o...)

}
func (r *request) GetMin(name string, o ...int) int {
	v := r.GetInt(name, o...)
	return types.GetMin(v, o...)
}
func (r *request) GetInt64(name string, def ...int64) int64 {
	v, _ := r.Get(name)
	return types.GetInt64(v, def...)
}
func (r *request) GetFloat32(name string, def ...float32) float32 {
	v, _ := r.Get(name)
	return types.GetFloat32(v, def...)
}
func (r *request) GetFloat64(name string, def ...float64) float64 {
	v, _ := r.Get(name)
	return types.GetFloat64(v, def...)
}
func (r *request) GetBool(name string, def ...bool) bool {
	v, _ := r.Get(name)
	return types.GetBool(v, def...)
}
func (r *request) GetDatetime(name string, format ...string) (time.Time, error) {
	v, _ := r.Get(name)
	return types.GetDatetime(v, format...)
}
func (r *request) IsEmpty(name string) bool {
	_, ok := r.Get(name)
	return ok
}

//GetTrace 获取trace信息
func (r *request) GetTrace() string {
	data, err := r.GetData()
	if err != nil {
		return err.Error()
	}
	if buff, err := json.Marshal(data); err == nil {
		return string(buff)
	}
	return ""

}
