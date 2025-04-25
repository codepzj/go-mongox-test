package dao

import (
	"context"
	"log"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (d *UserDao) UserFindOne() {
	findResult, err := d.coll.Finder().Filter(bson.M{"name": "Mingyong Chen"}).FindOne(context.Background(), options.FindOne().SetSkip(1))
	if err != nil {
		panic(err)
	}
	log.Println("查询单条数据成功", findResult)
}

func (d *UserDao) UserFindMany() {
	findResult, err := d.coll.Finder().Filter(query.NewBuilder().Eq("name", "Mingyong Chen").Build()).Find(context.Background(), options.Find().SetLimit(3).SetSkip(1).SetSort(bson.M{"age": 1}))
	if err != nil {
		panic(err)
	}
	log.Println("查询多条数据成功:3条")
	for _, v := range findResult {
		log.Println(v)
	}
}

func (d *UserDao) UserCount() {
	count, err := d.coll.Finder().Filter(query.NewBuilder().Eq("name", "Mingyong Chen").Gte("age", 18).Build()).Count(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("查询文档数量成功", count)
}

func (d *UserDao) UserDistinct() {
	var ageSlice []int
	err := d.coll.Finder().DistinctWithParse(context.Background(), "age", &ageSlice)
	if err != nil {
		panic(err)
	}
	log.Println("查询指定字段的唯一值成功", ageSlice)
}

func (d *UserDao) UserFindSort() {
	findResult, err := d.coll.Finder().Filter(query.NewBuilder().Gte("age", 18).Build()).Sort(bson.M{"age": 1}).Skip(1).Limit(3).Find(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("查询排序数据成功", findResult)
}
