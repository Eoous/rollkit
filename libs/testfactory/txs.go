package testfactory

import (
	mrand "math/rand"

	"github.com/cometbft/cometbft/types"
)

func GenerateRandomlySizedTxs(count, maxSize int) types.Txs {
	txs := make(types.Txs, count)
	for i := 0; i < count; i++ {
		size := mrand.Intn(maxSize) //nolint:gosec
		if size == 0 {
			size = 1
		}
		txs[i] = GenerateRandomTxs(1, size)[0]
	}
	return txs
}

func GenerateRandomTxs(count, size int) types.Txs {
	txs := make(types.Txs, count)
	for i := 0; i < count; i++ {
		tx := make([]byte, size)
		_, err := mrand.Read(tx) //nolint:gosec
		if err != nil {
			panic(err)
		}
		txs[i] = tx
	}
	return txs
}
