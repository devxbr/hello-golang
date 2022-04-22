package cotacao

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Currency string

type Cotacao struct {
	Active Currency
	Bid,
	Ask string
}

type CotacaoResponse struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

const (
	USDBTC Currency = "BTCUSDT"
	USDETH Currency = "ETHUSDT"
)

type RetornoCotacao struct {
	Cotacao  Cotacao
	HasError error
}

// RetornarCotacao retorna as ofertas de compra e venda disponível no book da binance
func RetornarCotacao(currency Currency) RetornoCotacao {
	res, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/depth?symbol=%s&limit=1", currency))

	if err != nil {
		return RetornoCotacao{Cotacao{currency, "", ""}, errors.New(fmt.Sprintf("Erro ao obter cotação - Detalhes %v", err))}
	}

	if status := res.StatusCode; status != http.StatusOK {
		return RetornoCotacao{Cotacao{currency, "", ""}, errors.New(fmt.Sprintf("Erro ao obter cotação - Detalhes %v", res.Status))}
	}

	// defer aguarda a execução de todo os blocos lógicos e quando tem mais de um observei um comportamento FILO
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return RetornoCotacao{Cotacao{currency, "", ""}, errors.New(fmt.Sprintf("Erro ao ler retorno - Detalhes %v", err))}
	}

	var result CotacaoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("can't unmarshal JSON", err)
	}

	return RetornoCotacao{Cotacao{Active: currency, Ask: result.Asks[0][0], Bid: result.Bids[0][0]}, nil}
}
