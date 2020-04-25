package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://47.244.184.242:27017"))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("my_db").Collection("col2")

	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	insertResult, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID.(primitive.ObjectID).Hex())
}
