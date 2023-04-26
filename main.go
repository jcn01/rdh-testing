package main

import (
	"context"
	"fmt"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	crypto "github.com/proximax-storage/go-xpx-crypto"
)

var (
	baseUrl             = "http://13.229.210.71:3000"
	xpxHolderPrivateKey = "3C7C91E82BF69B206A523E64DB21B07598834970065ABCFA2BB4212138637E0B"
	nodePublicKey       = "E92978122F00698856910664C480E8F3C2FDF0A733F42970FBD58A5145BD6F21"
	minHarvesterBalance = 100000000000
	client              *sdk.Client
	xpxHolderAcc        *sdk.Account
	nodeAcc             *sdk.PublicAccount
)

func main() {
	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		panic(err)
	}

	client = sdk.NewClient(nil, conf)

	xpxHolderAcc, _ = client.NewAccountFromPrivateKeyAndVersion(xpxHolderPrivateKey, 1)
	nodeAcc, _ = client.NewAccountFromPublicKey(nodePublicKey)
	// fmt.Println("xpxHolderAcc", xpxHolderAcc.Address.Address, xpxHolderAcc.PublicAccount.PublicKey, xpxHolderAcc.PrivateKey)
	// fmt.Println("nodeAcc", nodeAcc.Address.Address, nodeAcc.PublicKey)

	// Generate accounts
	// customerAcc1, _ := client.NewAccountFromVersion(2)
	// customerAccRemote1, _ := client.NewAccountFromVersion(2)
	// customerAccVrf1, _ := client.NewAccountFromVersion(2)
	// fmt.Println("customerAcc", customerAcc1.Address.Address, customerAcc1.PublicAccount.PublicKey, customerAcc1.PrivateKey)
	// fmt.Println("customerAccRemote", customerAccRemote1.Address.Address, customerAccRemote1.PublicAccount.PublicKey, customerAccRemote1.PrivateKey)
	// fmt.Println("customerAccVrf", customerAccVrf1.Address.Address, customerAccVrf1.PublicAccount.PublicKey, customerAccVrf1.PrivateKey)
	
	// customerAcc2, _ := client.NewAccountFromVersion(2)
	// customerAccRemote2, _ := client.NewAccountFromVersion(2)
	// customerAccVrf2, _ := client.NewAccountFromVersion(2)
	// fmt.Println("customerAcc", customerAcc2.Address.Address, customerAcc2.PublicAccount.PublicKey, customerAcc2.PrivateKey)
	// fmt.Println("customerAccRemote", customerAccRemote2.Address.Address, customerAccRemote2.PublicAccount.PublicKey, customerAccRemote2.PrivateKey)
	// fmt.Println("customerAccVrf", customerAccVrf2.Address.Address, customerAccVrf2.PublicAccount.PublicKey, customerAccVrf2.PrivateKey)

	// customerAcc3, _ := client.NewAccountFromVersion(2)
	// customerAccRemote3, _ := client.NewAccountFromVersion(2)
	// customerAccVrf3, _ := client.NewAccountFromVersion(2)
	// fmt.Println("customerAcc", customerAcc3.Address.Address, customerAcc3.PublicAccount.PublicKey, customerAcc3.PrivateKey)
	// fmt.Println("customerAccRemote", customerAccRemote3.Address.Address, customerAccRemote3.PublicAccount.PublicKey, customerAccRemote3.PrivateKey)
	// fmt.Println("customerAccVrf", customerAccVrf3.Address.Address, customerAccVrf3.PublicAccount.PublicKey, customerAccVrf3.PrivateKey)

	// customerAcc4, _ := client.NewAccountFromVersion(2)
	// customerAccRemote4, _ := client.NewAccountFromVersion(2)
	// customerAccVrf4, _ := client.NewAccountFromVersion(2)
	// fmt.Println("customerAcc", customerAcc4.Address.Address, customerAcc4.PublicAccount.PublicKey, customerAcc4.PrivateKey)
	// fmt.Println("customerAccRemote", customerAccRemote4.Address.Address, customerAccRemote4.PublicAccount.PublicKey, customerAccRemote4.PrivateKey)
	// fmt.Println("customerAccVrf", customerAccVrf4.Address.Address, customerAccVrf4.PublicAccount.PublicKey, customerAccVrf4.PrivateKey)


	// fmt.Println("======1st======")
	customerAcc1, _ := client.NewAccountFromPrivateKeyAndVersion("de8cb81cc7b23efe98891a255f06e3cd8a8de609a79a019aae31be7cdef498fa", 2)
	customerAccRemote1, _ := client.NewAccountFromPrivateKeyAndVersion("7cf77f22ea868a42165c4b4aa137469ae34cad774191db93f2fe1be3f5341be0", 2)
	customerAccVrf1, _ := client.NewAccountFromPrivateKeyAndVersion("67ec49008f6d6a5a137d0a26f53ee72fef2b75669131018c12bf8361c929c419", 2)

	// amount := uint64(minHarvesterBalance + 100000000)
	// AnnounceXPXTransferTransaction(client, customerAcc1, xpxHolderAcc, amount)
	// time.Sleep(time.Second * 30)

	AnnounceAccountLink(client, customerAcc1, customerAccRemote1, sdk.AccountLink)
	AnnounceVrfKeyLink(client, customerAcc1, customerAccVrf1, sdk.AccountLink)
	AnnounceNodeLink(client, customerAcc1, nodeAcc, sdk.AccountLink)
	time.Sleep(time.Second * 30)
	AnnounceTransferMessage(client, xpxHolderAcc, customerAccRemote1, customerAccVrf1)
	time.Sleep(time.Second * 30)
	GetNodeUnlockedAccounts(client)

	// fmt.Println("======2nd======")
	// customerAcc2, _ := client.NewAccountFromPrivateKeyAndVersion("a343c5aa6bbd87b9735834325f4b362f1e597bfd512812df2a727e9c08ed8cc8", 2)
	// customerAccRemote2, _ := client.NewAccountFromPrivateKeyAndVersion("5a83e24d422b261da53756a961a3762aa7085c77fcca08ef73d1490cbb8e33ae", 2)
	// customerAccVrf2, _ := client.NewAccountFromPrivateKeyAndVersion("11b6d19d195808b170ee75cac1aa74d51a4481a94943379decd8d4e44378ce9e", 2)

	// amount2 := uint64(minHarvesterBalance + 250000000)
	// AnnounceXPXTransferTransaction(client, customerAcc2, xpxHolderAcc, amount2)
	// time.Sleep(time.Second * 30)

	// AnnounceAccountLink(client, customerAcc2, customerAccRemote2, sdk.AccountUnlink)
	// AnnounceVrfKeyLink(client, customerAcc2, customerAccVrf2, sdk.AccountUnlink)
	// AnnounceNodeLink(client, customerAcc2, nodeAcc, sdk.AccountUnlink)
	// time.Sleep(time.Second * 30)
	// AnnounceTransferMessage(client, customerAcc2, customerAccRemote2, customerAccVrf2)
	// time.Sleep(time.Second * 30)
	// GetNodeUnlockedAccounts(client)

	// fmt.Println("======3rd======")
	// customerAcc3, _ := client.NewAccountFromPrivateKeyAndVersion("ffeff6e41c773fadcfca047a71e73b65fce8c0832c9c1099a9c1badce652525e", 2)
	// customerAccRemote3, _ := client.NewAccountFromPrivateKeyAndVersion("1058576da885e804e40bb219aaab1c6564cff334548dcf0ca375e584da3ba7ed", 2)
	// customerAccVrf3, _ := client.NewAccountFromPrivateKeyAndVersion("43294530614b14cbd222dfcac2fedf4e9824ea8a07e79702c8f51e8f2d70e59b", 2)

	// amount3 := uint64(minHarvesterBalance + 500000000)
	// AnnounceXPXTransferTransaction(client, customerAcc3, xpxHolderAcc, amount3)
	// time.Sleep(time.Second * 30)

	// AnnounceAccountLink(client, customerAcc3, customerAccRemote3, sdk.AccountUnlink)
	// AnnounceVrfKeyLink(client, customerAcc3, customerAccVrf3, sdk.AccountUnlink)
	// AnnounceNodeLink(client, customerAcc3, nodeAcc, sdk.AccountUnlink)
	// time.Sleep(time.Second * 30)
	// AnnounceTransferMessage(client, customerAcc3, customerAccRemote3, customerAccVrf3)
	// time.Sleep(time.Second * 45)
	// GetNodeUnlockedAccounts(client)

	// fmt.Println("====4th===")
	// customerAcc4, _ := client.NewAccountFromPrivateKeyAndVersion("2a21244180b181a4d1707f5b718e67c56447447ccdd0765a58f99a223c42ef6e", 2)
	// customerAccRemote4, _ := client.NewAccountFromPrivateKeyAndVersion("c1e64ffc23ef58e63c4e7be4e0affb2361e4779b8dbc232debb9807fd9008bc4", 2)
	// customerAccVrf4, _ := client.NewAccountFromPrivateKeyAndVersion("014aa020943abe54443e89f3f9892d334bdf8a5192b28595a081b936303578ab", 2)

	// amount4 := uint64(minHarvesterBalance + 750000000)
	// AnnounceXPXTransferTransaction(client, customerAcc4, xpxHolderAcc, amount4)
	// time.Sleep(time.Second * 30)

	// AnnounceAccountLink(client, customerAcc4, customerAccRemote4, sdk.AccountUnlink)
	// AnnounceVrfKeyLink(client, customerAcc4, customerAccVrf4, sdk.AccountUnlink)
	// AnnounceNodeLink(client, customerAcc4, nodeAcc, sdk.AccountUnlink)
	// time.Sleep(time.Second * 30)
	// AnnounceTransferMessage(client, customerAcc4, customerAccRemote4, customerAccVrf4)
	// time.Sleep(time.Second * 45)
	// GetNodeUnlockedAccounts(client)
}

