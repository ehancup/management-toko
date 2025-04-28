package main

import b "gin-boilerplate/src"

//	@title			Toko manajer
//	@version		1.0
//	@description	This is a Gin Boilerplate for Rest API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ehancup
//	@contact.url	http://www.swagger.io/support
//	@contact.email	rhanysuf24@gmail.com

//	@host	localhost:3000

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

func main() {
	b.BoostrapApp()
}