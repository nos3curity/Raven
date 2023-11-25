package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Comment struct {
	Id        int       `orm:"auto"`
	SystemIp  *System   `orm:"rel(fk);on_delete(cascade)"`
	Username  string    `orm:"type(varchar)"`
	Text      string    `orm:"type(text)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

func AddComment(systemIp string, username string, commentText string) error {
	o := orm.NewOrm()

	// Find the system by IP
	system := System{Ip: systemIp}
	if err := o.Read(&system, "Ip"); err != nil {
		return err // System not found
	}

	comment := &Comment{
		SystemIp: &system,
		Username: username, // Save the username with the comment
		Text:     commentText,
	}

	// Insert the comment into the database
	_, err := o.Insert(comment)
	return err
}

func GetCommentsBySystemIp(systemIp string) ([]Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	_, err := o.QueryTable("comment").Filter("SystemIp__Ip", systemIp).RelatedSel().All(&comments)
	return comments, err
}

// DeleteComment deletes a comment from the database
func DeleteComment(commentId int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Comment{Id: commentId})
	return err
}
