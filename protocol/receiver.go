package protocol

import (
	"github.com/google/uuid"
	"github.com/pandodao/mtg/mtgpack"
)

type MultisigReceiver struct {
	Version   uint8       `json:"version"`
	Members   []uuid.UUID `json:"members"`
	Threshold uint8       `json:"threshold"`
}

func (m *MultisigReceiver) DecodeMtg(d *mtgpack.Decoder) error {
	var err error
	m.Version, err = d.DecodeUint8()
	if err != nil {
		return err
	}

	count, err := d.DecodeUint8()
	if err != nil {
		return err
	}

	m.Threshold, err = d.DecodeUint8()
	if err != nil {
		return err
	}

	m.Members = make([]uuid.UUID, int(count))
	for i := range m.Members {
		m.Members[i], err = d.DecodeUUID()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m MultisigReceiver) EncodeMtg(e *mtgpack.Encoder) error {
	if err := e.EncodeUint8(m.Version); err != nil {
		return err
	}

	if err := e.EncodeUint8(uint8(len(m.Members))); err != nil {
		return err
	}

	if err := e.EncodeUint8(m.Threshold); err != nil {
		return err
	}

	for _, member := range m.Members {
		if err := e.EncodeUUID(member); err != nil {
			return err
		}
	}

	return nil
}
