package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Hash    string `json:"result"`
}

type CollectionMetaType struct {
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
	Src    string      `json:"src"`
}

type PackMetaType struct {
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:config`
	Src    string      `json:"src"`
}

type CardMetaType struct {
	Name   string `json:"name"`
	Tier   string `json:"tier"`
	Rarity string `json:"rarity"`
	Src    string `json:"src"`
}

type CollectionType struct {
	Metadata CollectionMetaType
	Id       int64
	TokenId  *big.Int
	Address  common.Address
	Packs    *[]PackType
}

type PackType struct {
	Metadata PackMetaType
	Id       int64
	TokenId  *big.Int
	Address  common.Address
	Cards    *[]CardType
}

type CardType struct {
	Metadata     CardMetaType
	Id           int64
	TokenId      *big.Int
	Type         string
	Address      common.Address
	CollectionId int64
	PackId       int64
}

type PackConfig []int64
type Packs []PackConfig

var (
	rpc       = "https://polygon-mumbai.g.alchemy.com/v2/PTi4_upQduxYasRED-Hy1EBldW7pBstt"
	walletKey = "fbc533ef71a592206fb846e1ce4f2ca6ee3cb08388d13b2ff555745c28eb0c74"

	value    = big.NewInt(0)
	gasLimit = uint64(20_000_000)
	chainId  = big.NewInt(80001)

	ownerAddr = common.HexToAddress("0xf22679BBaf587B9c776D0A25a05e64B214f19CAd")

	collectionAddr          = common.HexToAddress("0x52F1908459370CFa63C8891b1C080334541C57CD")
	packFactoryAddr         = common.HexToAddress("0x3a2d175B8228De11dF7B44959527D0114cEbbe62")
	cateogryFactoryAddr     = common.HexToAddress("0xf474b7B4d946e559f156c7d5d02e8f5Ab994bDB6")
	yearFactoryAddr         = common.HexToAddress("0x687Dcd759C323dBB76D53a9D5893F9D391b93358")
	daymonthFactoryAddr     = common.HexToAddress("0xb040c97359377478D51Bf0b3C5D459FC46250E43")
	craftingCardFactoryAddr = common.HexToAddress("0x3c4Ca73bFA4C4CCa9D67b7D79b23dBdc2703Fe5C")
	triggerFactoryAddr      = common.HexToAddress("0x78BD6D288237b0CEC3fF4Ee95Aa76A70F660A174")
	identityFactoryAddr     = common.HexToAddress("0x3b0e4F34873feD7d5680000D7ae1C4866A75d501")
	predictionFactoryAddr   = common.HexToAddress("0xe27B954f788FEA7C4BA36D12bA705D70017f8Bf3")

	packAddress            = common.HexToAddress("0xEddaC99edca0e45B374834D85a82c40a48C8B5e9")
	categoryAddress        common.Address
	yearAddress            common.Address
	dayMonthAddress        common.Address
	craftingCardAddress    common.Address
	triggerAddress         common.Address
	identityAddress        common.Address
	predictionAddress      common.Address
	collectionCreatedTopic = "0x0ba92c0f760258e3285b70eb8220efa72578d8bb68dfa3e26d48f3cecd2e10b7"

	collectionId = big.NewInt(0)
	PackAmts     = []*big.Int{big.NewInt(300), big.NewInt(200), big.NewInt(100)}

	quantity         = big.NewInt(1)
	standardPackTier = big.NewInt(0)

	packOpenerAddr = common.HexToAddress("0x5D6A8216590cEEeFd149f3CB26ad096B8516bB00")

	collectionInstance          *Collection
	cardPackFactoryInstance     *CardPackFactory
	categoryFactoryInstance     *CategoryFactory
	yearFactoryInstance         *YearFactory
	dayMonthFactoryInstance     *DayMonthFactory
	craftingCardFactoryInstance *CraftingFactory
	triggerFactoryInstance      *TriggerFactory
	identityFactoryInstance     *IdentityFactory
	predictionFactoryInstance   *PredictionFactory

	cardpackInstance     *CardPack
	categoryInstance     *Category
	yearInstance         *Year
	daymonthInstance     *Daymonth
	craftingcardInstance *CraftingCard
	triggerInstance      *Trigger
	PackopenerInstance   *Packopener
	IdentifyInstance     *Identity
	predictionInstance   *Prediction

	cards      []CardType
	pack       PackType
	collection CollectionType

	err error
)

func main() {
	// setCollectionMinterAndAdmin()
	// createCollection()
	getAddresses(collectionId)
	// setMinterAndCrafter()
	// createCard()
	openPack()
	// craftingIdentity()
	// craftingPrediction()
}

func setCollectionMinterAndAdmin() {
	client, _ := ethclient.Dial(rpc)
	var (
		auth *bind.TransactOpts
	)

	collectionInstance, err = NewCollection(collectionAddr, client)
	auth = getTransactor()
	collectionInstance.ChangeMinter(auth, ownerAddr)

	cardPackFactoryInstance, err = NewCardPackFactory(packFactoryAddr, client)
	auth = getTransactor()
	cardPackFactoryInstance.ChangeAdmin(auth, collectionAddr)

	categoryFactoryInstance, err = NewCategoryFactory(cateogryFactoryAddr, client)
	auth = getTransactor()
	categoryFactoryInstance.ChangeAdmin(auth, collectionAddr)

	yearFactoryInstance, err = NewYearFactory(yearFactoryAddr, client)
	auth = getTransactor()
	yearFactoryInstance.ChangeAdmin(auth, collectionAddr)

	dayMonthFactoryInstance, err = NewDayMonthFactory(daymonthFactoryAddr, client)
	auth = getTransactor()
	dayMonthFactoryInstance.ChangeAdmin(auth, collectionAddr)

	craftingCardFactoryInstance, err = NewCraftingFactory(craftingCardFactoryAddr, client)
	auth = getTransactor()
	craftingCardFactoryInstance.ChangeAdmin(auth, collectionAddr)

	triggerFactoryInstance, err = NewTriggerFactory(triggerFactoryAddr, client)
	auth = getTransactor()
	triggerFactoryInstance.ChangeAdmin(auth, collectionAddr)

	identityFactoryInstance, err = NewIdentityFactory(identityFactoryAddr, client)
	auth = getTransactor()
	identityFactoryInstance.ChangeAdmin(auth, collectionAddr)

	predictionFactoryInstance, err = NewPredictionFactory(predictionFactoryAddr, client)
	auth = getTransactor()
	predictionFactoryInstance.ChangeAdmin(auth, collectionAddr)
}
func createCollection() {

	client, _ := ethclient.Dial(rpc)
	auth := getTransactor()

	collectionInstance, err = NewCollection(collectionAddr, client)
	tx, _ := collectionInstance.CreateCollection(auth,
		"https://example.com/collection/0", PackAmts)

	fmt.Printf("create collection tx hash: %s", tx.Hash())

}

func getAddresses(collectionId *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, error) {
	client, _ := ethclient.Dial(rpc)

	var (
		logs []types.Log
	)

	collectionABI, err := CollectionMetaData.GetAbi()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{collectionAddr},
		Topics:    [][]common.Hash{{common.HexToHash(collectionCreatedTopic)}},
	}

	logs, err = client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, err
	}

	for _, vlog := range logs {
		topic := vlog.Topics[0].Hex()
		fmt.Println(topic)
		if topic == collectionCreatedTopic {

			data, _ := collectionABI.Unpack("CollectionCreated", vlog.Data)
			// spew.Dump("topic", data)

			tokenId := data[0].(*big.Int)
			fmt.Println(tokenId, collectionId, tokenId.Cmp(collectionId))
			if tokenId.Cmp(collectionId) == 0 {
				packAddress = data[1].(common.Address)
				categoryAddress = data[2].(common.Address)
				yearAddress = data[3].(common.Address)
				dayMonthAddress = data[4].(common.Address)
				craftingCardAddress = data[5].(common.Address)
				triggerAddress = data[6].(common.Address)
				identityAddress = data[7].(common.Address)
				predictionAddress = data[8].(common.Address)

				fmt.Println("card pack address", packAddress)
				fmt.Println("category address", categoryAddress)
				fmt.Println("year address", yearAddress)
				fmt.Println("daymonth address", dayMonthAddress)
				fmt.Println("trigger address", triggerAddress)
				fmt.Println("craftingcard address", craftingCardAddress)
				fmt.Println("identity address", identityAddress)
				fmt.Println("prediction address", predictionAddress)

				return packAddress, categoryAddress, yearAddress, dayMonthAddress, craftingCardAddress, triggerAddress, identityAddress, predictionAddress, err
			}
		}
	}
	return common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, err
}

func setMinterAndCrafter() {
	client, _ := ethclient.Dial(rpc)
	var (
		auth *bind.TransactOpts
	)

	cardpackInstance, err = NewCardPack(packAddress, client)
	auth = getTransactor()
	cardpackInstance.ChangeMinter(auth, ownerAddr)

	auth = getTransactor()
	cardpackInstance.ChangeOpener(auth, packOpenerAddr)

	categoryInstance, err = NewCategory(categoryAddress, client)
	auth = getTransactor()
	categoryInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	categoryInstance.ChangeCrafter(auth, identityAddress)

	yearInstance, err = NewYear(yearAddress, client)
	auth = getTransactor()
	yearInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	yearInstance.ChangeCrafter(auth, identityAddress)

	daymonthInstance, err = NewDaymonth(dayMonthAddress, client)
	auth = getTransactor()
	daymonthInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	daymonthInstance.ChangeCrafter(auth, identityAddress)

	craftingcardInstance, err = NewCraftingCard(craftingCardAddress, client)
	auth = getTransactor()
	craftingcardInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	craftingcardInstance.ChangeCrafter(auth, identityAddress)
	auth = getTransactor()
	craftingcardInstance.ChangePredictor(auth, predictionAddress)

	triggerInstance, err = NewTrigger(triggerAddress, client)
	auth = getTransactor()
	triggerInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	triggerInstance.ChangeCrafter(auth, predictionAddress)

	IdentifyInstance, err = NewIdentity(identityAddress, client)
	auth = getTransactor()
	IdentifyInstance.ChangeCrafter(auth, predictionAddress)

}

func createCard() {
	var (
		tokenURI = "https://example.com/collecition/0/cardpack/0"
	)
	client, _ := ethclient.Dial(rpc)
	cardpackInstance, err = NewCardPack(packAddress, client)
	auth := getTransactor()
	tx, _ := cardpackInstance.Mint(auth, ownerAddr, quantity, standardPackTier, tokenURI)
	fmt.Printf("card pack tx hash: %s", tx.Hash())
}

func openPack() {
	var (
		cardPackTokenId = big.NewInt(0)

		categoryTokenURIs     = []string{"https://example.com/collection/0/category/1"}
		yearTOkenURIs         = []string{"https://example.com/collection/0/year/1"}
		daymonthTokenURIs     = []string{"https://example.com/collection/0/daymonth/1"}
		craftingCardTokenURIs = []string{
			"https://example.com/collection/0/craftingcard/1",
			"https://example.com/collection/0/craftingcard/2",
			"https://example.com/collection/0/craftingcard/3",
		}
		triggerTokenURIs = []string{"https://example.com/collection/0/trigger/1"}
	)

	client, _ := ethclient.Dial(rpc)
	PackopenerInstance, err = NewPackopener(packOpenerAddr, client)
	auth := getTransactor()

	tx, _ := PackopenerInstance.OpenPack(
		auth,
		cardPackTokenId,
		packAddress,
		categoryAddress,
		yearAddress,
		dayMonthAddress,
		craftingCardAddress,
		triggerAddress,
		categoryTokenURIs,
		yearTOkenURIs,
		daymonthTokenURIs,
		craftingCardTokenURIs,
		triggerTokenURIs,
	)

	fmt.Printf("open pack tx hash: %s", tx.Hash())
}

func craftingIdentity() {
	client, _ := ethclient.Dial(rpc)
	var (
		categoryTokenId = big.NewInt(0)
		dayMonthTokenId = big.NewInt(0)
		yearTokenId     = big.NewInt(0)
		craftingTokenId = big.NewInt(1)

		identityTokenUrl = "https://example.com/collection/0/identity/0"
	)

	IdentifyInstance, err = NewIdentity(identityAddress, client)
	auth := getTransactor()
	tx, _ := IdentifyInstance.CraftCollection(auth,
		dayMonthAddress, yearAddress, categoryAddress, craftingCardAddress,
		dayMonthTokenId, yearTokenId, categoryTokenId, craftingTokenId, identityTokenUrl)

	fmt.Printf("crafting identify tx hash: %s", tx.Hash())
}

func craftingPrediction() {
	client, _ := ethclient.Dial(rpc)
	var (
		identityTokenId    = big.NewInt(0)
		triggerTokenIds    = []*big.Int{big.NewInt(0)}
		craftingTokenId    = big.NewInt(2)
		predictionTokenUrl = "https://example.com/collection/0/prediction/0"
	)

	predictionInstance, err = NewPrediction(predictionAddress, client)
	auth := getTransactor()
	tx, _ := predictionInstance.CraftCollection(auth,
		identityAddress, triggerAddress, craftingCardAddress,
		identityTokenId, triggerTokenIds, craftingTokenId, predictionTokenUrl)

	fmt.Printf("tx sent: %s", tx.Hash())
}

func getTransactor() *bind.TransactOpts {
	client, _ := ethclient.Dial(rpc)

	privatekey, _ := crypto.HexToECDSA(walletKey)

	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth := bind.NewKeyedTransactor(privatekey)
	auth.Nonce = big.NewInt(int64(getNonce(privatekey)))
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(chainId)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privatekey)
		if err != nil {
			fmt.Printf("Error signing: %v\n", err)
			os.Exit(1)
			return nil, nil
		}
		return tx.WithSignature(signer, signature)
	}
	auth.Signer = signfunc

	return auth
}

func getNonce(privatekey *ecdsa.PrivateKey) uint64 {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("from", fromAddress, nonce)

	return nonce
}
