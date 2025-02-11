package entities

import "time"

const (
	RoomStatusInactive = 0
	RoomStatusActive   = 1
	RoomStatusWait     = 2
	RoomStatusEnding
)

type RoomModel struct {
	ID        string    `json:"id" db:"id"`
	UserId1   string    `json:"user1_id" db:"user1_id"`
	UserId2   string    `json:"user2_id" db:"user2_id"`
	Status    int       `json:"status" db:"status"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
