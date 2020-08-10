package db_test

import (
	"cloudDesktop/util/conf"
	"cloudDesktop/util/db"
	"cloudDesktop/util/test/assert"
	"cloudDesktop/util/test/migrate"
	"context"
	"testing"
	"time"
)

func TestLastInsertID(t *testing.T) {
	ctx := context.Background()

	clean, err := migrate.DBMigrate(ctx, conf.Get("DB_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(clean)

	var lastInsertID int64
	c := db.Get(ctx, "default")
	// 插入department数据并返回id结果
	query := `INSERT INTO departments(name,parent,description,created,deleted) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	sql := db.SQLInsert("departments", query)
	err = c.QueryRowContext(ctx, sql, "test", 0, "test", time.Now().Unix(), 0).Scan(&lastInsertID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), lastInsertID)
}
