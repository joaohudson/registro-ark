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
	nameLen := len(adm.Name)
	if nameLen > util.MaxNameAdm {
		message := fmt.Sprintf("Nome do administrador muito longo! O tamanho máximo permitido é de %v caracteres.", util.MaxNameAdm)
		return util.ThrowApiError(message, http.StatusBadRequest)
	} else if nameLen == 0 {
		return util.ThrowApiError("Informe o nome do administrador!", http.StatusBadRequest)
	}

	passwordLen := len(adm.Password)
	if passwordLen > util.MaxNameAdm {
		message := fmt.Sprintf("Senha do administrador muito longa! O tamanho máximo permitido é de %v caracteres.", util.MaxPasswordAdm)
		return util.ThrowApiError(message, http.StatusBadRequest)
	} else if passwordLen == 0 {
		return util.ThrowApiError("Informe a senha do administrador!", http.StatusBadRequest)
	}

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

func (a *AdmService) GetAdm(idAdm uint64) (*models.Adm, *util.ApiError) {
	adm, err := a.admRepo.GetAdmById(idAdm)
	if err != nil {
		return nil, util.ThrowApiError("", http.StatusUnauthorized)
	}

	return adm, nil
}

func (a *AdmService) GetAdms(idAdm uint64) ([]models.Adm, *util.ApiError) {
	adm, err := a.admRepo.GetAdmById(idAdm)
	if err != nil {
		fmt.Println("Erro ao recuperar dados do administrador: ", err)
		return []models.Adm{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if adm == nil {
		fmt.Println("Adm não encontrado!")
		return []models.Adm{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	adms, err := a.admRepo.GetAdms()
	if err != nil {
		fmt.Println("Erro ao listar administradores: ", err)
		return []models.Adm{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return adms, nil
}
