package config

import (
	m "github.com/i-coder-robot/go-micro-action-user/middleware"
	PB "github.com/i-coder-robot/go-micro-action-user/user/proto/permission"
)

// Config 配置
type Config struct {
	Name        string
	Version     string
	Service     map[string]string
	Permissions []*PB.Permission
}

// Middleware 用户中间件初始化
func (srv *Config) Middleware() *m.Handler {
	return &m.Handler{
		UserService: srv.Service["user"],
		Permissions: srv.Permissions,
	}
}
