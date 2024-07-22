package httpserver

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"server/pkg/param"
)

// function check licence

func (h Handler) CheckLicense(ctx *fiber.Ctx) error {

	// bind request value
	clientInfo := new(param.ClientRequest)

	if err := json.Unmarshal(ctx.Body(), clientInfo); err != nil {
		log.Println("22", err)
		return fiber.NewError(400, err.Error())
	}

	// validate data. 1. check timestamp, 2. check random number

	// todo write function in service for this
	validate := validator.New()

	if err := validate.Struct(clientInfo); err != nil {
		return fiber.NewError(400, err.Error())
	}
	// end todo

	if err := h.clientSvc.ValidateTimestamp(clientInfo.TimeStamp); err != nil {
		log.Println("37", err)

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
			log.Println("57", err)

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
			log.Println("87", err)

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

	return nil
}

func (h Handler) UpdateClient(ctx *fiber.Ctx) error {

	return nil
}

func (h Handler) DeleteClient(ctx *fiber.Ctx) error { return nil }

func (h Handler) CreateClient(ctx *fiber.Ctx) error { return nil }
