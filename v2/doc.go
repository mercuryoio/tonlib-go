/*
Package v2 is the second version of tonlib-go, it's a simple library which simplifies usage of tonlibjson library from Golang.

Example of client initialization:
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
*/
package v2
