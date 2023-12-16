package controller

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func HandleStreamSocket(stream ports.StramingComputerService) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		id := c.Params("id")
		role := c.Params("role")
		defer func() {
			stream.Receiver(func() {
				stream.RemoveConnection(id, role)
			})
		}()

		stream.Receiver(func() {
			stream.AddConnection(id, role, c)
		})

		for {
			Command := &models.StreamEvent{}

			err := c.ReadJSON(Command)
			if err == websocket.ErrBadHandshake {
				log.Println("error de :", err.Error())
				break
			}
			if err == websocket.ErrCloseSent {
				log.Println("error de closesent:", err.Error())
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("loaded command: ", Command)

			stream.Receiver(func() {
				stream.Broadcast(Command)
			})
		}

	}
}
