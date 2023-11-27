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

func GetAllLoot(sortBy string, orderBy string) ([]Loot, error) {
	var lootItems []Loot
	o := orm.NewOrm()

	query := o.QueryTable(new(Loot))

	// Check if sortBy is provided and valid
	if sortBy != "" {
		// Build the order string
		orderString := ""
		if orderBy == "desc" {
			orderString = "-" + sortBy // Prefix with "-" for descending order
		} else {
			orderString = sortBy // Ascending order is default
		}

		// Apply ordering to the query
		query = query.OrderBy(orderString)
	}

	// Execute the query
	_, err := query.All(&lootItems)
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

func GetSystemLoot(systemIp string, sortBy string, orderBy string) (systemLoot []Loot, err error) {

	o := orm.NewOrm()

	query := o.QueryTable("loot").Filter("System__Ip", systemIp).RelatedSel()

	// Check if sortBy is provided and valid
	if sortBy != "" {
		// Build the order string
		orderString := ""
		if orderBy == "desc" {
			orderString = "-" + sortBy // Prefix with "-" for descending order
		} else {
			orderString = sortBy // Ascending order is default
		}

		// Apply ordering to the query
		query = query.OrderBy(orderString)
	}

	// Execute the query
	_, err = query.All(&systemLoot)
	if err != nil {
		return nil, err
	}

	return systemLoot, nil
}

func GetLootedSystems() (systemIps []string, err error) {

	// First get all the loot records
	lootItems, err := GetAllLoot("system_id", "asc")
	if err != nil {
		return nil, err
	}

	// Temporary map to store unique IPs
	uniqueIPs := make(map[string]bool)

	// Then loop over the loot array
	for _, loot := range lootItems {
		ip := loot.System.Ip
		if _, exists := uniqueIPs[ip]; !exists {
			uniqueIPs[ip] = true
			systemIps = append(systemIps, ip)
		}
	}

	return systemIps, nil
}

func GetLootedTeams() (teamIds []int, err error) {

	// First get all looted systems
	allLootedSystems, err := GetLootedSystems()
	if err != nil {
		return nil, err
	}

	// Use a map to track unique team IDs
	uniqueTeamIDs := make(map[int]struct{})

	// Loop over all looted systems
	for _, systemIp := range allLootedSystems {
		// Get the team of the system
		team, err := GetSystemsTeam(systemIp)
		if err != nil {
			return nil, err
		}

		// Check if the team ID is already in the map
		if _, exists := uniqueTeamIDs[team.Id]; !exists {
			// If not, add it to the map and the teamIds slice
			uniqueTeamIDs[team.Id] = struct{}{}
			teamIds = append(teamIds, team.Id)
		}
	}

	return teamIds, nil
}

func GetLootedTeamSystems(teamId int) (systemIps []string, err error) {
	// First get all looted systems
	allLootedSystems, err := GetLootedSystems()
	if err != nil {
		return nil, err
	}

	// Loop over all looted systems
	for _, systemIp := range allLootedSystems {

		// Get the team of the system
		team, err := GetSystemsTeam(systemIp)
		if err != nil {
			return nil, err
		}

		// Check if the system belongs to the specified team
		if team.Id == teamId {
			// Add the IP to the array
			systemIps = append(systemIps, systemIp)
		}
	}

	return systemIps, nil
}

func FilterLootByTag(lootItems []Loot, tag string) (filteredLoot []Loot, err error) {

	// Make sure the loot tag is valid
	if !LootTagValid(tag) {
		return nil, nil
	}

	// Loop over the loop items
	for _, loot := range lootItems {

		// Return only those that have the tag
		if loot.Tag == tag {
			filteredLoot = append(filteredLoot, loot)
		}
	}

	return filteredLoot, nil
}

func SortLoot(lootItems []Loot, sortBy string, order string) (sortedLoot []Loot, err error) {

	//
	return sortedLoot, nil
}
