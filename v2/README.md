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
import tonlib "github.com/mercuryoio/tonlib-go/v2"
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
- [-] wallet.getAccountState
- [-] wallet.sendGrams
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
    options, err := tonlib.ParseConfigFile("path/to/config.json")
    if err != nil {
        panic(err)
    }

    // make req
    req := tonlib.TonInitRequest{
        "init",
        *options,
    }

    tonClient, err = tonlib.NewClient(
    	&req, // init request
    	tonlib.Config{}, // config
    	10, // timeout in seconds for each (currently only QueryEstimateFees) tonlib.Client`s public method
    	true, // enable client`s logs
    	9, // logging level in ton lib.
    )
    if err != nil {
        panic(err)
    }
    defer cln.Destroy()
```
### Create new private key
```go
    // prepare data
    loc := SecureBytes("loc_pass")
    mem := SecureBytes("mem_pass")
    seed := SecureBytes("")

    // create new key
    pKey, err := cln.CreateNewKey(&loc, &mem, &seed)
    if err != nil {
       panic(err)
    }
```
### Get wallet address
```go
    addrr, err := cln.WalletGetAccountAddress(tonlib.NewWalletInitialAccountState("YourPublicKey"))
    if err != nil {
        panic(err)
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
