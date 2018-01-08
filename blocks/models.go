package blocks

type Block struct {
	Difficulty int
	ExtraData uint64
	GasLimit int
	GasUsed int
	Hash uint64
	LogsBloom string
	Miner uint64
	Number int
	ParentHash uint64
	ReceiptRoot uint64
	Sha3Uncles uint64
	Size int
	StateRoot uint64
	Timestamp int
	TotalDifficulty uint64
	Transactions []uint64
	TransactionsRoot uint64
	Uncles []uint64
}

//func LastFiveBlocks() ([]Block, error) {
//
//}