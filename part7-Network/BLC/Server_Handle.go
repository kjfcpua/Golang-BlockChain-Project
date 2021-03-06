package BLC

import (
	"encoding/gob"
	"bytes"
	"log"
)

/*
	处理version命令
	1.根据本地区块高度以及版本信息判断后续操作
	本地高度>对方高度 -> 向对方发送本地的version命令消息
	对方高度>本地高度 -> 向对方请求对方的区块链信息
 */
func handleVersion(request []byte, bc *Blockchain) {

	//1.从request中获取版本的数据：[]byte
	commandBytes := request[COMMAND_LENGTH:]

	//2.反序列化--->version
	var version Version

	decoder := gob.NewDecoder(bytes.NewReader(commandBytes))

	err := decoder.Decode(&version)
	if err != nil {
		log.Panic(err)
	}

	//3.操作bc，获取自己的最后block的height
	height := bc.GetBestHeight()
	foreignerBestHeight := version.BestHeight

	//4.根对方的比较, 相同则不做操作
	if height > foreignerBestHeight {
		//当前节点比对方节点高度高
		sendVersion(version.AddrFrom, bc)
	} else if foreignerBestHeight > height {
		//当前节点比对方节点高度低,向对方节点请求对方节点的blockchain hash集
		sendGetBlocksHash(version.AddrFrom)
	}

}

/*
	处理getblocks命令
	向对方发送本地的区块链hash集
 */
func handleGetBlocksHash(request []byte, bc *Blockchain) {
	//1.从request中获取版本的数据：[]byte
	commandBytes := request[COMMAND_LENGTH:]

	//2.反序列化--->version
	var getblocks GetBlocks

	decoder := gob.NewDecoder(bytes.NewReader(commandBytes))

	err := decoder.Decode(&getblocks)
	if err != nil {
		log.Panic(err)
	}

	blocksHashes := bc.getBlocksHashes()

	sendInv(getblocks.AddrFrom, BLOCK_TYPE, blocksHashes)
}

/*
	处理Inv命令
	1. block type :  如果本地区块

 */
func handleInv(request []byte, bc *Blockchain) {
	//1.从request中获取版本的数据：[]byte
	commandBytes := request[COMMAND_LENGTH:]

	//2.反序列化--->version
	var inv Inv

	decoder := gob.NewDecoder(bytes.NewReader(commandBytes))

	err := decoder.Decode(&inv)
	if err != nil {
		log.Panic(err)
	}

	if inv.Type == BLOCK_TYPE {
		//获取hashes中第一个hash,请求对方返回此hash对应的block
		hash := inv.Items[0]
		sendGetData(inv.AddrFrom, BLOCK_TYPE, hash)

		//保存items剩余未请求的hashes到变量blockArray(handleBlockData 方法会用到)
		if len(inv.Items) > 0 {
			blockArray = inv.Items[1:]
		}

	} else if inv.Type == TX_TYPE {

	}
}

func handleGetData(request []byte, bc *Blockchain) {
	//1.从request中获取版本的数据：[]byte
	commandBytes := request[COMMAND_LENGTH:]

	//2.反序列化--->version
	var getData GetData

	decoder := gob.NewDecoder(bytes.NewReader(commandBytes))

	err := decoder.Decode(&getData)
	if err != nil {
		log.Panic(err)
	}

	if getData.Type == BLOCK_TYPE {
		block := bc.GetBlockByHash(getData.Hash)
		sendBlock(getData.AddrFrom, block)
	} else if getData.Type == TX_TYPE {

	}
}

func handleGetBlockData(request []byte, bc *Blockchain) {
	//1.从request中获取版本的数据：[]byte
	commandBytes := request[COMMAND_LENGTH:]

	//2.反序列化--->version
	var getBlockData BlockData

	decoder := gob.NewDecoder(bytes.NewReader(commandBytes))

	err := decoder.Decode(&getBlockData)
	if err != nil {
		log.Panic(err)
	}

	blockBytes := getBlockData.Block
	//block := DeserializeBlock(blockBytes)
	var block Block
	gobDecode(blockBytes, &block)
	//fmt.Println(&block)
	bc.AddBlock(&block)

	if len(blockArray) == 0 {
		utxoSet := UTXOSet{bc}
		utxoSet.ResetUTXOSet()
	}

	if len(blockArray) > 0 {
		hash := blockArray[0]
		sendGetData(getBlockData.AddrFrom, BLOCK_TYPE, hash)
		blockArray = blockArray[1:]
	}

}


