package cpe_test

import (
	"testing"

	"github.com/randolphcyg/cpe"
	"github.com/stretchr/testify/assert"
)

func TestParseCPE22_1(t *testing.T) {
	cpeString := "cpe:/a:hiox_india:guest_book:4.0/"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "a", c.Part)
	assert.Equal(t, "hiox_india", c.Vendor)
	assert.Equal(t, "guest_book", c.Product)
	assert.Equal(t, "4.0", c.Version)
	assert.Equal(t, "", c.Update)
	assert.Equal(t, "", c.Edition)
	assert.Equal(t, "", c.Language)
}

func TestParseCPE22_2(t *testing.T) {
	cpeString := "cpe:/a:oracle:connector/j:5.1.27"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "a", c.Part)
	assert.Equal(t, "oracle", c.Vendor)
	assert.Equal(t, "connector/j", c.Product)
	assert.Equal(t, "5.1.27", c.Version)
	assert.Equal(t, "", c.Update)
	assert.Equal(t, "", c.Edition)
	assert.Equal(t, "", c.Language)
}

func TestParseCPE23_1(t *testing.T) {
	// TODO err str handle
	cpeString := "cpe:2.3:a:disney:where\\'s_my_perry?_free:1.5.1:*:*:*:*:android:*:*"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "a", c.Part)
	assert.Equal(t, "disney", c.Vendor)
	assert.Equal(t, "where's_my_perry?_free", c.Product)
	assert.Equal(t, "1.5.1", c.Version)
	assert.Equal(t, "android", c.TargetSw)
}

func TestParseCPE23_2(t *testing.T) {
	cpeString := "cpe:/a:adobe:flash_player:::~~~chrome~~"
	c, err := cpe.ParseCPE(cpeString)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "a", c.Part)
	assert.Equal(t, "adobe", c.Vendor)
	assert.Equal(t, "flash_player", c.Product)
	assert.Equal(t, "", c.Version)
	assert.Equal(t, "", c.Update)
	assert.Equal(t, "", c.Edition)
	assert.Equal(t, "", c.Language)

	assert.Equal(t, "chrome", c.TargetSw)
	assert.Equal(t, "", c.TargetHw)
	assert.Equal(t, "", c.SwEdition)
	assert.Equal(t, "", c.Other)
}

func TestRegionMatchsV22(t *testing.T) {
	cpeString := "cpe:2.3:a:disney:where\\'s_my_perry?_free:1.5.1:*:*:*:*:android:*:*"
	ret := cpe.RegionMatches(cpeString, false, 0, "cpe:2.3:", 0, 8)
	assert.Equal(t, true, ret)
}

func TestRegionMatchsV23(t *testing.T) {
	cpeString := "cpe:/a:adobe:flash_player:::~~~chrome~~"
	ret := cpe.RegionMatches(cpeString, false, 0, "cpe:/", 0, 5)
	assert.Equal(t, true, ret)
}

func TestInvalidCpe22Part(t *testing.T) {
	cpeString := "cpe:/t:vendor:product:1.0"

	_, err := cpe.ParseCPE(cpeString)
	assert.Equal(t, cpe.ErrInvalidPart, err)
}

func TestInvalidCpe23Part(t *testing.T) {
	cpeString := "cpe:/t:adobe:flash_player:::~~~chrome~~"

	_, err := cpe.ParseCPE(cpeString)
	assert.Equal(t, cpe.ErrInvalidPart, err)
}
