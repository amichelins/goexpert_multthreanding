package request

import (
    "encoding/json"
    "io"
    "net/http"

    "github.com/amichelins/goexpert_multthreanding/internal/dto"
)

// BrasilApi Recebe um CEP e consulta brasilapi para saber os dados do cep
//
// PARAMETERS
//
//     sCep string Cep para obter os dados
//
// RETURN
//
//     *dto.BrasilApiInput Dados do CEP
//
//     error Erro ocorrido ou nil
//
func BrasilApi(sCep string) (dto.BrasilApiInput, error) {
    var CepBrasil dto.BrasilApiInput

    request, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + sCep)

    if err != nil {
        return CepBrasil, err
    }

    if request.StatusCode == 404 {
        CepBrasil.Service = "SN" // Indica que não a dados

        return CepBrasil, nil
    }

    // Pegamos as resposta
    data, err := io.ReadAll(request.Body)

    if err != nil {
        return CepBrasil, err
    }
    defer request.Body.Close()

    err = json.Unmarshal(data, &CepBrasil)

    if err != nil {
        return CepBrasil, err
    }

    return CepBrasil, nil

}
