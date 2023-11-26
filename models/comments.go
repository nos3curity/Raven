package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Comment struct {
	Id        int       `orm:"auto"`
	System    *System   `orm:"rel(fk);on_delete(cascade)"`
	Username  string    `orm:"type(varchar)"`
	Text      string    `orm:"type(text)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

func FormatTime(t time.Time) string {

	localTime := t.Local()
	return localTime.Format("2006-01-02 15:04:05")
}

func AddComment(systemIp string, username string, commentText string) (err error) {

	o := orm.NewOrm()

	// Find the system by IP
	system := System{Ip: systemIp}
	if err := o.Read(&system, "Ip"); err != nil {
		return err // System not found
	}

	comment := &Comment{
		System:   &system,
		Username: username, // Save the username with the comment
		Text:     commentText,
	}

	// Insert the comment into the database
	_, err = o.Insert(comment)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

// DeleteComment deletes a comment from the database
func DeleteComment(commentId int) (err error) {

	o := orm.NewOrm()

	_, err = o.Delete(&Comment{Id: commentId})
	if err != nil {
		return err
	}

	return nil
}

func GetSystemComments(systemIp string) ([]Comment, error) {

	o := orm.NewOrm()

	var comments []Comment

	_, err := o.QueryTable("comment").Filter("System__Ip", systemIp).RelatedSel().All(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
