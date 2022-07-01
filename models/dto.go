package models

type DinoRegistryRequest struct {
	Name         string `json:"name"`
	Training     string `json:"training"`
	Utility      string `json:"utility"`
	RegionId     uint64 `json:"regionId"`
	LocomotionId uint64 `json:"locomotionId"`
	FoodId       uint64 `json:"foodId"`
}

type CategoryRegistryRequest struct {
	Name string `json:"name"`
}

type DinoFilter struct {
	Name         string
	RegionId     uint64
	LocomotionId uint64
	FoodId       uint64
}

type DinoCategoryResponse struct {
	Regions     []Category `json:"regions"`
	Locomotions []Category `json:"locomotions"`
	Foods       []Category `json:"foods"`
}
