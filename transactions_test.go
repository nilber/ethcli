package ethcli_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethcli "github.com/medeirosfalante/ethcli"
)

func TestGetByBlock(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	txs := ethcli.NewTransactions(client)

	tx, err := txs.GetBlock(9197771, 9197775, []string{"0xf35d75E2Ce765fD4aB1Da7b331eB03C56D4859c4"})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(tx) <= 0 {
		t.Errorf("invalid counter  need more than 0")
		return
	}

}
