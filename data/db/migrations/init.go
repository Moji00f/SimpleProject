package migrations

import (
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/data/model"
	"github.com/Moji00f/SimpleProject/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up() {
	database := db.GetDb()

	//if database == nil {
	//	logger.Fatal(logging.Postgres, logging.Migration, "faile to get database", nil)
	//	return
	//}

	//fmt.Println("✅ مقدار دیتابیس:", database)          // بررسی مقدار
	//fmt.Println("✅ وضعیت اتصال:", database.Migrator()) // بررسی امکان اجرای Migrator
	var tables []interface{}

	country := &model.Country{}
	city := &model.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	if len(tables) > 0 {
		err := database.Migrator().AutoMigrate(tables...)
		if err != nil {
			logger.Fatal(logging.Postgres, logging.Migration, "migration failed", nil)
			return
		}
		logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
	}

}

func Down() {
	database := db.GetDb()
	database.Migrator().DropTable(&model.City{})   // حذف جدول قبلی
	database.Migrator().AutoMigrate(&model.City{}) // ایجاد مجدد با تغییرات جدید
}

func AlterTableColumn() {

}
