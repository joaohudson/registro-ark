package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type AdmService struct {
	admRepo  *db.AdmRepository
	dinoRepo *db.DinoRepository
}

func NewAdmService(admRepo *db.AdmRepository, dinoRepo *db.DinoRepository) *AdmService {
	return &AdmService{admRepo: admRepo, dinoRepo: dinoRepo}
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

func (a *AdmService) DeleteAdm(idAdm uint64) *util.ApiError {

	adm, err := a.admRepo.GetAdmById(idAdm)
	if err != nil {
		fmt.Println("Erro ao buscar adm por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if adm == nil {
		return util.ThrowApiError("Esse adm não existe!", http.StatusPreconditionFailed)
	}
	if adm.MainAdm {
		return util.ThrowApiError("O administrador principal não pode ser deletado!", http.StatusPreconditionFailed)
	}

	dinos, err := a.dinoRepo.FindDinoByFilterForAdm(idAdm, models.DinoFilter{})
	if err != nil {
		fmt.Println("Erro ao buscar dinos por adm: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if len(dinos) > 0 {
		return util.ThrowApiError("Esse administrador possui dinos, sendo assim, não pode ser removido!", http.StatusPreconditionFailed)
	}

	err = a.admRepo.DeleteAdm(idAdm)
	if err != nil {
		fmt.Println("Erro ao deletar administrador: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func (a *AdmService) PutAdmPermissions(permissions models.AdmChangePermissionsRequest) *util.ApiError {
	adm, err := a.admRepo.GetAdmById(permissions.Id)
	if err != nil {
		fmt.Println("Erro ao recuperar dados do administrador: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if adm == nil {
		return util.ThrowApiError("Este administrador não existe!", http.StatusNotFound)
	}
	if adm.MainAdm {
		return util.ThrowApiError("O administrador principal não pode ter suas permissões alteradas!", http.StatusPreconditionFailed)
	}

	err = a.admRepo.PutPermissions(permissions)
	if err != nil {
		fmt.Println("Erro ao modificar permissões de administrador: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	return nil
}

func (a *AdmService) PutAdmCredentials(idAdm uint64, credentials models.AdmChangeCredentialsRequest) *util.ApiError {

	id, err := a.admRepo.GetAdmIdByCredentials(credentials.Name, credentials.Password)
	if err != nil {
		fmt.Println("Erro ao buscar administrador por credenciais:", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if id != idAdm {
		return util.ThrowApiError("Usuário ou senha inválidos!", http.StatusBadRequest)
	}

	err = a.admRepo.PutCredentials(idAdm, credentials.NewName, credentials.NewPassowrd)
	if err != nil {
		fmt.Println("Erro ao atualizar credenciais do administrador: ", err)
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

func (a *AdmService) GetAdms() ([]models.Adm, *util.ApiError) {
	adms, err := a.admRepo.GetAdms()
	if err != nil {
		fmt.Println("Erro ao listar administradores: ", err)
		return []models.Adm{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return adms, nil
}
