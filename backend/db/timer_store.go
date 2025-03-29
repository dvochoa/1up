package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dvochoa/1up/models"

	"github.com/jackc/pgx/v5"
)

type TimerStore struct {
	conn *pgx.Conn
}

// NewTimerStore creates an instance of a TimerStore
func NewTimerStore(ctx context.Context, connStr string) (*TimerStore, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		_ = fmt.Errorf("failed to ping database: %w", err)
		return nil, err
	}
	log.Println("Successfully connected to PostgreSQL database")

	return &TimerStore{conn: conn}, nil
}

// CloseConnection properly closes the database connection pool
func (store TimerStore) CloseConnection(ctx context.Context) {
	if store.conn != nil {
		log.Println("Closing database connection")
		if err := store.conn.Close(ctx); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}
}

// GetTimers returns all timers
func (store TimerStore) GetTimers(ctx context.Context, userId int64) ([]models.TimerOverview, error) {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	rows, _ := store.conn.Query(
		queryCtx,
		`SELECT 
			ts.id as id,
			ts.owner_id as ownerId,
			ts.title as title,
			COALESCE(SUM(tp.session_duration_in_seconds), 0) as totalTimeInSeconds
		 FROM (
		 	SELECT id, owner_id, title FROM timersettings WHERE owner_id = $1
		 ) ts
		 LEFT JOIN timersessions tp ON tp.timer_setting_id = ts.id
		 GROUP BY ts.id, ts.owner_id, ts.title;`,
		userId,
	)
	timers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TimerOverview])
	return timers, err
}

// GetTimerSessions returns the timer with matching id, if any
func (store TimerStore) GetTimerSessions(ctx context.Context, timerSettingId int) ([]models.TimerSession, error) {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	rows, _ := store.conn.Query(
		queryCtx,
		"SELECT id, session_duration_in_seconds, created_at FROM timerSessions WHERE timer_setting_id = $1",
		timerSettingId,
	)
	timerSessions, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TimerSession])
	return timerSessions, err
}

// AddTimerSession inserts a new session row into the timerSessions table
func (store TimerStore) AddTimerSession(
	ctx context.Context,
	timerSettingId int64,
	createTimerSessionRequest models.CreateTimerSessionRequest) (models.TimerSession, error) {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	timerSession := models.TimerSession{SessionDurationInSeconds: createTimerSessionRequest.SessionDurationInSeconds}
	err := store.conn.QueryRow(
		queryCtx,
		`INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds)
		 VALUES ($1, $2)
		 RETURNING id, created_at
		 `,
		timerSettingId, createTimerSessionRequest.SessionDurationInSeconds,
	).Scan(&timerSession.Id, &timerSession.CreatedAt)
	return timerSession, err
}

// CreateTimerSetting inserts a new timer into the timerSettings table
func (store TimerStore) CreateTimerSetting(
	ctx context.Context,
	createReq *models.CreateTimerRequest) (models.TimerOverview, error) {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	timer := models.TimerOverview{OwnerId: createReq.OwnerId, Title: createReq.Title}
	err := store.conn.QueryRow(
		queryCtx,
		"INSERT INTO timerSettings (owner_id, title) VALUES ($1, $2) RETURNING id",
		timer.OwnerId, timer.Title,
	).Scan(&timer.Id)
	return timer, err
}

// UpdateTimer replaces the timer keyed by id.
// Throws an error when no matching timer is found
func (store TimerStore) UpdateTimerSettings(ctx context.Context, timer *models.TimerOverview) error {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	_, err := store.conn.Exec(
		queryCtx,
		"UPDATE timerSettings SET title = $2 WHERE id = $1",
		timer.Id, timer.Title,
	)
	return err
}

// DeleteTimerSettings deletes the timer keyed by the specified id from the timers table
// Deletes cascade down to child tables.
func (store TimerStore) DeleteTimerSettings(ctx context.Context, id int) error {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	commandTag, err := store.conn.Exec(queryCtx, "DELETE FROM timerSettings WHERE id=$1", id)

	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}

	return nil
}

func getQueryCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 5*time.Second)
}
