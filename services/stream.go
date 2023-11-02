package services

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type Connection struct {
	ID   string `json:"id,omitempty"`
	Role string `json:"role,omitempty"`
	ws   *websocket.Conn
}

type ConnectionManager struct {
	connsAdmin   map[string]*Connection
	connsDesktop map[string]*Connection
	mailBox      chan func()
}

const (
	WS_ADMIN_ROLE   string = "admin"
	WS_DESKTOP_ROLE string = "desktop"
)

// events system
const (
	EVENT_DESKTOP_NOTIFICATION uint8 = 1 + iota // "notify"
	EVENT_ADMIN_LOADINFO                        // "load-info-system"
	EVENT_DESKTOP_LOADED                        // "loaded-info"
)

func NewConnectionWSmanager() ports.StramingComputerService {
	srv := &ConnectionManager{
		connsAdmin:   make(map[string]*Connection),
		connsDesktop: make(map[string]*Connection),
		mailBox:      make(chan func()),
	}
	go srv.Start()
	return srv
}

func (cm *ConnectionManager) Start() {
	for {
		select {
		case fn := <-cm.mailBox:
			fn()
		}
	}
}
func (cm *ConnectionManager) Receiver(fn func()) {
	cm.mailBox <- fn
}

func (cm *ConnectionManager) AddConnection(id, role string, w *websocket.Conn) {
	switch role {
	case WS_ADMIN_ROLE:
		cm.connsAdmin[id] = &Connection{
			ID:   id,
			ws:   w,
			Role: role,
		}
	case WS_DESKTOP_ROLE:
		cm.connsDesktop[id] = &Connection{
			ID:   id,
			ws:   w,
			Role: role,
		}
	}
}
func (cm *ConnectionManager) RemoveConnection(id, role string) {
	switch role {
	case WS_ADMIN_ROLE:
		delete(cm.connsAdmin, id)
	case WS_DESKTOP_ROLE:
		delete(cm.connsDesktop, id)
		cm.Broadcast(&models.StreamEvent{
			ID:     id,
			Status: "offline",
			Event:  EVENT_DESKTOP_NOTIFICATION, //"notify",
			Role:   role,
		})
	}
}
func (cm *ConnectionManager) eventAdmin(msg *models.StreamEvent) {
	switch msg.Event {
	case EVENT_ADMIN_LOADINFO:
		//buf, _ := json.Marshal(msg)
		cm.connsDesktop[msg.ID].ws.WriteJSON(msg)
	}
}
func (cm *ConnectionManager) eventDesktop(msg *models.StreamEvent) {
	switch msg.Event {
	case EVENT_DESKTOP_NOTIFICATION:
		for _, c := range cm.connsAdmin {
			c.ws.WriteJSON(msg)
		}
	case EVENT_DESKTOP_LOADED:
		cm.connsAdmin[msg.EventEmisorID].ws.WriteJSON(msg)
	}
}

func (cm *ConnectionManager) Broadcast(msg *models.StreamEvent) {
	switch msg.Role {
	case WS_ADMIN_ROLE:
		cm.eventAdmin(msg)
		log.Println("command send by admin")
	case WS_DESKTOP_ROLE:
		cm.eventDesktop(msg)
		log.Println("command send by desktop")
	}
}

func (cm *ConnectionManager) SendAllAdminNotifycation() {
	for _, c := range cm.connsAdmin {
		c.ws.WriteJSON("")
	}
}

func (cm *ConnectionManager) ListAllConnections(result chan<- models.StreamEvent) {
	signal := make(chan struct{}, 2)

	go func(s chan struct{}) {
		for _, c := range cm.connsAdmin {
			result <- models.StreamEvent{
				ID:     c.ID,
				Status: "online",
				Role:   c.Role,
			}
		}
		s <- struct{}{}
	}(signal)
	go func(s chan struct{}) {
		for _, c := range cm.connsDesktop {
			result <- models.StreamEvent{
				ID:     c.ID,
				Status: "online",
				Role:   c.Role,
			}

		}
		s <- struct{}{}
	}(signal)

	<-signal
	<-signal
	close(result)
}
