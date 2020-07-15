package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"
)

var _ conf.IPub = &Pub{}

//Pub 系统发布路径
type Pub struct {
	platName    string
	sysName     string
	serverType  string
	clusterName string
	serverID    string
	*varPath
}

//Split 根据主配置获取平台名称、系统名称、服务类型、集群名
func Split(mainConfPath string) (platName string, sysName string, serverType string, clusterName string) {
	sections := strings.Split(strings.Trim(mainConfPath, "/"), "/")
	return sections[0], sections[1], sections[2], sections[3]
}

//NewPub 构建服务发布路径信息
func NewPub(platName string, sysName string, serverType string, clusterName string) *Pub {
	return &Pub{
		platName:    platName,
		sysName:     sysName,
		serverType:  serverType,
		clusterName: clusterName,
		serverID:    global.GetMatchineCode() + time.Now().Format("0405"),
		varPath:     NewVarPath(platName),
	}
}

//GetMainPath 获取配置路径
func (c *Pub) GetMainPath() string {
	return registry.Join(c.platName, c.sysName, c.serverType, c.clusterName, "conf")
}

//GetSubConfPath 获取子配置路径
func (c *Pub) GetSubConfPath(name ...string) string {
	l := []string{c.GetMainPath()}
	l = append(l, name...)
	return registry.Join(l...)
}

//GetRPCServicePubPath 获取服务发布跟路径
func (c *Pub) GetRPCServicePubPath(svName string) string {
	return registry.Join(c.platName, "services", c.serverType, svName, "providers")
}

//GetServicePubPath 获取服务发布跟路径
func (c *Pub) GetServicePubPath() string {
	return registry.Join(c.platName, "services", c.serverType, "providers")
}

//GetDNSPubPath 获取DNS服务路径
func (c *Pub) GetDNSPubPath(svName string) string {
	return registry.Join("/dns", svName)
}

//GetServerPubPath 获取服务器发布的跟路径
func (c *Pub) GetServerPubPath(clustName ...string) string {
	if len(clustName) == 0 {
		return registry.Join(c.platName, c.sysName, c.serverType, c.clusterName, "servers")
	}
	return registry.Join(c.platName, c.sysName, c.serverType, clustName[0], "servers")
}

//GetServerID 获取当前服务的集群编号
func (c *Pub) GetServerID() string {
	return c.serverID
}

//GetPlatName 获取平台名称
func (c *Pub) GetPlatName() string {
	return c.platName
}

//GetSysName 获取系统名称
func (c *Pub) GetSysName() string {
	return c.sysName
}

//GetServerType 获取服务器类型
func (c *Pub) GetServerType() string {
	return c.serverType
}

//GetClusterName 获取集群名称
func (c *Pub) GetClusterName() string {
	return c.clusterName
}

//GetServerName 获取服务器名称
func (c *Pub) GetServerName() string {
	return fmt.Sprintf("%s(%s)", c.sysName, c.clusterName)
}

//AllowGray 是否允许灰度到其它集群
func (c *Pub) AllowGray() bool {
	return c.serverType == global.API ||
		c.serverType == global.Web
}
