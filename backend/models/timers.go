package models

type Timer struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type GetTimersResponse struct {
	Timers []Timer `json:"timers"`
}
