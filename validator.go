// Package validator содержит функции валидации пользовательских данных.
package validator

import (
	"errors"
	"strings"
	"unicode"
)

// Ошибки валидации пароля.
var (
	ErrPasswordEmpty      = errors.New("password is empty")                      //ec
	ErrPasswordTooShort   = errors.New("password is too short")                  //ec
	ErrPasswordTooLong    = errors.New("password is too long")                   //ec
	ErrPasswordNoUpper    = errors.New("password must contain uppercase letter") //ec
	ErrPasswordNoLower    = errors.New("password must contain lowercase letter") //ec
	ErrPasswordNoDigit    = errors.New("password must contain digit")            //ec
	ErrPasswordHasSpace   = errors.New("password must not contain spaces")       //ec
	ErrPasswordCommonWord = errors.New("password is too common")
)

const (
	minPasswordLen = 8
	maxPasswordLen = 64
)

// commonPasswords — список запрещённых очевидных паролей (регистр не учитывается).
var commonPasswords = map[string]struct{}{
	"password": {},
	"qwerty":   {},
	"12345678": {},
	"admin":    {},
	"letmein":  {},
}

// ValidatePassword проверяет пароль по набору правил:
//   - не пустой и не состоит только из пробелов;
//   - длина от 8 до 64 символов включительно;
//   - содержит хотя бы одну заглавную букву (любого алфавита);
//   - содержит хотя бы одну строчную букву;
//   - содержит хотя бы одну цифру;
//   - не содержит пробельных символов;
//   - не входит в список очевидных паролей (регистр не учитывается).
//
// Возвращает первую найденную ошибку или nil, если пароль валиден.
// Порядок проверок: пустота → длина → запрещённое слово → состав символов.
func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return ErrPasswordEmpty
	}

	if len(password) < minPasswordLen {
		return ErrPasswordTooShort
	}
	if len(password) > maxPasswordLen {
		return ErrPasswordTooLong
	}

	if _, ok := commonPasswords[strings.ToLower(password)]; ok {
		return ErrPasswordCommonWord
	}

	var hasUpper, hasLower, hasDigit bool
	for _, r := range password {
		switch {
		case unicode.IsSpace(r):
			return ErrPasswordHasSpace
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	if !hasUpper {
		return ErrPasswordNoUpper
	}
	if !hasLower {
		return ErrPasswordNoLower
	}
	if !hasDigit {
		return ErrPasswordNoDigit
	}

	return nil
}

/*func TestValidatePassword(t *testing.T) {
    tests := []struct {
        name     string
        password string
        wantErr  bool // true — ждем ошибку, false — ждем nil
    }{
        {
            name:     "valid strong password",
            password: "SecurePassword123!",
            wantErr:  false, // ожидаем nil
        },
        {
            name:     "too short password",
            password: "123",
            wantErr:  true, // ожидаем ошибку
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidatePassword(tt.password)

            if (err != nil) != tt.wantErr {
                t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}*/
