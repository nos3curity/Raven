package models

import "github.com/beego/beego/v2/client/orm"

type SystemPort struct {
	Id     int     `json:"id"`
	System *System `orm:"rel(fk);column(system_ip);on_delete(cascade)"`
	Port   *Port   `orm:"rel(fk);column(port_number);on_delete(cascade)"`
	Open   bool    `json:"open"`
}

func init() {
	orm.RegisterModel(new(SystemPort))
}

func AddOpenSystemPort(system *System, port *Port) (err error) {

	o := orm.NewOrm()

	systemPort := SystemPort{
		System: system,
		Port:   port,
		Open:   true,
	}

	// Check if there are any port associations for this system
	systemPorts, err := GetSystemPorts(system.Ip)
	if len(systemPorts) == 0 {

		// If no associations exist, create it
		_, err = o.Insert(&systemPort)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return err
		}
	} else {

		// See if this particular association exists
		systemPorts, err = GetSystemPort(system.Ip, port.PortNumber)
		if err != nil {
			return err
		}

		if len(systemPorts) == 0 {

			// If no associations exist, create it
			_, err = o.Insert(&systemPort)
			if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
				return err
			}
		} else {

			// If it exists, update it
			systemPort.Id = systemPorts[0].Id
			_, err = o.Update(&systemPort)
			if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
				return err
			}
		}
	}

	return nil
}

func GetSystemPort(systemIp string, port int) (systemPorts []SystemPort, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(SystemPort)).RelatedSel().Filter("System__Ip", systemIp).Filter("Port__PortNumber", port).All(&systemPorts)
	if err != nil {
		return nil, err
	}

	return systemPorts, nil
}

func SetAllSystemPortsClosedByIp(systemIp string) (err error) {

	o := orm.NewOrm()

	// Update Open field to false for all records matching the system IP
	_, err = o.QueryTable(new(SystemPort)).Filter("System__Ip", systemIp).Update(orm.Params{
		"Open": false,
	})

	if err != nil {
		return err
	}

	return nil
}
