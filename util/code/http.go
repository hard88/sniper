package code

const (
	// 成功code: 1xx
	// 失败的code请用util/errors/errors
	// 通用的成功，通常用于其他success无法表达的时候
	Success       int32 = 100
	AddSuccess          = 101 // 添加成功
	PatchSuccess        = 102 // 更新成功
	DeleteSuccess       = 103 // 删除成功
)
