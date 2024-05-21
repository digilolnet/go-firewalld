# go-firewalld

[![Go report](https://goreportcard.com/badge/github.com/digilolnet/go-firewalld)](https://goreportcard.com/report/github.com/digilolnet/go-firewalld)
[![GoDoc](https://godoc.org/github.com/digilolnet/go-firewalld?status.svg)](https://godoc.org/github.com/digilolnet/go-firewalld)
[![License](https://img.shields.io/github/license/digilolnet/go-firewalld.svg)](https://github.com/digilolnet/go-firewalld/blob/master/LICENSE.md)

Go wrapper for firewalld D-Bus interface.

[![Digilol offers managed hosting and software development](https://www.digilol.net/banner-hosting-development.png)](https://www.digilol.net)

## Usage

```go
package main

import (
	"log"

	"github.com/digilolnet/go-firewalld"
)

func main() {
	fw, err := firewalld.NewFirewalldClient()
	if err != nil {
		log.Println(err)
	}

	// Add new chain using permanent direct rule
	if err := fw.DirectAddChainPermanent("ipv4", "filter", "MYCHAIN"); err != nil {
		log.Println(err)
	}

	// Handle packets related to processes with owner UID 1000 at MYCHAIN
	if err := fw.DirectAddRulePermanent("ipv4", "filter", "OUTPUT", 0, "-m owner --uid-owner 1000 -j MYCHAIN"); err != nil {
		log.Println(err)
	}

	if err := fw.DirectAddRulePermanent("ipv4", "filter", "MYCHAIN", 0, "-m state --state ESTABLISHED,RELATED -j ACCEPT"); err != nil {
		log.Println(err)
	}

	if err := fw.DirectAddRulePermanent("ipv4", "filter", "MYCHAIN", 0, "-p udp --dport 53 -j ACCEPT"); err != nil {
		log.Println(err)
	}
	
	if err := fw.DirectAddRulePermanent("ipv4", "filter", "MYCHAIN", 0, "-j REJECT"); err != nil {
		log.Println(err)
	}

	// Reload for changes to take effect immediately
	if err := fw.Reload(); err != nil {
		log.Println(err)
	}
}
```
