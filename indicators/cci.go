package indicators

import (
	"container/list"
	"errors"
	"github.com/jaybutera/gotrade"
	"math"
)

// A Commodity Channel Index Indicator (Cci), no storage, for use in other indicators
type CciWithoutStorage struct {
	*baseIndicatorWithFloatBounds

	// private variables
	periodCounter          int
	typicalPriceAvg        *SmaWithoutStorage
	factor                 float64
	typicalPriceHistory    *list.List
	currentAvgTypicalPrice float64
	currentTypicalPrice    float64
	timePeriod             int
}

// NewCciWithoutStorage creates a Commodity Channel Index Indicator (Cci) without storage
func NewCciWithoutStorage(timePeriod int, valueAvailableAction ValueAvailableActionFloat) (indicator *CciWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	// the minimum timeperiod for a CCi indicator is 2
	if timePeriod < 2 {
		return nil, errors.New("timePeriod is less than the minimum (2)")
	}

	// check the maximum timeperiod
	if timePeriod > MaximumLookbackPeriod {
		return nil, errors.New("timePeriod is greater than the maximum (100000)")
	}

	lookback := timePeriod - 1
	ind := CciWithoutStorage{
		baseIndicatorWithFloatBounds: newBaseIndicatorWithFloatBounds(lookback, valueAvailableAction),
		factor:              0.015,
		periodCounter:       (timePeriod * -1),
		typicalPriceHistory: list.New(),
		timePeriod:          timePeriod,
	}

	ind.typicalPriceAvg, err = NewSmaWithoutStorage(timePeriod, func(dataItem float64, streamBarIndex int) {
		currentTypicalPriceAvg := dataItem

		var meanDeviation float64 = 0.0
		// calculate the mean deviation
		for e := ind.typicalPriceHistory.Front(); e != nil; e = e.Next() {
			value := e.Value.(float64)
			meanDeviation += math.Abs(value - currentTypicalPriceAvg)
		}
		meanDeviation /= float64(ind.timePeriod)

		result := ((ind.currentTypicalPrice - currentTypicalPriceAvg) / (ind.factor * meanDeviation))

		ind.UpdateIndicatorWithNewValue(result, streamBarIndex)
	})

	return &ind, err
}

// A Commodity Channel Index Indicator (Cci)
type Cci struct {
	*CciWithoutStorage

	// public variables
	Data []float64
}

// NewCci creates a Commodity Channel Index Indicator (Cci) for online usage
func NewCci(timePeriod int) (indicator *Cci, err error) {
	ind := Cci{}
	ind.CciWithoutStorage, err = NewCciWithoutStorage(timePeriod, func(dataItem float64, streamBarIndex int) {
		ind.Data = append(ind.Data, dataItem)
	})

	return &ind, err
}

// NewDefaultCci creates a Commodity Channel Index (Cci) for online usage with default parameters
//	- timePeriod: 14
func NewDefaultCci() (indicator *Cci, err error) {
	timePeriod := 14
	return NewCci(timePeriod)
}

// NewCciWithSrcLen creates a Commodity Channel Index (Cci) for offline usage
func NewCciWithSrcLen(sourceLength uint, timePeriod int) (indicator *Cci, err error) {
	ind, err := NewCci(timePeriod)

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDefaultCciWithSrcLen creates a Commodity Channel Index (Cci) for offline usage with default parameters
func NewDefaultCciWithSrcLen(sourceLength uint) (indicator *Cci, err error) {
	ind, err := NewDefaultCci()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewCciForStream creates a Commodity Channel Index (Cci) for online usage with a source data stream
func NewCciForStream(priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Cci, err error) {
	ind, err := NewCci(timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultCciForStream creates a Commodity Channel Index (Cci) for online usage with a source data stream
func NewDefaultCciForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Cci, err error) {
	ind, err := NewDefaultCci()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewCciForStreamWithSrcLen creates a Commodity Channel Index (Cci) for offline usage with a source data stream
func NewCciForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Cci, err error) {
	ind, err := NewCciWithSrcLen(sourceLength, timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultCciForStreamWithSrcLen creates a Commodity Channel Index (Cci) for offline usage with a source data stream
func NewDefaultCciForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Cci, err error) {
	ind, err := NewDefaultCciWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *Cci) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	ind.periodCounter += 1

	// calculate the typical price
	typicalPrice := (tickData.H() + tickData.L() + tickData.C()) / 3.0
	ind.currentTypicalPrice = typicalPrice

	// push it to the history
	ind.typicalPriceHistory.PushBack(typicalPrice)

	// trim the history
	if ind.typicalPriceHistory.Len() > ind.timePeriod {
		var first = ind.typicalPriceHistory.Front()
		ind.typicalPriceHistory.Remove(first)
	}

	// add it to the average
	ind.typicalPriceAvg.ReceiveTick(typicalPrice, streamBarIndex)
}
