package web_socket

import (
	"net/http"
	"time"

	common_data_repos "github.com/SabianF/htmx-playground/modules/common/data/repositories"
	pages "github.com/SabianF/htmx-playground/modules/web_socket/presentation/pages"

	gorilla_websocket "github.com/gorilla/websocket"
)

const ROUTE_PAGE string = "/web-socket"

func getPage(w http.ResponseWriter, r *http.Request) {
	pages.WebSocketPage().Render(r.Context(), w)
}

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

	// Listen and handle events
	sendFrequentMessages(ws)
	getAndRespondToMessages(ws)
}

func sendFrequentMessages(ws *gorilla_websocket.Conn) {
	for {
		time.Sleep(3 * time.Second)
		ws.WriteMessage(
			gorilla_websocket.TextMessage,
			[]byte("<div hx-swap-oob='beforeend:#chat'><article>Hello, there.</article></div>"),
		)
	}
}

func getAndRespondToMessages(ws *gorilla_websocket.Conn) {
	for {
		// Receive message from client
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			common_data_repos.Log("ws.ReadMessage error: %v\n", err)
			break
		}
		common_data_repos.Log("Web socket message: %v", msg)

		// Send received message back to client
		time.Sleep(2 * time.Second)
		err = ws.WriteMessage(msgType, msg)
		if err != nil {
			common_data_repos.Log("ws.WriteMessage error: %v\n", err)
			break
		}
	}
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_WS, serveWebSocket)
	mux.HandleFunc(ROUTE_PAGE, getPage)
}
