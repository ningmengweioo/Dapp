package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
本地验证私链
*/
func main() {
	// 连接到本地运行的以太坊节点，默认HTTP-RPC端口为8545
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("无法连接到以太坊节点: %v", err)
	}
	fmt.Println("成功连接到以太坊节点")
	defer client.Close()

	// 1. 获取最新区块号
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("获取区块号失败: %v", err)
	}
	fmt.Printf("最新区块号: %d\n", blockNumber)

	// 2. 获取最新区块详细信息
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatalf("获取区块信息失败: %v", err)
	}
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("区块大小: %d bytes\n", block.Size())
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	fmt.Printf("区块时间戳: %d\n", block.Time())

	// 3. 检查指定账户的余额
	// 这里使用一个示例地址，您可以替换为您节点上的实际账户地址
	account := common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("获取账户余额失败: %v", err)
	}
	// 将余额从Wei转换为ETH
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1000000000000000000))
	fmt.Printf("账户 %s 余额: %s ETH\n", account.Hex(), ethBalance.Text('f', 18))

	// 4. 获取建议的Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("获取Gas价格失败: %v", err)
	}
	fmt.Printf("建议的Gas价格: %s Wei\n", gasPrice.String())

	// 5. 检查节点是否在同步状态
	syncing, err := client.SyncProgress(context.Background())
	if err != nil {
		log.Fatalf("检查同步状态失败: %v", err)
	}
	if syncing != nil {
		fmt.Printf("节点正在同步 - 当前区块: %d, 最高区块: %d\n", syncing.CurrentBlock, syncing.HighestBlock)
	} else {
		fmt.Println("节点已同步完成")
	}
}
