package httpserver

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log/slog"
	"server/logger"
	"server/pkg/param"
)

const group = "httpserver"
const checkLicenseGroup = "checklicensegroup"

// function check licence

func (h Handler) CheckLicense(ctx *fiber.Ctx) error {

	clientHashParam := ctx.Params("hash", "none")

	customLog := logger.SaveClientLogsOnDatabase()

	// bind request value
	clientInfo := new(param.ClientRequest)

	if err := json.Unmarshal(ctx.Body(), clientInfo); err != nil {
		customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
			slog.String("request_body", string(ctx.Body())),
			slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	// validate data. 1. check timestamp, 2. check random number

	validate := validator.New()

	if err := validate.Struct(clientInfo); err != nil {
		customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
			slog.String("request_body", string(ctx.Body())),
			slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := h.clientSvc.ValidateTimestamp(clientInfo.TimeStamp); err != nil {
		customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
			slog.String("request_body", string(ctx.Body())),
			slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := h.clientSvc.CheckDuplicateRequests(clientInfo.RandomNumber); err != nil {
		customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
			slog.String("request_body", string(ctx.Body())),
			slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	// check if postgresql client send nil uuid

	switch clientInfo.ID {
	case "":

		uID := uuid.New()
		clientInfo.ID = uID.String()

		if err := h.clientSvc.AddNewClient(clientInfo); err != nil {
			customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
				slog.String("request_body", string(ctx.Body())),
				slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

			return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
		}

		clientRes := new(param.ClientResponse)
		clientRes.AuthKey = ""
		clientRes.ValidationStatus = true
		clientRes.ClientID = clientInfo.ID

		//customLog.WithGroup(checkLicenseGroup).With(slog.String("request_body", string(ctx.Body())),
		//	slog.Int("code", 200)).Info("ok")

		return ctx.Status(200).JSON(clientRes)

	default:

		// if not get postgresqlclient with uuid and check hash and ...

		isValid, err := h.clientSvc.ValidateClientHashInfo(clientInfo.ID, clientInfo.HardwareHash)
		if err != nil {
			customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
				slog.String("request_body", string(ctx.Body())),
				slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

			return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
		}

		if !isValid {
			customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
				slog.String("request_body", string(ctx.Body())),
				slog.Int("code", fiber.ErrBadRequest.Code)).Error("invalid postgresql client")

			return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "invalid postgresql client"})
		}

		// generate auth key
		authKey, err := h.clientSvc.GenerateAuthKey(clientInfo.TimeStamp, clientInfo.RandomNumber)
		if err != nil {
			customLog.WithGroup(checkLicenseGroup).With(slog.String("client_hash", clientHashParam),
				slog.String("request_body", string(ctx.Body())),
				slog.Int("code", fiber.ErrBadRequest.Code)).Error(err.Error())

			return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
		}

		clientRes := new(param.ClientResponse)
		clientRes.AuthKey = authKey
		clientRes.ValidationStatus = true

		//customLog.WithGroup(checkLicenseGroup).With(slog.String("request_body", string(ctx.Body())),
		//	slog.Int("code", 200)).Info("ok")

		return ctx.Status(200).JSON(clientRes)
	}
}

// todo change logs args slog.string.....
// panel httpserver

func (h Handler) ListClients(ctx *fiber.Ctx) error {

	req := new(param.ClientFilter)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	clients, err := h.clientSvc.ListClients(*req)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(clients)
}

func (h Handler) UpdateClient(ctx *fiber.Ctx) error {

	req := new(param.UpdateClientRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := h.clientSvc.UpdateClient(*req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "updated successfully"})
}

func (h Handler) UpdateActivateStatus(ctx *fiber.Ctx) error {

	req := new(param.ChangeActivateRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := h.clientSvc.ChangeActiveStatus(*req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "updated activate status successfully"})
}

func (h Handler) DeleteClient(ctx *fiber.Ctx) error {

	req := new(param.DeleteClientRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := h.clientSvc.DeleteClient(req.ID); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "deleted successfully"})
}

func (h Handler) GetClient(ctx *fiber.Ctx) error {

	req := new(param.GetClientRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	client, err := h.clientSvc.GetClient(req.ID)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(client)
}

func (h Handler) ListLogs(ctx *fiber.Ctx) error {

	req := new(param.GetLogsClientRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	logs, err := h.clientSvc.GetClientLogs(req.Hash)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(logs)
}

// todo implement if needed
// func (h Handler) CreateClient(ctx *fiber.Ctx) error { return nil }
