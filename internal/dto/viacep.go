package dto

import (
    "strings"
)

type ViaCepInput struct {
    Cep         string `json:"cep"`
    Logradouro  string `json:"logradouro"`
    Complemento string `json:"complemento"`
    Bairro      string `json:"bairro"`
    Localidade  string `json:"localidade"`
    Uf          string `json:"uf"`
    Ibge        string `json:"ibge"`
    Gia         string `json:"gia"`
    Ddd         string `json:"ddd"`
    Siafi       string `json:"siafi"`
}

func (v ViaCepInput) ToCmd() string {
    return "Cep: " + strings.ReplaceAll(v.Cep, "-", "") + "\n" +
        "Logradouro: " + v.Logradouro + "\n" +
        "Bairro: " + v.Bairro + "\n" +
        "Cidade: " + v.Localidade + "\n" +
        "Estado: " + v.Uf + "\n"

}

type ViaCepError struct {
    Error bool `json:"error"`
}
