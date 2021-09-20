package validator

import (
	"fmt"
	"strings"
	"testing"
)

func Test_ValidationErrors(t *testing.T) {
	var tests = []struct {
		envFile     string
		errorMesage string
		setEnvs     map[string]string
	}{
		{
			"VAR=flerp # required\nVAR2=notrequired\nVAR3=1234       #      required",
			`these variables are missing in the envionment (VAR,VAR3)
these variables have invalid formats ()`,
			nil,
		},
		{
			"VAR=2 # required,format=int\nVAR2=1.31415 #format=float\nVAR3=str #    format=str\nVAR4=bob@bobloblaw.com # required,format=email\nVAR5=https://www.bobloblawlaw.com #   format=url\nVAR6=1234556 # format=\\d+",
			`these variables are missing in the envionment ()
these variables have invalid formats (VAR,VAR2,VAR4,VAR5,VAR6)`,
			map[string]string{"VAR": "flerp", "VAR2": "derp", "VAR3": "str", "VAR4": "bobloblaw.com", "VAR5": "https//bobloblaw.com", "VAR6": "bob"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test invalid error messages %d", i), func(t *testing.T) {
			if tt.setEnvs != nil {
				for env, value := range tt.setEnvs {
					t.Setenv(env, value)
				}
			}

			err := processEnvFile(strings.NewReader(tt.envFile))
			if err.Error() != tt.errorMesage {
				t.Errorf("got %s want %s", err.Error(), tt.errorMesage)
			}
		})
	}
}

func Test_NoValidationErrors(t *testing.T) {
	var tests = []struct {
		envFile string
		setEnvs map[string]string
	}{
		{
			"VAR=flerp # required\nVAR2=notrequired",
			map[string]string{"VAR": "flerp"},
		},
		{
			"VAR=2 # required,format=int\nVAR2=1.31415 #format=float\nVAR3=str #    format=str\nVAR4=bob@bobloblaw.com # required,format=email\nVAR5=https://www.bobloblawlaw.com #   format=url\nVAR6=ABCD # format=[A-Z]+",
			map[string]string{"VAR": "2", "VAR2": "1.31415", "VAR3": "str", "VAR4": "bob@bobloblaw.com", "VAR5": "https://www.bobloblawlaw.com", "VAR6": "ABCD"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test no validation errors %d", i), func(t *testing.T) {
			if tt.setEnvs != nil {
				for env, value := range tt.setEnvs {
					t.Setenv(env, value)
				}
			}

			err := processEnvFile(strings.NewReader(tt.envFile))
			if err != nil {
				t.Errorf("got %v", err)
			}
		})
	}
}
