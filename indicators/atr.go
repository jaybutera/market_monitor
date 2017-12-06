package indicators

import (
	"errors"
	"github.com/jaybutera/gotrade"
)

// An Average True Range Indicator (Atr), no storage, for use in other indicators
type AtrWithoutStorage struct {
	*baseIndicatorWithFloatBounds

	// private variables
	trueRange            *TrueRangeWithoutStorage
	sma                  *SmaWithoutStorage
	previousAvgTrueRange float64
	multiplier           float64
	timePeriod           int
}

// NewAtrWithoutStorage creates an Average True Range Indicator (Atr) without storage
func NewAtrWithoutStorage(timePeriod int, valueAvailableAction ValueAvailableActionFloat) (indicator *AtrWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	// the minimum timeperiod for an Atr indicator is 1
	if timePeriod < 1 {
		return nil, errors.New("timePeriod is less than the minimum (1)")
	}

	// check the maximum timeperiod
	if timePeriod > MaximumLookbackPeriod {
		return nil, errors.New("timePeriod is greater than the maximum (100000)")
	}

	lookback := timePeriod
	ind := AtrWithoutStorage{
		baseIndicatorWithFloatBounds: newBaseIndicatorWithFloatBounds(lookback, valueAvailableAction),
		multiplier:                   float64(timePeriod - 1),
		previousAvgTrueRange:         -1,
		timePeriod:                   timePeriod,
	}

	ind.sma, err = NewSmaWithoutStorage(timePeriod, func(dataItem float64, streamBarIndex int) {
		ind.previousAvgTrueRange = dataItem

		ind.UpdateIndicatorWithNewValue(dataItem, streamBarIndex)
	})

	ind.trueRange, err = NewTrueRangeWithoutStorage(func(dataItem float64, streamBarIndex int) {

		if ind.previousAvgTrueRange == -1 {
			ind.sma.ReceiveTick(dataItem, streamBarIndex)
		} else {

			result := ((ind.previousAvgTrueRange * ind.multiplier) + dataItem) / float64(ind.timePeriod)

			ind.UpdateIndicatorWithNewValue(result, streamBarIndex)

			// update the previous true range for the next tick
			ind.previousAvgTrueRange = result
		}

	})
	return &ind, nil
}

// An Average True Range Indicator (Atr)
type Atr struct {
	*AtrWithoutStorage

	// public variables
	Data []float64
}

// NewAtr creates an Average True Range (Atr) for online usage
func NewAtr(timePeriod int) (indicator *Atr, err error) {
	ind := Atr{}
	ind.AtrWithoutStorage, err = NewAtrWithoutStorage(timePeriod, func(dataItem float64, streamBarIndex int) {
		ind.Data = append(ind.Data, dataItem)
	})

	return &ind, err
}

// NewDefaultAtr creates an Average True Range (Atr) for online usage with default parameters
//	- timePeriod: 14
func NewDefaultAtr() (indicator *Atr, err error) {
	timePeriod := 14
	return NewAtr(timePeriod)
}

// NewAtrWithSrcLen creates an Average True Range (Atr) for offline usage
func NewAtrWithSrcLen(sourceLength uint, timePeriod int) (indicator *Atr, err error) {
	ind, err := NewAtr(timePeriod)

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDefaultAtrWithSrcLen creates an Average True Range (Atr) for offline usage with default parameters
func NewDefaultAtrWithSrcLen(sourceLength uint) (indicator *Atr, err error) {
	ind, err := NewDefaultAtr()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewAtrForStream creates an Average True Range (Atr) for online usage with a source data stream
func NewAtrForStream(priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Atr, err error) {
	ind, err := NewAtr(timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultAtrForStream creates an Average True Range (Atr) for online usage with a source data stream
func NewDefaultAtrForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Atr, err error) {
	ind, err := NewDefaultAtr()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewAtrForStreamWithSrcLen creates an Average True Range (Atr) for offline usage with a source data stream
func NewAtrForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Atr, err error) {
	ind, err := NewAtrWithSrcLen(sourceLength, timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultAtrForStreamWithSrcLen creates an Average True Range (Atr) for offline usage with a source data stream
func NewDefaultAtrForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Atr, err error) {
	ind, err := NewDefaultAtrWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *AtrWithoutStorage) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	// update the current true range
	ind.trueRange.ReceiveDOHLCVTick(tickData, streamBarIndex)
}
