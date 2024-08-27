package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"server/logger"
	"server/pkg/param"
	"strconv"
)

// function check licence

func (h Handler) CheckLicense(ctx *fiber.Ctx) error {

	// todo encryption
	//encryptReq := new(param.EncryptedCheckLicenseRequest)
	//if err := json.Unmarshal(ctx.Body(), encryptReq); err != nil {
	//
	//	return ctx.SendStatus(fiber.StatusBadRequest)
	//}
	//
	//rawBody := encrypt.Decrypt(encryptReq.Body)

	req := new(param.CheckLicenseRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.customerProductSvc.ValidateTimestamp(req.TimeStamp); err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	cp, err := h.customerProductSvc.GetCustomerProductByID(&param.GetCustomerProductRequest{ID: req.ID})
	if err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if cp == nil {
		//logger.L().Error("cp nil")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// first confirm
	if cp.HardwareHash == "hash" || cp.FirstConfirmedAt == -1 || cp.LastConfirmedAt == -1 {

		updateCP := param.UpdateCustomerProductRequest{
			ID:               cp.ID,
			CustomerID:       cp.CustomerID,
			ProductID:        cp.ProductID,
			HardwareHash:     req.HardwareHash,
			LicenseType:      cp.LicenseType,
			IsActive:         &cp.IsActive,
			ExpireAt:         cp.ExpireAt,
			FirstConfirmedAt: req.TimeStamp,
			LastConfirmedAt:  req.TimeStamp,
			CreatedAt:        cp.CreatedAt,
		}

		if err := h.customerProductSvc.UpdateCustomerProduct(&updateCP); err != nil {
			//logger.L().Error(err.Error())
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		customerProductRestrictions, err := h.restrictionSvc.GetCustomersProductRestrictionByCustomersProductID(&param.GetCustomersProductRestrictionsByCustomerProductIDRequest{CustomersProductID: cp.ID})
		if err != nil {
			//logger.L().Error(err.Error())
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		type EncryptRestrictionBody struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		var ERB []EncryptRestrictionBody
		var encryptedRestriction string = ""

		if len(customerProductRestrictions) != 0 {

			for _, item := range customerProductRestrictions {
				restriction, err := h.restrictionSvc.GetRestrictionByID(&param.GetRestrictionRequest{ID: strconv.Itoa(int(item.RestrictionID))})
				if err != nil {
					//logger.L().Error(err.Error())
					return ctx.SendStatus(fiber.StatusBadRequest)
				}

				ERB = append(ERB, EncryptRestrictionBody{Key: restriction.Key, Value: item.Value})
			}
			erb, _ := json.Marshal(ERB)
			encryptedRestriction = string(erb)

			// todo encryption
			//encryptedRestriction, err = h.restrictionSvc.EncryptRestriction(ERB)
			//if err != nil {
			//	logger.L().Error(err.Error())
			//	return ctx.SendStatus(fiber.StatusBadRequest)
			//}
		}

		authKey, err := h.customerProductSvc.GenerateAuthKey(req.TimeStamp, req.RandomNumber)
		if err != nil {
			//logger.L().Error(err.Error())
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		checkLicenseResponse := param.CheckLicenseResponse{
			AuthKey:          authKey,
			Restriction:      encryptedRestriction,
			ValidationStatus: true,
		}

		// todo encryption
		//encryptResponse, err := h.customerProductSvc.EncryptResponse(checkLicenseResponse)
		//if err != nil {
		//
		//	return ctx.SendStatus(fiber.StatusBadRequest)
		//}

		return ctx.Status(fiber.StatusOK).JSON(checkLicenseResponse)
	}

	// after first confirm

	isValid, err := h.customerProductSvc.ValidateClientHashInfo(req.ID, req.HardwareHash)
	if err != nil {
		logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if !isValid {
		//logger.L().Error("not valid ")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// check times
	if !h.customerProductSvc.CheckTimesConditions(cp.ExpireAt, cp.FirstConfirmedAt, cp.LastConfirmedAt, req.TimeStamp) {
		//logger.L().Error("time stamps not valid")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	updateCP := param.UpdateCustomerProductRequest{
		ID:               cp.ID,
		CustomerID:       cp.CustomerID,
		ProductID:        cp.ProductID,
		HardwareHash:     cp.HardwareHash,
		LicenseType:      cp.LicenseType,
		IsActive:         &cp.IsActive,
		ExpireAt:         cp.ExpireAt,
		FirstConfirmedAt: cp.FirstConfirmedAt,
		LastConfirmedAt:  req.TimeStamp,
		CreatedAt:        cp.CreatedAt,
	}

	if err := h.customerProductSvc.UpdateCustomerProduct(&updateCP); err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	customerProductRestrictions, err := h.restrictionSvc.GetCustomersProductRestrictionByCustomersProductID(&param.GetCustomersProductRestrictionsByCustomerProductIDRequest{CustomersProductID: cp.ID})
	if err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	type EncryptRestrictionBody struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var ERB []EncryptRestrictionBody
	var encryptedRestriction string = ""

	if len(customerProductRestrictions) != 0 {

		for _, item := range customerProductRestrictions {
			restriction, err := h.restrictionSvc.GetRestrictionByID(&param.GetRestrictionRequest{ID: strconv.Itoa(int(item.RestrictionID))})
			if err != nil {
				//logger.L().Error(err.Error())
				return ctx.SendStatus(fiber.StatusBadRequest)
			}

			ERB = append(ERB, EncryptRestrictionBody{Key: restriction.Key, Value: item.Value})
		}
		erb, _ := json.Marshal(ERB)
		encryptedRestriction = string(erb)
		// todo encryption
		//encryptedRestriction, err = h.restrictionSvc.EncryptRestriction(ERB)
		//if err != nil {
		//	logger.L().Error(err.Error())
		//	return ctx.SendStatus(fiber.StatusBadRequest)
		//}
	}

	authKey, err := h.customerProductSvc.GenerateAuthKey(req.TimeStamp, req.RandomNumber)
	if err != nil {
		//logger.L().Error(err.Error())
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	checkLicenseResponse := param.CheckLicenseResponse{
		AuthKey:          authKey,
		Restriction:      encryptedRestriction,
		ValidationStatus: true,
	}
	// todo encryption
	//encryptResponse, err := h.customerProductSvc.EncryptResponse(checkLicenseResponse)
	//if err != nil {
	//
	//	return ctx.SendStatus(fiber.StatusBadRequest)
	//}

	return ctx.Status(fiber.StatusOK).JSON(checkLicenseResponse)
}
