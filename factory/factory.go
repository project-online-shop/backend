package factory

import (
	_userBusiness "project/e-comerce/features/users/business"
	_userData "project/e-comerce/features/users/data"
	_userPresentation "project/e-comerce/features/users/presentation"

	_authBusiness "project/e-comerce/features/auth/business"
	_authData "project/e-comerce/features/auth/data"
	_authPresentation "project/e-comerce/features/auth/presentation"

	_productBusiness "project/e-comerce/features/products/bussiness"
	_productData "project/e-comerce/features/products/data"
	_productPresentation "project/e-comerce/features/products/presentation"

	"gorm.io/gorm"
)

type Presenter struct{
	UserPresenter *_userPresentation.UserHandler
	AuthPresenter *_authPresentation.UserHandler
	ProductPresenter *_productPresentation.ProductHandler
}

func InitFactory(dbConn *gorm.DB) Presenter{

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewUserRepository(dbConn)
	authBusiness := _authBusiness.NewUserBusiness(authData)
	authPresentation := _authPresentation.NewUserHandler(authBusiness)

	productData := _productData.NewProductRepository(dbConn)
	productBusiness := _productBusiness.NewProductBusiness(productData)
	productPresentation := _productPresentation.NewProductHandler(productBusiness)


	return Presenter{
		UserPresenter: userPresentation,
		AuthPresenter: authPresentation,
		ProductPresenter: productPresentation,		
	}
}