package dao

import (
	"context"
	"log"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"
	"github.com/chenmingyong0423/go-mongox/v2/builder/aggregation"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (d *UserDao) UserIgnoreAge() {
	user, err := d.coll.Aggregator().Pipeline(aggregation.NewStageBuilder().Project(bson.M{"age": 0}).Build()).Aggregate(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("忽略age字段")
	for _, u := range user {
		log.Println(*u)
	}
}

// 分页查询
func (d *UserDao) UserPagination() {
	user, err := d.coll.Aggregator().Pipeline(aggregation.NewStageBuilder().Sort(bson.M{"age": -1}).Limit(10).Skip(0).Build()).Aggregate(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("分页查询, 查询10条数据, 跳过0条数据, 排序为age降序")
	for _, u := range user {
		log.Println(*u)
	}
}

// 分组求值
func (d *UserDao) UserGroup() {
	user, err := d.coll.Aggregator().Pipeline(aggregation.NewStageBuilder().Group("$name", aggregation.Sum("age", "$age")...).Project(bsonx.NewD().Add("age", 1).Add("name", 1).Add("_id", 0).Build()).Build()).Aggregate(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("分组求值, 分组字段为name, 求值字段为age, 求值方式为sum")
	for _, u := range user {
		log.Println(*u)
	}
}
