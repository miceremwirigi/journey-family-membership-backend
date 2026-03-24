package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/apis"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/common/middlewares"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/config"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/database"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/repository"
	"github.com/miceremwirigi/journey-family-membership-backend/pkg/services"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	app := fiber.New()

	// In your main.go, add before routes:
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Repositories
	memberRepo := repository.NewMemberRepository(db)
	familyRepo := repository.NewFamilyRepository(db)
	smallGroupRepo := repository.NewSmallGroupRepository(db)
	visitorRepo := repository.NewVisitorRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	eventRepo := repository.NewEventRepository(db)

	// Services
	memberService := services.NewMemberService(memberRepo)
	familyService := services.NewFamilyService(familyRepo)
	smallGroupService := services.NewSmallGroupService(smallGroupRepo)
	visitorService := services.NewVisitorService(visitorRepo)
	messageService := services.NewMessageService(messageRepo)
	eventService := services.NewEventService(eventRepo)

	// Handlers
	memberHandler := apis.NewMemberHandler(memberService)
	familyHandler := apis.NewFamilyHandler(familyService)
	smallGroupHandler := apis.NewSmallGroupHandler(smallGroupService)
	visitorHandler := apis.NewVisitorHandler(visitorService)
	messageHandler := apis.NewMessageHandler(messageService)
	eventHandler := apis.NewEventHandler(eventService)

	// API routes
	api := app.Group("/api")

	// Auth
	api.Post("/login", func(c *fiber.Ctx) error {
		return apis.Login(c, memberRepo, cfg)
	})

	// Members
	membersAPI := api.Group("/members", middlewares.AuthMiddleware(cfg))
	membersAPI.Get("/", middlewares.RoleAuthMiddleware("admin", "hr"), memberHandler.GetMembers)
	membersAPI.Get("/:id", middlewares.RoleAuthMiddleware("admin", "hr"), memberHandler.GetMember)
	membersAPI.Post("/", middlewares.RoleAuthMiddleware("admin", "hr"), memberHandler.CreateMember)
	membersAPI.Put("/:id", middlewares.RoleAuthMiddleware("admin", "hr"), memberHandler.UpdateMember)
	membersAPI.Delete("/:id", middlewares.RoleAuthMiddleware("admin"), memberHandler.DeleteMember)

	// Families
	familiesAPI := api.Group("/families", middlewares.AuthMiddleware(cfg))
	familiesAPI.Get("/", familyHandler.GetFamilies)
	familiesAPI.Get("/:id", familyHandler.GetFamily)
	familiesAPI.Post("/", middlewares.RoleAuthMiddleware("admin", "staff"), familyHandler.CreateFamily)
	familiesAPI.Put("/:id", middlewares.RoleAuthMiddleware("admin", "staff"), familyHandler.UpdateFamily)
	familiesAPI.Delete("/:id", middlewares.RoleAuthMiddleware("admin"), familyHandler.DeleteFamily)

	// Small Groups
	smallGroupsAPI := api.Group("/small-groups", middlewares.AuthMiddleware(cfg))
	smallGroupsAPI.Get("/", smallGroupHandler.GetSmallGroups)
	smallGroupsAPI.Get("/:id", smallGroupHandler.GetSmallGroup)
	smallGroupsAPI.Post("/", middlewares.RoleAuthMiddleware("admin", "staff"), smallGroupHandler.CreateSmallGroup)
	smallGroupsAPI.Put("/:id", middlewares.RoleAuthMiddleware("admin", "staff"), smallGroupHandler.UpdateSmallGroup)
	smallGroupsAPI.Delete("/:id", middlewares.RoleAuthMiddleware("admin"), smallGroupHandler.DeleteSmallGroup)

	// Visitors
	visitorsAPI := api.Group("/visitors", middlewares.AuthMiddleware(cfg))
	visitorsAPI.Get("/", middlewares.RoleAuthMiddleware("admin", "staff"), visitorHandler.GetVisitors)
	visitorsAPI.Get("/:id", middlewares.RoleAuthMiddleware("admin", "staff"), visitorHandler.GetVisitor)
	visitorsAPI.Post("/", middlewares.RoleAuthMiddleware("admin", "staff"), visitorHandler.CreateVisitor)
	visitorsAPI.Put("/:id", middlewares.RoleAuthMiddleware("admin", "staff"), visitorHandler.UpdateVisitor)
	visitorsAPI.Delete("/:id", middlewares.RoleAuthMiddleware("admin"), visitorHandler.DeleteVisitor)

	// Messages
	messagesAPI := api.Group("/messages", middlewares.AuthMiddleware(cfg))
	messagesAPI.Get("/", middlewares.RoleAuthMiddleware("admin", "staff", "hr", "finance"), messageHandler.GetMessages)
	messagesAPI.Get("/:id", middlewares.RoleAuthMiddleware("admin", "staff", "hr", "finance"), messageHandler.GetMessage)
	messagesAPI.Post("/", middlewares.RoleAuthMiddleware("admin", "staff"), messageHandler.CreateMessage)
	messagesAPI.Put("/:id", middlewares.RoleAuthMiddleware("admin", "staff"), messageHandler.UpdateMessage)
	messagesAPI.Delete("/:id", middlewares.RoleAuthMiddleware("admin"), messageHandler.DeleteMessage)

	// Events
	eventsAPI := api.Group("/events", middlewares.AuthMiddleware(cfg))
	eventsAPI.Get("/", eventHandler.GetEvents)
	eventsAPI.Get("/:id", eventHandler.GetEvent)
	eventsAPI.Post("/", middlewares.RoleAuthMiddleware("admin", "staff"), eventHandler.CreateEvent)
	eventsAPI.Put("/:id", middlewares.RoleAuthMiddleware("admin", "staff"), eventHandler.UpdateEvent)
	eventsAPI.Delete("/:id", middlewares.RoleAuthMiddleware("admin"), eventHandler.DeleteEvent)

	log.Fatal(app.Listen(":3002"))
}
