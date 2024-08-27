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
	"server/pkg/hash"
	"server/pkg/param"
	"strconv"
	"strings"
	"time"
)

func (h Handler) Login(ctx *fiber.Ctx) error {

	//log.Println("start", time.Now().UnixNano())

	lr := new(param.LoginRequest)
	if err := json.Unmarshal(ctx.Body(), lr); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("data", string(ctx.Body()))).Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalidRequestBody"})
	}

	//log.Println("before validation", time.Now().UnixNano())

	validate := validator.New()
	if err := validate.Struct(lr); err != nil {

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
			slog.Uint64("reqID", ctx.Context().ID())).Error(fmt.Sprintf("validation Errors: [%s]", strings.Join(errStrings, "; ")))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": VErrors})
	}
	//log.Println("before get user by name ", time.Now().UnixNano())

	user, err := h.userSvc.GetUserByUsername(&param.GetUserRequest{Username: lr.Username})
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("user", lr.Username)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if user == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("user", lr.Username)).Error("user not found")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userNotFound"})
	}
	//log.Println("before check hash password", time.Now().UnixNano())

	if !hash.CheckHash(lr.Password, user.Password) {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("user", lr.Username)).Error("password is wrong")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "passwordIsWrong"})
	}
	//log.Println("before generate token", time.Now().UnixNano())

	token, err := h.jwtSvc.GenerateJWTAccessToken(user.Username, user.Role)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("user", strconv.Itoa(user.ID))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}
	//log.Println("before set cookie", time.Now().UnixNano())

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(150000000 * time.Minute)
	//log.Println(time.Now().Unix())

	ctx.Cookie(cookie)
	//log.Println("before create activity log", time.Now().UnixNano())

	_ = h.logSvc.CreateActivityLog("login", []any{user.Name, user.Username})
	//log.Println("before send status ok", time.Now().UnixNano())

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Logout(ctx *fiber.Ctx) error {

	ctx.Cookie(&fiber.Cookie{
		Name: "token",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 200)),
		HTTPOnly: true,
		SameSite: "lax",
	})

	_ = h.logSvc.CreateActivityLog("logout", []any{ctx.Locals("userID")})

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) CreateUser(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.CreateUserRequest)
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
	exist, err := h.userSvc.CheckExist(req.Username)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", req.Username)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if exist {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", req.Username)).Error("user already exist")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userAlreadyExist"})
	}

	// create user
	if err := h.userSvc.AddUser(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", req.Username)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("createUser", []any{req.Name, req.Username, ctx.Locals("userID")})

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) UpdateUser(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.UpdateUserRequest)
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

	// update user

	if err := h.userSvc.UpdateUser(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", req.Username)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("updateUser", []any{req.Name, req.Username, ctx.Locals("userID")})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) DeleteUser(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.DeleteUserRequest)
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

	// delete user

	if err := h.userSvc.DeleteUser(req); err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", strconv.Itoa(req.ID))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	_ = h.logSvc.CreateActivityLog("deleteUser", []any{ctx.Locals("userID"), req.ID})

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h Handler) ListUsers(ctx *fiber.Ctx) error {

	// get users
	users, err := h.userSvc.GetAll()
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if len(users) == 0 {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string))).Error("not found any user")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "usersNotFound"})
	}

	_ = h.logSvc.CreateActivityLog("listUser", []any{ctx.Locals("userID"), len(users)})

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (h Handler) GetUser(ctx *fiber.Ctx) error {

	// get req body
	req := new(param.GetUserRequest)
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

	// get user
	user, err := h.userSvc.GetUserByUsername(req)
	if err != nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", req.Username)).Error(err.Error())

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "databaseError"})
	}

	if user == nil {

		logger.L().With(
			slog.Uint64("reqID", ctx.Context().ID()),
			slog.String("userID", ctx.Locals("userID").(string)),
			slog.String("user", req.Username)).Error("user not found")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userNotFound"})
	}

	_ = h.logSvc.CreateActivityLog("getUser", []any{user.Name, user.Username, ctx.Locals("userID")})

	return ctx.Status(fiber.StatusOK).JSON(user)
}
