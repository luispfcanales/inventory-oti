package ports

import (
	"github.com/gofiber/websocket/v2"
	"github.com/luispfcanales/inventory-oti/models"
)

// services computers
type StorageComputerService interface {
	SelectComputers() []*models.CPU
}
type ComputerService interface {
	ListComputers() []*models.CPU
}

type StramingComputerService interface {
	AddConnection(id, role string, c *websocket.Conn)
	RemoveConnection(id, role string)
	Broadcast(*models.StreamEvent)
	ListAllConnections(chan<- models.StreamEvent)
	Receiver(func())
}
