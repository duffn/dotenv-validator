package validator

import (
	"bufio"
	"fmt"
	"io"
	"net/mail"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ValidateWithFilename(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return processEnvFile(file)
}

func Validate() error {
	return ValidateWithFilename(".env.sample")
}

func processEnvFile(file io.Reader) error {
	var missingVariables []string
	var invalidFormats []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " #")

		var config string
		if len(lineParts) > 1 {
			config = lineParts[1]
		}

		variableParts := strings.Split(lineParts[0], "=")
		variableName := variableParts[0]

		value := os.Getenv(variableName)
		if value == "" && config != "" && strings.Contains(config, "required") {
			missingVariables = append(missingVariables, variableName)
			continue
		}

		r := regexp.MustCompile("format=(?P<Format>.*)?")
		match := r.FindStringSubmatch(config)
		if match == nil {
			continue
		}

		valid, err := checkValueFormat(value, match[1])
		if err != nil {
			return err
		}

		if !valid {
			invalidFormats = append(invalidFormats, variableName)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if missingVariables != nil || invalidFormats != nil {
		return fmt.Errorf("these variables are missing in the envionment (%s)\nthese variables have invalid formats (%s)", strings.Join(missingVariables, ","), strings.Join(invalidFormats, ","))
	}

	return nil
}

func checkValueFormat(value string, format string) (bool, error) {
	switch format {
	case "int", "integer":
		if _, err := strconv.ParseInt(value, 10, 64); err == nil {
			return true, nil
		}
	case "float":
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			return true, nil
		}
	case "str", "string":
		return true, nil
	case "email":
		_, err := mail.ParseAddress(value)
		return err == nil, nil
	case "url", "uri":
		_, err := url.ParseRequestURI(value)
		return err == nil, nil
	default:
		userRegex, err := regexp.Compile(format)
		if err != nil {
			return false, err
		}

		userMatch := userRegex.FindString(value)
		fmt.Println(userMatch)
		if userMatch != "" {
			return true, nil
		}
	}

	return false, nil
}
