package request

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBrasilApi(t *testing.T) {
    CepBrail, err := BrasilApi("91040300")

    assert.Nil(t, err)
    assert.Equal(t, "91040300", CepBrail.Cep)
    assert.Contains(t, CepBrail.City, "Alegre")
    assert.Equal(t, "RS", CepBrail.State)

}

func TestBrasilApiNaoExiste(t *testing.T) {
    CepBrail, err := BrasilApi("00000000")

    assert.Nil(t, err)
    assert.Equal(t, "", CepBrail.Cep)
    assert.Equal(t, "", CepBrail.City)
    assert.Equal(t, "", CepBrail.State)
    assert.Equal(t, "SN", CepBrail.Service)
}
