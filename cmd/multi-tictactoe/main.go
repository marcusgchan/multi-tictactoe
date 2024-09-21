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
	GameId   string
	PlayerId string
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

func (self *Game) AddPlayer(ws *websocket.Conn, id PlayerId) {
	player := Player{ID: id, Connection: ws}
	self.Players = append(self.Players, player)
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

func initGameState() *Game {
	players := []Player{}
	return &Game{Started: false, Players: players}
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

	var playerId PlayerId

	playerCookie, err := r.Cookie("playerId")
	if err != nil {
		playerId = PlayerId(uuid.New().String())
	} else {
		playerId = PlayerId(playerCookie.Value)
	}

	gameId := GameId(r.PathValue("id"))
	existingGame, ok := games[gameId]
	if !ok {
		// Initialize game state
		playerId := uuid.New().String()
		newGameState := *initGameState()

		newGameState.AdminID = PlayerId(playerId)

		newGameState.AddPlayer(ws, PlayerId(playerId))

		games[gameId] = newGameState
	} else {
		existingGame.AddPlayer(ws, PlayerId(playerId))
	}

	// Add player to game

	// for {
	// }
}
