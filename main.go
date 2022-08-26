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

	//variáveis de ambiente
	port, err := util.GetEnv("PORT")
	if err != nil {
		fmt.Println("Erro ao recuperar porta do ambiente: ", err)
		return
	}
	secret, err := util.GetEnv("SECRET")
	if err != nil {
		fmt.Println("Erro ao recuperar secret do ambiente: ", err)
		return
	}

	//repositórios
	loginRepo := db.NewLoginRepository(database)
	dinoRepo := db.NewDinoRepository(database)
	locomotionRepo := db.NewLocomotionRepository(database)
	regionRepo := db.NewRegionRepository(database)
	foodRepo := db.NewFoodRepository(database)
	admRepo := db.NewAdmRepository(database)
	imageRepo := db.NewImageRepository(database)

	//services
	dinoService := service.NewDinoService(dinoRepo, locomotionRepo, regionRepo, foodRepo, admRepo, imageRepo)
	locomotionService := service.NewLocomotionService(locomotionRepo, dinoRepo)
	regionService := service.NewRegionService(regionRepo, dinoRepo)
	foodService := service.NewFoodService(foodRepo, dinoRepo)
	admService := service.NewAdmService(admRepo, dinoRepo)
	loginService := service.NewLoginService(secret, admRepo, loginRepo)

	//controllers
	dinoController := controller.NewDinoController(dinoService, loginService)
	admController := controller.NewAdmController(admService, loginService)
	loginController := controller.NewLoginController(loginService)
	categoryController := controller.NewCategoryController(loginService, locomotionService, regionService, foodService)

	//rotas públicas
	router.Post("/api/adm/login", loginController.Login)
	router.Get("/api/dino", dinoController.FindDinoById)
	router.Get("/api/dinos", dinoController.FindDinoByFilter)
	router.Get("/api/dino/categories", categoryController.DinoCategories)
	router.Get("/api/dino/image", dinoController.GetImage)

	//rotas privadas
	router.Post("/api/dino", dinoController.CreateDino)
	router.Put("/api/dino", dinoController.PutDino)
	router.Delete("/api/dino", dinoController.DeleteDino)
	router.Get("/api/adm/dinos", dinoController.FindDinoByFilterForAdm)
	router.Post("/api/adm/logout", loginController.Logout)
	router.Post("/api/adm", admController.CreateAdm)
	router.Delete("/api/adm", admController.DeleteAdm)
	router.Put("/api/adm/permissions", admController.PutAdmPermissions)
	router.Put("/api/adm/credentials", admController.PutAdmCredentials)
	router.Get("/api/adm", admController.GetAdm)
	router.Get("/api/adms", admController.GetAdms)
	router.Post("/api/dino/category/food", categoryController.CreateFood)
	router.Delete("/api/dino/category/food", categoryController.DeleteFood)
	router.Post("/api/dino/category/locomotion", categoryController.CreateLocomotion)
	router.Delete("/api/dino/category/locomotion", categoryController.DeleteLocomotion)
	router.Post("/api/dino/category/region", categoryController.CreateRegion)
	router.Delete("/api/dino/category/region", categoryController.DeleteRegion)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}
