package models

import "time"

type Dino struct {
	Id           uint64    `json:"id"`
	Name         string    `json:"name"`
	Region       string    `json:"region"`
	Locomotion   string    `json:"locomotion"`
	Food         string    `json:"food"`
	Training     string    `json:"training"`
	Utility      string    `json:"utility"`
	CreationDate time.Time `json:"creationDate"`
}

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
