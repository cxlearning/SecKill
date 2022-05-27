package setup

import (
	"myProject/SecKill/sk_admin/app/dao"
	"myProject/SecKill/sk_admin/conf"
	"myProject/SecKill/sk_admin/library/db"
)

func InitDB(configPath string) {
	conf.Init(configPath)
	db.Init()

	create := db.GetDBInstance().Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_bin")
	create.AutoMigrate(
		&dao.Activity{},
		&dao.Product{},
	)
}
