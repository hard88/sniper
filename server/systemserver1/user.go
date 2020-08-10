package systemserver1

import (
	pb "cloudDesktop/rpc/system/v1"
	"context"
)

// UserServer 实现 /twirp/system.v1.User 服务
// FIXME 服务必须写注释
type UserServer struct{}

// Echo 实现 /twirp/system.v1.User/Login 接口
// FIXME 接口必须写注释
//
// 这里的行尾注释 sniper:foo 有特殊含义，是可选的
// 框架会将此处冒号后面的值(foo)注入到 ctx 中，
// 用户可以使用 twirp.MethodOption(ctx) 查询，并执行不同的逻辑
// 这个 sniper 前缀可以通过 --twirp_out=option_prefix=sniper:. 自定义
func (s *UserServer) Login(ctx context.Context, req *pb.UserLoginReq) (resp *pb.UserLoginResp, err error) {
	// 处理入参

	// 调用接口逻辑

	// 返回出参
	return &pb.UserLoginResp{Msg: req.Msg}, nil
}
