# TONLIB Golang library
![](https://github.com/mercuryoio/tonlib-go/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mercuryoio/tonlib-go)](https://goreportcard.com/report/github.com/mercuryoio/tonlib-go) 
[![GoDoc](https://godoc.org/github.com/mercuryoio/tonlib-go?status.svg)](https://godoc.org/github.com/mercuryoio/tonlib-go) 



TONLIB Golang library for accessing [Telegram Open Network](https://test.ton.org) with liteclient protocol, which is based itself on [tdlib](https://github.com/tdlib/td) library.
**Warning:** this repository is under active development, not ready for production use
## Install
```sh
$ go get -u github.com/mercuryoio/tonlib-go
```
## Usage
```go
import "github.com/mercuryoio/tonlib-go"
```
## Supported methods
- [x] createNewKey
- [x] deleteKey
- [x] exportKey
- [x] exportPemKey
- [x] exportEncryptedKey
- [x] importKey
- [x] importPemKey
- [x] importEncryptedKey
- [x] changeLocalPassword
- [x] unpackAccountAddress
- [x] packAccountAddress
- [x] wallet.init
- [x] wallet.getAccountAddress
- [x] wallet.getAccountState
- [x] wallet.sendGrams
- [x] raw.sendMessage
- [x] raw.getTransactions
- [x] raw.getAccountState
- [x] generic.sendGrams
- [x] getLogStream
- [x] sync
- [x] CreateAndSendMessage
- [ ] generic.createSendGramsQuery
- [ ] query.send
- [ ] query.forge
- [ ] query.estimateFees
- [ ] query.getInfo
- [ ] smc.load
- [ ] smc.getCode
- [ ] smc.getData
- [ ] smc.getState
- [ ] smc.runGetMethod

## Examples
Create new client 
```go
    cln, err := NewClient(getTestConfig(), Config{})
    if err != nil {
        t.Errorf("Init client error: %v. ", err)
    }
    defer cln.Destroy()
```
### Create new private key
```go
    _, err = cln.CreatePrivateKey([]byte(TEST_PASSWORD))
    if err != nil {
        t.Errorf("Ton create key error: %v. ", err)
    }
```
### Get wallet address
```go
    _, err = cln.WalletGetAddress(pKey.PublicKey)
    if err != nil {
        t.Errorf("Ton get wallet address error: %v. ", err)
    }
```
## CLI:
To install sample cli application:
```sh
$ go get -u github.com/mercuryoio/tonlib-go/cmd/tongo
```
To run sample cli app your have to set LD_LIBRARY_PATH:

For linux `export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:<path2repository>/lib/linux`

For MacOS `export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:<path2repository>/lib/darwin`
## Code generation from new *.tl files released by TON team
If you need to update structures and add new methods based on a fresh release of TON`s client you can do it by using code
 generation command. In order to perform such operation - run the command bellow and provide path of *.tl file to the running command 
 as in the example bellow. 
```sh
$ go run github.com/mercuryoio/tonlib-go/cmd/tlgenerator /path/to/repos/ton/tl/generate/scheme/tonlib_api.tl
```
## Developers
[Mercuryo.io](https://mercuryo.io)
## Contribute
PRs are welcome!
