package note

import "time"

type Note struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int       `json:"user_id"`
}

type CreateNoteRequest struct {
	Text string `json:"text"`
}
