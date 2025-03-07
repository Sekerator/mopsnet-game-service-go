package main

import (
	"game-service/config"
	"game-service/internal/db/repositories"
	"game-service/internal/handlers"
	"game-service/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var clients = make(map[string]*websocket.Conn)
var mu sync.Mutex

func main() {
	logFile, err := os.OpenFile("logs/requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла для логов: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("Ошибка загрузки данных: %v", err)
		return
	}

	conn, err := sqlx.Connect("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
		return
	}
	defer conn.Close()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger,
		NoColor: true,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	gameRepository := repositories.NewGameRepo(conn)
	gameService := services.NewGameService(gameRepository)
	gameHandler := handlers.NewGameHandler(gameService, &clients, &mu)

	r.Route("/", func(r chi.Router) {
		r.Get("/ws", gameHandler.Ws)
	})

	log.Println("Сервер запущен по адресу: ", cfg.ListenAddrAndPort())
	log.Println("Версия: ", cfg.Env)
	logger.Println("Сервер запущен по адресу: ", cfg.ListenAddrAndPort())
	logger.Println("Версия: ", cfg.Env)

	if cfg.Env == "local" {
		err = http.ListenAndServe(cfg.ListenAddrAndPort(), r)
	} else {
		go func() {
			log.Println("Redirecting HTTP to HTTPS")
			_ = http.ListenAndServe(cfg.ListenAddrAndPort(), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
			}))
		}()

		err = http.ListenAndServeTLS(cfg.ListenAddrAndPort(),
			"/etc/letsencrypt/live/game.mopsnet.ru/fullchain.pem",
			"/etc/letsencrypt/live/game.mopsnet.ru/privkey.pem",
			r)
	}

	if err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
		return
	}

	log.Println("Сервер отключен")
	logger.Println("Сервер отключен")
}
