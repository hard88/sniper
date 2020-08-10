package systemserver1

import (
	"cloudDesktop/service/system/department"
	"cloudDesktop/util/code"
	"cloudDesktop/util/errors"
	"cloudDesktop/util/log"
	"context"
	"github.com/bilibili/twirp"

	pb "cloudDesktop/rpc/system/v1"
)

// DepartmentServer 实现 /twirp/system.v1.Department 服务
// 实现部门的相关接口
type DepartmentServer struct{}

// List 实现 /twirp/system.v1.Department/List 接口
// 实现获取部门列表
func (s *DepartmentServer) List(ctx context.Context, req *pb.DepartmentListReq) (resp *pb.DepartmentListResp, err error) {
	v, ok := twirp.MethodOption(ctx)
	log.Get(ctx).Infoln("method option", v, ok)
	// 处理入参数
	if req.Size <= 0 {
		req.Size = 20
	}
	if req.Page--; req.Page < 0 {
		req.Page = 0
	}
	// 请求服务
	departments, err := department.List(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &pb.DepartmentListResp{Code: code.Success, Msg: "Success", Data: departments}, nil
}

// Add 实现 /twirp/system.v1.Department/Add 接口
// 实现添加部门
func (s *DepartmentServer) Add(ctx context.Context, req *pb.DepartmentAddReq) (resp *pb.DepartmentAddResp, err error) {
	d := req.GetDepartment()
	if len(d.Name) == 0 {
		return nil, errors.InvalidArgumentError("name", "不能为空")
	}
	id, err := department.Add(ctx, d)
	if err != nil {
		return nil, err
	}
	return &pb.DepartmentAddResp{
		Code: code.AddSuccess,
		Msg:  "添加部门成功",
		Data: id, // Todo last id
	}, nil
}

// Add 实现 /twirp/system.v1.Department/Add 接口
// 实现删除部门
func (s *DepartmentServer) Delete(ctx context.Context, req *pb.DepartmentDeleteReq) (resp *pb.DepartmentDeleteResp, err error) {
	if req.Id == 0 {
		return nil, errors.InvalidArgumentError("id", "不能为空")
	}
	err = department.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DepartmentDeleteResp{
		Code: code.DeleteSuccess,
		Msg:  "删除部门成功",
	}, nil
}
