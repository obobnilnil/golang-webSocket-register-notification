package webSockets

import (
	"log"
	"net/http"
	"webSocket_git/register/models"
	"webSocket_git/register/transactions"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]models.UserInfo) // Store UserInfo for each connection

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all CORS for now
	},
}

func HandleMessages() {
	for {
		// Assuming you have a broadcast channel where you receive messages
		msg, ok := <-transactions.Broadcast
		if !ok {
			return // Channel closed
		}

		// Extract CompanyID from msg, which is of type primitive.D
		var companyID string
		for _, item := range msg {
			if item.Key == "companyID" {
				companyID, ok = item.Value.(string)
				if !ok {
					log.Println("Error asserting companyID to string")
					return
				}
				break
			}
		}

		// Broadcast the message to all clients with matching role and companyID
		for client, info := range clients {
			if info.Role == "admin" && info.CompanyID == companyID {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("WebSocket error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	companyID := r.URL.Query().Get("companyID")

	if role == "" || companyID == "" {
		http.Error(w, "Missing role or companyID", http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer ws.Close()

	// Save the new client with its role and companyID
	clients[ws] = models.UserInfo{Role: role, CompanyID: companyID}

	// Read messages from the client (this is optional, depends on your use case)
	for {
		// Create a variable to read incoming messages into
		var msg interface{} // Use the appropriate type for your application
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}
		// Process the message as needed
	}

	// Remove the client when the loop is exited
	delete(clients, ws)
}
