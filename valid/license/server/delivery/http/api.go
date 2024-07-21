package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"server/param"
	"server/service"
)

// function check licence

func CheckLicense(ctx *fiber.Ctx) error {

	// bind request value
	clientInfo := new(param.ClientRequest)

	if err := json.Unmarshal(ctx.Body(), clientInfo); err != nil {
		log.Println("22", err)
		return fiber.NewError(400, err.Error())
	}

	// validate data. 1. check timestamp, 2. check random number

	validate := validator.New()

	if err := validate.Struct(clientInfo); err != nil {
		log.Println("31", err, "\n", clientInfo)

		return fiber.NewError(400, err.Error())
	}

	if err := service.ValidateTimestamp(clientInfo.TimeStamp); err != nil {
		log.Println("37", err)

		return fiber.NewError(400, err.Error())
	}

	if err := service.CheckDuplicateRequests(clientInfo.RandomNumber); err != nil {
		log.Println("43", err)

		return fiber.NewError(400, err.Error())
	}

	// check if client send nil uuid

	switch clientInfo.ID {
	case "":

		uID := uuid.New()
		clientInfo.ID = uID.String()

		if err := service.AddNewClient(clientInfo); err != nil {
			log.Println("57", err)

			return fiber.NewError(400, err.Error())
		}

		clientRes := new(param.ClientResponse)
		clientRes.AuthKey = ""
		clientRes.ValidationStatus = true
		clientRes.ClientID = clientInfo.ID

		return ctx.Status(200).JSON(clientRes)

	default:

		// if not get client with uuid and check hash and ...
		isValid, err := service.ValidateClientHashInfo(clientInfo.ID, clientInfo.HardwareHash)
		if err != nil {

			return fiber.NewError(400, err.Error())
		}
		if !isValid {
			return fiber.NewError(400, "invalid client")
		}

		// generate auth key
		authKey, err := service.GenerateAuthKey(clientInfo.TimeStamp, clientInfo.RandomNumber)
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

// panel api

func ListClients(ctx *fiber.Ctx) error {

	return nil
}

func UpdateClient(ctx *fiber.Ctx) error {

	return nil
}

func Login(ctx *fiber.Ctx) error {

	return nil
}
