package blocks

type Block struct {
	Difficulty int
	ExtraData float64
	GasLimit int
	GasUsed int
	Hash float64
	LogsBloom string
	Miner float64
	Number int
	ParentHash float64
	ReceiptRoot float64
	Sha3Uncles float64
	Size int
	StateRoot float64
	Timestamp int
	TotalDifficulty float64
	Transactions []float64
	TransactionsRoot float64
	Uncles []float64
}

//func LastFiveBlocks() ([]Block, error) {
//
//}