// minHarvesterBalance = 100'000'000'000
func AnnounceXPXTransferTransaction(client *sdk.Client, customerAcc *sdk.Account, xpxHolderAcc *sdk.Account, amount uint64) {
	transferTransaction, err := client.NewTransferTransaction(
		sdk.NewDeadline(time.Hour),
		customerAcc.Address,
		[]*sdk.Mosaic{sdk.Xpx(amount)},
		sdk.NewPlainMessage("transfer xpx"),
	)
	if err != nil {
		panic(fmt.Errorf("transfer xpx transaction returned error: %s", err))
	}

	signedTx, err := xpxHolderAcc.Sign(transferTransaction)
	if err != nil {
		panic(fmt.Errorf("transfer xpx transaction signing returned error: %s", err))
	}

	restTx, err := client.Transaction.Announce(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("XpxTransfer TxHash: %s\n", restTx)

	time.Sleep(time.Second * 30)
	tx, _ := client.Transaction.GetTransaction(context.Background(), sdk.Confirmed, restTx)
	fmt.Println(tx)
}

// main - remote : AccountKeyLinkTransaction
func AnnounceAccountLink(client *sdk.Client, customerAcc *sdk.Account, customerAccRemote *sdk.Account, linkAction sdk.AccountLinkAction) {
	accountKeyLinkTransaction, _ := client.NewAccountLinkTransaction(
		sdk.NewDeadline(time.Hour*1),
		customerAccRemote.PublicAccount,
		linkAction)

	signedAccountLinKTransaction, err := customerAcc.Sign(accountKeyLinkTransaction)

	if err != nil {
		panic(fmt.Errorf("account link transaction signing returned error: %s", err))
	}
	restTx, err := client.Transaction.Announce(context.Background(), signedAccountLinKTransaction)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AccountKeyLink TxHash: %s\n", restTx)
}

// main - vrf: VrfKeyLinkTransaction
func AnnounceVrfKeyLink(client *sdk.Client, customerAcc *sdk.Account, customerAccVrf *sdk.Account, linkAction sdk.AccountLinkAction) {
	vrfKeyLinkTransaction, _ := client.NewVrfKeyLinkTransaction(
		sdk.NewDeadline(time.Hour*1),
		customerAccVrf.PublicAccount,
		linkAction)

	signedVrfKeyLinkTransaction, err := customerAcc.Sign(vrfKeyLinkTransaction)
	if err != nil {
		panic(fmt.Errorf("vrf Key Link transaction signing returned error: %s", err))
	}

	restTx, err := client.Transaction.Announce(context.Background(), signedVrfKeyLinkTransaction)
	if err != nil {
		panic(fmt.Errorf("cannot announce vrf key link: %s", err))
	}
	fmt.Printf("VrfKeyLink TxHash: %s\n", restTx)
}

// main - node : NodeKeyLinkTransaction
func AnnounceNodeLink(client *sdk.Client, customerAcc *sdk.Account, nodeAcc *sdk.PublicAccount, linkAction sdk.AccountLinkAction) {
	nodeLinkTransaction, _ := client.NewNodeKeyLinkTransaction(
		sdk.NewDeadline(time.Hour*1),
		nodeAcc.PublicKey,
		linkAction)

	signedNodeLinkTransaction, err := customerAcc.Sign(nodeLinkTransaction)

	if err != nil {
		panic(fmt.Errorf("node link transaction signing returned error: %s", err))
	}

	restTx, err := client.Transaction.Announce(context.Background(), signedNodeLinkTransaction)
	if err != nil {
		panic(fmt.Errorf("cannot announce node link: %s", err))
	}
	fmt.Printf("NodeLink TxHash: %s\n", restTx)
}

// remote - node :  Persistent Delegation Request Transaction
func AnnounceTransferMessage(client *sdk.Client, customerAcc *sdk.Account, customerAccRemote *sdk.Account, customerAccVrf *sdk.Account) {
	nodePubKey, err := crypto.NewPublicKeyfromHex("E92978122F00698856910664C480E8F3C2FDF0A733F42970FBD58A5145BD6F21")
	if err != nil {
		panic(err)
	}
	message, _ := sdk.NewPersistentHarvestingDelegationMessageFromPlainText(customerAccRemote.PrivateKey, customerAccVrf.PrivateKey, nodePubKey)
	persistentDelegationLinkTransaction, _ := client.NewTransferTransaction(
		sdk.NewDeadline(time.Hour*1),
		nodeAcc.Address,
		[]*sdk.Mosaic{},
		message)
	signedPersistentDelegationLinkTransaction, err := customerAcc.Sign(persistentDelegationLinkTransaction)

	if err != nil {
		panic(fmt.Errorf("transfer message transaction signing returned error: %s", err))
	}

	restTx, err := client.Transaction.Announce(context.Background(), signedPersistentDelegationLinkTransaction)
	if err != nil {
		panic(fmt.Errorf("transfer message transaction announcing returned error: %s", err))
	}
	fmt.Printf("PersistentDelegationLink TxHash: %s\n", restTx)
}

func GetNodeUnlockedAccounts(client *sdk.Client) {
	key, err := client.Node.GetNodeUnlockedAccounts(context.Background())
	if err != nil {
		panic(fmt.Errorf("cannot retrieve unlocked accounts: %s", err))
	}
	fmt.Printf("%v\n", key)
}
