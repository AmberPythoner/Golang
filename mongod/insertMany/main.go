package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

type BeforeTime struct {
	Before int64 `bson:"$lt"`
}

type Dels struct {
	Before BeforeTime `bson:"timePoint.startTime"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://47.244.184.242:27017"))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("my_db").Collection("col2")
	//插入对象
	//s := &Students{
	//	//s := &Students{
	//	//	Name: "xiaoli",
	//	//	Age:  10,
	//	//}
	//	//: "xiaoli",
	//	Age:  10,
	//}
	//
	//students := []interface{}{
	//	s, s,
	//}
	//var insertResult *mongo.InsertManyResult
	//if insertResult, err = collection.InsertMany(context.TODO(), students); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(insertResult)

	//find 对象
	//finds := FindStudents{Name: "xiaoli"}
	//var cursor *mongo.Cursor
	//ops := &options.FindOptions{}
	//ops.SetSkip(2)
	//ops.SetLimit(3)
	//if cursor, err = collection.Find(context.TODO(), finds, ops); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer cursor.Close(context.TODO())
	//for cursor.Next(context.TODO()) {
	//	s := &Students{}
	//	cursor.Decode(s)
	//	fmt.Println(s.Age)
	//}

	//	删除
	d := &Dels{Before: BeforeTime{Before: time.Now().Unix()}}
	var dels *mongo.DeleteResult
	if dels, err = collection.DeleteMany(context.TODO(), d); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dels.DeletedCount)
}
