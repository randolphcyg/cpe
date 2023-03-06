package cpe_test

import (
	"fmt"
	"testing"

	"github.com/randolphcyg/cpe"
)

func TestParseCPE22(t *testing.T) {
	cpeString := "cpe:/a:teamspeak:teamspeak2:2.0.23.19:tes /t:test2/"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	fmt.Println("Part:", c.Part)
	fmt.Println("Vendor:", c.Vendor)
	fmt.Println("Product:", c.Product)
	fmt.Println("Version:", c.Version)
	fmt.Println("Update:", c.Update)
	fmt.Println("Edition:", c.Edition)
	fmt.Println("Language:", c.Language)
}

func TestParseCPE23(t *testing.T) {
	//cpeString := "cpe:/a:teamspeak:teamspeak2:2.0.23.19:tes /t:test2/"
	cpeString := "cpe:2.3:o:microsoft:windows_server_2012:r2:-:-:*:standard:*:*:*"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	fmt.Println("Part:", c.Part)
	fmt.Println("Vendor:", c.Vendor)
	fmt.Println("Product:", c.Product)
	fmt.Println("Version:", c.Version)
	fmt.Println("Update:", c.Update)
	fmt.Println("Edition:", c.Edition)
	fmt.Println("Language:", c.Language)

	fmt.Println("TargetSw:", c.TargetSw)
	fmt.Println("TargetHw:", c.TargetHw)
	fmt.Println("SwEdition:", c.SwEdition)
	fmt.Println("Other:", c.Other)
}
