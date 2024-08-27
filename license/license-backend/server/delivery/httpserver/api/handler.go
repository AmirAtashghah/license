package api

import (
	"server/pkg/jwt"
	"server/service/customer_product_service"
	"server/service/customer_service"
	"server/service/log_service"
	"server/service/product_service"
	"server/service/restriction_service"
	"server/service/user_service"
)

type Handler struct {
	customerProductSvc *customer_product_service.Service
	restrictionSvc     *restriction_service.Service
	customerSvc        *customer_service.Service
	productSvc         *product_service.Service
	userSvc            *user_service.Service
	jwtSvc             *jwt.Service
	logSvc             *log_service.Service
}

func NewHandler(
	customerProductSvc *customer_product_service.Service,
	restrictionSvc *restriction_service.Service,
	customerSvc *customer_service.Service,
	productSvc *product_service.Service,
	userSvc *user_service.Service,
	jwtSvc *jwt.Service,
	logSvc *log_service.Service,

) *Handler {
	return &Handler{
		customerProductSvc: customerProductSvc,
		restrictionSvc:     restrictionSvc,
		customerSvc:        customerSvc,
		productSvc:         productSvc,
		userSvc:            userSvc,
		jwtSvc:             jwtSvc,
		logSvc:             logSvc,
	}
}
