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
func (store TimerStore) GetTimers(ctx context.Context, user_id int64) ([]models.Timer, error) {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	rows, _ := store.conn.Query(
		queryCtx,
		`SELECT ts.id as id, ts.title as title, COALESCE(SUM(tp.session_duration_in_seconds), 0) as totalTime
		 FROM (
		 	SELECT id, title FROM timersettings WHERE owner_id = $1
		 ) ts
		 LEFT JOIN timerprogress tp ON tp.timer_setting_id = ts.id
		 GROUP BY ts.id, ts.title;`,
		user_id,
	)
	timers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Timer])
	return timers, err
}

// TODO: Update this to return TimerSessions
// GetTimerById returns the timer with matching id, if any
func (store TimerStore) GetTimerById(ctx context.Context, id int) (models.Timer, error) {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	var timer models.Timer
	err := store.conn.QueryRow(
		queryCtx,
		"SELECT id, title FROM timers WHERE id = $1",
		id,
	).Scan(&timer.Id, &timer.Title)
	return timer, err
}

// TODO: Update
// CreateTimer inserts a new timer into the timers table
func (store TimerStore) CreateTimer(ctx context.Context, timer *models.Timer) error {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	err := store.conn.QueryRow(
		queryCtx,
		"INSERT INTO timers (title) VALUES ($1) RETURNING id",
		timer.Title,
	).Scan(&timer.Id)
	return err
}

// TODO: Update this to refer to TimerSettings
// UpdateTimer replaces the timer keyed by id.
// Throws an error when no matching timer is found
func (store TimerStore) UpdateTimer(ctx context.Context, id int, timer *models.Timer) error {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	_, err := store.conn.Exec(
		queryCtx,
		"INSERT INTO timers (id, title) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET title = $2 RETURNING id",
		timer.Id, timer.Title,
	)
	return err
}

// TODO: Update this to refer to TimerSettings, need to delete from both tables
// DeleteTimer deletes the timer matching the specified int from the timers table
func (store TimerStore) DeleteTimer(ctx context.Context, id int) error {
	queryCtx, cancel := getQueryCtx(ctx)
	defer cancel()

	commandTag, err := store.conn.Exec(queryCtx, "DELETE FROM timers WHERE id=$1", id)

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
