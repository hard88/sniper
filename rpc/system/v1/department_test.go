package system_v1

import (
	"cloudDesktop/cmd/job"
	"cloudDesktop/util/conf"
	"cloudDesktop/util/test/assert"
	"context"
	"log"
	"testing"
)

func TestDepartment(t *testing.T) {
	ctx := context.Background()
	// set db
	conf.Set("DB_DEFAULT_DSN", conf.Get("DB_TEST_DSN"))

	// migrate database table form
	err := job.Install_test(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// run test
	s := new(departmentServer)
	resp, err := s.List(ctx, &DepartmentListReq{Page: 0, Size: 20})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, &DepartmentListResp{Code: 100, Msg: ""}, resp)
	// add two

	// list
	// delete
	// list

}
