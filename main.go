package main

import (
	"backend/config"
	"backend/db/seeders"
	"backend/router"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	seed := flag.Bool("seed", false, "seed")
	flag.Parse()

	config.LoadConfig()
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	if *seed {
		seeders.UserSeeder(db)
		fmt.Println("Seeder User berhasil dijalankan!")
		os.Exit(0)
	}

	userController, friendController, chatController := config.DependencyInjection(config.DB)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept,Authorization",
		AllowCredentials: true,
	}))

	router.SetUpRoutes(app, userController, friendController, chatController)

	fmt.Println("success")

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
