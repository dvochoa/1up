package models

import "time"

type TimerOverview struct {
	Id                 int64  `json:"id"`
	OwnerId            int64  `json:"ownerId"`
	Title              string `json:"title"`
	TotalTimeInSeconds int64  `json:"totalTimeInSeconds"`
}

type GetTimersResponse struct {
	TimerOverviews []TimerOverview `json:"timerOverviews"`
}

type CreateTimerRequest struct {
	OwnerId int64  `json:"ownerId"`
	Title   string `json:"title"`
}

type CreateTimerSessionRequest struct {
	SessionDurationInSeconds int32 `json:"sessionDurationInSeconds"`
}

type TimerSession struct {
	Id                       int64     `json:"id"`
	CreatedAt                time.Time `json:"createdAt"`
	SessionDurationInSeconds int32     `json:"sessionDurationInSeconds"`
}

type GetTimerHistoryResponse struct {
	TimerSessions []TimerSession `json:"timerSessions"`
}
