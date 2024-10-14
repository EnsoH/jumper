package jumper_bridge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

type Jumper struct{}

func New() *Jumper {
	return &Jumper{}
}

type RoutesParams struct {
	FromAddress  string
	FromAmount   string
	FromChId     int
	FromTokenAdd string
	ToChId       int
	ToTokenAdd   string
}

func (j *Jumper) Routes(client *http.Client, params *RoutesParams) (*RespRoute, error) {
	const ApiEndpoint = "https://li.quest/v1/advanced/routes"

	data := Route{
		FromAddress:      params.FromAddress,
		FromAmount:       params.FromAmount,
		FromChainID:      params.FromChId,
		FromTokenAddress: params.FromTokenAdd,
		ToChainID:        params.ToChId,
		ToTokenAddress:   params.ToTokenAdd,
		Options: struct {
			Integrator       string  `json:"integrator"`
			Order            string  `json:"order"`
			Slippage         float64 `json:"slippage"`
			MaxPriceImpact   float64 `json:"maxPriceImpact"`
			AllowSwitchChain bool    `json:"allowSwitchChain"`
		}{
			Integrator:       "jumper.exchange",
			Order:            "CHEAPEST",
			Slippage:         0.005,
			MaxPriceImpact:   0.4,
			AllowSwitchChain: true,
		},
	}

	dataB, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, ApiEndpoint, bytes.NewBuffer(dataB))
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://jumper.exchange")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://jumper.exchange/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	req.Header.Set("x-lifi-integrator", "jumper.exchange")
	req.Header.Set("x-lifi-sdk", "3.2.3")
	req.Header.Set("x-lifi-widget", "3.6.2")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respData *RespRoute
	err = json.Unmarshal(body, &respData)

	info := respData.Routes[0]

	fmt.Println("ROUTE ID:", info.Id)
	fmt.Println("FROM CHAIN:", info.FromChainId)
	fmt.Println("TO CHAIN:", info.ToChainId)
	fmt.Println("FROM AMOUNT USD:", info.FromAmountUSD)
	fmt.Println("FROM AMOUNT:", info.FromAmount)
	fmt.Println("TO AMOUNT USD:", info.ToAmountUSD)
	fmt.Println("TO AMOUNT:", info.ToAmount)

	for _, tag := range info.Tags {
		fmt.Println("TAGS", tag)
	}

	return respData, nil
}

func (j *Jumper) Transactions(client *http.Client, params *RoutesParams) {
	const ApiEndpoint = "https://api.jumper.exchange/v1/wallets/transactions"

	route, err := j.Routes(client, params)
	if err != nil {
		fmt.Println(err)
	}

	sessionId := uuid.New()

	gasCostUSD, err := strconv.ParseFloat(route.Routes[0].GasCostUSD, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	}

	fromAmountUSD, err := strconv.ParseFloat(route.Routes[0].FromAmountUSD, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	}

	toAmountUSD, err := strconv.ParseFloat(route.Routes[0].ToAmountUSD, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	}

	gasCost := new(big.Int)
	if _, ok := gasCost.SetString(route.Routes[0].Steps[0].Estimate.GasCosts[0].Amount, 10); !ok {
		log.Fatal("Invalid gas cost amount")
	}

	fromAmount := new(big.Int)
	if _, ok := fromAmount.SetString(route.Routes[0].FromAmount, 10); !ok {
		log.Fatal("Invalid from amount")
	}

	toAmount := new(big.Int)
	if _, ok := toAmount.SetString(route.Routes[0].ToAmount, 10); !ok {
		log.Fatal("Invalid to amount")
	}

	data := RTrans{
		SessionID:          sessionId,
		RouteID:            route.Routes[0].Steps[0].Id,
		Integrator:         route.Routes[0].Steps[0].Integrator,
		Action:             "execution_start",
		Type:               "CROSS_CHAIN",
		FromToken:          route.Routes[0].FromToken.Address,
		ToToken:            route.Routes[0].ToToken.Address,
		StepNumber:         1,
		Exchange:           route.Routes[0].Steps[0].ToolDetails.Key,
		TransactionStatus:  "STARTED",
		IsFinal:            false,
		GasCost:            gasCost,
		GasCostUSD:         gasCostUSD,
		FromAmountUSD:      fromAmountUSD,
		ToAmountUSD:        toAmountUSD,
		FromAmount:         fromAmount,
		FromChainID:        route.Routes[0].FromChainId,
		ToAmount:           toAmount,
		ToChainID:          route.Routes[0].ToChainId,
		TransactionHash:    "",
		WalletAddress:      params.FromAddress,
		WalletProvider:     "Rabby Wallet",
		URL:                fmt.Sprintf("https://jumper.exchange/?fromChain=%d&fromToken=%s&toChain=%d&toToken=%s", route.Routes[0].FromChainId, route.Routes[0].FromToken.Address, route.Routes[0].ToChainId, route.Routes[0].ToToken.Address),
		BrowserFingerprint: "unknown",
	}

	dataB, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, ApiEndpoint, bytes.NewBuffer(dataB))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://jumper.exchange")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://jumper.exchange/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var tranResp *TranResp
	err = json.Unmarshal(body, &tranResp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("STATUS", tranResp.Status)
	fmt.Println("MESSAGE", tranResp.Message)

	j.stepTransaction(client, route)
}

func (j *Jumper) stepTransaction(client *http.Client, route *RespRoute) {
	const ApiEndpoint = "https://li.quest/v1/advanced/stepTransaction"

	dataB, err := json.Marshal(route.Routes[0].Steps[0])
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, ApiEndpoint, bytes.NewBuffer(dataB))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://jumper.exchange")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://jumper.exchange/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	req.Header.Set("x-lifi-integrator", "jumper.exchange")
	req.Header.Set("x-lifi-sdk", "3.2.3")
	req.Header.Set("x-lifi-widget", "3.6.2")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data *StepTxResp
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DATA TX", data.TransactionRequest.Data)
}
