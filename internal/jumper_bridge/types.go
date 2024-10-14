package jumper_bridge

import (
	"github.com/google/uuid"
	"math/big"
	"time"
)

type Route struct {
	FromAddress      string `json:"fromAddress"`
	FromAmount       string `json:"fromAmount"`
	FromChainID      int    `json:"fromChainId"`
	FromTokenAddress string `json:"fromTokenAddress"`
	ToChainID        int    `json:"toChainId"`
	ToTokenAddress   string `json:"toTokenAddress"`
	Options          struct {
		Integrator       string  `json:"integrator"`
		Order            string  `json:"order"`
		Slippage         float64 `json:"slippage"`
		MaxPriceImpact   float64 `json:"maxPriceImpact"`
		AllowSwitchChain bool    `json:"allowSwitchChain"`
	} `json:"options"`
}

type RespRoute struct {
	Routes []struct {
		Id            string `json:"id"`
		FromChainId   int    `json:"fromChainId"`
		FromAmountUSD string `json:"fromAmountUSD"`
		FromAmount    string `json:"fromAmount"`
		FromToken     struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"fromToken"`
		FromAddress string `json:"fromAddress"`
		ToChainId   int    `json:"toChainId"`
		ToAmountUSD string `json:"toAmountUSD"`
		ToAmount    string `json:"toAmount"`
		ToAmountMin string `json:"toAmountMin"`
		ToToken     struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"toToken"`
		ToAddress           string `json:"toAddress"`
		GasCostUSD          string `json:"gasCostUSD"`
		ContainsSwitchChain bool   `json:"containsSwitchChain"`
		Steps               []struct {
			Type        string `json:"type"`
			Id          string `json:"id"`
			Tool        string `json:"tool"`
			ToolDetails struct {
				Key     string `json:"key"`
				Name    string `json:"name"`
				LogoURI string `json:"logoURI"`
			} `json:"toolDetails"`
			Action struct {
				FromToken struct {
					Address  string `json:"address"`
					ChainId  int    `json:"chainId"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					CoinKey  string `json:"coinKey"`
					LogoURI  string `json:"logoURI"`
					PriceUSD string `json:"priceUSD"`
				} `json:"fromToken"`
				FromAmount string `json:"fromAmount"`
				ToToken    struct {
					Address  string `json:"address"`
					ChainId  int    `json:"chainId"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					CoinKey  string `json:"coinKey"`
					LogoURI  string `json:"logoURI"`
					PriceUSD string `json:"priceUSD"`
				} `json:"toToken"`
				FromChainId int     `json:"fromChainId"`
				ToChainId   int     `json:"toChainId"`
				Slippage    float64 `json:"slippage"`
				FromAddress string  `json:"fromAddress"`
				ToAddress   string  `json:"toAddress"`
			} `json:"action"`
			Estimate struct {
				Tool            string `json:"tool"`
				ApprovalAddress string `json:"approvalAddress"`
				ToAmountMin     string `json:"toAmountMin"`
				ToAmount        string `json:"toAmount"`
				FromAmount      string `json:"fromAmount"`
				FeeCosts        []struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					Token       struct {
						Address  string `json:"address"`
						ChainId  int    `json:"chainId"`
						Symbol   string `json:"symbol"`
						Decimals int    `json:"decimals"`
						Name     string `json:"name"`
						CoinKey  string `json:"coinKey"`
						LogoURI  string `json:"logoURI"`
						PriceUSD string `json:"priceUSD"`
					} `json:"token"`
					Amount     string `json:"amount"`
					AmountUSD  string `json:"amountUSD"`
					Percentage string `json:"percentage"`
					Included   bool   `json:"included"`
				} `json:"feeCosts"`
				GasCosts []struct {
					Type      string `json:"type"`
					Price     string `json:"price"`
					Estimate  string `json:"estimate"`
					Limit     string `json:"limit"`
					Amount    string `json:"amount"`
					AmountUSD string `json:"amountUSD"`
					Token     struct {
						Address  string `json:"address"`
						ChainId  int    `json:"chainId"`
						Symbol   string `json:"symbol"`
						Decimals int    `json:"decimals"`
						Name     string `json:"name"`
						CoinKey  string `json:"coinKey"`
						LogoURI  string `json:"logoURI"`
						PriceUSD string `json:"priceUSD"`
					} `json:"token"`
				} `json:"gasCosts"`
				ExecutionDuration float64 `json:"executionDuration"`
				FromAmountUSD     string  `json:"fromAmountUSD"`
				ToAmountUSD       string  `json:"toAmountUSD"`
			} `json:"estimate"`
			IncludedSteps []struct {
				Id     string `json:"id"`
				Type   string `json:"type"`
				Action struct {
					FromChainId int    `json:"fromChainId"`
					FromAmount  string `json:"fromAmount"`
					FromToken   struct {
						Address  string `json:"address"`
						ChainId  int    `json:"chainId"`
						Symbol   string `json:"symbol"`
						Decimals int    `json:"decimals"`
						Name     string `json:"name"`
						CoinKey  string `json:"coinKey"`
						LogoURI  string `json:"logoURI"`
						PriceUSD string `json:"priceUSD"`
					} `json:"fromToken"`
					ToChainId int `json:"toChainId"`
					ToToken   struct {
						Address  string `json:"address"`
						ChainId  int    `json:"chainId"`
						Symbol   string `json:"symbol"`
						Decimals int    `json:"decimals"`
						Name     string `json:"name"`
						CoinKey  string `json:"coinKey"`
						LogoURI  string `json:"logoURI"`
						PriceUSD string `json:"priceUSD"`
					} `json:"toToken"`
					Slippage                  float64 `json:"slippage"`
					FromAddress               string  `json:"fromAddress"`
					DestinationGasConsumption string  `json:"destinationGasConsumption"`
				} `json:"action"`
				Estimate struct {
					FromAmount        string  `json:"fromAmount"`
					ToAmount          string  `json:"toAmount"`
					ToAmountMin       string  `json:"toAmountMin"`
					ApprovalAddress   string  `json:"approvalAddress"`
					ExecutionDuration float64 `json:"executionDuration"`
					FeeCosts          []struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						Token       struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"token"`
						Amount     string `json:"amount"`
						AmountUSD  string `json:"amountUSD"`
						Percentage string `json:"percentage"`
						Included   bool   `json:"included"`
					} `json:"feeCosts"`
					GasCosts []struct {
						Type      string `json:"type"`
						Price     string `json:"price"`
						Estimate  string `json:"estimate"`
						Limit     string `json:"limit"`
						Amount    string `json:"amount"`
						AmountUSD string `json:"amountUSD"`
						Token     struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"token"`
					} `json:"gasCosts"`
					Tool string `json:"tool"`
				} `json:"estimate"`
				Tool        string `json:"tool"`
				ToolDetails struct {
					Key     string `json:"key"`
					Name    string `json:"name"`
					LogoURI string `json:"logoURI"`
				} `json:"toolDetails"`
			} `json:"includedSteps"`
			Integrator string `json:"integrator"`
		} `json:"steps"`
		Tags []string `json:"tags"`
	} `json:"routes"`
	UnavailableRoutes struct {
		FilteredOut []struct {
			OverallPath string `json:"overallPath"`
			Reason      string `json:"reason"`
		} `json:"filteredOut"`
		Failed []struct {
			OverallPath string `json:"overallPath"`
			Subpaths    struct {
				ETHMayanWH8453ETH []struct {
					ErrorType string `json:"errorType"`
					Code      string `json:"code"`
					Action    struct {
						FromChainId int    `json:"fromChainId"`
						FromAmount  string `json:"fromAmount"`
						FromToken   struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"fromToken"`
						ToChainId int `json:"toChainId"`
						ToToken   struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"toToken"`
						Slippage                  float64 `json:"slippage"`
						FromAddress               string  `json:"fromAddress"`
						DestinationGasConsumption string  `json:"destinationGasConsumption"`
					} `json:"action"`
					Tool    string `json:"tool"`
					Message string `json:"message"`
				} `json:"42161:ETH-mayanWH-8453:ETH,omitempty"`
				ETHMayanMCTP8453ETH []struct {
					ErrorType string `json:"errorType"`
					Code      string `json:"code"`
					Action    struct {
						FromChainId int    `json:"fromChainId"`
						FromAmount  string `json:"fromAmount"`
						FromToken   struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"fromToken"`
						ToChainId int `json:"toChainId"`
						ToToken   struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"toToken"`
						Slippage                  float64 `json:"slippage"`
						FromAddress               string  `json:"fromAddress"`
						DestinationGasConsumption string  `json:"destinationGasConsumption"`
					} `json:"action"`
					Tool    string `json:"tool"`
					Message string `json:"message"`
				} `json:"42161:ETH-mayanMCTP-8453:ETH,omitempty"`
				ETHStargate8453ETH []struct {
					ErrorType string `json:"errorType"`
					Code      string `json:"code"`
					Action    struct {
						FromChainId int    `json:"fromChainId"`
						FromAmount  string `json:"fromAmount"`
						FromToken   struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"fromToken"`
						ToChainId int `json:"toChainId"`
						ToToken   struct {
							Address  string `json:"address"`
							ChainId  int    `json:"chainId"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							Name     string `json:"name"`
							CoinKey  string `json:"coinKey"`
							LogoURI  string `json:"logoURI"`
							PriceUSD string `json:"priceUSD"`
						} `json:"toToken"`
						Slippage                  float64 `json:"slippage"`
						FromAddress               string  `json:"fromAddress"`
						DestinationGasConsumption string  `json:"destinationGasConsumption"`
					} `json:"action"`
					Tool    string `json:"tool"`
					Message string `json:"message"`
				} `json:"42161:ETH-stargate-8453:ETH,omitempty"`
			} `json:"subpaths"`
		} `json:"failed"`
	} `json:"unavailableRoutes"`
}

