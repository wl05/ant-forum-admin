## start
启动
```bash
go run main.go
```
生成api文档
```
swag init
```
访问 http://localhost:8080/swagger/index.html



## 使用air启动
1. 安装air
```bash
go get -v github.com/cosmtrek/air/cmd/...
```

2. 创建.air.conf文件

可以复制 https://github.com/cosmtrek/air 里面的air.conf.example内容

3. 启动

根目录下

```bash
air
```

## 遇到的一些问题

1. 如何方式gorm save的时候覆盖create_at 字段

db.Save(&user)会一同更新created_at。为了使created_at字段的值不被重置，可以使用Omit忽略字段。即：db.Omit("created_at").Save(&user)