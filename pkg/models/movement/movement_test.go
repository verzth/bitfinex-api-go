package movement_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/verzth/bitfinex-api-go/pkg/models/movement"
)

func TestNewMovementFromRaw(t *testing.T) {
	t.Run("insufficient arguments", func(t *testing.T) {
		payload := []interface{}{138797990}

		w, err := movement.FromRaw(payload)
		require.NotNil(t, err)
		require.Nil(t, w)
	})

	t.Run("valid arguments", func(t *testing.T) {
		payload := []interface{}{
			138797990,
			"BTC",
			nil,
			1.5948918e+12,
			nil,
			0.001,
			30.002,
			nil,
			"Transfer of 0.001 BTC from wallet Exchange to Trading on wallet margin",
		}

		w, err := movement.FromRaw(payload)
		require.Nil(t, err)

		expected := &movement.Movement{
			ID:          138797990,
			Currency:    "BTC",
			MTS:         1594891800000,
			Amount:      0.001,
			Balance:     30.002,
			Description: "Transfer of 0.001 BTC from wallet Exchange to Trading on wallet margin",
		}

		assert.Equal(t, expected, w)
	})
}

func TestMovementSnapshotFromRaw(t *testing.T) {
	t.Run("rest success", func(t *testing.T) {
		payload := []interface{}{
			[]interface{}{
				138797990,
				"BTC",
				nil,
				1.5948918e+12,
				nil,
				0.001,
				30.002,
				nil,
				"Transfer of 0.001 BTC from wallet Exchange to Trading on wallet margin",
			},
			[]interface{}{
				138797710,
				"BTC",
				nil,
				1.5948919e+12,
				nil,
				-0.001,
				39.9988,
				nil,
				"Transfer of 0.001 BTC from wallet Exchange to Trading on wallet exchange",
			},
		}

		got, err := movement.SnapshotFromRaw(payload, movement.FromRaw)
		require.Nil(t, err)

		expected := &movement.Snapshot{
			Snapshot: []*movement.Movement{
				{
					ID:          138797990,
					Currency:    "BTC",
					MTS:         1594891800000,
					Amount:      0.001,
					Balance:     30.002,
					Description: "Transfer of 0.001 BTC from wallet Exchange to Trading on wallet margin",
				},
				{
					ID:          138797710,
					Currency:    "BTC",
					MTS:         1594891900000,
					Amount:      -0.001,
					Balance:     39.9988,
					Description: "Transfer of 0.001 BTC from wallet Exchange to Trading on wallet exchange",
				},
			},
		}

		assert.Equal(t, expected, got)
	})

	t.Run("rest fail", func(t *testing.T) {
		payload := []interface{}{
			[]interface{}{138797990},
			[]interface{}{
				138797710,
				"BTC",
				nil,
				1.5948919e+12,
				nil,
				-0.001,
				39.9988,
				nil,
				"Transfer of 0.001 BTC from wallet Exchange to Trading on wallet exchange",
			},
		}

		got, err := movement.SnapshotFromRaw(payload, movement.FromRaw)
		require.NotNil(t, err)
		require.Nil(t, got)
	})
}
