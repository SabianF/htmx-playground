package web_socket

import (
	"net/http"

	common_data_repos "github.com/SabianF/htmx-playground/modules/common/data/repositories"

	gorilla_websocket "github.com/gorilla/websocket"
)

const ROUTE_WS string = "/ws"

var upgrader = gorilla_websocket.Upgrader{}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {

	// Initialize web socket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		common_data_repos.Log("serveWebSocket upgrader.Upgrade error: %v\n", err)
		return
	}
	defer ws.Close()

	common_data_repos.Log("Web socket client connected: %v", ws.RemoteAddr())

	// Listen for events
	for {
		// Receive message from client
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			common_data_repos.Log("ws.ReadMessage error: %v\n", err)
			break
		}
		common_data_repos.Log("Web socket message: %v", msg)

		// Send received message back to client
		err = ws.WriteMessage(msgType, msg)
		if err != nil {
			common_data_repos.Log("ws.WriteMessage error: %v\n", err)
			break
		}
	}
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_WS, serveWebSocket)
}
