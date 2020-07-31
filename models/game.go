package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Game struct {
	Id      int       `orm:"column(id);auto"`
	User1   *Users    `orm:"column(user1);rel(fk)"`
	User2   *Users    `orm:"column(user2);rel(fk)"`
	Status  int       `orm:"column(status);null"`
	Timer   string    `orm:"column(timer);size(45);null"`
	Result  int       `orm:"column(result);null"`
	Started time.Time `orm:"column(started);type(datetime);null;auto_now_add"`
	Ended   time.Time `orm:"column(ended);type(datetime);null"`
}

func (t *Game) TableName() string {
	return "game"
}

func init() {
	orm.RegisterModel(new(Game))
}

// AddGame insert a new Game into database and returns
// last inserted Id on success.
func AddGame(m *Game) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGameById retrieves Game by Id. Returns error if
// Id doesn't exist
func GetGameById(id int) (v *Game, err error) {
	o := orm.NewOrm()
	v = &Game{Id: id}
	if err = o.QueryTable(new(Game)).Filter("id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func CheckUsersAlreadyInGame(user1, user2 int) (bool, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT id from game where status = 1 and (user1 = ? or user1 = ? or user2 = ? or user2 = ?)", user1, user2, user1, user2).Values(&maps)
	if err == nil && num > 0 {
		return true, nil
	}
	return false, err
}

// UpdateGame updates Game by Id and returns error if
// the record to be updated doesn't exist
func UpdateGameById(m *Game) (err error) {
	o := orm.NewOrm()
	v := Game{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
