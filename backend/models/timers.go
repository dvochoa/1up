package models

type Timer struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	TotalTime int64  `json:"totalTime"`
}

type TimerSession struct {
	Id              int64 `json:"id"`
	Timestamp       int64 `json:"timestamp"`
	SessionDuration int32 `json:"sessionDuration"`
}

type GetTimersResponse struct {
	Timers []Timer `json:"timers"`
}

type GetTimerSessionsResponse struct {
	TimerSessions []TimerSession `json:"timerSessions"`
}
