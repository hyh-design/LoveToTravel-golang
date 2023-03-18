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

type Plan struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserId     string             `json:"userId,omitempty" bson:"userId,omitempty"`
	Budget     string             `json:"budget,omitempty" bson:"budget,omitempty"`
	Depart     string             `json:"depart,omitempty" bson:"depart,omitempty"`
	Start      string             `json:"start,omitempty" bson:"start,omitempty"`
	End        string             `json:"end,omitempty" bson:"end,omitempty"`
	SubPlans   []SubPlan          `json:"subPlans,omitempty" bson:"subPlans,omitempty"`
	Deleted    string             `json:"deleted,omitempty" bson:"deleted,omitempty"`
	CreateTime string             `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime string             `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

type SubPlan struct {
	CityId    string `json:"cityId,omitempty" bson:"cityId,omitempty"`
	City      string `json:"city,omitempty" bson:"city,omitempty"`
	Budget    string `json:"budget,omitempty" bson:"budget,omitempty"`
	Days      []Days `json:"days,omitempty" bson:"days,omitempty"`
	DayLength int    `json:"dayLength,omitempty" bson:"dayLength,omitempty"`
}

type Days struct {
	Route []Route `json:"route,omitempty" bson:"route,omitempty"`
}

type Route struct {
	Origin     []float32 `json:"origin,omitempty" bson:"origin,omitempty"`
	OriginName string    `json:"originName,omitempty" bson:"originName,omitempty"`
	DepartTime int       `json:"departTime,omitempty" bson:"departTime,omitempty"`
	Vehicle    string    `json:"vehicle,omitempty" bson:"vehicle,omitempty"`
}

var (
	planCollection *mongo.Collection
)

func GetPlanCollection() (*mongo.Collection, error) {
	client := config.NewMongoClient()
	planCollection = client.Database("travelservice").Collection("chat")
	return planCollection, nil
}

func (service *Plan) GetPlanById(id string) serializer.Response {
	planCollection, _ = GetPlanCollection()

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}
	var result map[string]interface{}
	err := planCollection.FindOne(context.TODO(), filter).Decode(&result)
	result["id"] = result["_id"]
	delete(result, "_id")
	if err != nil {
		return serializer.Error(err.Error())
	}
	return serializer.Success(result)
}

func (service *Plan) GetPlanList() serializer.Response {
	planCollection, _ = GetPlanCollection()

	filter := bson.D{{Key: "deleted", Value: "0"}}
	var result []map[string]interface{}
	cur, err := planCollection.Find(context.Background(), filter)
	if err != nil {
		return serializer.Success(err.Error())
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &result)
	_ = cur.Close(context.Background())
	return serializer.Success(result)
}

func (service *Plan) GetPlanPage(p vo.Page) serializer.Response {
	planCollection, _ = GetPlanCollection()

	filter := bson.D{{Key: "deleted", Value: "0"}}
	var findOptions *options.FindOptions = &options.FindOptions{}

	limit := int64(p.PageSize)
	skip := int64((p.PageSize * p.PageNum) - p.PageSize)

	if p.PageSize > 0 {
		findOptions.SetLimit(limit)
		findOptions.SetSkip(skip)
	}
	var result []map[string]interface{}
	cur, err := planCollection.Find(context.TODO(), filter, findOptions)
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

func (service *Plan) GetPlanPageFuzzy(p vo.Page) serializer.Response {
	planCollection, _ = GetPlanCollection()

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
	cur, err := planCollection.Find(context.TODO(), filter, findOptions)
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

func (service *Plan) CreatePlan() serializer.Response {
	planCollection, _ = GetPlanCollection()

	plan := Plan{
		ID:         primitive.NewObjectID(),
		UserId:     service.UserId,
		Budget:     service.Budget,
		Depart:     service.Depart,
		Start:      service.Start,
		End:        service.End,
		SubPlans:   service.SubPlans,
		Deleted:    "0",
		CreateTime: time.Now().String(),
		UpdateTime: time.Now().String(),
	}
	objId, err := planCollection.InsertOne(context.TODO(), plan)
	if err != nil {
		return serializer.Success(err.Error())
	}
	return serializer.Success(objId.InsertedID)
}

func (service *Plan) UpdatePlan() serializer.Response {
	planCollection, _ = GetPlanCollection()

	objectId := service.ID
	fmt.Println(objectId)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "userId", Value: service.UserId},
			{Key: "budget", Value: service.Budget},
			{Key: "depart", Value: service.Depart},
			{Key: "start", Value: service.Start},
			{Key: "end", Value: service.End},
			{Key: "subPlans", Value: service.SubPlans},
			{Key: "updateTime", Value: time.Now().String()},
		}},
	}
	result, err := planCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return serializer.Success(err.Error())
	}
	return serializer.Success(result)
}

func (service *Plan) DeletePlanById(id string) serializer.Response {
	planCollection, _ = GetPlanCollection()

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}, {Key: "deleted", Value: "0"}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deleted", Value: "1"},
			{Key: "updateTime", Value: time.Now().String()},
		}},
	}
	result, err := planCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return serializer.Success(err.Error())
	}
	return serializer.Success(result)
}
