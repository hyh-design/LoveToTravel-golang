package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"ltt-gc/config"
	"ltt-gc/model"
	"ltt-gc/model/vo"
	"ltt-gc/serializer"
)

type NoteService struct {
	Id         primitive.ObjectID
	UserId     string
	UserName   string
	Title      string
	PlanId     string
	Url        string
	Content    string
	Comment    interface{}
	View       interface{}
	Star       interface{}
	Trip       interface{}
	Deleted    string
	CreateTime string
	UpdateTime string
}

func (service *NoteService) GetNoteById(id string) serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	// 新版ObjectId转换方法
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}
	var result map[string]interface{}
	err := noteCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return serializer.Error(serializer.NoteNotExist)
	}
	return serializer.Success(result)
}

func (service *NoteService) GetNoteList() serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	filter := bson.D{{Key: "deleted", Value: "0"}}
	var result []*model.Note
	cur, err := noteCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &result)
	_ = cur.Close(context.Background())
	return serializer.Success(result)
}

func (service *NoteService) GetNotePage(p vo.Page) serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	filter := bson.D{{Key: "deleted", Value: "0"}}
	var findOptions *options.FindOptions = &options.FindOptions{}

	limit := int64(p.PageSize)
	skip := int64((p.PageSize * p.PageNum) - p.PageSize)

	if p.PageSize > 0 {
		findOptions.SetLimit(limit)
		findOptions.SetSkip(skip)
	}
	var result []map[string]interface{}
	cur, err := noteCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return serializer.Success(serializer.ServerError)
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &result)
	_ = cur.Close(context.Background())
	return serializer.Success(result)
}

func (service *NoteService) GetNotePageFuzzy(p vo.Page) serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	filter := bson.M{
		"content": primitive.Regex{
			Pattern: p.QueryStr,
			Options: "i",
		},
		"deleted": primitive.Regex{
			Pattern: "0",
			Options: "i",
		},
	}
	var findOptions *options.FindOptions = &options.FindOptions{}

	limit := int64(p.PageSize)
	skip := int64((p.PageSize * p.PageNum) - p.PageSize)

	if p.PageSize > 0 {
		findOptions.SetLimit(limit)
		findOptions.SetSkip(skip)
	}
	var result []map[string]interface{}
	cur, err := noteCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return serializer.Success(serializer.ServerError)
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &result)
	_ = cur.Close(context.Background())
	return serializer.Success(result)
}
