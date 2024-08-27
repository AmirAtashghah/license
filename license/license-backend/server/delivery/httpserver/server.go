package httpserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/delivery/httpserver/api"
	"server/logger"
	"server/pkg/jwt"
	"server/service/customer_product_service"
	"server/service/customer_service"
	"server/service/log_service"
	"server/service/product_service"
	"server/service/restriction_service"
	"server/service/user_service"
)

type Config struct {
	Host       string `yaml:"host" env:"HOST" `
	Port       int    `yaml:"port" env:"PORT" `
	AllowPanel bool   `yaml:"allow_panel" env:"ALLOW_PANEL" `
}

type Server struct {
	config  Config
	Handler *api.Handler
	Router  *fiber.App
}

const group = "httpserver"

func New(config Config, customerProductSvc *customer_product_service.Service,
	restrictionSvc *restriction_service.Service,
	customerSvc *customer_service.Service,
	productSvc *product_service.Service,
	userSvc *user_service.Service,
	jwtSvc *jwt.Service,
	logSvc *log_service.Service) Server {
	return Server{
		config: config,

		Handler: api.NewHandler(customerProductSvc,
			restrictionSvc,
			customerSvc,
			productSvc,
			userSvc,
			jwtSvc,
			logSvc),

		Router: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
	}
}

func (s Server) Serve() {

	// Routes
	s.Handler.SetRoutes(s.config.AllowPanel, s.Router)

	// Start server
	address := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	fmt.Printf("start server on %s\n", address)
	if err := s.Router.Listen(address); err != nil {
		logger.L().WithGroup(group).Error("router start error", "error", err.Error())
		fmt.Println("router start error", err)
	}
}
