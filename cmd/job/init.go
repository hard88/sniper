package job

func init() {
	// 注册所有安装脚本
	install()

	//manual("foo", func(ctx context.Context) error {
	//	fmt.Printf("manual run foo with args: %+v\n", onceArgs)
	//	return nil
	//})

	//cron("bar", "@every 1m", func(ctx context.Context) error {
	//	fmt.Printf("run bar @%v\n", time.Now())
	//	return nil
	//})
	//
	//http("baz", "0 18-23 * * *", func(ctx context.Context) error {
	//	fmt.Printf("run http task @%v\n", time.Now())
	//	return nil
	//})
}
