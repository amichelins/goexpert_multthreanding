package request

import (
    "encoding/json"
    "io"
    "net/http"
    "strings"

    "github.com/amichelins/goexpert_multthreanding/internal/dto"
)

// ViaCep Recebe um CEP e consulta VIACEP para saber os dados do cep
//
// PARAMETERS
//
//     sCep string Cep para obter os dados
//
// RETURN
//
//     *dto.ViaCepInput Dados do CEP
//
//     error Erro ocorrido ou nil
//
func VaiCep(sCep string) (dto.ViaCepInput, error) {
    var CepBrasil dto.ViaCepInput

    request, err := http.Get("https://viacep.com.br/ws/" + sCep + "/json/")

    if err != nil {
        return CepBrasil, err
    }

    // Pegamos as resposta
    data, err := io.ReadAll(request.Body)

    if err != nil {
        return CepBrasil, err
    }
    defer request.Body.Close()

    if strings.Contains(string(data), `"erro"`) {
        CepBrasil.Siafi = "SN"

        return CepBrasil, nil
    }

    err = json.Unmarshal(data, &CepBrasil)

    if err != nil {
        return CepBrasil, err
    }

    return CepBrasil, nil

}
