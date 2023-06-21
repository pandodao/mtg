package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/pandodao/mtg/mtgpack"
	"github.com/pandodao/mtg/protocol"
	"github.com/shopspring/decimal"
)

var (
	decodeFlag            = flag.String("d", "", "decode")
	decodeOmitMmsigFlag   = flag.Bool("om", false, "decode omit mmsig")
	decodeParamsTypesFlag = flag.String("pts", "", "decode params types, example: [\"decimal\", \"uuid\", false, 0, \"int8\"]")

	encodeFlag = flag.String("e", "", "encode")
)

type EncodeData struct {
	protocol.Header
	Mmsig  *protocol.MultisigReceiver `json:"mmsig,omitempty"`
	Params []any                      `json:"params,omitempty"`
}

func main() {
	flag.Parse()

	var (
		result string
		err    error
	)
	switch {
	case *decodeFlag != "":
		result, err = Decode(*decodeFlag, *decodeOmitMmsigFlag, *decodeParamsTypesFlag)
	case *encodeFlag != "":
		result, err = Encode(*encodeFlag)
	default:
		flag.PrintDefaults()
		return
	}

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
func Decode(v string, omitMmsig bool, paramsTypes string) (string, error) {
	r, err := decode(v, omitMmsig, paramsTypes)
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decode(ds string, omitMmsig bool, paramsTypesStr string) (*EncodeData, error) {
	data, err := base64.StdEncoding.DecodeString(ds)
	if err != nil {
		return nil, fmt.Errorf("base64 decode memo failed: %v", err)
	}

	dec := mtgpack.NewDecoder(data)

	result := &EncodeData{}
	if err := dec.DecodeValue(&result.Header); err != nil {
		return nil, fmt.Errorf("decode header failed: %v", err)
	}
	if !omitMmsig {
		result.Mmsig = &protocol.MultisigReceiver{}
		if err := dec.DecodeValue(result.Mmsig); err != nil {
			return nil, fmt.Errorf("decode mmsig failed: %v", err)
		}
	}

	if paramsTypesStr != "" {
		var paramsTypes []any
		if err := json.Unmarshal([]byte(paramsTypesStr), &paramsTypes); err != nil {
			return nil, fmt.Errorf("decode params types failed: %v", err)
		}

		for _, param := range paramsTypes {
			var (
				decodeValue any
				prFunc      func() any
			)

			if paramStr, ok := param.(string); ok {
				switch paramStr {
				case "", "string":
					v := ""
					decodeValue = &v
					prFunc = func() any { return "string:" + v }
				case "decimal":
					v := decimal.Zero
					decodeValue = &v
					prFunc = func() any { return "decimal:" + v.String() }
				case "uuid":
					v := uuid.UUID{}
					decodeValue = &v
					prFunc = func() any { return "uuid:" + v.String() }
				case "int", "int64":
					v := int64(0)
					decodeValue = &v
					prFunc = func() any { return "int64:" + strconv.FormatInt(v, 10) }
				case "int8":
					v := int8(0)
					decodeValue = &v
					prFunc = func() any { return "int8:" + strconv.FormatInt(int64(v), 10) }
				case "int16":
					v := int16(0)
					decodeValue = &v
					prFunc = func() any { return "int16:" + strconv.FormatInt(int64(v), 10) }
				case "int32":
					v := int32(0)
					decodeValue = &v
					prFunc = func() any { return "int32:" + strconv.FormatInt(int64(v), 10) }
				case "uint", "uint64":
					v := uint64(0)
					decodeValue = &v
					prFunc = func() any { return "uint64:" + strconv.FormatUint(v, 10) }
				case "uint8":
					v := uint8(0)
					decodeValue = &v
					prFunc = func() any { return "uint8:" + strconv.FormatUint(uint64(v), 10) }
				case "uint16":
					v := uint16(0)
					decodeValue = &v
					prFunc = func() any { return "uint16:" + strconv.FormatUint(uint64(v), 10) }
				case "uint32":
					v := uint32(0)
					decodeValue = &v
					prFunc = func() any { return "uint32:" + strconv.FormatUint(uint64(v), 10) }
				case "bool":
					v := false
					decodeValue = &v
					prFunc = func() any { return "bool:" + strconv.FormatBool(v) }
				default:
					return nil, fmt.Errorf("invalid param type: %s", paramStr)
				}
			} else {
				decodeValue = &param
				prFunc = func() any { return param }
			}

			if err := dec.DecodeValue(decodeValue); err != nil {
				return nil, fmt.Errorf("decode param failed: %v, param: %v", err, param)
			}

			result.Params = append(result.Params, prFunc())
		}
	}

	return result, nil
}

func Encode(es string) (string, error) {
	ed := &EncodeData{}
	if err := json.Unmarshal([]byte(es), ed); err != nil {
		return "", fmt.Errorf("unmarshal encode data failed: %v", err)
	}

	enc := mtgpack.NewEncoder()
	if err := mtgpack.EncodeValue(enc, ed.Header); err != nil {
		return "", fmt.Errorf("encode header failed: %v", err)
	}

	if ed.Mmsig != nil {
		if err := mtgpack.EncodeValue(enc, *ed.Mmsig); err != nil {
			return "", fmt.Errorf("encode mmsig failed: %v", err)
		}
	}

	for _, param := range ed.Params {
		var (
			v   any
			err error
		)
		if paramStr, ok := param.(string); ok {
			tv := strings.SplitN(paramStr, ":", 2)
			if len(tv) == 1 {
				v = tv[0]
			} else {
				switch tv[0] {
				case "decimal":
					v, err = decimal.NewFromString(tv[1])
				case "uuid":
					v, err = uuid.Parse(tv[1])
				case "int", "int64":
					v, err = strconv.ParseInt(tv[1], 10, 64)
				case "int8":
					var vi int64
					vi, err = strconv.ParseInt(tv[1], 10, 64)
					v = int8(vi)
				case "int16":
					var vi int64
					vi, err = strconv.ParseInt(tv[1], 10, 64)
					v = int16(vi)
				case "int32":
					var vi int64
					vi, err = strconv.ParseInt(tv[1], 10, 64)
					v = int32(vi)
				case "uint", "uint64":
					v, err = strconv.ParseUint(tv[1], 10, 64)
				case "uint8":
					var vi uint64
					vi, err = strconv.ParseUint(tv[1], 10, 64)
					v = uint8(vi)
				case "uint16":
					var vi uint64
					vi, err = strconv.ParseUint(tv[1], 10, 64)
					v = uint16(vi)
				case "uint32":
					var vi uint64
					vi, err = strconv.ParseUint(tv[1], 10, 64)
					v = uint32(vi)
				case "bool":
					v, err = strconv.ParseBool(tv[1])
				case "string":
					v = tv[1]
				default:
					err = fmt.Errorf("unknown type: %s", tv[0])
				}
			}
		} else {
			v = param
		}
		if err != nil {
			return "", fmt.Errorf("parse param failed: %v, param: %v", err, param)
		}

		if err := mtgpack.EncodeValue(enc, v); err != nil {
			return "", fmt.Errorf("encode param failed: %v, param: %v", err, param)
		}
	}

	return base64.StdEncoding.EncodeToString(enc.Bytes()), nil
}
