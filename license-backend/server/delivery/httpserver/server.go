package httpserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/logger"
	"server/pkg/jwt"
	"server/service/clientservice"
	"server/service/userservice"
)

type Server struct {
	//config  config.Config
	Handler *Handler
	Router  *fiber.App
}

// todo add confing
func New(userSvc *userservice.Service,
	clientSvc *clientservice.Service,
	jwtSvc *jwt.Service) Server {
	return Server{
		Router: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
		//config:  config,
		Handler: NewHandler(userSvc, clientSvc, jwtSvc),
	}
}

func (s Server) Serve() {

	// Routes
	s.Handler.SetRoutes(s.Router)

	// Start server
	address := fmt.Sprintf(":%d", 3000) // s.config.HTTPServer.Port
	fmt.Printf("start server on %s\n", address)
	if err := s.Router.Listen(address); err != nil {
		logger.L().WithGroup(group).Error("router start error", "error", err.Error())
		fmt.Println("router start error", err)
	}
}
