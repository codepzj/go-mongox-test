package dao

import (
	"context"
	"go-mongox-test/model"
	"log"
)

func (d *UserDao) UserInsertOne() {
	insertOneResult, err := d.coll.Creator().InsertOne(context.Background(), &model.User{Name: "Mingyong Chen", Age: 18})
	if err != nil {
		panic(err)
	}
	if insertOneResult.InsertedID == nil {
		panic("insertOneResult.InsertedID is nil")
	}
	log.Println("插入单条数据成功", insertOneResult.InsertedID)
}

func (d *UserDao) UserInsertMany() {
	insertManyResult, err := d.coll.Creator().InsertMany(context.Background(), []*model.User{
		{Name: "Mingyong Chen", Age: 18},
		{Name: "codepzj", Age: 18},
	})
	if err != nil {
		panic(err)
	}
	log.Println("插入多条数据成功", insertManyResult.InsertedIDs)
}
