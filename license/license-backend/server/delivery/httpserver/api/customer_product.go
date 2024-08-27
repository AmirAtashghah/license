package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"server/entity"
	"server/logger"
	"server/pkg/consts"
	"server/pkg/param"
	"strings"
)

func (h Handler) CreateCustomerProduct(ctx *fiber.Ctx) error {

	req := new(param.CreateCustomerProductRequest)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	//log.Println(req)

	validate := validator.New()
	if err := validate.Struct(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
		}

		var VErrors []param.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			VErrors = append(VErrors, param.ValidationErr{Field: err.Field(), Message: consts.ValidationErrors[err.Field()]})
		}

		var errStrings []string
		for _, er := range VErrors {
			errStrings = append(errStrings, fmt.Sprintf("Feild: %s, Message: %s", er.Field, er.Message))
		}

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": VErrors})
	}

	// check customer and product exits // done

	customer, err := h.customerSvc.GetCustomerByID(&param.GetCustomerRequest{ID: req.CustomerID})
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.CustomerID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if customer == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("customer not exist : %s", req.CustomerID))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customerNotFound"})
	}

	product, err := h.productSvc.GetProductByID(&param.GetProductRequest{ID: req.ProductID})
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ProductID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if product == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("product not exist : %s", req.ProductID))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "productNotFound"})
	}

	exist, err := h.customerProductSvc.CheckExist(req.ProductID, req.CustomerID)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", fmt.Sprintf("%s:%s", req.CustomerID, req.ProductID))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ProductID)).Error(fmt.Sprintf("customerProduct already exist : %s--%s", req.ProductID, req.ProductID))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customerProductAlreadyExist"})
	}

	customersProductID, err := h.customerProductSvc.AddCustomerProduct(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", fmt.Sprintf("%s:%s", req.CustomerID, req.ProductID))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if req.Restrictions != "" {
		err := h.restrictionSvc.AddCustomersProductRestriction(&param.CreateCustomersProductRestrictionRequest{CustomersProductID: customersProductID, RestrictionIDAndValues: req.Restrictions})
		if err != nil {
			logger.L().With(
				slog.Uint64("reqID", ctx.Context().ID()),
				slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
		}
	}
	_ = h.logSvc.CreateActivityLog("createCustomerProduct", []any{ctx.Locals("userID"), req.CustomerID, req.ProductID})

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) UpdateCustomerProduct(ctx *fiber.Ctx) error {

	req := new(param.UpdateCustomerProductRequest)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
		}

		var VErrors []param.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			VErrors = append(VErrors, param.ValidationErr{Field: err.Field(), Message: consts.ValidationErrors[err.Field()]})
		}

		var errStrings []string
		for _, er := range VErrors {
			errStrings = append(errStrings, fmt.Sprintf("Feild: %s, Message: %s", er.Field, er.Message))
		}

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": VErrors})
	}

	if err := h.customerProductSvc.UpdateCustomerProduct(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if req.Restrictions != "" {
		err := h.restrictionSvc.AddCustomersProductRestriction(&param.CreateCustomersProductRestrictionRequest{CustomersProductID: req.ID, RestrictionIDAndValues: req.Restrictions})
		if err != nil {
			logger.L().With(
				slog.Uint64("reqID", ctx.Context().ID()),
				slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
		}
	}

	_ = h.logSvc.CreateActivityLog("updateCustomerProduct", []any{ctx.Locals("userID"), req.CustomerID, req.ProductID})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) DeleteCustomerProduct(ctx *fiber.Ctx) error {

	req := new(param.DeleteCustomerProductRequest)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
		}

		var VErrors []param.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			VErrors = append(VErrors, param.ValidationErr{Field: err.Field(), Message: consts.ValidationErrors[err.Field()]})
		}

		var errStrings []string
		for _, er := range VErrors {
			errStrings = append(errStrings, fmt.Sprintf("Feild: %s, Message: %s", er.Field, er.Message))
		}

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": VErrors})
	}

	err := h.restrictionSvc.DeleteCustomersProductRestriction(&param.DeleteCustomersProductRestrictionRequest{CustomersProductID: req.ID})
	if err != nil {
		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if err := h.customerProductSvc.DeleteCustomerProduct(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("deleteCustomerProduct", []any{ctx.Locals("userID"), req.ID})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) ListCustomerProducts(ctx *fiber.Ctx) error {

	req := new(param.GetAllCustomerProductsRequest)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
		}

		var VErrors []param.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			VErrors = append(VErrors, param.ValidationErr{Field: err.Field(), Message: consts.ValidationErrors[err.Field()]})
		}

		var errStrings []string
		for _, er := range VErrors {
			errStrings = append(errStrings, fmt.Sprintf("Feild: %s, Message: %s", er.Field, er.Message))
		}

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": VErrors})
	}

	customerProducts, err := h.customerProductSvc.GetAll(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON("databaseError")
	}

	if len(customerProducts) == 0 {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error("not found any customerProducts")

		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customersProductsNotFound"})
		emptyList := make([]string, 0)
		return ctx.Status(fiber.StatusOK).JSON(emptyList)
	}

	_ = h.logSvc.CreateActivityLog("listCustomerProduct", []any{ctx.Locals("userID"), len(customerProducts)})

	return ctx.Status(fiber.StatusOK).JSON(customerProducts)
}

