package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type DinoService struct {
	dinoRepo       *db.DinoRepository
	locomotionRepo *db.LocomotionRepository
	regionRepo     *db.RegionRepository
	foodRepo       *db.FoodRepository
	admRepo        *db.AdmRepository
	imageRepo      *db.ImageRepository
}

func NewDinoService(
	dinoRepo *db.DinoRepository,
	locomotionRepo *db.LocomotionRepository,
	regionRepo *db.RegionRepository,
	foodRepo *db.FoodRepository,
	admRepo *db.AdmRepository,
	imageRepo *db.ImageRepository) *DinoService {

	return &DinoService{
		dinoRepo:       dinoRepo,
		locomotionRepo: locomotionRepo,
		regionRepo:     regionRepo,
		foodRepo:       foodRepo,
		admRepo:        admRepo,
		imageRepo:      imageRepo,
	}
}

func (s *DinoService) validateRegistryDino(dino models.DinoRegistryRequest) *util.ApiError {
	nameLen := len(dino.Name)
	if nameLen > util.MaxDinoName {
		message := fmt.Sprintf("Nome do dino muito longo!\nO tamanho máximo permitido é de %v caracteres.", util.MaxDinoName)
		return util.ThrowApiError(message, http.StatusBadRequest)
	} else if nameLen == 0 {
		return util.ThrowApiError("Informe o nome do dino!", http.StatusBadRequest)
	}

	utilityLen := len(dino.Utility)
	if utilityLen > util.MaxDinoUtility {
		message := fmt.Sprintf("Utilidade do dino muito longa!\nO tamanho máximo permitido é de %v caracteres.", util.MaxDinoUtility)
		return util.ThrowApiError(message, http.StatusBadRequest)
	} else if utilityLen == 0 {
		return util.ThrowApiError("Informe a utilidade do dino!", http.StatusBadRequest)
	}

	if len(dino.Training) == 0 {
		return util.ThrowApiError("Informe a descrição do dino!", http.StatusBadRequest)
	}

	existsLocomotion, err2 := s.locomotionRepo.ExistsLocomotionById(dino.LocomotionId)
	if err2 != nil {
		fmt.Printf("Erro ao verificar locomoção para id %v: %v\n", dino.LocomotionId, err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsLocomotion {
		return util.ThrowApiError("Informe uma locomoção válida!", http.StatusBadRequest)
	}

	existsRegion, err3 := s.regionRepo.ExistsRegionById(dino.RegionId)
	if err3 != nil {
		fmt.Printf("Erro ao verificar região para id %v: %v\n", dino.RegionId, err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsRegion {
		return util.ThrowApiError("Informe uma região válida!", http.StatusBadRequest)
	}

	existsFood, err4 := s.foodRepo.ExistsFoodById(dino.FoodId)
	if err4 != nil {
		fmt.Printf("Erro ao verificar alimentação para id %v: %v\n", dino.RegionId, err4)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsFood {
		return util.ThrowApiError("Informe um tipo de alimentação válida!", http.StatusBadRequest)
	}

	if len(dino.ImageBase64) > util.MaxImageSize {
		return util.ThrowApiError("O tamanho máximo da imagem é de "+util.MaxImageSizeLabel+"!", http.StatusPreconditionFailed)
	}

	return nil
}

func (s *DinoService) CreateDino(idAdm uint64, dino models.DinoRegistryRequest) *util.ApiError {

	apiErr := s.validateRegistryDino(dino)
	if apiErr != nil {
		return apiErr
	}

	existsDino, err := s.dinoRepo.ExistsDinoByName(dino.Name)
	if err != nil {
		fmt.Println("Erro ao verificar dino por nome: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if existsDino {
		return util.ThrowApiError("Já existe um dino com esse nome!", http.StatusPreconditionFailed)
	}

	idDino, err := s.dinoRepo.CreateDino(idAdm, dino)
	if err != nil {
		fmt.Println("Erro ao inserir dados no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusBadRequest)
	}

	if dino.ImageBase64 != "" {
		err = s.imageRepo.Insert(imageNameByDino(idDino), dino.ImageBase64)
		if err != nil {
			fmt.Println("Erro ao salvar imagem do dino: ", err)
			return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
		}
	}

	return nil
}

func (s *DinoService) PutDino(idDino uint64, idAdm uint64, dino models.DinoRegistryRequest) *util.ApiError {

	apiErr := s.validateRegistryDino(dino)
	if apiErr != nil {
		return apiErr
	}

	existsDino, err := s.dinoRepo.ExistsDinoById(idDino)
	if err != nil {
		fmt.Println("Erro ao buscar dino por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsDino {
		return util.ThrowApiError("Dino não encontrado!", http.StatusNotFound)
	}

	dinoIdWithName, err := s.dinoRepo.FindDinoIdByName(dino.Name)
	if err != nil {
		fmt.Println("Erro ao buscar id de dino por nome: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if dinoIdWithName != 0 && dinoIdWithName != idDino {
		return util.ThrowApiError("Já existe um dino com esse nome!", http.StatusPreconditionFailed)
	}

	if dino.ImageBase64 != "" {
		err = s.imageRepo.Put(imageNameByDino(idDino), dino.ImageBase64)
		if err != nil {
			fmt.Println("Erro ao atualizar imagem do dino: ", err)
			return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
		}
	}

	err = s.dinoRepo.PutDino(idDino, idAdm, dino)
	if err != nil {
		fmt.Println("Erro ao atualizar dados no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusBadRequest)
	}

	return nil
}

func (s *DinoService) DeleteDino(idAdm uint64, id uint64) *util.ApiError {
	dino, err := s.dinoRepo.FindDinoById(id)
	if err != nil {
		fmt.Println("Erro ao verificar dino por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if dino == nil {
		return util.ThrowApiError("Esse dino não existe!", http.StatusNotFound)
	}

	mainAdmId, err2 := s.admRepo.GetMainAdmId()
	if err2 != nil {
		fmt.Println("Erro ao buscar adm principal: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if dino.IdAdm != idAdm && idAdm != mainAdmId {
		return util.ThrowApiError("", http.StatusForbidden)
	}

	err = s.imageRepo.Delete(imageNameByDino(dino.Id))
	if err != nil {
		fmt.Println("Erro ao deletar imagem durante remoção do dino: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	err3 := s.dinoRepo.DeleteDino(id)
	if err3 != nil {
		fmt.Println("Erro ao deletar dino: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func (s *DinoService) FindDinoById(id uint64) (models.Dino, *util.ApiError) {
	dino, err := s.dinoRepo.FindDinoById(id)
	if err != nil {
		fmt.Println("Erro ao recuperar dino por id: ", err.Error())
		return models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if dino == nil {
		return models.Dino{}, util.ThrowApiError("Dino não encontrado!", http.StatusBadRequest)
	}

	return *dino, nil
}

func (s *DinoService) FindDinoByFilter(filter models.DinoFilter) ([]models.Dino, *util.ApiError) {
	dinos, err := s.dinoRepo.FindDinoByFilter(filter)
	if err != nil {
		fmt.Println("Erro ao buscar dinos por filtro: ", err)
		return []models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return dinos, nil
}

func (s *DinoService) FindDinoByFilterForAdm(idAdm uint64, filter models.DinoFilter) ([]models.Dino, *util.ApiError) {

	idMainAdm, err := s.admRepo.GetMainAdmId()
	if err != nil {
		fmt.Println("Erro ao buscar adm principal: ", err)
		return []models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	var dinos []models.Dino
	var err2 error

	if idMainAdm == idAdm {
		dinos, err2 = s.dinoRepo.FindDinoByFilterForMainAdm(filter)
	} else {
		dinos, err2 = s.dinoRepo.FindDinoByFilterForAdm(idAdm, filter)
	}

	if err2 != nil {
		fmt.Println("Erro ao buscar dinos por filtro (adm): ", err2)
		return []models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return dinos, nil
}

func (s *DinoService) GetImage(dinoId uint64) (models.Image, *util.ApiError) {
	dino, err := s.dinoRepo.FindDinoById(dinoId)
	if err != nil {
		fmt.Println("Erro ao buscar dino por id: ", err)
		return models.Image{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	base64, err := s.imageRepo.Get(imageNameByDino(dino.Id))
	if err != nil {
		fmt.Println("Erro ao buscar imagem por dino: ", err)
		return models.Image{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if base64 == "" {
		return models.Image{}, util.ThrowApiError("Esse dino não possui imagem!", http.StatusNotFound)
	}

	return models.Image{Base64: base64}, nil
}

func imageNameByDino(dinoId uint64) string {
	return fmt.Sprintf("dino-%d", dinoId)
}
