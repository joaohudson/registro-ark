package models

type Dino struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	Locomotion string `json:"locomotion"`
	Food       string `json:"food"`
	Training   string `json:"training"`
	Utility    string `json:"utility"`
}

type DinoCategoryResponse struct {
	Regions     []string `json:"regions"`
	Locomotions []string `json:"locomotions"`
	Foods       []string `json:"foods"`
}
