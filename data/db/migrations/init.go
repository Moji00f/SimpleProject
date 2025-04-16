package migrations

import (
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/constant"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/data/models"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up() {
	database := db.GetDb()

	createTables(database)
	createDefaultUserInformation(database)

}

func Down() {
	database := db.GetDb()
	database.Migrator().DropTable(&models.City{})   // حذف جدول قبلی
	database.Migrator().AutoMigrate(&models.City{}) // ایجاد مجدد با تغییرات جدید
}

func AlterTableColumn() {

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	//Basic
	tables = addNewTable(database, &models.Country{}, tables)
	tables = addNewTable(database, &models.City{}, tables)

	//user
	tables = addNewTable(database, &models.User{}, tables)
	tables = addNewTable(database, &models.Role{}, tables)
	tables = addNewTable(database, &models.UserRole{}, tables)

	if len(tables) > 0 {
		err := database.Migrator().AutoMigrate(tables...)
		if err != nil {
			logger.Fatal(logging.Postgres, logging.Migration, "migration failed", nil)
			return
		}
		logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
	}

}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}

	return tables
}

func createDefaultUserInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constant.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constant.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	user := models.User{Username: constant.DefaultUserName,
		FirstName:    "Test",
		LastName:     "Test",
		MobileNumber: "09125888188", Email: "admin@gmail.com"}

	pass := "123"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	createAdminUserIfNotExists(database, &user, adminRole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.Model(&models.Role{}).Select("1").Where("name=?", r.Name).First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.Model(&models.User{}).Select("1").Where("username=?", u.Username).First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}

}