func (h Handler) GetCustomerProduct(ctx *fiber.Ctx) error {

	req := new(param.GetCustomerProductRequest)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
		}

		var VErrors []param.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			VErrors = append(VErrors, param.ValidationErr{Field: err.Field(), Message: consts.ValidationErrors[err.Field()]})
		}

		var errStrings []string
		for _, er := range VErrors {
			errStrings = append(errStrings, fmt.Sprintf("Feild: %s, Message: %s", er.Field, er.Message))
		}

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": VErrors})
	}

	customerProduct, err := h.customerProductSvc.GetCustomerProductByID(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if customerProduct == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", req.ID)).Error(fmt.Sprintf("not found any customerProducts : %s", req.ID))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customersProductNotFound"})
	}

	customersProductRestriction, err := h.restrictionSvc.GetCustomersProductRestrictionByCustomersProductID(&param.GetCustomersProductRestrictionsByCustomerProductIDRequest{CustomersProductID: customerProduct.ID})
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customerProduct", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("getCustomerProduct", []any{ctx.Locals("userID"), customerProduct.CustomerID, customerProduct.ProductID})

	result := struct {
		ID               string                                `json:"id"`
		CustomerID       string                                `json:"customer_id"`
		ProductID        string                                `json:"product_id"`
		Restrictions     []*entity.CustomersProductRestriction `json:"restrictions"`
		HardwareHash     string                                `json:"hardware_hash"`
		LicenseType      string                                `json:"license_type"`
		IsActive         bool                                  `json:"is_active"`
		ExpireAt         int64                                 `json:"expire_at"`
		FirstConfirmedAt int64                                 `json:"first_confirmed_at"`
		LastConfirmedAt  int64                                 `json:"last_confirmed_at"`
		CreatedAt        int64                                 `json:"created_at"`
		UpdatedAt        int64                                 `json:"updated_at"`
	}{
		ID:               customerProduct.ID,
		CustomerID:       customerProduct.CustomerID,
		ProductID:        customerProduct.ProductID,
		Restrictions:     customersProductRestriction,
		HardwareHash:     customerProduct.HardwareHash,
		LicenseType:      customerProduct.LicenseType,
		IsActive:         customerProduct.IsActive,
		ExpireAt:         customerProduct.ExpireAt,
		FirstConfirmedAt: customerProduct.FirstConfirmedAt,
		LastConfirmedAt:  customerProduct.LastConfirmedAt,
		CreatedAt:        customerProduct.CreatedAt,
		UpdatedAt:        customerProduct.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
