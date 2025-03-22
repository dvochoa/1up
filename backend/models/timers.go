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

type GetTimersResponse struct {
	Timers []Timer `json:"timers"`
}

// TODO: Think about how I want the interface to look, ultimately I am imagining two types of charts:
// 1. totalTime over time (monotomically increasing)
// 2. sessionDuration by day, week, month, year. Can hover each bucket to see how many sessions contributed
// TODO: What if the UI wants to view not just individual sessions but aggregations (like one per month)?
type GetTimerHistoryResponse struct {
	TimerSessions []TimerSession `json:"timerSessions"`
}
