package repositories

import (
	"game-service/internal/db/entities"
	"github.com/jmoiron/sqlx"
)

type RoomRepo struct {
	conn *sqlx.DB
}

func NewRoomRepo(conn *sqlx.DB) RoomRepository {
	return &RoomRepo{conn: conn}
}

func (ur *RoomRepo) GetRoomByID(id string, status int) (*entities.RoomModel, error) {
	var room entities.RoomModel

	sql := "SELECT * FROM room WHERE id=$1"

	err := ur.conn.Get(&room, sql, id)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (ur *RoomRepo) GetRoomsByUserId(userId string, statuses ...int) (*[]entities.RoomModel, error) {
	var room []entities.RoomModel

	sql := "SELECT * FROM room WHERE user1_id=$1 OR user2_id=$1"

	for i := 0; i < len(statuses); i++ {
		sql += " AND status=$2"
	}

	err := ur.conn.Get(&room, sql, userId, statuses)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (ur *RoomRepo) GetWaitingRooms() (*entities.RoomModel, error) {
	var room entities.RoomModel

	sql := "SELECT * FROM room WHERE status=$1 LIMIT 1"

	err := ur.conn.Get(&room, sql, entities.RoomStatusWait)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (ur *RoomRepo) CreateRoom(room *entities.RoomModel) (*entities.RoomModel, error) {
	sql := `INSERT INTO room (user1_id, status) 
            VALUES (:user1_id, :status)`

	_, err := ur.conn.NamedExec(sql, room)
	if err != nil {
		return nil, err
	}

	sql = "SELECT * FROM room WHERE user1_id=$1 AND status=$2"

	err = ur.conn.Get(room, sql, room.UserId1, room.Status)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (ur *RoomRepo) UpdateRoom(room *entities.RoomModel) (*entities.RoomModel, error) {
	sql := `UPDATE room
            SET user1_id = :user1_id, user2_id = :user2_id, status = :status)`

	_, err := ur.conn.NamedExec(sql, room)
	if err != nil {
		return nil, err
	}

	sql = "SELECT * FROM room WHERE user1_id=$1 AND status=$2"

	err = ur.conn.Get(room, sql, room.UserId1, room.Status)
	if err != nil {
		return nil, err
	}

	return room, nil
}
