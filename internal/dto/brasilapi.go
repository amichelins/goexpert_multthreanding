package dto

type BrasilApiInput struct {
    Cep          string `json:"cep"`
    State        string `json:"state"`
    City         string `json:"city"`
    Neighborhood string `json:"neighborhood"`
    Street       string `json:"street"`
    Service      string `json:"service"`
}

func (b BrasilApiInput) ToCmd() string {
    return "Cep: " + b.Cep + "\n" +
        "Logradouro: " + b.Street + "\n" +
        "Bairro: " + b.Neighborhood + "\n" +
        "Cidade: " + b.City + "\n" +
        "Estado: " + b.State + "\n"

}
