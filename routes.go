package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/luispfcanales/inventory-oti/api"
	"github.com/luispfcanales/inventory-oti/controller"
	"github.com/luispfcanales/inventory-oti/middle"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
	"github.com/luispfcanales/inventory-oti/services"
	"github.com/luispfcanales/inventory-oti/storage/clouddeta"
)

// any service implement ports
var (
	REPOSITORY_USER ports.StorageUserService
	REPOSITORY_CPU  ports.StorageComputerService

	AUTH_SRV    ports.AuthService
	CPU_SRV     ports.ComputerService
	USER_SRV    ports.UserService
	API_DNI_SRV ports.ApiService

	STREAMING_SRV ports.StramingComputerService
)

// initialized all services
func init() {
	REPOSITORY := clouddeta.NewCloudDetaStorage("e0ytsyfs3et_F3KZDz938AnuKc62WXBdzjt1WnKrNHh8")

	REPOSITORY_USER = REPOSITORY
	REPOSITORY_CPU = REPOSITORY

	AUTH_SRV = services.NewAuth(REPOSITORY_USER)
	CPU_SRV = services.NewComputer(REPOSITORY_CPU)
	USER_SRV = services.NewUser(REPOSITORY_USER)

	API_DNI_SRV = services.NewApiDni()

	STREAMING_SRV = services.NewConnectionWSmanager()
}

// ConfigRoutes setting routes to api and controllers routes
func ConfigRoutes(app *fiber.App) {
	app.Static("/public", "public")

	RegisterRoutesController(app)
	CreateWebsockets(app)
	CreateApiRoutes(app)
}

// RegisterRoutesController set the routes
func RegisterRoutesController(app *fiber.App) {
	//create session storage
	services.CreateStoreSession()

	app.Use(middle.NoCacheMiddleware)

	//routing application
	app.Get("/", controller.Index)

	//authentication application render view
	app.Get("/login", middle.RedirectAppIsValidSession(controller.PageLogin))
	app.Post("/login", controller.LoginPost(AUTH_SRV))
	app.Get("/exit", controller.LoginExit)

	//session value
	app.Get("/session/:id", controller.SessionValue)
	//app.Get("/test", controller.PageAdminUserSystem)
	app.Get("/test", controller.PageAdminOnlineComputers)

	//pages administrator
	admin := app.Group("/admin")
	admin.Use(middle.CheckSession())
	admin.Get("", controller.PageAdmin)
	admin.Get("/computers/online", nil)
	admin.Get("/computers/registered", controller.PageAdminRegisteredComputers)
	admin.Get("/users-system", controller.PageAdminUserSystem)

	//render page not found if not exist url only use by group /admin
	admin.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./views/404.html")
	})
}

// CreateApiRoutes create new routes to /api/anyroutes
func CreateApiRoutes(app *fiber.App) {
	rest := app.Group("/api")
	rest.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//a.Get("", api.Documentation)
	//a.Post("/login", api.Login(AUTH_SRV))
	//a.GET("/user/:id", middle.CheckHeaderToken(api.GetUserInfoByID(AUTH_SRV)))

	usersApi := rest.Group("/users")
	usersApi.Get("", api.GetAllUsers(USER_SRV))
	usersApi.Post("", api.CreateUser(USER_SRV))
	usersApi.Put("", api.UpdateUser(USER_SRV))

	rest.Get("/computers", api.GetComputers(CPU_SRV))
	rest.Get("/dni/:dni<int>?", api.GetDataDni(API_DNI_SRV))
}

// Create Websockets to realtime application
func CreateWebsockets(app *fiber.App) {
	streamWS := app.Group("/stream")

	streamWS.Get("", api.GetAllConnectionStream(STREAMING_SRV))
	streamWS.Get("/:role/computer/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		role := c.Params("role")
		defer func() {
			STREAMING_SRV.Receiver(func() {
				STREAMING_SRV.RemoveConnection(id, role)
			})
		}()

		STREAMING_SRV.Receiver(func() {
			STREAMING_SRV.AddConnection(id, role, c)
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

			STREAMING_SRV.Receiver(func() {
				STREAMING_SRV.Broadcast(Command)
			})
		}
	}))
}
