package main

import (
	"go-mongox-test/dao"
)

func main() {
	// 初始化dao
	dao.InitDao()
	// 创建用户(creator操作)
	// dao.TestUserCreate()

	// 查询用户(finder操作)
	// dao.TestUserFind()

	// 更新用户(updator操作)
	// dao.TestUserUpdate()

	// 删除用户(deletor操作)
	dao.TestUserDelete()
}
