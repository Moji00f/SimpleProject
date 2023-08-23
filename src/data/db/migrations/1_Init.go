package migrations

import (
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/constants"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/data/models"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	createTables(database)
	createDefaultInformation(database)

}
func createTables(database *gorm.DB) {

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Tables Created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}

	return tables
}

func createDefaultInformation(database *gorm.DB) {

	adminrole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminrole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{UserName: constants.DefaultUserName, FirstName: "Test", LastName: "Test", MobileNumber: "09195870197", Email: "admin@admin.com"}
	pass := "12345678"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashPassword)

	createAdminUserIfNotExists(database, &u, adminrole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.Model(&models.Role{}).Select("1").Where("name=?", constants.DefaultRoleName).First(&exists)

	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.Model(&models.User{}).Select("1").Where("user_name=?", constants.DefaultRoleName).First(&exists)

	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down_1() {

}
