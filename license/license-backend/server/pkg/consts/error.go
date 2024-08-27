package consts

const (
	DB_ERORROR        = "database error"
	UNEXPECTED_ERROR  = "unexpected error"
	VALIDATION_ERROR  = "validation error"
	BAD_REQUEST_ERROR = "bad request error"
	FORBIDDEN_ERROR   = "forbidden error"
)

var ValidationErrors = map[string]string{
	"Username":         "usernameIsRequired",
	"Password":         "passwordIsRequired",
	"Name":             "nameIsRequired",
	"Role":             "roleIsRequired",
	"ID":               "idIsRequired",
	"CreatedAt":        "createdAtIsRequired",
	"ProductID":        "productIDIsRequired",
	"Key":              "keyIsRequired",
	"Value":            "valueIsRequired",
	"Title":            "titleIsRequired",
	"Version":          "versionIsRequired",
	"CustomerID":       "customerIDIsRequired",
	"LicenseType":      "licenseTypeIsRequired",
	"IsActive":         "isActiveIsRequired",
	"ExpireAt":         "expireAtIsRequired",
	"HardwareHash":     "hardwareHashIsRequired",
	"FirstConfirmedAt": "firstConfirmedAtIsRequired",
	"LastConfirmedAt":  "lastConfirmedAtIsRequired",
	"Email":            "emailIsRequired",
	"Phone":            "phoneIsRequired",
}
