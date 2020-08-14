package main

import (
	"github.com/ont-bizsuite/ddxf-api-sdk"
	"fmt"
	"github.com/ont-bizsuite/ddxf-sdk"
	"github.com/zhiqiangxu/ddxf-api/pkg/common"
	"io"
	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/token"
)

const (
	MainNet  = "http://dappnode1.ont.io:20336"
	TestNet  = "http://polaris2.ont.io:20336"
	LocalNet = "http://127.0.0.1:20336"
)


func main() {

	ddxfContractSdk := ddxf_sdk.NewDdxfSdk(TestNet)
	wallet,_ := ddxfContractSdk.GetOntologySdk().OpenWallet("./wallet.dat")
	pwd := []byte("111111")

	id := common.GenerateOntId()
	acc,_ := wallet.NewDefaultSettingAccount(pwd)

	tokenSDK := sdk.NewTokenSdk("http://127.0.0.1:8080",TestNet,nil)

	createTokenTemplate := true
	if createTokenTemplate {
		input := token.CreateTokenTemplateInput{
		}
		tokenTemplateId, err := tokenSDK.CreateTokenTemplate(id, acc, input, acc)
		if err != nil {
			fmt.Println("CreateTokenTemplate failed:", err)
			return
		}
	}
}
