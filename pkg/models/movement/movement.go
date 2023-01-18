package movement

import (
	"fmt"

	"github.com/verzth/bitfinex-api-go/pkg/convert"
)

type Movement struct {
	ID           int64
	Currency     string
	CurrencyName string
	// placeholder
	// placeholder
	MTSStarted int64
	MTSUpdated int64
	// placeholder
	// placeholder
	Status string
	// placeholder
	// placeholder
	Amount float64
	Fees   float64
	// placeholder
	// placeholder
	DestinationAddress string
	// placeholder
	// placeholder
	// placeholder
	TransactionId           string
	WithdrawTransactionNote string
}

type Snapshot struct {
	Snapshot []*Movement
}

type transformerFn func(raw []interface{}) (w *Movement, err error)

// FromRaw takes the raw list of values as returned from the websocket
// service and tries to convert it into an Movement.
func FromRaw(raw []interface{}) (o *Movement, err error) {
	if len(raw) < 9 {
		return o, fmt.Errorf("data slice too short for movement: %#v", raw)
	}

	o = &Movement{
		ID:                      convert.I64ValOrZero(raw[0]),
		Currency:                convert.SValOrEmpty(raw[1]),
		CurrencyName:            convert.SValOrEmpty(raw[2]),
		MTSStarted:              convert.I64ValOrZero(raw[5]),
		MTSUpdated:              convert.I64ValOrZero(raw[6]),
		Status:                  convert.SValOrEmpty(raw[9]),
		Amount:                  convert.F64ValOrZero(raw[12]),
		Fees:                    convert.F64ValOrZero(raw[13]),
		DestinationAddress:      convert.SValOrEmpty(raw[16]),
		TransactionId:           convert.SValOrEmpty(raw[20]),
		WithdrawTransactionNote: convert.SValOrEmpty(raw[21]),
	}

	return
}

// SnapshotFromRaw takes a raw list of values as returned from the websocket
// service and tries to convert it into an Snapshot.
func SnapshotFromRaw(raw []interface{}, t transformerFn) (s *Snapshot, err error) {
	if len(raw) == 0 {
		return s, fmt.Errorf("data slice too short for movements: %#v", raw)
	}

	lss := make([]*Movement, 0)
	switch raw[0].(type) {
	case []interface{}:
		for _, v := range raw {
			if l, ok := v.([]interface{}); ok {
				o, err := t(l)
				if err != nil {
					return s, err
				}
				lss = append(lss, o)
			}
		}
	default:
		return s, fmt.Errorf("not an movement snapshot")
	}
	s = &Snapshot{Snapshot: lss}
	return
}
