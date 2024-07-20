package http

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"server/pkg/param"
	"server/service"
)

// function check licence

func CheckLicense(ctx *fiber.Ctx) error {

	// bind request value
	clientInfo := new(param.ClientRequest)

	if err := json.Unmarshal(ctx.Body(), clientInfo); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// validate data. 1. check timestamp, 2. check random number

	validate := validator.New()

	if err := validate.Struct(clientInfo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := service.ValidateTimestamp(clientInfo.TimeStamp); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := service.CheckDuplicateRequests(clientInfo.RandomNumber); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// check if client send nil uuid

	switch clientInfo.ID {
	case "":
		uID := uuid.New()
		clientInfo.ID = uID.String()

		if err := service.AddNewClient(clientInfo); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

	default:
		// if not get client with uuid and check hash and ...
		if err := service.ValidateClientHashInfo(clientInfo.HardwareHash); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// generate auth key
		authKey, err := service.GenerateAuthKey(clientInfo.TimeStamp, clientInfo.RandomNumber)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		clientRes := new(param.ClientResponse)
		clientRes.AuthKey = authKey

	}

	return nil
}
