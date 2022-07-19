package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Dino struct {
	Id           uint64    `json:"id"`
	IdAdm        uint64    `json:"-"`
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

type Adm struct {
	Id                        uint64 `json:"id"`
	Name                      string `json:"name"`
	PermissionManagerDino     bool   `json:"permissionManagerDino"`
	PermissionManagerCategory bool   `json:"permissionManagerCategory"`
	PermissionManagerAdm      bool   `json:"permissionManagerAdm"`
	MainAdm                   bool   `json:"mainAdm"`
}

type Claims struct {
	Id        uint64    `json:"id"`
	DateLogin time.Time `json:"dateLogin"`
	jwt.StandardClaims
}
