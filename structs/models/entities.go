package models

import "time"

type Tp struct {
	ID                     int       `gorm:"primary_key" json:"id"`
	Name                   string    `gorm:"type:varchar(255);not null;" json:"name"`
	AuthSuccessRedirectUrl string    `gorm:"type:varchar(255);not null;" json:"auth_success_redirect_url"`
	ClientId               string    `gorm:"type:varchar(255);not null;unique_index" json:"client_id"`
	ClientKey              string    `gorm:"type:varchar(255);not null;" json:"client_key"`
	DecodeKey              string    `gorm:"type:varchar(255);not null;" json:"decode_key"`
	Ticket                 string    `gorm:"type:varchar(255);" json:"ticket"`
	AccessToken            string    `gorm:"type:varchar(255);" json:"access_token"`
	ExpiredAt              time.Time `gorm:"default:null" json:"expired_at"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type App struct {
	ID           int       `gorm:"primary_key" json:"id"`
	TpId         int       `gorm:"type:int;" json:"tp_id"`
	AppId        string    `gorm:"default:null;type:varchar(255);unique_index" json:"app_id"`
	AppName      string    `gorm:"default:null;type:varchar(255);" json:"app_name"`
	AccessToken  string    `gorm:"type:varchar(255);" json:"access_token"`
	RefreshToken string    `gorm:"type:varchar(255);" json:"refresh_token"`
	Data         string    `gorm:"default:null;type:text;" json:"data"`
	ExpiredAt    time.Time `gorm:"default:null" json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Tp           Tp        `json:"tp"`
}
