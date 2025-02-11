package handlers

import (
	"fmt"
	"game-service/internal/services"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type GameHand struct {
	GameServ services.GameServices
	Clients  *map[string]*websocket.Conn
	Mu       *sync.Mutex
}

func NewGameHandler(gameServ services.GameServices, clients *map[string]*websocket.Conn, mu *sync.Mutex) GameHandler {
	return &GameHand{
		gameServ,
		clients,
		mu,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (uh *GameHand) Ws(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer ws.Close()

	uh.Mu.Lock()
	(*uh.Clients)["SS"] = ws
	uh.Mu.Unlock()
	fmt.Println("Новый клиент подключился")

	var msg []byte

	for {
		_, msg, err = ws.ReadMessage()
		if err != nil {
			fmt.Println("Клиент отключился")
			uh.Mu.Lock()
			delete(*uh.Clients, "SS")
			uh.Mu.Unlock()
			break
		}
		fmt.Printf("Получено: %s\n", msg)
	}

	uh.Mu.Lock()
	for id, client := range *uh.Clients {
		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			client.Close()
			delete(*uh.Clients, id)
		}
	}
	uh.Mu.Unlock()
}
