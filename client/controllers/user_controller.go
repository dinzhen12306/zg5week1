package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"time"
	"week1/client/svc"
	"week1/client/util"
	user "week1/server/userrpc"
)

type UserController struct {
	svc.UserSvc
}

func (c *UserController) Login(ctx *gin.Context) {

}

func (c *UserController) Register(ctx *gin.Context) {
	data := struct {
		Username   string    `json:"username"`
		Password   string    `json:"password"`
		Sex        int64     `json:"sex"`
		CreateTime time.Time `json:"create_time"`
		Text       string    `json:"text"`
		School     string    `json:"school"`
	}{}
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数获取失败",
		})
	}
	if !util.CheckText(data.Username) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名不合规",
		})
	}
	//校验出生日期
	var a int64 = 150 * 365 * 24 * 24 * 60 //150年的时间戳
	if data.CreateTime.Unix()-time.Now().Unix() > a {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "出生日期错误",
		})
	}
	client, err := c.UserRpcConn(":8081")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "rpc服务连接失败",
		})
	}
	createUser, err := client.CreateUser(ctx, &user.CreateUserReq{
		User: &user.UserInfo{
			Username:   data.Username,
			Password:   data.Password,
			Sex:        user.Sex(data.Sex),
			CreateTime: data.CreateTime.Unix(),
			Text:       data.Text,
			School:     data.School,
			UID:        int64(uuid.New().ID()),
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "rpc服务调用失败",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": createUser,
	})
}

func (c *UserController) UpLoad(ctx *gin.Context) {
	id := ctx.GetInt64("id")
	file, src, err := ctx.Request.FormFile("img")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件上传失败",
		})
	}
	if src.Size > 5*1024*1024 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件过大",
		})
	}
	ok := false
	imgType := []string{"image/jpg", "image/png"}
	for _, v := range imgType {
		if v == src.Header.Get("content-Type") {
			ok = true
		}
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件类型错误",
		})
	}
	create, err := os.Create("./static/" + src.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件创建失败",
		})
	}
	_, err = io.Copy(create, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件拷贝失败",
		})
	}
	client, err := c.UserRpcConn(":8081")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "rpc服务连接失败",
		})
	}
	updateUser, err := client.UpdateUser(ctx, &user.UpdateUserReq{
		User: &user.UserInfo{
			Id:    id,
			Title: "./static/" + src.Filename,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "rpc服务调用失败",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "文件上传成功",
		"data": updateUser,
	})
}
