package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Loot struct {
	Id         int       `orm:"auto"`
	Tag        string    `orm:"type(text)"`
	Uploader   string    `orm:"type(varchar)"`
	Name       string    `orm:"type(varchar)"`
	Filename   string    `orm:"type(varchar)"`
	System     *System   `orm:"rel(fk);on_delete(cascade)"`
	UploadedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// Allowed loot tags
var LootTags = []string{
	"pii",
	"config",
	"script",
	"database",
	"key",
	"other",
}

func init() {
	orm.RegisterModel(new(Loot))
}

func LootTagValid(tag string) bool {

	for _, t := range LootTags {
		if tag == t {
			return true
		}
	}

	return false
}

func AddLoot(loot *Loot) (err error) {

	o := orm.NewOrm()

	_, err = o.Insert(loot)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func DeleteLoot(lootId int) (err error) {

	o := orm.NewOrm()

	loot := Loot{Id: lootId}

	_, err = o.Delete(&loot)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func GetLoot(lootId int) (loot Loot, err error) {

	o := orm.NewOrm()

	loot = Loot{Id: lootId}
	err = o.Read(&loot, "Id")
	if err != nil {
		return Loot{}, err
	}

	return loot, nil
}

func GetAllLoot() (lootItems []Loot, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(Loot)).All(&lootItems)
	if err != nil {
		return nil, err
	}

	return lootItems, nil
}

func GetLootPath(lootId int) (filePath string, err error) {

	loot, err := GetLoot(lootId)
	if err != nil {
		return "", err
	}

	filePath = "uploads/loot/" + loot.Filename

	return filePath, nil
}

func GetLootName(lootId int) (lootName string, err error) {

	loot, err := GetLoot(lootId)
	if err != nil {
		return "", err
	}

	lootName = loot.Name

	return lootName, nil
}
