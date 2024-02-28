package server

import (
	"bytes"
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"week1/server/models/mysql"
	"week1/server/models/redis"
	user "week1/server/userrpc"
)

func CreateUser(info *user.UserInfo) (*user.UserInfo, error) {
	users := ProtoToMysql(info)
	err := users.CreateUser()
	if err != nil {
		return nil, err
	}
	return MysqlToProto(users), nil
}
func DeleteUser(id int64) error {
	users := mysql.NewUser()
	err := users.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(info *user.UserInfo) (*user.UserInfo, error) {
	users := ProtoToMysql(info)
	err := users.UpdateUser()
	if err != nil {
		return nil, err
	}
	return MysqlToProto(users), nil
}

func GetUser(where map[string]string) (*user.UserInfo, error) {
	users := mysql.NewUser()
	wheres := make(map[string]interface{})
	for k, v := range where {
		wheres[k] = v
	}
	err := users.GetUser(wheres)
	if err != nil {
		return nil, err
	}
	return MysqlToProto(users), nil
}

func GetUsers(limit, offset int64) ([]*user.UserInfo, int64, error) {
	//先从redis获取数据
	redisKey := bytes.Buffer{}
	redisKey.WriteString("userInfo")
	if !redis.RedisKeyExists(redisKey.String()) {
		list, err := redis.GetList(redisKey.String(), offset, -1)
		if err != nil {
			return nil, 0, err
		}
		var data []*user.UserInfo
		for _, v := range list {
			var a *user.UserInfo
			proto.Unmarshal([]byte(v), a)
			data = append(data, a)
		}
		return data, int64(len(list)), err
	} else {
		users := mysql.NewUser()
		getUsers, i, err := users.GetUsers(int(limit), int(offset))
		if err != nil {
			return nil, 0, err
		}
		var data []*user.UserInfo
		for _, v := range getUsers {
			//将数据存入redis
			marshal, _ := json.Marshal(v)
			redis.List(redisKey.String(), float64(v.Id), marshal)
			data = append(data, MysqlToProto(&v))
		}

		return data, i, nil
	}
}

func MysqlToProto(u *mysql.User) *user.UserInfo {
	return &user.UserInfo{
		Id:         u.Id,
		Username:   u.Username,
		Password:   u.Password,
		Sex:        user.Sex(u.Sex),
		CreateTime: u.CreateTime,
		Text:       u.Text,
		School:     u.School,
		UID:        u.UID,
	}
}
func ProtoToMysql(u *user.UserInfo) *mysql.User {
	return &mysql.User{
		Id:         u.Id,
		Username:   u.Username,
		Password:   u.Password,
		Sex:        int64(u.Sex),
		CreateTime: u.CreateTime,
		Text:       u.Text,
		School:     u.School,
		UID:        u.UID,
	}
}
