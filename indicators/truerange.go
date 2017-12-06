package indicators

import (
	"github.com/jaybutera/gotrade"
	"math"
)

// TrueHigh = Max(High[0], Close[-1])
// TrueLow = Min(Low[0], Close[-1])
// TrueRange = TrueHigh = TrueLow

// A True Range Indicator (TrueRange), no storage, for use in other indicators
type TrueRangeWithoutStorage struct {
	*baseIndicatorWithFloatBounds

	// private variables
	periodCounter int
	previousClose float64
}

// NewTrueRangeWithoutStorage creates a True Range Indicator (TrueRange) without storage
func NewTrueRangeWithoutStorage(valueAvailableAction ValueAvailableActionFloat) (indicator *TrueRangeWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	lookback := 1
	ind := TrueRangeWithoutStorage{
		baseIndicatorWithFloatBounds: newBaseIndicatorWithFloatBounds(lookback, valueAvailableAction),
		periodCounter:                -1,
		previousClose:                0.0,
	}
	return &ind, nil
}

// A True Range Indicator (TrueRange)
type TrueRange struct {
	*TrueRangeWithoutStorage

	// public variables
	Data []float64
}

// NewTrueRange creates a True Range Indicator (TrueRange) for online usage
func NewTrueRange() (indicator *TrueRange, err error) {
	ind := TrueRange{}
	ind.TrueRangeWithoutStorage, err = NewTrueRangeWithoutStorage(func(dataItem float64, streamBarIndex int) {
		ind.Data = append(ind.Data, dataItem)
	})
	return &ind, err
}

// NewTrueRangeWithSrcLen creates a True Range Indicator (TrueRange) for offline usage
func NewTrueRangeWithSrcLen(sourceLength uint) (indicator *TrueRange, err error) {
	ind, err := NewTrueRange()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewTrueRangeForStream creates a True Range Indicator (TrueRange) for online usage with a source data stream
func NewTrueRangeForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *TrueRange, err error) {
	ind, err := NewTrueRange()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewTrueRangeForStreamWithSrcLen creates a True Range Indicator (TrueRange) for offline usage with a source data stream
func NewTrueRangeForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *TrueRange, err error) {
	ind, err := NewTrueRangeWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *TrueRangeWithoutStorage) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	ind.periodCounter += 1

	if ind.periodCounter > 0 {

		high := math.Max(tickData.H(), ind.previousClose)
		low := math.Min(tickData.L(), ind.previousClose)
		result := high - low

		ind.UpdateIndicatorWithNewValue(result, streamBarIndex)
	}

	ind.previousClose = tickData.C()
}
