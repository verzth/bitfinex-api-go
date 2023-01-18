package rest

import (
	"fmt"
	"path"

	"github.com/verzth/bitfinex-api-go/pkg/models/common"
	"github.com/verzth/bitfinex-api-go/pkg/models/ledger"
)

// LedgerService manages the Ledgers endpoint.
type LedgerService struct {
	requestFactory
	Synchronous
}

const (
	CATEGORY_EXCHANGE                = 5
	CATEGORY_POSITION                = 22
	CATEGORY_POSITION_CLAIM          = 23
	CATEGORY_POSITION_TRANSFER       = 25
	CATEGORY_POSITION_SWAP           = 26
	CATEGORY_POSITION_CHARGED        = 27
	CATEGORY_MARGIN_INTEREST_PAYMENT = 28
	CATEGORY_DERIVATIVE_EVENT        = 29
	CATEGORY_SETTLEMENT              = 31
	CATEGORY_TRANSFER                = 51
	CATEGORY_DEPOSIT                 = 101
	CATEGORY_WITHDRAWAL              = 104
	CATEGORY_CANCELLED_WITHDRAWAL    = 105
	CATEGORY_DEPOSIT_FEE             = 251
	CATEGORY_WITHDRAWAL_FEE          = 254
	CATEGORY_WITHDRAWAL_EXPRESS_FEE  = 255
)

var maxLimit int32 = 2500

// Ledgers - all of the past ledger entreies
// see https://docs.bitfinex.com/reference#ledgers for more info
func (s *LedgerService) Ledgers(currency string, category int, start int64, end int64, max int32) (*ledger.Snapshot, error) {
	if max > maxLimit {
		return nil, fmt.Errorf("Max request limit:%d, got: %d", maxLimit, max)
	}

	payload := map[string]interface{}{"category": category, "start": start, "end": end, "limit": max}
	req, err := s.requestFactory.NewAuthenticatedRequestWithData(common.PermissionRead, path.Join("ledgers", currency, "hist"), payload)
	if err != nil {
		return nil, err
	}

	raw, err := s.Request(req)
	if err != nil {
		return nil, err
	}

	lss, err := ledger.SnapshotFromRaw(raw, ledger.FromRaw)
	if err != nil {
		return nil, err
	}

	return lss, nil
}
