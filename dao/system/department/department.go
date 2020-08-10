package department

import (
	"cloudDesktop/util/db"
	"cloudDesktop/util/errors"
	"cloudDesktop/util/log"
	"context"
	"time"
)

// 部门
type Department struct {
	// id 唯一标志
	ID int64
	// 名字
	Name string
	// 所属department
	Parent int64
	//
	Description string
	// 创建时间
	Created int64
	// 删除时间，大于零为删除
	Deleted int64
}

func CreateTable(ctx context.Context) error {
	c := db.Get(ctx, "default")

	sql := `CREATE TABLE departments (
	id BIGSERIAL,
    name VARCHAR(255) NOT NULL UNIQUE,
	parent BIGINT NOT NULL,
    description TEXT,
	created BIGINT,
    deleted BIGINT
	);`

	q := db.SQLCreate("department", sql)
	_, err := c.ExecContext(ctx, q)
	return err
}

func Get(ctx context.Context, id int64) (*Department, error) {
	c := db.Get(ctx, "default")
	sql := db.SQLSelect("department", `select * from departments where id = $1 AND deleted = 0`)
	result := c.QueryRowContext(ctx, sql, id)

	d := new(Department)
	if err := result.Scan(&d.ID, &d.Name, &d.Parent, &d.Created, &d.Parent); err != nil {
		if db.IsNoRowsErr(err) {
			log.Get(ctx).Debug("not exiest")
			return d, nil
		}
		return nil, errors.Wrap(err, "scan")
	}

	return d, nil
}

// Todo: 查询条件, limit,page,order by
func List(ctx context.Context, page, size int32) ([]*Department, error) {
	c := db.Get(ctx, "default")

	departments := make([]*Department, 0, size)

	query := `
SELECT *
FROM departments
WHERE deleted = 0
LIMIT $1 OFFSET $2
`

	sql := db.SQLSelect("department", query)
	results, err := c.QueryContext(ctx, sql, size, page*size)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		department := new(Department)
		err := results.Scan(&department.ID, &department.Name, &department.Parent, &department.Description, &department.Created, &department.Deleted)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}
	return departments, nil
}

func Add(ctx context.Context, department *Department) (int64, error) {
	c := db.Get(ctx, "default")
	var lastInsertID int64
	query := ` INSERT INTO departments(name,parent,description,created,deleted) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	sql := db.SQLInsert("department", query)
	err := c.QueryRowContext(ctx, sql,
		department.Name, department.Parent, department.Description, time.Now().Unix(), 0).Scan(&lastInsertID)
	if db.IsDuplicateObject(err) {
		return 0, errors.DuplicateObjectError
	}

	return lastInsertID, err
}

//func Patch(ctx context.Context) error {
//
//}

// 删除部门，把department标记为deleted
func Delete(ctx context.Context, id int64) error {
	c := db.Get(ctx, "default")
	query := `UPDATE departments SET deleted = $1 WHERE id = $2`
	sql := db.SQLDelete("department", query)
	_, err := c.ExecContext(ctx, sql, time.Now().Unix(), id)
	return err
}
