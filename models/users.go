package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int    `orm:"column(id);auto"`
	UserName string `orm:"column(userName);size(45)"`
	EmailId  string `orm:"column(emailId);size(45);null"`
	State    string `orm:"column(state);null"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetUserByuserName ...
func GetUserByuserName(userName string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{}
	if err = o.QueryTable(new(Users)).Filter("userName", userName).One(v); err == nil {
		return v, nil
	}
	return nil, err
}
