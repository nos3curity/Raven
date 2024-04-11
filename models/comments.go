package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
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

func normalizeTime(t time.Time) time.Time {
	// Adjust time to defined timezone in conf
	timezone, _ := beego.AppConfig.String("timezone")
	loc, _ := time.LoadLocation(timezone)
	normalizedTime := t.In(loc)

	return normalizedTime
}

func FormatTime(t time.Time) string {
	// Format the time according to the desired format
	normalizedTime := normalizeTime(t)
	formattedTime := normalizedTime.Format("Jan. _2, 2006 03:04PM")

	return formattedTime
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

	// Reverse the order of comments
	for i, j := 0, len(comments)-1; i < j; i, j = i+1, j-1 {
		comments[i], comments[j] = comments[j], comments[i]
	}

	return comments, nil
}
