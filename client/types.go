package main

import (
	"client/location"
	"time"
)

type Client struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type SearchArea struct {
	Start location.Location `json:"start" redis:"start"`
	End   location.Location `json:"end" redis:"end"`
}

type Scooter struct {
	UUID       string    `json:"uuid" redis:"uuid"`
	Occupied   bool      `json:"occupied" redis:"occupied"`
	Lat        float64   `json:"lat" redis:"lat"`
	Lon        float64   `json:"lon" redis:"lon"`
	LastUpdate time.Time `json:"last_update" redis:"last_update"`
}

type Trip struct {
	UUID        string            `json:"uuid" redis:"uuid"`
	ScooterUUID string            `json:"scooter_uuid" redis:"scooter_uuid"`
	UserUUID    string            `json:"user_uuid" redis:"user_uuid"`
	Start       location.Location `json:"start"`
	End         location.Location `json:"end"`
}
