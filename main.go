package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	controller "github.com/joaohudson/registro-ark/controlllers"
	"github.com/joaohudson/registro-ark/db"
	service "github.com/joaohudson/registro-ark/services"
	"github.com/joaohudson/registro-ark/util"
)

func main() {

	database := db.CreatePostgresDatabase()
	defer database.Close()

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)

	//repositórios
	loginRepo := db.NewLoginRepository(database)
	dinoRepo := db.NewDinoRepository(database)
	locomotionRepo := db.NewLocomotionRepository(database)
	regionRepo := db.NewRegionRepository(database)
	foodRepo := db.NewFoodRepository(database)
	admRepo := db.NewAdmRepository(database)

	//services
	dinoService := service.NewDinoService(dinoRepo, locomotionRepo, regionRepo, foodRepo, admRepo)
	locomotionService := service.NewLocomotionService(locomotionRepo, dinoRepo)
	regionService := service.NewRegionService(regionRepo, dinoRepo)
	foodService := service.NewFoodService(foodRepo, dinoRepo)
	admService := service.NewAdmService(admRepo)
	loginService := service.NewLoginService(admRepo, loginRepo)

	//controllers
	dinoController := controller.NewDinoController(dinoService, locomotionService, regionService, foodService, loginService)
	admController := controller.NewAdmController(admService, loginService)
	loginController := controller.NewLoginController(loginService)

	//rotas públicas
	router.Post("/api/adm/login", loginController.Login)
	router.Get("/api/dino", dinoController.FindDinoById)
	router.Get("/api/dinos", dinoController.FindDinoByFilter)
	router.Get("/api/dino/categories", dinoController.DinoCategories)

	//rotas privadas
	router.Post("/api/dino", dinoController.CreateDino)
	router.Delete("/api/dino", dinoController.DeleteDino)
	router.Get("/api/adm/dinos", dinoController.FindDinoByFilterForAdm)
	router.Post("/api/adm/logout", loginController.Logout)
	router.Post("/api/adm", admController.CreateAdm)
	router.Put("/api/adm/permissions", admController.PutAdmPermissions)
	router.Put("/api/adm/credentials", admController.PutAdmCredentials)
	router.Get("/api/adm", admController.GetAdm)
	router.Get("/api/adms", admController.GetAdms)
	router.Post("/api/dino/category/food", dinoController.CreateFood)
	router.Delete("/api/dino/category/food", dinoController.DeleteFood)
	router.Post("/api/dino/category/locomotion", dinoController.CreateLocomotion)
	router.Delete("/api/dino/category/locomotion", dinoController.DeleteLocomotion)
	router.Post("/api/dino/category/region", dinoController.CreateRegion)
	router.Delete("/api/dino/category/region", dinoController.DeleteRegion)

	port, err := util.GetEnv("PORT")
	if err != nil {
		fmt.Println("Erro ao recuperar porta do ambiente: ", err)
		return
	}

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}
