package BLC

import "fmt"

type TXInput struct {
	TxID    []byte // 1. 交易的Hash
	Vout      int    //2. 存储TXOutput在Vout里面的索引(第几个交易)
	ScriptSig string // 3. 用户名花费的是谁的钱(解锁脚本,包含数字签名)
}


//判断TXInput是否指定的address消费
func (txInput *TXInput) UnlockWithAddress(address string) bool {
	return txInput.ScriptSig == address
}

//格式化输出
func (tx *TXInput) String() string {
	return fmt.Sprintf("\n\t\t\tTxInput_TXID: %x, Vout: %v, ScriptSig: %v", tx.TxID, tx.Vout, tx.ScriptSig)
}