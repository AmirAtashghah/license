package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"server/logger"
	"server/pkg/consts"
	"server/pkg/param"
	"strings"
)

func (h Handler) CreateCustomer(ctx *fiber.Ctx) error {

	req := new(param.CreateCustomerRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

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

	exist, err := h.customerSvc.CheckExist(req.Name, req.Email)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", fmt.Sprintf("%s:%s", req.Name, req.Email))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.Name)).Error("customer already exist")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customerAlreadyExist"})
	}

	if err := h.customerSvc.AddCustomer(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.Name)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("createCustomer", []any{ctx.Locals("userID"), req.Name})

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) UpdateCustomer(ctx *fiber.Ctx) error {

	req := new(param.UpdateCustomerRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

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

	if err := h.customerSvc.UpdateCustomer(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.Name)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("updateCustomer", []any{ctx.Locals("userID"), req.Name})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) DeleteCustomer(ctx *fiber.Ctx) error {

	req := new(param.DeleteCustomerRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

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

	//   if there is customersProduct customer can not delete

	exist, err := h.customerProductSvc.CheckExistByCustomerID(req.ID)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.ID)).Error("can not delete customer, because customer have one or more product first delete customersProduct")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "canNotDeleteCustomerThatHasProduct"})
	}

	if err := h.customerSvc.DeleteCustomer(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("deleteCustomer", []any{ctx.Locals("userID"), req.ID})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) ListCustomers(ctx *fiber.Ctx) error {

	req := new(param.GetAllCustomersRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

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

	customers, err := h.customerSvc.GetAll(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if len(customers) == 0 {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error("not found any customer")

		emptyList := make([]string, 0)
		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customersNotFound"})
		return ctx.Status(fiber.StatusOK).JSON(emptyList)
	}

	_ = h.logSvc.CreateActivityLog("listCustomer", []any{ctx.Locals("userID"), len(customers)})

	return ctx.Status(fiber.StatusOK).JSON(customers)
}

func (h Handler) GetCustomer(ctx *fiber.Ctx) error {

	req := new(param.GetCustomerRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

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

	customer, err := h.customerSvc.GetCustomerByID(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("customer", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if customer == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error("customer not found")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customerNotFound"})
	}

	_ = h.logSvc.CreateActivityLog("getCustomer", []any{ctx.Locals("userID"), customer.Name})

	return ctx.Status(fiber.StatusOK).JSON(customer)

}

//
