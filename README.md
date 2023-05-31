# mtg

## Protocol

### Header

The `Header` struct represents the header of a binary message protocol:

```golang
type Header struct {
    Version    uint8     `json:"version"`
    ProtocolID uint8     `json:"protocol_id"`
    FollowID   uuid.UUID `json:"follow_id"`
    Action     uint16    `json:"action"`
}
```

- `Version` (uint8): The version of the binary message protocol. It is always set to 1.
- `ProtocolID` (uint8): The ID of the protocol in use. It can be one of the following constants:
    - `ProtocolFswap` (value 1)
    - `ProtocolLeaf` (value 2)
    - `ProtocolRings` (value 3)
- `FollowID` (uuid.UUID): An optional UUID field that can be used to follow a sequence of binary messages.
- `Action` (uint16): The ID of the action to be performed.

Example usage:

```
// Create a new Header with ProtocolID set to ProtocolLeaf and Action set to 42
h := Header{
    Version:    1,
    ProtocolID: ProtocolLeaf,
    FollowID:   uuid.New(),
    Action:     42,
}
```

### MultisigReceiver

The `MultisigReceiver` type represents information about a multi-signature receiving account.

```golang
type MultisigReceiver struct {
    Version   uint8       `json:"version"`
    Members   []uuid.UUID `json:"members"`
    Threshold uint8       `json:"threshold"`
}
```

- `Version`: A `uint8` representing the version of the multi-signature account. This is always set to `1`.
- `Members`: An array of type `uuid.UUID` representing the members of the multi-signature account.
- `Threshold`: A `uint8` representing the minimum number of signatures required to authorize a transaction for the multi-signature account.

## Mtgpack

A binary encoding and decoding package. It supports encoding and decoding of commonly used data types such as integers, strings, floats, bools, and even more complex types like uuid and decimal.

In addition to the basic data types, this library also provides encoding and decoding methods for the `interface{}` type using Go's built-in reflection functionality. This allows for the encoding and decoding of any type that implements the `interface{}` interface.

### Features

- Encoding and decoding of basic data types like integers, strings, floats, and bools
- Encoding and decoding of complex data types like uuid and decimal
- Encoding and decoding of the `interface{}` type using reflection

### Example

generate memo for 4swap trade

```go
import (
  "encoding/base64"

  "github.com/google/uuid"
  "github.com/pandodao/mtg/mtgpack"
  "github.com/pandodao/mtg/protocol"
  "github.com/shopspring/decimal"
)

func generateSwapMemo() string {
  // protocol header
  header := protocol.Header{
    Version:    1,
    ProtocolID: protocol.ProtocolFswap,
    FollowID:   uuid.New(),
    Action:     3,
  }

  // receiver
  userID := "a539e9e7-0a9f-4871-b1d7-69568d8a5347" // replace with your mixin id

  receiver := protocol.MultisigReceiver{
    Version:   1,
    Members:   []uuid.UUID{uuid.MustParse(userID)},
    Threshold: 1,
  }

  // id of the asset you want to get
  assetID := "c6d0c728-2624-429b-8e0d-d9d19b6592fa"

  // minimum amount of the asset you want to get
  min, _ := decimal.NewFromString("0.1")

  // route of the asset you want to get
  route := "xvgf"

  // encode memo
  enc := mtgpack.NewEncoder()
  if err := enc.EncodeValues(header, receiver, uuid.MustParse(assetID), route, min); err != nil {
    panic(err)
  }

  return base64.StdEncoding.EncodeToString(enc.Bytes())
}
```






