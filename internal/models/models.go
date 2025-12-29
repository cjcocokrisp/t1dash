package models

import (
	"time"

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

// Struct that represents a session
type Session struct {
	Id        pgxuuid.UUID
	UserId    pgxuuid.UUID `db:"user_id"` // Associated user id
	CreatedAt time.Time    `db:"created_at"`
	ExpiresAt time.Time    `db:"expires_at"`
	LastSeen  time.Time    `db:"last_seen"`
	Valid     bool
	Ip        string
}
