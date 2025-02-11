package repositories

import (
	"game-service/internal/db/entities"
	"github.com/jmoiron/sqlx"
)

type GameRepo struct {
	conn *sqlx.DB
}

func NewGameRepo(conn *sqlx.DB) GameRepository {
	return &GameRepo{conn: conn}
}

func (ur *GameRepo) GetGameByRoomId(roomId string) (*entities.GameModel, error) {
	var user entities.GameModel

	sql := "SELECT * FROM game WHERE room_id=$1"

	err := ur.conn.Get(&user, sql, roomId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *GameRepo) GetGameByUser(roomId string, userId string) (*entities.GameModel, error) {
	var user entities.GameModel

	sql := "SELECT * FROM game WHERE room_id=$1 AND user_id=$2"

	err := ur.conn.Get(&user, sql, roomId, userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
