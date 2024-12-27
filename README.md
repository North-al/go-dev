# gin-admin

## 项目介绍
> 采用go + gin框架，实现后台管理功能

## 项目文档
1. 生成swagger.json文件
```
swag init -g ./cmd/main.go --v3.1
```
2. 启动项目
```
go run ./cmd/main.go
```
3. 访问swagger文档
```
http://localhost:3333/docs
```
4. 访问swagger.json文件
```
http://localhost:3333/docs.json
```