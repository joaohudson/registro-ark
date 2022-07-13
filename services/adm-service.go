package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type AdmService struct {
	admRepo *db.AdmRepository
}

func NewAdmService(admRepo *db.AdmRepository) *AdmService {
	return &AdmService{admRepo: admRepo}
}

func (a *AdmService) CreateAdm(adm models.AdmRegistryRequest) *util.ApiError {

	existsAdm, err := a.admRepo.ExistsAdmByName(adm.Name)
	if err != nil {
		fmt.Println("Erro ao buscar adm por nome: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if existsAdm {
		return util.ThrowApiError("Já existe um administrador com este nome!", http.StatusPreconditionFailed)
	}

	err2 := a.admRepo.CreateAdm(adm)
	if err2 != nil {
		fmt.Println("Erro ao criar conta de adminstrador: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
