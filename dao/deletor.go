package dao

import (
	"context"
	"log"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (d *UserDao) UserDeleteOne() {
	bid, _ := bson.ObjectIDFromHex("680b03b88c4fc25daa7f312d")
	deleteResult, err := d.coll.Deleter().Filter(query.In("_id", bid)).DeleteOne(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("删除单个文档成功", deleteResult)
}

func (d *UserDao) UserDeleteMany() {
	deleteResult, err := d.coll.Deleter().Filter(query.In("name", "codepzj")).DeleteMany(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("删除多个文档成功", deleteResult)
}