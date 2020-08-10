package server

import (
	"net/http"

	"cloudDesktop/cmd/server/hook"

	"github.com/bilibili/twirp"

	system_v1 "cloudDesktop/rpc/system/v1"
	"cloudDesktop/server/systemserver1"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewCheckLogin(),
	hook.NewLog(),
)

func initMux(mux *http.ServeMux, isInternal bool) {
	if isInternal {
		// 只注册内部服务
		return
	}
	{
		server := &systemserver1.DepartmentServer{}
		handler := system_v1.NewDepartmentServer(server, hooks)
		mux.Handle(system_v1.DepartmentPathPrefix, handler)
	}
	{
		server := &systemserver1.UserServer{}
		handler := system_v1.NewUserServer(server, hooks)
		mux.Handle(system_v1.UserPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
