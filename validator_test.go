package validator

import "testing"

func TestValidatePassword(t *testing.T) {
	test := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "password is empty",
			input:   "",
			wantErr: true,
		}, {
			name:    "password is empty",
			input:   "        ",
			wantErr: true,
		},
		{
			name:    "password is too short",
			input:   "123",
			wantErr: true,
		},
		{
			name:    "password is too long",
			input:   "123123123123123123123123123123123123123123123123123123123123123123123123123123",
			wantErr: true,
		}, {
			name:    "password is too common",
			input:   "password",
			wantErr: true,
		}, {
			name:    "password is too common",
			input:   "qwerty",
			wantErr: true,
		}, {
			name:    "password is too common",
			input:   "12345678",
			wantErr: true,
		}, {
			name:    "password is too common",
			input:   "admin",
			wantErr: true,
		}, {
			name:    "password is too common",
			input:   "letmein",
			wantErr: true,
		}, {
			name:    "password must not contain spaces",
			input:   "my secret password",
			wantErr: true,
		}, {
			name:    "password must contain uppercase letter",
			input:   "abcabcabc",
			wantErr: true,
		}, {
			name:    "password must contain lowercase letter",
			input:   "ABCABCABC",
			wantErr: true,
		}, {
			name:    "password must contain digit",
			input:   "MYpasswordtop",
			wantErr: true,
		}, {
			name:    "strong password",
			input:   "Q109fgOFDAw_123d",
			wantErr: false,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePassword(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}

}
