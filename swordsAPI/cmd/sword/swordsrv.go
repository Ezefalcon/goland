package main

import (
	"flag"
	"fmt"
	"github.com/Ezefalcon/golang/swordsAPI/internal/config"
	"github.com/Ezefalcon/golang/swordsAPI/internal/database"
	"github.com/Ezefalcon/golang/swordsAPI/internal/service"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	cfg := initConfig()

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Println(cfg.DB.Driver)
	//fmt.Println(cfg.Version)
	database.PopulateDatabase(db,"./test/mock.sql")

	swordService, _ := service.NewSwordService(cfg, db)

	makeEndpoints(swordService)

}

func makeEndpoints(swordService service.SwordService) {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	sword := v1.Group("/sword")

	sword.GET("/", func(c *gin.Context) {
		c.JSON(200, swordService.FindAll())
	})

	sword.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(200, swordService.FindById(int(parseInt)))
	})

	sword.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		sword := new(service.Sword)
		err := c.Bind(sword)
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		updateSword, err := swordService.UpdateSword(*sword, int(parseInt))
		c.JSON(200, updateSword)
	})

	sword.POST("/", func(c *gin.Context) {
		sword := new(service.Sword)
		err := c.Bind(sword)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
		}
		updateSword, err := swordService.AddSword(*sword)
		c.JSON(200, updateSword)
	})


	router.Run(":8080")
}

func initConfig() *config.Config {
	configFile := flag.String("config", "./configs/config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}