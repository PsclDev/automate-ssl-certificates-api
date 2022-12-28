package models

type Certificate struct {
	Name string `json:"name" validate:"required"`
	DNS string `json:"dns" validate:"required"`
	IP string `json:"ip" validate:"required"`
}