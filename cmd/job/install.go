package job

import (
	"cloudDesktop/dao/system/department"
	"cloudDesktop/util/log"
	"context"
)

//版本格式：主版本号.次版本号.修订号，版本号递增规则如下：
//
//主版本号：当你做了不兼容的 API 修改，
//次版本号：当你做了向下兼容的功能性新增，
//修订号：当你做了向下兼容的问题修正。
// example: 1.0.0

// install 注册所有安装脚本
func install() {
	// 首次部署时调用, 执行所有安装脚本
	manual("init", install_init)
	// v0.0.0
	manual("0.0.0", install_000)
	// v0.1.0
	manual("0.1.0", install_010)
}

// 用于测试时的生成表结构
func Install_test(ctx context.Context) error {
	// Todo: 修改环境变量
	log.Get(ctx).Infoln("install test")
	// create table
	if err := install_010(ctx); err != nil {
		return err
	}
	// drop table

	return nil
}

func install_init(ctx context.Context) error {
	log.Get(ctx).Infoln("install init")
	if err := install_000(ctx); err != nil {
		return err
	}
	if err := install_010(ctx); err != nil {
		return err
	}

	return nil
}

// v0.0.0
// 初始化项目
// 设置环境变量，配置文件参数等
func install_000(ctx context.Context) error {
	log.Get(ctx).Infoln("install version 0.0.0", "初始化项目")
	// 初始化项目
	// 设置环境变量

	return nil
}

// system 模块相关
func install_010(ctx context.Context) error {
	log.Get(ctx).Infoln("install version 0.1.0", "安装departments表")

	err := department.CreateTable(ctx)

	// init user table

	return err
}

// Todo: Up to date 更新至最新版本
//  初步思想，用table记录当前版本号，执行up to date时 顺序执行旧版本和新版本之间的安装脚本

// 更新版本
func upToDate() {

}
