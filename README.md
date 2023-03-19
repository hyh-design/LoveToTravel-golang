# 爱旅游项目 - Go语言重构版本
这是一个使用Go语言实现的Web后端项目，用于展示旅游相关信息。该项目使用MySQL和MongoDB数据库来存储数据。

## 安装和启动
1. 克隆或下载该项目的代码：
```
git clone https://github.com/hyh315/ltt-gc.git
```
2. 安装依赖：
```
go mod tidy
```
3. 创建并配置数据库：
该项目需要使用MySQL和MongoDB数据库，需要先安装和启动这两个数据库，并在config/config.yaml文件中配置数据库连接信息。
4. 启动项目：
```
go run main.go
```
或者使用编译后的可执行文件启动项目：
```
go build main.go
./main
```

## 依赖版本
+ Go version: 1.16 或更高版本
+ MySQL version: 5.7 或更高版本
+ MongoDB version: 4.2 或更高版本

## 技术栈
+ Go
+ Gin
+ Gorm
+ MySQL
+ MongoDB

## 作者
+ hyh315