type RTrans struct {
	SessionID          uuid.UUID `json:"sessionId"`
	RouteID            string    `json:"routeId"`
	Integrator         string    `json:"integrator"`
	Action             string    `json:"action"`
	Type               string    `json:"type"`
	FromToken          string    `json:"fromToken"`
	ToToken            string    `json:"toToken"`
	StepNumber         int       `json:"stepNumber"`
	Exchange           string    `json:"exchange"`
	TransactionStatus  string    `json:"transactionStatus"`
	IsFinal            bool      `json:"isFinal"`
	GasCost            *big.Int  `json:"gasCost"`
	GasCostUSD         float64   `json:"gasCostUSD"`
	FromAmountUSD      float64   `json:"fromAmountUSD"`
	ToAmountUSD        float64   `json:"toAmountUSD"`
	FromAmount         *big.Int  `json:"fromAmount"`
	FromChainID        int       `json:"fromChainId"`
	ToAmount           *big.Int  `json:"toAmount"`
	ToChainID          int       `json:"toChainId"`
	TransactionHash    string    `json:"transactionHash"`
	WalletAddress      string    `json:"walletAddress"`
	WalletProvider     string    `json:"walletProvider"`
	URL                string    `json:"url"`
	BrowserFingerprint string    `json:"browserFingerprint"`
}

type TranResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Meta    struct {
		Timestamp time.Time `json:"timestamp"`
		Path      string    `json:"path"`
		Method    string    `json:"method"`
	} `json:"meta"`
}

type StepTxResp struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Tool        string `json:"tool"`
	ToolDetails struct {
		Key     string `json:"key"`
		Name    string `json:"name"`
		LogoURI string `json:"logoURI"`
	} `json:"toolDetails"`
	Action struct {
		FromToken struct {
			Address  string `json:"address"`
			ChainID  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"fromToken"`
		FromAmount string `json:"fromAmount"`
		ToToken    struct {
			Address  string `json:"address"`
			ChainID  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"toToken"`
		FromChainID int     `json:"fromChainId"`
		ToChainID   int     `json:"toChainId"`
		Slippage    float64 `json:"slippage"`
		FromAddress string  `json:"fromAddress"`
		ToAddress   string  `json:"toAddress"`
	} `json:"action"`
	Estimate struct {
		Tool            string `json:"tool"`
		ApprovalAddress string `json:"approvalAddress"`
		ToAmountMin     string `json:"toAmountMin"`
		ToAmount        string `json:"toAmount"`
		FromAmount      string `json:"fromAmount"`
		FeeCosts        []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Token       struct {
				Address  string `json:"address"`
				ChainID  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"token"`
			Amount     string `json:"amount"`
			AmountUSD  string `json:"amountUSD"`
			Percentage string `json:"percentage"`
			Included   bool   `json:"included"`
		} `json:"feeCosts"`
		GasCosts []struct {
			Type      string `json:"type"`
			Price     string `json:"price"`
			Estimate  string `json:"estimate"`
			Limit     string `json:"limit"`
			Amount    string `json:"amount"`
			AmountUSD string `json:"amountUSD"`
			Token     struct {
				Address  string `json:"address"`
				ChainID  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"token"`
		} `json:"gasCosts"`
		ExecutionDuration int    `json:"executionDuration"`
		FromAmountUSD     string `json:"fromAmountUSD"`
		ToAmountUSD       string `json:"toAmountUSD"`
	} `json:"estimate"`
	IncludedSteps []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Action struct {
			FromChainID int    `json:"fromChainId"`
			FromAmount  string `json:"fromAmount"`
			FromToken   struct {
				Address  string `json:"address"`
				ChainID  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"fromToken"`
			ToChainID int `json:"toChainId"`
			ToToken   struct {
				Address  string `json:"address"`
				ChainID  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"toToken"`
			Slippage                  float64 `json:"slippage"`
			FromAddress               string  `json:"fromAddress"`
			DestinationGasConsumption string  `json:"destinationGasConsumption"`
			ToAddress                 string  `json:"toAddress"`
		} `json:"action"`
		Estimate struct {
			FromAmount        string `json:"fromAmount"`
			ToAmount          string `json:"toAmount"`
			ToAmountMin       string `json:"toAmountMin"`
			ApprovalAddress   string `json:"approvalAddress"`
			ExecutionDuration int    `json:"executionDuration"`
			FeeCosts          []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Token       struct {
					Address  string `json:"address"`
					ChainID  int    `json:"chainId"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					CoinKey  string `json:"coinKey"`
					LogoURI  string `json:"logoURI"`
					PriceUSD string `json:"priceUSD"`
				} `json:"token"`
				Amount     string `json:"amount"`
				AmountUSD  string `json:"amountUSD"`
				Percentage string `json:"percentage"`
				Included   bool   `json:"included"`
			} `json:"feeCosts"`
			GasCosts []struct {
				Type      string `json:"type"`
				Price     string `json:"price"`
				Estimate  string `json:"estimate"`
				Limit     string `json:"limit"`
				Amount    string `json:"amount"`
				AmountUSD string `json:"amountUSD"`
				Token     struct {
					Address  string `json:"address"`
					ChainID  int    `json:"chainId"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					CoinKey  string `json:"coinKey"`
					LogoURI  string `json:"logoURI"`
					PriceUSD string `json:"priceUSD"`
				} `json:"token"`
			} `json:"gasCosts"`
			Tool string `json:"tool"`
		} `json:"estimate"`
		Tool        string `json:"tool"`
		ToolDetails struct {
			Key     string `json:"key"`
			Name    string `json:"name"`
			LogoURI string `json:"logoURI"`
		} `json:"toolDetails"`
	} `json:"includedSteps"`
	Integrator         string `json:"integrator"`
	TransactionRequest struct {
		Data     string `json:"data"`
		To       string `json:"to"`
		Value    string `json:"value"`
		From     string `json:"from"`
		ChainID  int    `json:"chainId"`
		GasPrice string `json:"gasPrice"`
		GasLimit string `json:"gasLimit"`
	} `json:"transactionRequest"`
}
