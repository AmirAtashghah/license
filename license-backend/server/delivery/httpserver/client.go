package httpserver

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"server/pkg/param"
)

// function check licence

func (h Handler) CheckLicense(ctx *fiber.Ctx) error {

	// bind request value
	clientInfo := new(param.ClientRequest)

	if err := json.Unmarshal(ctx.Body(), clientInfo); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	// validate data. 1. check timestamp, 2. check random number

	validate := validator.New()

	if err := validate.Struct(clientInfo); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if err := h.clientSvc.ValidateTimestamp(clientInfo.TimeStamp); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if err := h.clientSvc.CheckDuplicateRequests(clientInfo.RandomNumber); err != nil {
		return fiber.NewError(400, err.Error())
	}

	// check if postgresqlclient send nil uuid

	switch clientInfo.ID {
	case "":

		uID := uuid.New()
		clientInfo.ID = uID.String()

		if err := h.clientSvc.AddNewClient(clientInfo); err != nil {
			return fiber.NewError(400, err.Error())
		}

		clientRes := new(param.ClientResponse)
		clientRes.AuthKey = ""
		clientRes.ValidationStatus = true
		clientRes.ClientID = clientInfo.ID

		return ctx.Status(200).JSON(clientRes)

	default:

		// if not get postgresqlclient with uuid and check hash and ...

		isValid, err := h.clientSvc.ValidateClientHashInfo(clientInfo.ID, clientInfo.HardwareHash)
		if err != nil {

			return fiber.NewError(400, err.Error())
		}
		if !isValid {
			return fiber.NewError(400, "invalid postgresqlclient")
		}

		// generate auth key
		authKey, err := h.clientSvc.GenerateAuthKey(clientInfo.TimeStamp, clientInfo.RandomNumber)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		clientRes := new(param.ClientResponse)
		clientRes.AuthKey = authKey
		clientRes.ValidationStatus = true

		return ctx.Status(200).JSON(clientRes)
	}
}

// panel httpserver

func (h Handler) ListClients(ctx *fiber.Ctx) error {

	req := new(param.ClientFilter)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		return fiber.NewError(400, err.Error())
	}

	clients, err := h.clientSvc.ListClients(*req)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return ctx.Status(200).JSON(clients)
}

func (h Handler) UpdateClient(ctx *fiber.Ctx) error {

	req := new(param.UpdateClientRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if err := h.clientSvc.UpdateClient(*req); err != nil {
		return fiber.NewError(400, err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "updated successfully"})
}

func (h Handler) UpdateActivateStatus(ctx *fiber.Ctx) error {

	req := new(param.ChangeActivateRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if err := h.clientSvc.ChangeActiveStatus(*req); err != nil {
		return fiber.NewError(400, err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "updated activate status successfully"})
}

func (h Handler) DeleteClient(ctx *fiber.Ctx) error {

	req := new(param.DeleteClientRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if err := h.clientSvc.DeleteClient(req.ID); err != nil {
		return fiber.NewError(400, err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "deleted successfully"})
}

// todo implement if needed
// func (h Handler) CreateClient(ctx *fiber.Ctx) error { return nil }
