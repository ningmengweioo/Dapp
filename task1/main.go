package main

import (
	"context"
	"crypto/ecdsa"
	"dapp_task1/counter"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("警告：无法加载.env文件，将使用环境变量: %v", err)
	}

	// 获取配置
	privateKey := os.Getenv("PRIVATE_KEY")
	sepoliaUrl := os.Getenv("SEPOLIA_URL")
	accountAddr := os.Getenv("ACCOUNT_ADDR")

	// 验证必要的配置是否存在
	if privateKey == "" || sepoliaUrl == "" || accountAddr == "" {
		log.Fatal("错误：缺少必要的环境配置")
	}

	done := make(chan struct{})

	//区块查询和转账
	go task1_transfer(done, privateKey, sepoliaUrl, accountAddr)

	//合约代码
	//go task2_constract(done, privateKey, sepoliaUrl, accountAddr)

	<-done
	fmt.Println("任务完成，程序退出")
}

// 查询区块和发送交易
func task1_transfer(done chan struct{}, PrivateKey string, sepoliaUrl string, accountAddr string) {
	defer close(done)
	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	// 查询区块
	blocknum := big.NewInt(8472842)
	block, err := client.BlockByNumber(context.Background(), blocknum)
	if err != nil {
		log.Fatal(err)
	}
	// 输出查询结果
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("交易树根哈希: %s (用于验证交易列表完整性)\n", block.TxHash().Hex())
	fmt.Printf("回执树根哈希: %s (用于验证交易结果完整性)\n", block.ReceiptHash().Hex())
	fmt.Printf("父区块哈希: %s (用于验证区块链接)\n", block.ParentHash().Hex())
	fmt.Printf("当前区块高度: %d\n", block.Number().Uint64())
	fmt.Printf("交易数量: %d\n", block.Transactions().Len())
	fmt.Printf("区块时间戳: %d\n", block.Time())

	// 1. 从私钥获取发送方地址
	/*privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 2. 设置接收方地址和转账金额
	toAddress := common.HexToAddress(accountAddr) // 替换为实际的接收方地址
	amount := big.NewInt(10000000000000000)       // in wei (0.01 eth)                                   // 1 ETH = 10^18 wei

	// 3. 获取发送方的nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 4. 设置Gas价格和Gas上限
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	gasPrice := big.NewInt(int64(64579546054))

	//gasLimit := uint64(21000) // 标准转账交易的Gas上限
	gasLimit := uint64(800000)

	// 5. 创建交易
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	// 6. 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 7. 发送交易
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal(err)
	}

	// 8. 输出交易哈希值 0xb50ba091ec1b123e91de3fcfc97f76dcabd4db914ae3a07be9e1bb178c297440
	fmt.Printf("交易哈希值: %s\n", signedTx.Hash().Hex())
	*/
	done <- struct{}{}
}

func task2_constract(done chan struct{}, PrivateKey string, sepoliaUrl string, accountAddr string) {
	defer close(done)

	// 1. 连接到Sepolia测试网络
	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatalf("无法连接到Sepolia测试网络: %v", err)
	}
	fmt.Println("成功连接到Sepolia测试网络")

	// 2. 从私钥创建交易选项
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatalf("无法解析私钥: %v", err)
	}

	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("无法获取链ID: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 创建交易选项
	// auth, err := types.NewKeyedTransactorWithChainID(privateKey, chainID)
	// if err != nil {
	// 	log.Fatalf("无法创建交易选项: %v", err)
	// }

	auth := bind.NewKeyedTransactor(privateKey, chainID)
	// 设置Gas价格和Gas上限
	auth.GasLimit = uint64(3000000) // 部署合约需要更多Gas
	auth.GasPrice = big.NewInt(64579546054)

	// 设置nonce
	//client.PendingNonceAt(context.Background(), fromAddress)
	nounce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("无法获取nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nounce))

	// 3. 使用DeployCounter函数部署智能合约
	fmt.Println("开始部署智能合约...")
	contractAddress, tx, contractInstance, err := counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatalf("部署智能合约失败: %v", err)
	}

	fmt.Printf("智能合约部署成功！\n")
	time.Sleep(10 * time.Second)
	fmt.Printf("合约地址: %s\n", contractAddress.Hex())
	fmt.Printf("部署交易哈希: %s\n", tx.Hash().Hex())

	// 4. 等待交易确认（可选，但建议）
	fmt.Println("等待交易确认...")
	// 注意：在实际应用中，应该等待交易被打包进区块后再进行后续操作
	// 这里可以添加等待交易确认的代码

	// 5. 调用合约的只读方法GetCount()获取当前计数值
	currentCount, err := contractInstance.GetCount(nil)
	if err != nil {
		log.Fatalf("调用GetCount方法失败: %v", err)
	}
	fmt.Printf("调用GetCount成功，当前计数值: %d\n", currentCount)

	// 6. 调用合约的写入方法Increment()增加计数值
	// 更新交易选项的nonce（因为之前已经用了一个nonce部署合约）
	// auth.Nonce, err = client.PendingNonceAt(context.Background(), auth.From)
	// if err != nil {
	// 	log.Fatalf("无法获取最新nonce: %v", err)
	// }
	// fmt.Printf("调用Increment前的nonce值: %d\n", auth.Nonce.Uint64())
	auth.Nonce = new(big.Int).SetUint64(auth.Nonce.Uint64() + 1)
	tx, err = contractInstance.Increment(auth)
	if err != nil {
		log.Fatalf("调用Increment方法失败: %v", err)
	}
	fmt.Printf("调用Increment成功，交易哈希: %s\n", tx.Hash().Hex())

	// 7. 等待交易确认并再次查询计数值
	fmt.Println("等待Increment交易确认...")
	time.Sleep(5 * time.Second)

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, tx.Hash())
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}
	fmt.Printf("交易已确认，区块高度: %d\n", receipt.BlockNumber.Uint64())

	// 再次查询计数值
	newCount, err := contractInstance.GetCount(nil)
	if err != nil {
		log.Fatalf("确认后调用GetCount方法失败: %v", err)
	}
	fmt.Printf("Increment交易确认后，当前计数值: %d\n", newCount)

	done <- struct{}{}
}
