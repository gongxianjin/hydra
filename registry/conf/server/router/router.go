package router

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/registry/conf"
)

type Routers struct {
	Routers []*Router `json:"srvs,omitempty"`
	RPCS    []*Router `json:"gwts,omitempty"`
}

//Router 路由信息
type Router struct {
	Path    string   `json:"path,omitempty" valid:"ascii,required"`
	Action  []string `json:"action,omitempty" valid:"uppercase,in(GET|POST|PUT|DELETE|HEAD|TRACE|OPTIONS)"`
	Service string   `json:"service" valid:"ascii,required"`
	Disable bool     `json:"disable,omitempty"`
}

//NewRouters 构建路由
func NewRouters() *Routers {
	r := &Routers{
		Routers: make([]*Router, 0),
		RPCS:    make([]*Router, 0),
	}
	r.Routers = append(r.Routers, &Router{Action: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		Path: "/*name", Service: "/@name"})
	return r
}

//AppendRPCProxy 添加proxy信息
func (h *Routers) AppendRPCProxy(path string, service string) *Routers {
	router := &Router{Path: path, Service: service}
	h.RPCS = append(h.RPCS, router)
	return h
}

//Append 添加路由信息
func (h *Routers) Append(path string, service string, action ...string) *Routers {
	h.Routers = append(h.Routers, &Router{
		Path:    path,
		Service: service,
		Action:  action,
	})
	return h
}

//GetConf 设置路由
func GetConf(cnf conf.IMainConf) (router *Routers) {
	_, err := cnf.GetSubObject("router", &router)
	if err == conf.ErrNoSetting || len(router.Routers) == 0 {
		router = NewRouters()
	}
	if err != nil && err != conf.ErrNoSetting {
		panic(fmt.Errorf("获取路由(%s)失败:%w", cnf.GetMainPath(), err))
	}
	if b, err := govalidator.ValidateStruct(&router); !b {
		panic(fmt.Errorf("路由(%s)配置有误:%w", cnf.GetMainPath(), err))
	}
	return router
}
