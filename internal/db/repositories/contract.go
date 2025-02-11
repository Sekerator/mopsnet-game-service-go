package repositories

import "game-service/internal/db/entities"

type GameRepository interface {
	GetGameByUser(roomId string, userId string) (*entities.GameModel, error)
	GetGameByRoomId(roomId string) (*entities.GameModel, error)
}

type RoomRepository interface {
	UpdateRoom(room *entities.RoomModel) (*entities.RoomModel, error)
	CreateRoom(room *entities.RoomModel) (*entities.RoomModel, error)
	GetWaitingRooms() (*entities.RoomModel, error)
	GetRoomsByUserId(userId string, statuses ...int) (*[]entities.RoomModel, error)
	GetRoomByID(id string, status int) (*entities.RoomModel, error)
}
