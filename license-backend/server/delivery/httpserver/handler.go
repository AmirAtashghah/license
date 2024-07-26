package httpserver

import (
	"server/pkg/jwt"
	"server/service/clientservice"
	"server/service/userservice"
)

type Handler struct {
	userSvc   *userservice.Service
	clientSvc *clientservice.Service
	jwtSvc    *jwt.Service
}

func NewHandler(
	userSvc *userservice.Service,
	clientSvc *clientservice.Service,
	jwtSvc *jwt.Service,
) *Handler {
	return &Handler{
		userSvc:   userSvc,
		clientSvc: clientSvc,
		jwtSvc:    jwtSvc,
	}
}
