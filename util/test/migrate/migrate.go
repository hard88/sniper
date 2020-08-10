package migrate

import (
	"cloudDesktop/cmd/job"
	"cloudDesktop/util/conf"
	"cloudDesktop/util/db"
	"cloudDesktop/util/log"
	"context"
)

// DBMigrate 连接目标DSN并执行数据库安装脚本, 通常dsn为test
// 并不是严格意义上的db_default 迁移至 db_test
func DBMigrate(ctx context.Context, DSN string) (clean func(), err error) {
	conf.Set("DB_DEFAULT_DSN", DSN)

	err = job.Install_test(ctx)
	if err != nil {
		return nil, err
	}

	// clean 清空数据库表结构
	clean = func() {
		c := db.Get(ctx, "default")
		log.Get(ctx).Infoln("clean up")
		query := db.SQLUpdate("all", `
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

GRANT ALL ON SCHEMA public TO wrh;
GRANT ALL ON SCHEMA public TO public;
`)
		// Todo: 修改user
		//  这里要返还grant to user，这里user为wrh，以后统一为postgres
		_, err := c.ExecContext(ctx, query)
		if err != nil {
			log.Get(ctx).Errorf("drop table: %v", err)
		}
	}

	return clean, err
}
