package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"google.golang.org/grpc/examples/data"
	"log"
	"week1/server/api"
	"week1/server/config"
	"week1/server/middleware"
	"week1/server/models/mysql"
	"week1/server/models/redis"
)

var port = flag.Int("port", 8081, "rpc server")

func main() {
	conn, err := config.NewConfigConn(&config.NacosCnfig{
		DataId: "users",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		log.Fatal("nacos配置读取失败", err)
	}

	mysql.XDB, err = mysql.XormConn(&conn.Mysql)
	if err != nil {
		log.Fatal("mysql连接失败", err)
	}
	log.Println("mysql连接成功")

	err = mysql.Migrator()
	if err != nil {
		log.Fatal("数据表迁移失败", err)
	}
	_, err = redis.RedisConn(&conn.Redis)
	if err != nil {
		log.Fatal("redis连接失败", err)
	}
	go func() {
		for i := 0; i < 10000; i++ {
			user := mysql.NewUser()
			user.Username = "测试" + uuid.New().String()
			user.CreateUser()
		}
	}()
	err = api.OpenGrpcServer(*port, api.RpcServerRegister, middleware.NewTLS(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem")))
	if err != nil {
		log.Fatal("服务开启失败", err)
	}
}
