package db

import (
	"context"
	"encoding/json"

	"github.com/cjcocokrisp/t1dash/internal/models"
	"github.com/cjcocokrisp/t1dash/pkg/util"

	//pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateUser inserts a user struct into the database and returns it's id
// All fields must match the state need for insert
// id is handled by the database and role is ignored and should be updated after creation
func CreateUser(user *models.User) (*pgtype.UUID, error) {
	if DBPool == nil {
		return nil, util.NullDBConnection()
	}

	query := "INSERT INTO users (username, firstname, lastname, password, avatar, settings, connections)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	settingsBytes, err := json.Marshal(user.Settings)
	if err != nil {
		return nil, err
	}

	connectionsBytes, err := json.Marshal(user.Connections)
	if err != nil {
		return nil, err
	}

	var id pgtype.UUID
	err = DBPool.QueryRow(context.Background(), query, user.Username, user.Firstname, user.Lastname, user.Password,
		user.Avatar, settingsBytes, connectionsBytes).Scan(&id)

	return &id, err
}

// GetUserByUsername reads a user by username and then returns a user struct for that user
func GetUserByUsername(username string) (*models.User, error) {
	if DBPool == nil {
		return nil, util.NullDBConnection()
	}

	query := "SELECT * FROM users WHERE username=$1"

	rows, err := DBPool.Query(context.Background(), query, username)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUsersName updates a users first and last name if you are keeping one the same just pass the original value
func UpdateUsersName(username string, firstname string, lastname string) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE users SET firstname=$1, lastname=$2 WHERE username=$3"

	_, err := DBPool.Exec(context.Background(), query, firstname, lastname, username)

	return err
}

// UpdateUserPassword updates a users password, password should be hashed with salt in db
func UpdateUserPassword(username string, password string) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE users SET password=$1 WHERE username=$2"

	_, err := DBPool.Exec(context.Background(), query, password, username)

	return err

}

// UpdateUserRole updates the users role
func UpdateUserRole(username string, role string) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE users SET role=$1 WHERE username=$2"

	_, err := DBPool.Exec(context.Background(), query, role, username)

	return err
}

// UpdateUserAvatar updates the users avatar path
func UpdateUserAvatar(username string, avatar string) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE users SET avatar=$1 WHERE username=$2"

	_, err := DBPool.Exec(context.Background(), query, avatar, username)

	return err
}

// UpdateUserSettings updates the users settings
func UpdateUserSettings(username string, settings models.UserSettings) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	bytes, err := json.Marshal(settings)
	if err != nil {
		return nil
	}

	query := "UPDATE users SET settings=$1 WHERE username=$2"

	_, err = DBPool.Exec(context.Background(), query, bytes, username)

	return err
}

// UpdateUserConnections updates the users connections
func UpdateUserConnections(username string, connections models.UserConnections) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	bytes, err := json.Marshal(connections)
	if err != nil {
		return nil
	}

	query := "UPDATE users SET connections=$1 WHERE username=$2"

	_, err = DBPool.Exec(context.Background(), query, bytes, username)

	return err
}

// DeleteUserByUsername deletes the specified user from the table
func DeleteUserByUsername(username string) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "DELETE FROM users WHERE username=$1"

	_, err := DBPool.Exec(context.Background(), query, username)

	return err
}

// CheckIfUsersExist returns a bool of whether or not users exist in the users table
func CheckIfUsersExist() (bool, error) {
	if DBPool == nil {
		return false, util.NullDBConnection()
	}

	query := "SELECT count(*) FROM users"

	var count int
	err := DBPool.QueryRow(context.TODO(), query).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
