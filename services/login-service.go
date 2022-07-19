package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type LoginService struct {
	admRepo   *db.AdmRepository
	loginRepo *db.LoginRepository
}

func NewLoginService(admRepo *db.AdmRepository, loginRepo *db.LoginRepository) *LoginService {
	return &LoginService{admRepo: admRepo, loginRepo: loginRepo}
}

func (l *LoginService) Login(credentials models.LoginRequest) (string, *util.ApiError) {
	idAdm, err := l.admRepo.GetAdmIdByCredentials(credentials.Name, credentials.Password)
	if err != nil {
		fmt.Println("Erro ao recuperar id do usuário no banco: ", err)
		return "", util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	} else if idAdm == 0 {
		return "", util.ThrowApiError("Usuário ou senha inválidos!", http.StatusBadRequest)
	}

	cl := &models.Claims{
		Id:        idAdm,
		DateLogin: time.Now(),
	}

	err = l.loginRepo.SaveClaims(cl)
	if err != nil {
		return "", util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	secret := GetJwtSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)

	tokenStr, err2 := token.SignedString(secret)
	if err2 != nil {
		fmt.Println("Erro ao obter token: ", err2)
		return "", util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return tokenStr, nil
}

func (l *LoginService) GetIdByToken(token string) (uint64, *util.ApiError) {
	secret := GetJwtSecret()
	cl := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(token, cl, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return 0, util.ThrowApiError("", http.StatusUnauthorized)
	}
	if !tkn.Valid {
		return 0, util.ThrowApiError("", http.StatusUnauthorized)
	}

	existsToken, err := l.loginRepo.ExistsClaims(cl)
	if err != nil {
		return 0, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsToken {
		return 0, util.ThrowApiError("", http.StatusUnauthorized)
	}

	return cl.Id, nil
}

func (l *LoginService) CheckPermissionManagerDino(clientId uint64) *util.ApiError {
	adm, err := l.admRepo.GetAdmById(clientId)
	if err != nil {
		fmt.Println("Erro ao recuperar permissões do administrador: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if adm.PermissionManagerDino {
		return nil
	}

	return util.ThrowApiError("", http.StatusForbidden)
}

func (l *LoginService) CheckPermissionManagerCategory(clientId uint64) *util.ApiError {
	adm, err := l.admRepo.GetAdmById(clientId)
	if err != nil {
		fmt.Println("Erro ao recuperar permissões do administrador: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if adm.PermissionManagerCategory {
		return nil
	}

	return util.ThrowApiError("", http.StatusForbidden)
}

func (l *LoginService) CheckPermissionManagerAdm(clientId uint64) *util.ApiError {
	adm, err := l.admRepo.GetAdmById(clientId)
	if err != nil {
		fmt.Println("Erro ao recuperar permissões do administrador: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if adm.PermissionManagerAdm {
		return nil
	}

	return util.ThrowApiError("", http.StatusForbidden)
}

func GetJwtSecret() []byte {
	return []byte("AABBCC")
}
