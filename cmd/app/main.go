package main

import (
	"fmt"
	"jumper/internal/jumper_bridge"
	"jumper/internal/web3"
	"jumper/pkg"
)

const (
	ETH    = 1      // Ethereum Mainnet
	BSC    = 56     // Binance Smart Chain
	OP     = 10     // Optimism
	ARB    = 42161  // Arbitrum One
	MATIC  = 137    // Polygon (Matic)
	BASE   = 8453   // Base Mainnet
	BLAST  = 246    // Blast Network
	SOLANA = ""     // Solana
	AVAX   = 43114  // Avalanche C-Chain
	SCROLL = 534351 // Scroll Alpha
	LINEA  = 59144  // Linea Mainnet
	ZKSYNC = 324    // zkSync Era Mainnet
)

// Basic info
const (
	fromAddress = "0xd6Bf7a05538a620C6a880AA77Ae8d0bdaf27FAAA"

	fromChainId      = ARB
	fromTokenAddress = "0x0000000000000000000000000000000000000000"

	toChainId      = BASE
	toTokenAddress = "0x0000000000000000000000000000000000000000"

	amountUsd = 0
	amountEth = "0.002"

	RPC        = "https://rpc.ankr.com/eth"
	PrivateKey = "0x11a2a0200e408146f754610fe9114f2916c887b44255927ad526b4fc05b8acab"
)

func main() {
	// TODO: реальзовать отправку транзакции

	client := pkg.New()
	web, err := web3.New(RPC, PrivateKey)
	if err != nil {
		fmt.Println(err)
	}
	jump := jumper_bridge.New()

	bigFl, err := web.ParseBigFloat(amountEth)
	if err != nil {
		fmt.Println(err)
	}

	weiVal := web.EtherToWei(bigFl)

	payload := jumper_bridge.RoutesParams{
		FromAddress:  fromAddress,
		FromAmount:   weiVal.String(),
		FromChId:     fromChainId,
		FromTokenAdd: fromTokenAddress,
		ToChId:       toChainId,
		ToTokenAdd:   toTokenAddress,
	}

	jump.Transactions(client, &payload)

}
