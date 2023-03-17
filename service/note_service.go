package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ltt-gc/config"
	"ltt-gc/model"
	"ltt-gc/model/vo"
	"ltt-gc/serializer"
	"time"
)

type NoteService struct {
	ID         string
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
	result["id"] = result["_id"]
	delete(result, "_id")
	if err != nil {
		return serializer.Error(serializer.NoteNotExist)
	}
	return serializer.Success(result)
}

// 数据量多时不建议使用
func (service *NoteService) GetNoteList() serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	filter := bson.D{{Key: "deleted", Value: "0"}}
	var result []*model.Note
	cur, err := noteCollection.Find(context.TODO(), filter)
	if err != nil {
		return serializer.Success(serializer.ServerError)
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

	for i := 0; i < len(result); i++ {
		result[i]["id"] = result[i]["_id"]
		delete(result[i], "_id")
	}

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

	for i := 0; i < len(result); i++ {
		result[i]["id"] = result[i]["_id"]
		delete(result[i], "_id")
	}

	return serializer.Success(result)
}

func (service *NoteService) CreateNote() serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	note := model.Note{
		UserId:     service.UserId,
		UserName:   service.UserName,
		Title:      service.Title,
		PlanId:     service.PlanId,
		Url:        service.Url,
		Content:    service.Content,
		Comment:    0,
		View:       0,
		Star:       0,
		Trip:       service.Trip,
		Deleted:    "0",
		CreateTime: time.Now().String(),
		UpdateTime: time.Now().String(),
	}
	objId, err := noteCollection.InsertOne(context.TODO(), note)
	if err != nil {
		return serializer.Success(serializer.ServerError)
	}
	return serializer.Success(objId.InsertedID)
}

// 需要传入完整数据
func (service *NoteService) UpdateNote() serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	objectId, _ := primitive.ObjectIDFromHex(service.ID)
	fmt.Println(objectId)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "userName", Value: service.UserName},
			{Key: "title", Value: service.Title},
			{Key: "planId", Value: service.PlanId},
			{Key: "url", Value: service.Url},
			{Key: "content", Value: service.Content},
			{Key: "trip", Value: service.Trip},
			{Key: "updateTime", Value: time.Now().String()},
		}},
	}
	result, err := noteCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return serializer.Success(serializer.ServerError)
	}
	return serializer.Success(result)
}

func (service *NoteService) DeleteNoteById(id string) serializer.Response {
	client := config.NewMongoClient()
	noteCollection := client.Database("travelservice").Collection("note")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deleted", Value: "1"},
			{Key: "updateTime", Value: time.Now().String()},
		}},
	}
	result, err := noteCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return serializer.Success(serializer.ServerError)
	}
	return serializer.Success(result)
}
