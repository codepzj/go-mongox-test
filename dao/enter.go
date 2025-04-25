package dao

import (
	"go-mongox-test/global"
	"go-mongox-test/model"

	"github.com/chenmingyong0423/go-mongox/v2"
)

var userDao *UserDao

type UserDao struct {
	coll *mongox.Collection[model.User]
}

func NewUserDao(db *mongox.Database) *UserDao {
	return &UserDao{coll: mongox.NewCollection[model.User](db, "user")}
}

func InitDao() {
	userDao = NewUserDao(global.DB)
}

func TestUserCreate() {
	userDao.UserInsertOne()
	userDao.UserInsertMany()
}

func TestUserFind() {
	userDao.UserFindOne()
	userDao.UserFindMany()
	userDao.UserCount()
	userDao.UserDistinct()
}

func TestUserUpdate() {
	userDao.UserUpdateOne()
	userDao.UserUpdateOneWithStruct()
	userDao.UserUpdateMany()
	userDao.UserUpdateSet()
}

func TestUserDelete() {
	userDao.UserDeleteOne()
	userDao.UserDeleteMany()
}

