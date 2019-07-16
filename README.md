# go-planetlab
[![Build Status](https://travis-ci.org/axelniklasson/go-planetlab.svg?branch=dev)](https://travis-ci.org/axelniklasson/go-planetlab)


Go client library for accessing the [PlanetLab API](https://www.planet-lab.org/doc/plc_api). Note that this library only provides wrappers for the methods *not* requiring admin or PI access to the API. Furthermore, no delete methods have been implemented aside from `DeleteKey`, since deletion may have huge consequences.

## Usage
Simply import `"github.com/axelniklasson/go-planetlab/planetlab"`, instantiate an Auth object and start calling methods. Basic example can be found below.

```golang
package main

import (
  "fmt"
  "log"

  "github.com/axelniklasson/go-planetlab/planetlab"
)

func main() {
  auth := planetlab.GetClientAuth("username", "password")
  sliceFilter := struct{
    Name string `xmlrpc:"name"`  
  }{"chalmersple_2018_10_29"}

  slices, err := planetlab.GetSlices(auth, sliceFilter)
  if err != nil {
    log.Fatal(err)
  }

  // maybe do something more interesting with slices than printing it
  fmt.Println(slices[0].SliceID) // Output: 1578
}

```
