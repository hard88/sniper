package department

import (
	"cloudDesktop/dao/system/department"
	pb "cloudDesktop/rpc/system/v1"
	"context"
)

// 获取部门列表数据处理
func List(ctx context.Context, page, size int32) ([]*pb.Dep, error) {
	departments := make([]*pb.Dep, 0, size)

	ds, err := department.List(ctx, page, size)
	if err != nil {
		return nil, err
	}
	for _, d := range ds {
		if d.Deleted > 0 {
			continue
		}
		dep := &pb.Dep{
			Id:          d.ID,
			Name:        d.Name,
			Parent:      d.Parent,
			Created:     d.Created,
			Description: d.Description,
		}
		departments = append(departments, dep)
	}
	return departments, nil
}

//  添加部门操作 数据处理
func Add(ctx context.Context, dep *pb.Dep) (int64, error) {
	// Todo:
	//  判断parent是否存在
	d := &department.Department{
		Name:        dep.Name,
		Parent:      dep.Parent,
		Description: dep.Description,
	}
	id, err := department.Add(ctx, d)
	return id, err
}

//  删除部门操作 数据处理
func Delete(ctx context.Context, id int64) error {
	// Todo: 判断部门下是否有子部门
	//  判断部门下是否有user
	return department.Delete(ctx, id)
}
