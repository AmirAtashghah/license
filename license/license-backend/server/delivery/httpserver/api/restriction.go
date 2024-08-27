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
	"strconv"
	"strings"
)

func (h Handler) CreateRestriction(ctx *fiber.Ctx) error {

	req := new(param.CreateRestrictionRequest)
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

	exist, err := h.restrictionSvc.CheckExist(req.Key)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restrictionValue", req.Key)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restrictionValue", req.Key)).Error("restriction already exist")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "restrictionAlreadyExist"})
	}

	if err := h.restrictionSvc.AddRestriction(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restrictionValue", req.Key)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("createRestriction", []any{ctx.Locals("userID"), req.Key})

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) UpdateRestriction(ctx *fiber.Ctx) error {

	req := new(param.UpdateRestrictionRequest)
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

	if err := h.restrictionSvc.UpdateCustomer(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restriction", strconv.Itoa(int(req.ID)))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("updateRestriction", []any{ctx.Locals("userID"), req.Key})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) DeleteRestriction(ctx *fiber.Ctx) error {

	req := new(param.DeleteRestrictionRequest)
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

	if err := h.restrictionSvc.DeleteRestriction(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restriction", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("deleteRestriction", []any{ctx.Locals("userID"), req.ID})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) ListRestrictions(ctx *fiber.Ctx) error {
	//log.Println("hre")

	req := new(param.GetAllRestrictionsRequest)
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

	restrictions, err := h.restrictionSvc.GetAllRestrictions(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if len(restrictions) == 0 {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error("not found any restrictions")

		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "restrictionsNotFound"})
		emptyList := make([]string, 0)
		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customersNotFound"})
		return ctx.Status(fiber.StatusOK).JSON(emptyList)
	}

	_ = h.logSvc.CreateActivityLog("listRestriction", []any{ctx.Locals("userID"), len(restrictions)})

	//log.Println(restrictions)

	return ctx.Status(fiber.StatusOK).JSON(restrictions)
}

func (h Handler) GetRestriction(ctx *fiber.Ctx) error {

	req := new(param.GetRestrictionRequest)
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

	restriction, err := h.restrictionSvc.GetRestrictionByID(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restriction", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if restriction == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("restriction", req.ID)).Error("restriction not found")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "restrictionNotFound"})
	}

	_ = h.logSvc.CreateActivityLog("getRestriction", []any{ctx.Locals("userID"), restriction.Key})

	return ctx.Status(fiber.StatusOK).JSON(restriction)

}
