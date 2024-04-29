package request

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestVaiCep(t *testing.T) {
    CepBrail, err := VaiCep("91040300")

    assert.Nil(t, err)
    assert.Equal(t, "91040300", strings.ReplaceAll(CepBrail.Cep, "-", ""))
    assert.Contains(t, CepBrail.Localidade, "Alegre")
    assert.Equal(t, "RS", CepBrail.Uf)

}
