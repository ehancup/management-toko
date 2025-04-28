package src

import (
	"gin-boilerplate/src/app/auth"
	"gin-boilerplate/src/app/product"
	"gin-boilerplate/src/app/store"
)

func GetAuthService() *auth.Service {
	return &auth.Service{}
}

func GetStoreService() *store.Service {
	return &store.Service{}
}

func GetProductService() *product.Service{
	return &product.Service{}
}