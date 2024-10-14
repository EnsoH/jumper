package web3

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	jb "jumper/internal/jumper_bridge"
	"math/big"
	"strings"
)

type Client struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func New(rpc string, privateKey string) (*Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("can't connect to RPC: %w", err)
	}

	pKey, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return nil, fmt.Errorf("can't convert private key: %w", err)
	}

	wallet := &Client{
		Client:     client,
		PrivateKey: pKey,
		Address:    crypto.PubkeyToAddress(pKey.PublicKey),
	}

	return wallet, nil
}

func (c *Client) MakeBridgeTx(params *jb.RoutesParams) {
	// TODO: понять как отправить транзакцию
}

func (c *Client) WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236)
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func (c *Client) EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func (c *Client) ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

func (c *Client) GetNonce(ctx context.Context) (uint64, error) {
	nonce, err := c.Client.PendingNonceAt(ctx, c.Address)
	if err != nil {
		return 0, fmt.Errorf("error getting nonce: %w", err)
	}

	return nonce, err
}

func (c *Client) GetTipCap(ctx context.Context) (*big.Int, error) {
	tipCap, err := c.Client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting gas tip: %w", err)
	}

	return tipCap, nil
}

func (c *Client) GetBaseFee(ctx context.Context) (*big.Int, error) {
	header, err := c.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting block header: %w", err)
	}

	return header.BaseFee, nil
}

func (c *Client) GetMaxFeePerGas(ctx context.Context) (*big.Int, error) {
	tipCap, err := c.GetTipCap(ctx)
	if err != nil {
		return nil, err
	}

	baseFee, err := c.GetBaseFee(ctx)
	if err != nil {
		return nil, err
	}

	maxFeePerGas := new(big.Int).Add(baseFee, tipCap)

	return maxFeePerGas, nil
}

func (c *Client) GetGasLimit(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := c.Client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, fmt.Errorf("error when estimate gas limit: %w", err)
	}

	return gasLimit, nil
}
