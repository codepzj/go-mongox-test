package dao

import (
	"context"
	"go-mongox-test/model"
	"log"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/chenmingyong0423/go-mongox/v2/builder/update"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (d *UserDao) UserUpdateOne() {
	updateResult, err := d.coll.Updater().Filter(query.In("_id", "680b03b88c4fc25daa7f312d")).Updates(update.Set("name", "chenmingyong0423")).UpdateOne(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("更新单个文档成功", updateResult)
}

func (d *UserDao) UserUpdateOneWithStruct() {
	bid, _ := bson.ObjectIDFromHex("680b03b88c4fc25daa7f312d")
	updateResult, err := d.coll.Updater().Filter(query.In("_id", bid)).Updates(update.SetFields(&model.User{
		Name: "chengmingyong666",
	})).UpdateOne(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("更新单个文档成功", updateResult)
}

func (d *UserDao) UserUpdateMany() {
	updateResult, err := d.coll.Updater().Filter(query.In("name", "Mingyong Chen")).Updates(update.Set("name", "codepzj")).UpdateMany(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("更新多个文档成功", updateResult)
}

func (d *UserDao) UserUpdateSet() {
	updateResult, err := d.coll.Updater().Filter(query.In("name", "codepzj")).Updates(update.Set("name", "chengmingyong666")).Upsert(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("更新或插入一个文档成功", updateResult)
}