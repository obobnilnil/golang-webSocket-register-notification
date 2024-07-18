package webSocket

import (
	// userWebSockets "chicCRM/modules/finalCode/webSockets"
	adminWebSockets "webSocket_git/register/webSockets"
)

func RunWebSocketHandlers() {
	// go webSockets.HandleMessages()     // old webSockets at handler
	// go webSockets.HandleMessagesUser() // Approve for notification
	// go webSockets.HandleMessagesManager()
	// go webSockets.HandleMessagesManagerRequestNotification() // notification request manager
	go adminWebSockets.HandleMessages() // Register for notification
	// go userWebSockets.HandleMessagesUser() // Approve for notification
}
