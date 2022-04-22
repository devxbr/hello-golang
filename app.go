package main

import (
	"fmt"
	"hello-golang/cotacao"
	"time"
)

func main() {
	const UM_MINUTO = time.Second * 60

	for {
		app := cotacao.RetornarCotacao(cotacao.USDBTC)
		if app.HasError != nil {
			fmt.Printf("Deu ruim! %v", app.HasError)
		} else {
			fmt.Printf("Pair: %s  Ask: %s Bid: %s \n", app.Cotacao.Active, app.Cotacao.Ask, app.Cotacao.Bid)
		}

		time.Sleep(UM_MINUTO)
	}
}
