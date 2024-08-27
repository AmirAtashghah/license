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

func (h Handler) GetLicenseCheckActivityLog(ctx *fiber.Ctx) error {

	req := new(param.GetLicenseCheckActivityLogsByTitleRequest)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

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

	logs, err := h.logSvc.GetLicenseCheckActivityLogByTitle(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(logs)
}

func (h Handler) ListActivityLogsByTitle(ctx *fiber.Ctx) error {

	req := new(param.GetActivityLogsByTitleRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

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

	logs, err := h.logSvc.GetActivityLogByTitle(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(logs)
}

func (h Handler) ListActivityLogs(ctx *fiber.Ctx) error {

	req := new(param.GetActivityLogsRequest)

	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

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

	logs, err := h.logSvc.GetActivityLogs(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	return ctx.Status(fiber.StatusOK).JSON(logs)
}

func (h Handler) CountRecords(ctx *fiber.Ctx) error {

	cu, pr, us := h.logSvc.CountRecord()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"customer": cu, "product": pr, "user": us})
}
