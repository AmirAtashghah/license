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

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return ctx.SendStatus(fiber.StatusBadRequest)
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
			slog.Uint64("reqID", ctx.Context().ID())).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.customerProductSvc.ValidateTimestamp(req.TimeStamp); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("timestamp", strconv.FormatInt(req.TimeStamp, 10))).Error(err.Error())

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	cp, err := h.customerProductSvc.GetCustomerProductByID(&param.GetCustomerProductRequest{ID: req.ID})
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error(err.Error())

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if cp == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error("customerProduct not found")

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

			logger.L().With(
				slog.Uint64("reqID", ctx.Context().ID()),
				slog.String("customerProductID", req.ID)).Error(err.Error())

			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		customerProductRestrictions, err := h.restrictionSvc.GetCustomersProductRestrictionByCustomersProductID(&param.GetCustomersProductRestrictionsByCustomerProductIDRequest{CustomersProductID: cp.ID})
		if err != nil {

			logger.L().With(
				slog.Uint64("reqID", ctx.Context().ID()),
				slog.String("customerProductID", req.ID)).Error(err.Error())

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

					logger.L().With(
						slog.Uint64("reqID", ctx.Context().ID()),
						slog.String("customerProductID", req.ID),
						slog.String("RestrictionID", strconv.Itoa(int(item.RestrictionID)))).Error(err.Error())

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

			logger.L().With(
				slog.Uint64("reqID", ctx.Context().ID()),
				slog.String("customerProductID", req.ID)).Error(err.Error())

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

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error(err.Error())

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if !isValid {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error("invalid client hash info")

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// check times
	if !h.customerProductSvc.CheckTimesConditions(cp.ExpireAt, cp.FirstConfirmedAt, cp.LastConfirmedAt, req.TimeStamp) {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error("invalid times conditions")

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

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error(err.Error())

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	customerProductRestrictions, err := h.restrictionSvc.GetCustomersProductRestrictionByCustomersProductID(&param.GetCustomersProductRestrictionsByCustomerProductIDRequest{CustomersProductID: cp.ID})
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error(err.Error())

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

				logger.L().With(
					slog.Uint64("reqID", ctx.Context().ID()),
					slog.String("customerProductID", req.ID),
					slog.String("RestrictionID", strconv.Itoa(int(item.RestrictionID)))).Error(err.Error())

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

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("customerProductID", req.ID)).Error(err.Error())

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
