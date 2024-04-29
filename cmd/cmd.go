package main

import (
    "log"
    "os"
    "regexp"
    "time"

    "github.com/amichelins/goexpert_multthreanding/internal/dto"
    "github.com/amichelins/goexpert_multthreanding/internal/request"
)

func main() {
    var sCep string

    if len(os.Args) > 1 {
        sCep = os.Args[1]
        sCep = regexp.MustCompile(`[^0-9]`).ReplaceAllString(sCep, "")

        if len(sCep) != 8 {
            log.Fatal("Parâmetro errado. O parâmetro debe ser um CEP valid ( 8 dígitos númericos )")
        }

        cViaCep := make(chan dto.ViaCepInput)

        cBrasilApi := make(chan dto.BrasilApiInput)

        go func() {
            cep, err := request.VaiCep(sCep)

            if err == nil {
                cViaCep <- cep
            }

        }()

        go func() {
            cep, err := request.BrasilApi(sCep)

            if err == nil {
                cBrasilApi <- cep
            }

        }()

        select {
        // Se BRASILAPI enviou primeiro mostramos os dados
        case msg := <-cBrasilApi:
            println("URL BrasilAPI\n")
            if msg.Service == "SN" {
                println("CEP: " + sCep + " inexistente!!!")
            } else {
                println(msg.ToCmd())
            }

        // Se VIACEP enviou primeiro mostramos os dados
        case msg := <-cViaCep:
            println("URL ViaCep\n")
            if msg.Siafi == "SN" {
                println("CEP: " + sCep + " inexistente!!!")
            } else {
                println(msg.ToCmd())
            }
        // Se passou mais do que 1 segundo emitimos timeout
        case <-time.After(time.Second * 1):
            log.Fatal("Timeout após 1 Segungo!!!")
        }
    } else {
        println("Favor informar o CEP a ser procurado.")
        println("Para rodas:")
        println("  cmd {cep}")
        os.Exit(0)
    }
}
