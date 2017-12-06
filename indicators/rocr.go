package indicators

import (
	"container/list"
	"errors"
	"github.com/jaybutera/gotrade"
)

// A Rate of Change Ratio Indicator (RocR), no storage, for use in other indicators
type RocRWithoutStorage struct {
	*baseIndicatorWithFloatBounds

	// private variables
	periodCounter int
	periodHistory *list.List
	timePeriod    int
}

// NewRocRWithoutStorage creates a Rate of Change Ratio Indicator (RocR) without storage
func NewRocRWithoutStorage(timePeriod int, valueAvailableAction ValueAvailableActionFloat) (indicator *RocRWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	// the minimum timeperiod for this indicator is 1
	if timePeriod < 1 {
		return nil, errors.New("timePeriod is less than the minimum (1)")
	}

	// check the maximum timeperiod
	if timePeriod > MaximumLookbackPeriod {
		return nil, errors.New("timePeriod is greater than the maximum (100000)")
	}

	lookback := timePeriod
	ind := RocRWithoutStorage{
		baseIndicatorWithFloatBounds: newBaseIndicatorWithFloatBounds(lookback, valueAvailableAction),
		periodCounter:                (timePeriod * -1),
		periodHistory:                list.New(),
		timePeriod:                   timePeriod,
	}

	return &ind, err
}

// A Rate of Change Ratio Indicator (RocR)
type RocR struct {
	*RocRWithoutStorage
	selectData gotrade.DOHLCVDataSelectionFunc

	// public variables
	Data []float64
}

// NewRocR creates a Rate of Change Ratio Indicator (RocR) for online usage
func NewRocR(timePeriod int, selectData gotrade.DOHLCVDataSelectionFunc) (indicator *RocR, err error) {
	if selectData == nil {
		return nil, ErrDOHLCVDataSelectFuncIsNil
	}

	ind := RocR{
		selectData: selectData,
	}

	ind.RocRWithoutStorage, err = NewRocRWithoutStorage(timePeriod,
		func(dataItem float64, streamBarIndex int) {
			ind.Data = append(ind.Data, dataItem)
		})

	return &ind, err
}

// NewDefaultRocR creates a Rate of Change Ratio Indicator (RocR) for online usage with default parameters
//	- timePeriod: 10
func NewDefaultRocR() (indicator *RocR, err error) {
	timePeriod := 10
	return NewRocR(timePeriod, gotrade.UseClosePrice)
}

// NewRocRWithSrcLen creates a Rate of Change Ratio Indicator (RocR) for offline usage
func NewRocRWithSrcLen(sourceLength uint, timePeriod int, selectData gotrade.DOHLCVDataSelectionFunc) (indicator *RocR, err error) {
	ind, err := NewRocR(timePeriod, selectData)

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDefaultRocRWithSrcLen creates a Rate of Change Ratio Indicator (RocR) for offline usage with default parameters
func NewDefaultRocRWithSrcLen(sourceLength uint) (indicator *RocR, err error) {
	ind, err := NewDefaultRocR()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewRocRForStream creates a Rate of Change Ratio Indicator (RocR) for online usage with a source data stream
func NewRocRForStream(priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int, selectData gotrade.DOHLCVDataSelectionFunc) (indicator *RocR, err error) {
	ind, err := NewRocR(timePeriod, selectData)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultRocRForStream creates a Rate of Change Ratio Indicator (RocR) for online usage with a source data stream
func NewDefaultRocRForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *RocR, err error) {
	ind, err := NewDefaultRocR()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewRocRForStreamWithSrcLen creates a Rate of Change Ratio Indicator (RocR) for offline usage with a source data stream
func NewRocRForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int, selectData gotrade.DOHLCVDataSelectionFunc) (indicator *RocR, err error) {
	ind, err := NewRocRWithSrcLen(sourceLength, timePeriod, selectData)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultRocRForStreamWithSrcLen creates a Rate of Change Ratio Indicator (RocR) for offline usage with a source data stream
func NewDefaultRocRForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *RocR, err error) {
	ind, err := NewDefaultRocRWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *RocR) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	var selectedData = ind.selectData(tickData)
	ind.ReceiveTick(selectedData, streamBarIndex)
}

func (ind *RocRWithoutStorage) ReceiveTick(tickData float64, streamBarIndex int) {
	ind.periodCounter += 1
	ind.periodHistory.PushBack(tickData)

	if ind.periodCounter > 0 {

		//    RocR = (price/previousPrice - 1) * 100
		previousPrice := ind.periodHistory.Front().Value.(float64)

		var result float64
		if previousPrice != 0 {
			result = (tickData / previousPrice)
		} else {
			result = 0.0
		}

		ind.UpdateIndicatorWithNewValue(result, streamBarIndex)
	}

	if ind.periodHistory.Len() > ind.timePeriod {
		first := ind.periodHistory.Front()
		ind.periodHistory.Remove(first)
	}
}
