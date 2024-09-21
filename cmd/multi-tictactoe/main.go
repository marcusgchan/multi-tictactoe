package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type (
	GameId   = string
	PlayerId = string
)

type Player struct {
	ID         PlayerId
	Connection *websocket.Conn
}

type Game struct {
	Turn    PlayerId
	Started bool
	AdminID PlayerId
	Players []Player
}

var games = make(map[GameId]Game)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	mux.HandleFunc("GET /", handleHome)
	mux.HandleFunc("GET /game/{id}", handleGame)

	http.ListenAndServe(":8000", mux)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
}

func handleCreateLobby(w http.ResponseWriter, r *http.Request) {
	lobbyId := uuid.New().String()
	playerId := uuid.New().String()

	player := Player{ID: playerId}
	players := []Player{}
	newGame := Game{Turn: "", Started: false, AdminID: playerId}
}

func handleLobby(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/lobby.html")
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	gameId := r.PathValue("id")
	game, ok := games[gameId]
	if !ok {
		// Initialize game state
	} else {
		// Add player to game if exist
	}

	for {
	}
}
