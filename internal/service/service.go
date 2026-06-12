package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	isMorse := true

	for _, ch := range input {
		if ch != '.' && ch != '-' && ch != ' ' && ch != '\n' && ch != '\r' && ch != '\t' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}
