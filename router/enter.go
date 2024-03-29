package router

import (
	"wenku/router/autocode"
	"wenku/router/example"
	"wenku/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
