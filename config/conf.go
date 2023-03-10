package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func LoadConfig() {
	viper.SetConfigName("application")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("err nil")
	}
	fmt.Println(viper.Get("mysql"))
}

func LoadMysql() {

	//SQL日志自定义
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //开启彩色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println(DB)
	fmt.Println(DB)
	fmt.Println(DB)

}

func LoadMongo() *mongo.Client {
	fmt.Println(viper.GetString("mongodb.dns"))
	fmt.Println(viper.GetString("mongodb.dns"))
	fmt.Println(viper.GetString("mongodb.dns"))
	clientOptions := options.Client().ApplyURI(viper.GetString("mongo.dns"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检测连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTED to MONGODB SUCCESSFULLY!")

	return client
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

func Init() {
	LoadConfig()
	LoadMysql()
}
