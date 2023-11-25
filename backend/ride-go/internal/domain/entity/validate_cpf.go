package entity

import (
	"regexp"
	"strconv"
	"strings"
)

func ValidateCpf(cpf string) bool {
	if cpf == "" {
		return false
	}

	cpf = cleanSpecialCaracteres(cpf)

	if isInvalidLength(cpf) {
		return false
	}

	if allDigitsAreTheSame(cpf) {
		return false
	}

	digit1 := calculateDigit(cpf, 10)
	digit2 := calculateDigit(cpf, 11)

	return extractCheckDigit(cpf) == (strconv.Itoa(digit1) + strconv.Itoa(digit2))
}

func cleanSpecialCaracteres(cpf string) string {
	regex := regexp.MustCompile(`\d+`)
	return strings.Join(regex.FindAllString(cpf, -1), "")
}

func isInvalidLength(cpf string) bool {
	return len(cpf) != 11
}

func allDigitsAreTheSame(cpf string) bool {
	caracteres := make(map[rune]int)

	for _, char := range cpf {
		caracteres[char]++
	}

	for _, count := range caracteres {
		if count <= 1 {
			return false
		}
	}

	return true
}

func calculateDigit(cpf string, factor int) int {
	sum := 0

	for _, caracterer := range cpf {
		if factor > 1 {
			digit, _ := strconv.Atoi(string(caracterer))
			sum += digit * factor
			factor--
		}
	}

	rest := sum % 11

	if rest < 2 {
		return 0
	}

	return 11 - rest
}

func extractCheckDigit(cpf string) string {
	return cpf[9:]
}
