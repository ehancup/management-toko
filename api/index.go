package handler

import (
	// "fmt"
	"flag"
	"gin-boilerplate/config"
	"gin-boilerplate/src"
	"gin-boilerplate/src/database"
	"gin-boilerplate/src/utils/logger"
	"net/http"
	"strings"

	// "time"

	log "github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "gin-boilerplate/docs"

	"github.com/AlecAivazis/survey/v2"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// testing

	err := godotenv.Load()

	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	/*
		user commmented code below to use database
		dont forget to import "gin-boilerplate/src/database" and "flag"
	*/

	database.ConnectDatabase()

	isMigrate := flag.Bool("migrate", false, "set to migrate database.")
	flag.Parse()

	if *isMigrate {
		p := ""
		descriptions := map[string]string{
			"Normal migrate":  "just doing a normal migration (no data will lose)",
			"Reset migrate": "recommend for delete or edit column name (you will lose all data)",
		}
		prompt := &survey.Select{
			Message: "Select migration type",
			Options: []string{"Normal migrate", "Reset migrate"},
			Description: func(value string, index int) string {
				if desc, ok := descriptions[value]; ok {
					return desc
				}
				return ""
			},
		}
		survey.AskOne(prompt, &p)
		if strings.HasPrefix(p, "Reset") {
			sure := false
			pr := &survey.Confirm{
				Message: "Are you sure?",
			}
			survey.AskOne(pr, &sure)
			if sure {
				database.Migrate(true)
				return
			} else {return}
		} else {
			database.Migrate(false)
			return
		}
	} 

	
	appCfg := config.GetConfig().App
	if appCfg.Mode == "debug" {
		
        logger.Log.SetLevel(log.DebugLevel)
		logger.Debug("debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	
	app := gin.New()

	app.Use(gin.Recovery())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	src.InitRoutes(app)
	app.Static("/public", "./public")

	logger.Info("Check documentation ->", "url", appCfg.Url + "/swagger/index.html")

	// if err := app.Run(appCfg.Port); err != nil {
	// 	logger.Fatal("[ERR] fail starting servr :", "err", err)
	// }

	app.ServeHTTP(w, r)
}
