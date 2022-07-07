package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	controller "github.com/joaohudson/registro-ark/controlllers"
	"github.com/joaohudson/registro-ark/db"
)

func main() {

	database := db.CreatePostgresDatabase()
	defer database.Close()

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)

	dinoController := controller.NewDinoController(database)

	router.Post("/api/dino", dinoController.CreateDino)
	router.Delete("/api/dino", dinoController.DeleteDino)
	router.Get("/api/dino", dinoController.FindDinoById)
	router.Get("/api/dinos", dinoController.FindDinoByFilter)

	router.Get("/api/dino/categories", dinoController.DinoCategories)
	router.Post("/api/dino/category/food", dinoController.CreateFood)
	router.Post("/api/dino/category/locomotion", dinoController.CreateLocomotion)
	router.Post("/api/dino/category/region", dinoController.CreateRegion)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}
