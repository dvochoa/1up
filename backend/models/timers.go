package models

import "time"

type Timer struct {
	Id        int64  `json:"id"`
	OwnerId   int64  `json:"ownerId"`
	Title     string `json:"title"`
	TotalTime int64  `json:"totalTime"`
}

type TimerSession struct {
	Id                       int64     `json:"id"`
	SessionTimestamp         time.Time `json:"sessionTimestamp"`
	SessionDurationInSeconds int32     `json:"sessionDurationInSeconds"`
}

type CreateTimerRequest struct {
	OwnerId int64  `json:"ownerId"`
	Title   string `json:"title"`
}

type GetTimersResponse struct {
	Timers []Timer `json:"timers"`
}

type GetTimerHistoryResponse struct {
	TimerSessions []TimerSession `json:"timerSessions"`
}
