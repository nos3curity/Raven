package models

import (
	"crypto/rand"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type Config struct {
	Key   string `json:"key" orm:"pk"`
	Value string `json:"value"`
}

func init() {
	orm.RegisterModel(new(Config))
}

func SetConfig(key string, value string) (Config, error) {

	o := orm.NewOrm()

	config := Config{Key: key}

	// Try to read the existing config
	err := o.Read(&config, "Key")

	// Handle different cases based on whether the config exists
	switch {
	case err == orm.ErrNoRows:
		// Config doesn't exist, so create it
		config.Value = value
		_, err = o.Insert(&config)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return Config{}, err
		}
	case err == nil:
		// Config exists, so update it
		config.Value = value
		_, err = o.Update(&config)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return Config{}, err
		}
	default:
		// Some other error occurred
		return Config{}, err
	}

	return config, nil
}

func GetConfig(key string) (config Config, err error) {

	o := orm.NewOrm()

	config = Config{
		Key: key,
	}

	if err := o.Read(&config, "Key"); err != nil {
		return Config{}, err
	}

	return config, nil
}

func AddConfig(key string, value string) (config Config, err error) {

	o := orm.NewOrm()

	config = Config{
		Key:   key,
		Value: value,
	}

	_, err = o.Insert(&config)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return Config{}, err
	}

	return config, nil
}

func DeleteConfig(key string) (err error) {

	o := orm.NewOrm()

	config := Config{Key: key}

	_, err = o.Delete(&config)
	if err != nil {
		return err
	}

	return nil
}

func UpdateConfig(key string, value string) (err error) {

	o := orm.NewOrm()

	config := Config{
		Key:   key,
		Value: value,
	}

	_, err = o.Update(&config)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func RandomString(length int) string {

	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
