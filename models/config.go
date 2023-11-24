package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Config struct {
	Key   string `json:"key" orm:"pk"`
	Value string `json:"value"`
}

func init() {
	orm.RegisterModel(new(Config))
}

func SetConfig(key string, value string) (config Config, err error) {

	o := orm.NewOrm()

	// Get the config by its key
	config, err = GetConfig(key)
	if err != nil {
		return Config{}, err
	}

	if config.Key == "" {

		// If the config does not exist, create it
		_, err = o.Insert(&config)
		if err != nil {
			return Config{}, err
		}
	} else {

		// If the config exists, update it
		_, err = o.Update(&config)
		if err != nil {
			return Config{}, err
		}
	}

	return config, err
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
	if err != nil {
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
	if err != nil {
		return err
	}

	return nil
}
