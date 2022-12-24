package models

type DomainName struct {
	Country string `json:"country" validate:"required"`
	State string `json:"state" validate:"required"`
	Location string `json:"location" validate:"required"`
	Domain string `json:"domain" validate:"required"`
	Tld string `json:"tld" validate:"required"`
}