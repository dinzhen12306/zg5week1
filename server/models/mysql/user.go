package mysql

import "time"

// 用户表
type User struct {
	Id         int64     `xorm:"pk autoincr"`
	Username   string    `xorm:"unique"`
	Password   string    `xorm:"varchar(32)"`
	Sex        int64     `xorm:"tinyint(1)"`
	CreateTime int64     `xorm:"comment('出生时间戳')"`
	Text       string    `xorm:"varchar(255) comment('描述')"`
	School     string    `xorm:"varchar(50) comment('学校')"`
	UID        int64     `xorm:"unique"`
	Title      string    `xorm:"varchar(255)"`
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
	Deleted    time.Time `xorm:"deleted"`
}

func NewUser() *User {
	return new(User)
}

func (u *User) GetUser(where map[string]interface{}) (err error) {
	_, err = XDB.Where(where).Get(u)
	return
}
func (u *User) GetUsers(limit, offset int) (users []User, content int64, err error) {
	err = XDB.Limit(limit, offset).Find(&users)
	if err != nil {
		return nil, 0, err
	}
	total, err := XDB.Count(u)
	return users, total, err
}

func (u *User) CreateUser() (err error) {
	_, err = XDB.Insert(u)
	return
}

func (u *User) UpdateUser() (err error) {
	_, err = XDB.Where("id = ?", u.Id).Update(u)
	return
}

func (u *User) DeleteUser(id int64) (err error) {
	_, err = XDB.Table(u).Where("id = ?", id).Delete()
	return
}
