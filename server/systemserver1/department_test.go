package systemserver1

import (
	pb "cloudDesktop/rpc/system/v1"
	"cloudDesktop/util/conf"
	"cloudDesktop/util/log"
	"cloudDesktop/util/test/assert"
	"cloudDesktop/util/test/migrate"
	"context"
	"testing"
)

func TestDepartment(t *testing.T) {
	ctx := context.Background()
	log.Get(ctx).Infoln("start testing department")

	// init test db
	clean, err := migrate.DBMigrate(ctx, conf.Get("DB_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(clean)

	// run test
	// empty list not error
	s := new(DepartmentServer)
	{
		req := &pb.DepartmentListReq{Page: 0, Size: 20}
		want := &pb.DepartmentListResp{Code: 100, Msg: "Success", Data: []*pb.Dep{}}
		testDepartmentList(t, ctx, s, req, want)
	}
	{
		// add two
		dep1 := &pb.DepartmentAddReq{Department: &pb.Dep{Name: "test1", Description: "test1"}}
		want := &pb.DepartmentAddResp{Code: 100}
		testDepartmentAdd(t, ctx, s, dep1, want)

		dep2 := &pb.DepartmentAddReq{Department: &pb.Dep{Name: "test2", Description: "test2"}}
		testDepartmentAdd(t, ctx, s, dep2, want)
	}
	// list
	{
		req := &pb.DepartmentListReq{Page: 0, Size: 20}
		want := &pb.DepartmentListResp{Code: 100, Msg: "Success", Data: []*pb.Dep{
			{Id: 1, Name: "test1", Description: "test1"},
			{Id: 2, Name: "test2", Description: "test2"},
		}}
		testDepartmentList(t, ctx, s, req, want)
	}
	// delete
	{
		req := &pb.DepartmentDeleteReq{Id: 1}
		want := &pb.DepartmentDeleteResp{Code: 100}
		testDepartmentDelete(t, ctx, s, req, want)
	}
	// list
	{
		req := &pb.DepartmentListReq{Page: 0, Size: 20}
		want := &pb.DepartmentListResp{Code: 100, Msg: "Success", Data: []*pb.Dep{
			{Id: 2, Name: "test2", Description: "test2"},
		}}
		testDepartmentList(t, ctx, s, req, want)
	}
}

func testDepartmentList(t *testing.T, ctx context.Context, s *DepartmentServer, req *pb.DepartmentListReq, want *pb.DepartmentListResp) {
	resp, err := s.List(ctx, req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	if !assert.Equal(t, len(resp.Data), len(want.Data)) {
		t.Fatalf("list department get %d, want: %d", len(resp.Data), len(want.Data))
	}
}

// Todo: 表单验证测试
func testDepartmentAdd(t *testing.T, ctx context.Context, s *DepartmentServer, req *pb.DepartmentAddReq, want *pb.DepartmentAddResp) {
	resp, err := s.Add(ctx, req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	if !assert.Equal(t, resp.Code, want.Code) {
		t.Fatalf("add department get %d, want :%d", resp.Code, want.Code)
	}
}

func testDepartmentDelete(t *testing.T, ctx context.Context, s *DepartmentServer, req *pb.DepartmentDeleteReq, want *pb.DepartmentDeleteResp) {
	resp, err := s.Delete(ctx, req)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	if !assert.Equal(t, resp.Code, want.Code) {
		t.Fatalf("delete department get %d, want :%d", resp.Code, want.Code)
	}
}
