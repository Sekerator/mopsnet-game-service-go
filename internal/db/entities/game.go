package entities

type GameModel struct {
	ID     string `json:"id" db:"id"`
	RoomId string `json:"room_id" db:"room_id"`
	UserId string `json:"user_id" db:"user_id"`
	Hp     int    `json:"hp" db:"hp"`
	Mp     int    `json:"mp" db:"mp"`
	Speed  int    `json:"speed" db:"speed"`
}
