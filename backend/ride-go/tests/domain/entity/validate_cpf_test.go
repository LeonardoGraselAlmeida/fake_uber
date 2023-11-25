package entity_test

import (
	"testing"

	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func Test_If_It_Is_Valid_When_Pass_Cpf_Valid(t *testing.T) {
	cpf := "714.287.938-60"
	assert.True(t, entity.ValidateCpf(cpf))
}

func Test_If_It_Is_Valid_When_Pass_Cpf_Empty(t *testing.T) {
	cpf := ""
	assert.False(t, entity.ValidateCpf(cpf))
}

func Test_If_It_Is_Valid_When_Pass_Cpf_Invalid(t *testing.T) {
	cpf := "714.287.938-61"
	assert.False(t, entity.ValidateCpf(cpf))
}

func Test_If_It_Is_Valid_When_Pass_Cpf_With_All_Digits_Are_The_Same(t *testing.T) {
	cpf := "111.111.111-11"
	assert.False(t, entity.ValidateCpf(cpf))
}

func Test_If_It_Is_Valid_When_Pass_Cpf_Bigger(t *testing.T) {
	cpf := "714.287.938-600"
	assert.False(t, entity.ValidateCpf(cpf))
}

func Test_If_It_Is_Valid_When_Pass_Cpf_Smaller(t *testing.T) {
	cpf := "714.287.938-6"
	assert.False(t, entity.ValidateCpf(cpf))
}
