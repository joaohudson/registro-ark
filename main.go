package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	controller "github.com/joaohudson/registro-ark/controlllers"
	"github.com/joaohudson/registro-ark/db"
	service "github.com/joaohudson/registro-ark/services"
)

func main() {

	database := db.CreatePostgresDatabase()
	defer database.Close()

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)

	//repositórios
	dinoRepo := db.NewDinoRepository(database)
	locomotionRepo := db.NewLocomotionRepository(database)
	regionRepo := db.NewRegionRepository(database)
	foodRepo := db.NewFoodRepository(database)
	admRepo := db.NewAdmRepository(database)

	//services
	dinoService := service.NewDinoService(dinoRepo, locomotionRepo, regionRepo, foodRepo)
	locomotionService := service.NewLocomotionService(locomotionRepo, dinoRepo)
	regionService := service.NewRegionService(regionRepo, dinoRepo)
	foodService := service.NewFoodService(foodRepo, dinoRepo)
	admService := service.NewAdmService(admRepo)

	//controllers
	dinoController := controller.NewDinoController(dinoService, locomotionService, regionService, foodService)
	admController := controller.NewAdmController(admService)

	//rotas públicas
	router.Post("/api/dino", dinoController.CreateDino)
	router.Delete("/api/dino", dinoController.DeleteDino)
	router.Get("/api/dino", dinoController.FindDinoById)
	router.Get("/api/dinos", dinoController.FindDinoByFilter)
	router.Get("/api/dino/categories", dinoController.DinoCategories)
	router.Post("/api/dino/category/food", dinoController.CreateFood)
	router.Delete("/api/dino/category/food", dinoController.DeleteFood)
	router.Post("/api/dino/category/locomotion", dinoController.CreateLocomotion)
	router.Delete("/api/dino/category/locomotion", dinoController.DeleteLocomotion)
	router.Post("/api/dino/category/region", dinoController.CreateRegion)
	router.Delete("/api/dino/category/region", dinoController.DeleteRegion)

	//rotas privadas
	router.Post("/api/adm", admController.CreateAdm)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}
