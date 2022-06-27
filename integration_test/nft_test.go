package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestNft() {
	cases := []SubTest{
		{
			"MintNFT",
			MintNFT,
		},
		{
			"MintNFTWithEthPubKey",
			MintNFTWithEthPubKey,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func MintNFT(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0aae040aab040a172f697269736d6f642e6e66742e4d73674d696e744e4654128f040a166e6674313530393934343436313831313039333530361204686568651a12e88095e5b0b8e6978fe7868ae5b08fe59ba22218687474703a2f2f7777772e6875696665696a7579612e636e2ae8027b2264617461223a7b226f776e6572223a22e590bee69cace7868ae58583e5ae87e5ae99222c22746f6b656e436f6465223a2238333834222c22636f6465223a2236506e7759756d307a525a354c7237477369714c74754432736e6c5476744775222c22707265666978223a22776278222c22736f7274223a312c22696d6755726c223a2268747470733a2f2f6e66746d616c6c2e6f73732d636e2d6368656e6764752e616c6979756e63732e636f6d2f32303232303430312f583667314c73704b4b517766464f45397378383338342e706e67222c226170706964223a22686568653230323230333331777562656e78696f6e67222c226e616d65223a22e88095e5b0b8e6978fe7868ae5b08fe59ba2222c226964223a313530393934343436313831313039333530362c22617474726962757465223a22227d2c227265736f75726365223a22776278222c226964223a2231353039393434343631383131303933353036227d322a696161316678743030746a37356665776663756b6437333232373466336a3276756d617270326c6334323a2a696161316678743030746a37356665776663756b6437333232373466336a3276756d617270326c63343212640a4c0a400a192f636f736d6f732e63727970746f2e736d322e5075624b657912230a21022c0db0445d53bcce38224bf11137ce3fdc6b438019fb8f0b3153a0cd03797f9612040a02080118b5b20512140a0e0a0475676173120632303030303010c09a0c1a408f9c4dfb1d5ecc27ca7f321da446ad82b4b8d060e13d50a2af7b81b0bfa468f4bcdfdce29a69ed250f801d31af77bc4c6a7402771307118cc6e8677eed5b77d7")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func MintNFTWithEthPubKey(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0add010ada010a172f697269736d6f642e6e66742e4d73674d696e744e465412be010a106e6674706e706c343437343235333034120e796e7363343431353231313435381a15e4b8ade69687e5908de5ad97e58fafe4bba5e4b988221d687474703a2f2f7777772e796f756e6973686f7563616e672e636f6d2f2a0ce8bf99e698afe695b0e68dae322a69616131686d326376397a633865796a68706c6772643565326c6175357037736336796a397872717a343a2a69616131686d326376397a633865796a68706c6772643565326c6175357037736336796a397872717a3412710a590a4f0a282f65746865726d696e742e63727970746f2e76312e657468736563703235366b312e5075624b657912230a2102189b27bd97bdcf021b2d15505e39ba7bbb8c0e1d2c95c5461d506f896f3953d612040a020801180112140a0e0a0475676173120632303030303010c09a0c1a41a590a823ad0e6b01cb344ffed9362a2b053461923f61011a3caf589aea38cfff403c3a59cb3125b8583ef03fc126f2ff54685569c0dd0a454cc06968e52ea2621b")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
