package main

import (
	"golangwebsocketchatapp/handlers"
	"golangwebsocketchatapp/websockets"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {

	// Create views engine
	viewsEngine := html.New("./views", ".html")

	// Start new fiber instance
	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	// Static route and directory
	app.Static("/static/", "./static")
	//create handlers
	appHandler := handlers.NewAppHandler()
	app.Get("/", appHandler.HandleGetIndex)

	// create new webscoket
	server := websockets.NewWebSockethand()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessages()

	app.Listen(":3000")
}
