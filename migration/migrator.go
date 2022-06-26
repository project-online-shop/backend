package migration

import (
	_mProduct "project/e-comerce/features/products/data"
	_mUser "project/e-comerce/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB){
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mProduct.Product{})
}
