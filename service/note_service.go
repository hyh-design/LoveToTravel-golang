package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ltt-gc/config"
	"ltt-gc/model/vo"
	"ltt-gc/serializer"
	"time"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserId     string             `json:"user_id,omitempty" bson:"userId,omitempty"`
	UserName   string             `json:"user_name,omitempty" bson:"userName,omitempty"`
	Title      string             `json:"title,omitempty" bson:"title,omitempty"`
	PlanId     string             `json:"plan_id,omitempty" bson:"planId,omitempty"`
	Url        string             `json:"url,omitempty" bson:"url,omitempty"`
	Content    string             `json:"content,omitempty" bson:"content,omitempty"`
	Comment    interface{}        `json:"comment,omitempty" bson:"comment,omitempty"`
	View       interface{}        `json:"view,omitempty" bson:"view,omitempty"`
	Star       interface{}        `json:"star,omitempty" bson:"star,omitempty"`
	Trip       interface{}        `json:"trip,omitempty" bson:"trip,omitempty"`
	Deleted    string             `json:"deleted,omitempty" bson:"deleted,omitempty"`
	CreateTime string             `json:"create_time,omitempty" bson:"createTime,omitempty"`
	UpdateTime string             `json:"update_time,omitempty" bson:"updateTime,omitempty"`
}

var (
	noteCollection *mongo.Collection
)

// GetNoteCollection 获取note操作集合
func GetNoteCollection() (*mongo.Collection, error) {
	client := config.NewMongoClient()
	noteCollection = client.Database("travelservice").Collection("note")
	return noteCollection, nil
}

func (service *Note) GetNoteById(id string) serializer.Response {
	noteCollection, _ = GetNoteCollection()

	// 新版ObjectId转换方法
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}
	var result map[string]interface{}
	err := noteCollection.FindOne(context.TODO(), filter).Decode(&result)
	result["id"] = result["_id"]
	delete(result, "_id")
	if err != nil {
		return serializer.Error(err.Error())
	}
	return serializer.Success(result)
}

// 数据量多时不建议使用
func (service *Note) GetNoteList() serializer.Response {
	noteCollection, _ = GetNoteCollection()

	filter := bson.D{{Key: "deleted", Value: "0"}}
	var result []map[string]interface{}
	cur, err := noteCollection.Find(context.TODO(), filter)
	if err != nil {
		return serializer.Success(err.Error())
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &result)
	_ = cur.Close(context.Background())
	return serializer.Success(result)
}

func (service *Note) GetNotePage(p vo.Page) serializer.Response {
	noteCollection, _ = GetNoteCollection()

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
		return serializer.Success(err.Error())
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

func (service *Note) GetNotePageFuzzy(p vo.Page) serializer.Response {
	noteCollection, _ = GetNoteCollection()

	filter := bson.M{
		"content": primitive.Regex{
			Pattern: p.QueryStr,
			Options: "i",
		},
		"deleted": "0",
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
		return serializer.Success(err.Error())
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

func (service *Note) CreateNote() serializer.Response {
	noteCollection, _ = GetNoteCollection()

	note := Note{
		ID:         primitive.NewObjectID(),
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
		fmt.Println(err)
		return serializer.Success(err.Error())
	}
	return serializer.Success(objId.InsertedID)
}

// 需要传入完整数据
func (service *Note) UpdateNote() serializer.Response {
	noteCollection, _ = GetNoteCollection()

	objectId := service.ID
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
		return serializer.Success(err.Error())
	}
	return serializer.Success(result)
}

func (service *Note) DeleteNoteById(id string) serializer.Response {
	noteCollection, _ = GetNoteCollection()

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
		return serializer.Success(err.Error())
	}
	return serializer.Success(result)
}
