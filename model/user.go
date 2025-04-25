package model

import "github.com/chenmingyong0423/go-mongox/v2"

type User struct {
	mongox.Model `bson:"inline"`
	Name         string `bson:"name,omitempty"`
	Age          int    `bson:"age,omitempty"`
}
