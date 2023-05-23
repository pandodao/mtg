package protocol

import (
	"github.com/google/uuid"
	"github.com/pandodao/mtg/mtgpack"
)

const (
	ProtocolFswap      uint8 = 1
	ProtocolLeaf       uint8 = 2
	ProtocolRings      uint8 = 3
	ProtocolPool       uint8 = 4
	ProtocolTradingLab uint8 = 5
)

type Header struct {
	Version    uint8     `json:"version"`
	ProtocolID uint8     `json:"protocol_id"`
	FollowID   uuid.UUID `json:"follow_id"`
	Action     uint16    `json:"action"`
}

func (h Header) HasFollowID() bool {
	return h.FollowID != uuid.Nil
}

func (h *Header) DecodeMtg(d *mtgpack.Decoder) error {
	var err error
	h.Version, err = d.DecodeUint8()
	if err != nil {
		return err
	}

	h.ProtocolID, err = d.DecodeUint8()
	if err != nil {
		return err
	}

	hasFollowID, err := d.DecodeBool()
	if err != nil {
		return err
	}

	if hasFollowID {
		h.FollowID, err = d.DecodeUUID()
	}

	h.Action, err = d.DecodeUint16()
	if err != nil {
		return err
	}

	return nil
}

func (h Header) EncodeMtg(e *mtgpack.Encoder) error {
	if err := e.EncodeUint8(h.Version); err != nil {
		return err
	}

	if err := e.EncodeUint8(h.ProtocolID); err != nil {
		return err
	}

	hasFollowID := h.FollowID != uuid.Nil
	if err := e.EncodeBool(hasFollowID); err != nil {
		return err
	}

	if hasFollowID {
		if err := e.EncodeUUID(h.FollowID); err != nil {
			return err
		}
	}

	if err := e.EncodeUint16(h.Action); err != nil {
		return err
	}

	return nil
}
