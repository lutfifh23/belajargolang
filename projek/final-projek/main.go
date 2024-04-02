package main

import (
	"final-projek/config"
	"final-projek/controllers"
	"final-projek/routers"
)

func main() {
	db := config.InitDB()

	userRepo := controllers.UserRepo{DB: db}
	photoRepo := controllers.PhotoRepo{DB: db}
	commentRepo := controllers.CommentRepo{DB: db}
	mediaRepo := controllers.MediaRepo{DB: db}

	r := routers.StartApp(userRepo, photoRepo, commentRepo, mediaRepo)
	r.Run(":8080")
}
