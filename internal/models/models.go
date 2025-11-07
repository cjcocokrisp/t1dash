package models

import (
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
)

// Struct that represents a user
type User struct {
	Id          pgxuuid.UUID // UUID
	Username    string       // username must be unique
	Firstname   string
	Lastname    string
	Password    string // Hashed with salt
	Avatar      string // File path to avatar
	Role        string // Current possible values: user or admin
	Settings    UserSettings
	Connections UserConnections
}

type UserSettings struct {
	TestField string `json:"test"`
}

type UserConnections struct {
	TestField string `json:"test"`
}
