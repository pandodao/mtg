package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/google/uuid"
	"github.com/pandodao/mtg/mtgpack"
	"github.com/pandodao/mtg/protocol"
	"github.com/shopspring/decimal"
)

var (
	number = flag.Int("n", 10, "number of messages to generate")
)

func main() {
	flag.Parse()

	messages := make([]interface{}, 0, *number)

	for i := 0; i < *number; i++ {
		enc := mtgpack.NewEncoder()
		msg := make(map[string]any)

		header := generateHeader()
		if err := mtgpack.EncodeValue(enc, header); err != nil {
			log.Fatalf("encode header failed: %v", err)
		}

		msg["header"] = headerToMap(header)

		mmsig := generateMultisigReceivers()
		if err := mtgpack.EncodeValue(enc, mmsig); err != nil {
			log.Fatalf("encode multisig failed: %v", err)
		}
		msg["mmsig"] = mmsig

		params := generateParams(header.Action)
		paramsMap := make(map[string]any)
		for idx := 0; idx < len(params)-1; idx += 2 {
			name := params[idx].(string)
			value := params[idx+1]

			if err := mtgpack.EncodeValue(enc, value); err != nil {
				log.Fatalf("encode %s failed: %v", name, err)
			}

			paramsMap[name] = value
		}

		msg["params"] = paramsMap

		data := enc.Bytes()
		data = append(data, checksum(data)...)
		msg["memo"] = base64.StdEncoding.EncodeToString(data)
		messages = append(messages, msg)
	}

	if err := json.NewEncoder(os.Stdout).Encode(messages); err != nil {
		log.Fatalf("encode json failed: %v", err)
	}
}

func generateHeader() protocol.Header {
	header := protocol.Header{
		Version:    2,
		ProtocolID: protocol.ProtocolFswap,
	}

	if rand.Int()%2 == 0 {
		header.FollowID = uuid.New()
	}

	header.Action = uint16(rand.Intn(3)) + 1
	return header
}

func headerToMap(h protocol.Header) map[string]any {
	return map[string]any{
		"version":       h.Version,
		"protocol_id":   h.ProtocolID,
		"follow_id":     h.FollowID,
		"has_follow_id": h.FollowID != uuid.Nil,
		"action":        h.Action,
	}
}

func generateMultisigReceivers() protocol.MultisigReceiver {
	r := protocol.MultisigReceiver{
		Version: 1,
		Members: make([]uuid.UUID, 0),
	}

	for i := 0; i < rand.Intn(5); i++ {
		r.Members = append(r.Members, uuid.New())
	}

	if len(r.Members) > 0 {
		r.Threshold = uint8(rand.Intn(len(r.Members))) + 1
	}

	return r
}

func generateParams(action uint16) []interface{} {
	var params []interface{}

	switch action {
	case 1:
		params = append(params, "asset", uuid.New())
		params = append(params, "slippage", decimal.NewFromFloat(rand.Float64()).Truncate(8))
		params = append(params, "exp", int16(rand.Intn(1000)))
	case 2:
	case 3:
		params = append(params, "asset", uuid.New())
		params = append(params, "route", generateRoute())
		params = append(params, "min", generateDecimal())
	}

	return params
}

func generateRoute() string {
	b := make([]byte, rand.Intn(10))
	_, _ = crand.Read(b)

	return hex.EncodeToString(b)
}

func generateDecimal() decimal.Decimal {
	return decimal.New(int64(rand.Intn(math.MaxInt16)), -int32(rand.Intn(8)))
}

func checksum(data []byte) []byte {
	sum := sha256.Sum256(data)
	sum = sha256.Sum256(sum[:])
	return sum[:4]
}
