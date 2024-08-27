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

func (h Handler) CreateProduct(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.CreateProductRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	// validate req
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

	// check if exist
	exist, err := h.productSvc.CheckExist(req.Name, req.Version)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("productName", req.Name),
			slog.String("productVersion", req.Version)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("productName", req.Name),
			slog.String("productVersion", req.Version)).Error("product already exist")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "productAlreadyExist"})
	}

	// create product
	if err := h.productSvc.AddNewProduct(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("productName", req.Name),
			slog.String("productVersion", req.Version)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("createProduct", []any{ctx.Locals("userID"), req.Name})

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) UpdateProduct(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.UpdateProductRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	// validate req

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

	// update product

	if err := h.productSvc.UpdateProduct(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("updateProduct", []any{ctx.Locals("userID"), req.Name})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) DeleteProduct(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.DeleteProductRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	// validate req

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

	//   if there is customersProduct product can not delete

	exist, err := h.customerProductSvc.CheckExistByProductID(req.ID)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ID)).Error("can not delete product, because product used by customers first delete customersProduct")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "canNotDeleteProductThatUsedByCustomers"})
	}

	//res, err := h.restrictionSvc.GetRestrictionsByProductID(&param.GetRestrictionsByProductRequest{ProductID: req.ID})
	//if err != nil {
	//
	//	logger.L().With(
	//		slog.Uint64("reqID", ctx.Context().ID()),
	//		slog.String("userID", ctx.Locals("userID").(string)),
	//		slog.String("product", req.ID)).Error(err.Error())
	//
	//	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "databaseError"})
	//}
	//
	//if res != nil {
	//
	//	logger.L().With(
	//		slog.Uint64("reqID", ctx.Context().ID()),
	//		slog.String("userID", ctx.Locals("userID").(string)),
	//		slog.String("product", req.ID)).Error("can not delete product, because product have restrictions first delete restrictions")
	//
	//	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "canNotDeleteProductThatHasRestrictions"})
	//}

	// delete product

	if err := h.productSvc.DeleteProduct(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("deleteProduct", []any{ctx.Locals("userID"), req.ID})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) ListProducts(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.GetAllProductsRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	// validate req

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

	// get products
	products, err := h.productSvc.GetAll(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if len(products) == 0 {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error("not found any product")

		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "productsNotFound"})
		emptyList := make([]string, 0)
		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customersNotFound"})
		return ctx.Status(fiber.StatusOK).JSON(emptyList)
	}

	_ = h.logSvc.CreateActivityLog("listProduct", []any{ctx.Locals("userID"), len(products)})

	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (h Handler) GetProduct(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.GetProductRequest)
	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	// validate req

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

	// get product
	product, err := h.productSvc.GetProductByID(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ID)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if product == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("product", req.ID)).Error("product not found")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "productNotFound"})
	}

	_ = h.logSvc.CreateActivityLog("getProduct", []any{ctx.Locals("userID"), product.Name})

	return ctx.Status(fiber.StatusOK).JSON(product)
}
