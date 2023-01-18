package rest

import (
	"fmt"
	"path"

	"github.com/verzth/bitfinex-api-go/pkg/models/common"
	"github.com/verzth/bitfinex-api-go/pkg/models/movement"
)

// MovementService manages the Movements endpoint.
type MovementService struct {
	requestFactory
	Synchronous
}

// Movements - all of the past movement entreies
// see https://docs.bitfinex.com/reference#movements for more info
func (s *MovementService) Movements(currency string, start int64, end int64, max int32) (*movement.Snapshot, error) {
	if max > 1000 {
		return nil, fmt.Errorf("max request limit:%d, got: %d", 1000, max)
	}

	payload := map[string]interface{}{"start": start, "end": end, "limit": max}
	req, err := s.requestFactory.NewAuthenticatedRequestWithData(common.PermissionRead, path.Join("movements", currency, "hist"), payload)
	if err != nil {
		return nil, err
	}

	raw, err := s.Request(req)
	if err != nil {
		return nil, err
	}

	lss, err := movement.SnapshotFromRaw(raw, movement.FromRaw)
	if err != nil {
		return nil, err
	}

	return lss, nil
}
