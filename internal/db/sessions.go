package db

import (
	"context"
	"time"

	"github.com/cjcocokrisp/t1dash/internal/config"
	"github.com/cjcocokrisp/t1dash/internal/models"
	"github.com/cjcocokrisp/t1dash/pkg/util"

	//pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateSession(userId pgtype.UUID, ip string) (*pgtype.UUID, error) {
	if DBPool == nil {
		return nil, util.NullDBConnection()
	}

	query := "INSERT INTO sessions (user_id, created_at, expires_at, last_seen, valid, ip)" +
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	now := time.Now()
	expireAt := now.Add(time.Duration(config.AppCfg.SessionTTL) * time.Hour)

	var id pgtype.UUID
	err := DBPool.QueryRow(context.Background(), query, userId, now, expireAt, now, true, ip).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func GetSessionById(sessionId pgtype.UUID) (*models.Session, error) {
	if DBPool == nil {
		return nil, util.NullDBConnection()
	}

	query := "SELECT * FROM sessions WHERE id=$1"

	rows, err := DBPool.Query(context.Background(), query, sessionId)

	session, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Session])
	if err != nil {
		return nil, err
	}

	return &session, err
}

func UpdateLastSeen(sessionId pgtype.UUID, lastSeen time.Time) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE sessions SET last_seen=$1 WHERE id=$2"

	_, err := DBPool.Exec(context.Background(), query, lastSeen, sessionId)

	return err
}

func InvalidateSession(sessionId pgtype.UUID) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE sessions SET valid=$1 WHERE id=$2"

	_, err := DBPool.Exec(context.Background(), query, false, sessionId)

	return err
}

func InvalidateAllSessionsForUser(userId pgtype.UUID) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "UPDATE sessions SET valid=$1 WHERE user_id=$2"

	_, err := DBPool.Exec(context.Background(), query, false, userId)

	return err
}

func DeleteSession(sessionId pgtype.UUID) error {
	if DBPool == nil {
		return util.NullDBConnection()
	}

	query := "DELETE FROM sessions WHERE id=$1"

	_, err := DBPool.Exec(context.Background(), query, sessionId)

	return err
}
