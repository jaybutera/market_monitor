package indicators

import (
	"container/list"
	"errors"
	"github.com/jaybutera/gotrade"
)

// An Average Directional Index Rating (Adxr), no storage
type AdxrWithoutStorage struct {
	*baseIndicatorWithFloatBounds

	// private variables
	periodCounter int
	periodHistory *list.List
	adx           *AdxWithoutStorage
	timePeriod    int
}

// NewAdxrWithoutStorage creates an Average Directional Index Rating (Adxr) without storage
func NewAdxrWithoutStorage(timePeriod int, valueAvailableAction ValueAvailableActionFloat) (indicator *AdxrWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	// the minimum timeperiod for an Adxr indicator is 2
	if timePeriod < 2 {
		return nil, errors.New("timePeriod is less than the minimum (2)")
	}

	// check the maximum timeperiod
	if timePeriod > MaximumLookbackPeriod {
		return nil, errors.New("timePeriod is greater than the maximum (100000)")
	}

	ind := AdxrWithoutStorage{
		periodCounter: 0,
		periodHistory: list.New(),
		timePeriod:    timePeriod,
	}

	ind.adx, err = NewAdxWithoutStorage(timePeriod, func(dataItem float64, streamBarIndex int) {
		ind.periodHistory.PushBack(dataItem)

		if ind.periodCounter > ind.GetLookbackPeriod() {
			adxN := ind.periodHistory.Front().Value.(float64)
			result := (dataItem + adxN) / 2.0

			ind.UpdateIndicatorWithNewValue(result, streamBarIndex)
		}

		if ind.periodHistory.Len() >= timePeriod {
			first := ind.periodHistory.Front()
			ind.periodHistory.Remove(first)
		}
	})

	var lookback int = 3
	if timePeriod > 1 {
		lookback = timePeriod - 1 + ind.adx.GetLookbackPeriod()
	}
	ind.baseIndicatorWithFloatBounds = newBaseIndicatorWithFloatBounds(lookback, valueAvailableAction)

	return &ind, nil
}

// A Directional Movement Indicator Rating (Adxr)
type Adxr struct {
	*AdxrWithoutStorage

	// public variables
	Data []float64
}

// NewAdxr creates an Average Directional Index Rating (Adxr) for online usage
func NewAdxr(timePeriod int) (indicator *Adxr, err error) {
	ind := Adxr{}
	ind.AdxrWithoutStorage, err = NewAdxrWithoutStorage(timePeriod,
		func(dataItem float64, streamBarIndex int) {
			ind.Data = append(ind.Data, dataItem)
		})

	return &ind, err
}

// NewDefaultAdxr creates an Average Directional Index Rating (Adxr) for online usage with default parameters
//	- timePeriod: 14
func NewDefaultAdxr() (indicator *Adxr, err error) {
	timePeriod := 14
	return NewAdxr(timePeriod)
}

// NewAdxrWithSrcLen creates an Average Directional Index Rating (Adxr) for offline usage
func NewAdxrWithSrcLen(sourceLength uint, timePeriod int) (indicator *Adxr, err error) {
	ind, err := NewAdxr(timePeriod)

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDefaultAdxrWithSrcLen creates an Average Directional Index Rating (Adxr) for offline usage with default parameters
func NewDefaultAdxrWithSrcLen(sourceLength uint) (indicator *Adxr, err error) {
	ind, err := NewDefaultAdxr()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewAdxrForStream creates an Average Directional Rating Index (Adxr) for online usage with a source data stream
func NewAdxrForStream(priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Adxr, err error) {
	ind, err := NewAdxr(timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultAdxrForStream creates an Average Directional Index Rating (Adxr) for online usage with a source data stream
func NewDefaultAdxrForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Adxr, err error) {
	ind, err := NewDefaultAdxr()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewAdxrForStreamWithSrcLen creates an Average Directional Index Rating (Adxr) for offline usage with a source data stream
func NewAdxrForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Adxr, err error) {
	ind, err := NewAdxrWithSrcLen(sourceLength, timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultAdxrForStreamWithSrcLen creates an Average Directional Index Rating (Adxr) for offline usage with a source data stream
func NewDefaultAdxrForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Adxr, err error) {
	ind, err := NewDefaultAdxrWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *AdxrWithoutStorage) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	ind.periodCounter += 1
	ind.adx.ReceiveDOHLCVTick(tickData, streamBarIndex)
}
