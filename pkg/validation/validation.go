package validation

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en" // English
	ut "github.com/go-playground/universal-translator"
	libvalidator "github.com/go-playground/validator/v10"
	english "github.com/go-playground/validator/v10/translations/en"
)

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResultError struct {
	Message string
	Errors  []validationError
}

func (v ValidationResultError) Error() string {
	return v.Message
}

type Validator struct {
	validator  *libvalidator.Validate
	translator ut.Translator
}

// NewValidator membuat instance dari Validator
func NewValidator() Validator {
	validatorInstance := libvalidator.New(libvalidator.WithRequiredStructEnabled())
	uni := ut.New(en.New(), en.New())
	translator, _ := uni.GetTranslator("en")
	english.RegisterDefaultTranslations(validatorInstance, translator)

	return Validator{validator: validatorInstance, translator: translator}
}

// Validate memvalidasi struct yang akan mengembalikan ValidationResult jika ada validasi error.
// Jika tidak ada validasi error, maka akan mengembalikan nil.
// ValidationResult kompatibel dengan tipe error
func (v *Validator) Validate(request interface{}) *ValidationResultError {
	var validationErrors []validationError

	// Validate struct
	if err := v.validator.Struct(request); err != nil {
		// Check tipe error apakah ValidationErrors
		var validationErrs libvalidator.ValidationErrors
		if errors.As(err, &validationErrs) {
			// Looping lalu translate pesan error
			for _, validationErr := range validationErrs {
				translatedMessage := validationErr.Translate(v.translator)
				validationErrors = append(validationErrors, validationError{
					Field:   getJSONKeyFromTag(request, validationErr.Field()),
					Message: translatedMessage,
				})
			}

			return &ValidationResultError{
				Message: strings.Join(extractMessages(validationErrors), "|"),
				Errors:  validationErrors,
			}
		}
	}
	return nil
}

// extractMessages mengambil pesan error dari slice validationError
func extractMessages(errors []validationError) []string {
	messages := make([]string, len(errors))
	for i, err := range errors {
		messages[i] = err.Message
	}
	return messages
}

// getJSONKeyFromTag mengambil JSON key dari struct tag
func getJSONKeyFromTag(structInstance interface{}, fieldName string) string {
	structType := reflect.TypeOf(structInstance)

	field, exists := structType.FieldByName(fieldName)
	if !exists {
		return fieldName
	}

	// Ambil tag "json"
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return field.Name // Jika tidak ada json tag, return nama field
	}

	// Coba ambil json key nya
	jsonKey := strings.Split(jsonTag, ",")[0]
	if jsonKey == "" {
		return field.Name // Jika tidak ada, return nama field
	}

	return jsonKey
}
