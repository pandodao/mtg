package main

import (
	"encoding/base64"
	"testing"

	"github.com/google/uuid"
	"github.com/pandodao/mtg/mtgpack"
	"github.com/pandodao/mtg/protocol"
	"github.com/shopspring/decimal"
)

func TestDecodeMemo(t *testing.T) {
	const memo = "AQEBecIa9NuuSuqFjx6vl4I8swADAQAIrowoFSlDh7MN7WVBRYfkAAAAAAJPkCcGNWRjZmIy"
	b, err := base64.StdEncoding.DecodeString(memo)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(b)
	dec := mtgpack.NewDecoder(b)

	var (
		header protocol.Header
		mmsig  protocol.MultisigReceiver
	)

	if err := mtgpack.DecodeValues(dec, &header, &mmsig); err != nil {
		t.Fatal(err)
	}

	t.Logf("header: %+v", header)
	t.Logf("mmsig: %+v", mmsig)

	switch header.Action {
	case 1:
		var (
			asset    uuid.UUID
			slippage decimal.Decimal
			exp      int16
		)

		if err := mtgpack.DecodeValues(dec, &asset, &slippage, &exp); err != nil {
			t.Fatal(err)
		}

		t.Logf("asset: %s", asset)
		t.Logf("slippage: %s", slippage)
		t.Logf("exp: %d", exp)
	case 3:
		var (
			asset uuid.UUID
			route string
			min   decimal.Decimal
		)

		if err := mtgpack.DecodeValues(dec, &asset, &route, &min); err != nil {
			t.Fatal(err)
		}

		t.Logf("asset: %s", asset)
		t.Logf("route: %s", route)
		t.Logf("min: %s", min)
	}
}
