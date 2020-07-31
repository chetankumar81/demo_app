package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
)

type Cards struct {
	Id         int       `orm:"column(id);auto"`
	GameId     *Game     `orm:"column(gameId);rel(fk)"`
	UserId     *Users    `orm:"column(userId);rel(fk)"`
	Card       *CardMap  `orm:"column(card);rel(fk)"`
	PickedTime time.Time `orm:"column(pickedTime);type(datetime);null"`
}

type PickedCards struct {
	Id         int
	UserId     int
	Card       int
	PickedTime time.Time
}

func (t *Cards) TableName() string {
	return "cards"
}

func init() {
	orm.RegisterModel(new(Cards))
}

// AddCards insert a new Cards into database and returns
// last inserted Id on success.
func AddCards(m *Cards) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCardsById retrieves Cards by Id. Returns error if
// Id doesn't exist
func GetCardsById(id int) (v *Cards, err error) {
	o := orm.NewOrm()
	v = &Cards{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetLast3CardValue
func GetLast3CardValue(gameId, userId int) (v []int, err error) {
	o := orm.NewOrm()
	var a []orm.Params
	_, err = o.Raw("SELECT cardVal from cards c inner join cardMap cm on cm.id = c.card where gameId = ? and userId = ? order by c.id desc limit 3", gameId, userId).Values(&a)

	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	for index := range a {
		v = append(v, cast.ToInt(a[index]["cardVal"]))
	}

	return v, nil
}

//GetCardPicksByGameId ...
func GetCardPicksByGameId(gameId int) (pickedCards []PickedCards, err error) {
	o := orm.NewOrm()
	var v []*Cards
	_, err = o.QueryTable(new(Cards)).Filter("gameId", gameId).OrderBy("-id").All(&v)
	if err == nil {
		for _, card := range v {
			picked := PickedCards{}
			picked.Id = card.Id
			picked.UserId = card.UserId.Id
			picked.Card = card.Card.Id
			picked.PickedTime = card.PickedTime
			pickedCards = append(pickedCards, picked)

		}
		return pickedCards, nil
	}
	return nil, err
}
