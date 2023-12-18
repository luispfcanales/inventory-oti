package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/luispfcanales/inventory-oti/api"
	"github.com/luispfcanales/inventory-oti/controller"
	"github.com/luispfcanales/inventory-oti/middle"
	"github.com/luispfcanales/inventory-oti/ports"
	"github.com/luispfcanales/inventory-oti/services"
	"github.com/luispfcanales/inventory-oti/storage/postgre"
)

// any service implement ports
var (
	REPOSITORY_USER    ports.StorageUserService
	REPOSITORY_CPU     ports.StorageComputerService
	REPOSITORY_PERSON  ports.StoragePersonService
	REPOSITORY_NETWORK ports.StorageNetworkService
	REPOSITORY_CAMPUS  ports.StorageCampusService
	REPOSITORY_ZONE    ports.StorageZoneService

	AUTH_SRV    ports.AuthService
	CPU_SRV     ports.ComputerService
	USER_SRV    ports.UserService
	PERSON_SRV  ports.PersonService
	NETWORK_SRV ports.NetworkService
	CAMPUS_SRV  ports.CampusService
	ZONE_SRV    ports.ZoneService

	API_DNI_SRV ports.ApiService

	STREAMING_SRV ports.StramingComputerService
)

// initialized all services
func init() {
	//REPOSITORY := clouddeta.NewCloudDetaStorage("e0ytsyfs3et_F3KZDz938AnuKc62WXBdzjt1WnKrNHh8")
	REPOSITORY := postgre.NewPostgreStorage()

	REPOSITORY_USER = REPOSITORY
	REPOSITORY_CPU = REPOSITORY
	REPOSITORY_PERSON = REPOSITORY
	REPOSITORY_NETWORK = REPOSITORY
	REPOSITORY_CAMPUS = REPOSITORY
	REPOSITORY_ZONE = REPOSITORY

	AUTH_SRV = services.NewAuth(REPOSITORY_USER)
	CPU_SRV = services.NewComputer(REPOSITORY_CPU)
	USER_SRV = services.NewUser(REPOSITORY_USER)
	PERSON_SRV = services.NewPerson(REPOSITORY_PERSON)
	NETWORK_SRV = services.NewNetwork(REPOSITORY_NETWORK)
	CAMPUS_SRV = services.NewCampus(REPOSITORY_CAMPUS)
	ZONE_SRV = services.NewZone(REPOSITORY_ZONE)

	API_DNI_SRV = services.NewApiDni()

	STREAMING_SRV = services.NewConnectionWSmanager()
}

// ConfigRoutes setting routes to api and controllers routes
func ConfigRoutes(app *fiber.App) {
	app.Static("/public", "public")

	//RegisterRoutesController(app)
	CreateWebsockets(app)
	CreateApiRoutes(app)
}

// CreateApiRoutes create new routes to /api/anyroutes
func CreateApiRoutes(app *fiber.App) {
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowCredentials: true,
		ExposeHeaders:    "Authorization",
		MaxAge:           3600,
	}))
	app.Post("/login", api.Login(AUTH_SRV))
	rest := app.Group("/api")
	//a.Get("", api.Documentation)
	//rest.Use(middle.CheckToken)

	usersApi := rest.Group("/users", middle.CheckToken, cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowCredentials: true,
		ExposeHeaders:    "Authorization",
		MaxAge:           3600,
	}))
	usersApi.Get("/all", api.GetAllUsers(USER_SRV))
	usersApi.Post("", api.CreateUser(USER_SRV))
	usersApi.Put("", api.UpdateUser(USER_SRV))

	personApi := rest.Group("/person")
	personApi.Get("/all", api.HdlGetAllPersons(PERSON_SRV))
	personApi.Get("/:dni<int>?", api.HdlGetPerson(PERSON_SRV))
	personApi.Post("", api.HdlPostPerson(PERSON_SRV))
	personApi.Put("", api.HdlPutPerson(PERSON_SRV))
	personApi.Delete("/:dni<int>?", api.HdlDeletePerson(PERSON_SRV))

	networkApi := rest.Group("/network")
	networkApi.Get("/all", api.HdlGetAllNetworks(NETWORK_SRV))
	networkApi.Get("/all/resume", api.HdlGetResumeNetworks(NETWORK_SRV))

	campusApi := rest.Group("/campus")
	campusApi.Get("/all", api.HdlGetAllCampus(CAMPUS_SRV))
	campusApi.Get("/:id", api.HdlGetCampus(CAMPUS_SRV))
	campusApi.Put("", api.HdlPutCampus(CAMPUS_SRV))
	campusApi.Post("", api.HdlPostCampus(CAMPUS_SRV))
	campusApi.Delete("/:id", api.HdlDeleteCampus(CAMPUS_SRV))

	zoneApi := rest.Group("/zone")
	zoneApi.Get("/all", api.HdlGetAllZone(ZONE_SRV))
	zoneApi.Get("/:id", api.HdlGetZone(ZONE_SRV))

	rest.Get("/computers", api.GetComputers(CPU_SRV))

	rest.Get("/dni/:dni<int>?", api.GetDataDni(API_DNI_SRV))
}

// Create Websockets to realtime application
func CreateWebsockets(app *fiber.App) {
	streamWS := app.Group("/stream")

	streamWS.Get("", api.GetAllConnectionStream(STREAMING_SRV))
	streamWS.Get("/:role/computer/:id", websocket.New(controller.HandleStreamSocket(STREAMING_SRV)))
}
