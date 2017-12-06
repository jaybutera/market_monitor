package indicators

// DX = ( (+DI)-(-DI) ) / ( (+DI) + (-DI) )

import (
	"errors"
	"github.com/jaybutera/gotrade"
	"math"
)

// An Directional Movement Index Indicator (Dx), no storage, for use in other indicators
type DxWithoutStorage struct {
	*baseIndicatorWithFloatBounds

	// private variables
	minusDI        *MinusDi
	plusDI         *PlusDi
	currentPlusDi  float64
	currentMinusDi float64
	timePeriod     int
}

// NewDxWithoutStorage creates a Directional Movement Index Indicator (Dx) without storage
func NewDxWithoutStorage(timePeriod int, valueAvailableAction ValueAvailableActionFloat) (indicator *DxWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	// the minimum timeperiod for this indicator is 2
	if timePeriod < 2 {
		return nil, errors.New("timePeriod is less than the minimum (2)")
	}

	// check the maximum timeperiod
	if timePeriod > MaximumLookbackPeriod {
		return nil, errors.New("timePeriod is greater than the maximum (100000)")
	}

	lookback := 2
	if timePeriod > 1 {
		lookback = timePeriod
	}

	ind := DxWithoutStorage{
		baseIndicatorWithFloatBounds: newBaseIndicatorWithFloatBounds(lookback, valueAvailableAction),
		currentPlusDi:                0.0,
		currentMinusDi:               0.0,
		timePeriod:                   timePeriod,
	}

	ind.minusDI, err = NewMinusDi(timePeriod)

	ind.minusDI.valueAvailableAction = func(dataItem float64, streamBarIndex int) {
		ind.currentMinusDi = dataItem
	}

	ind.plusDI, err = NewPlusDi(timePeriod)

	ind.plusDI.valueAvailableAction = func(dataItem float64, streamBarIndex int) {
		ind.currentPlusDi = dataItem

		var result float64
		tmp := ind.currentMinusDi + ind.currentPlusDi
		if tmp != 0.0 {
			result = 100.0 * (math.Abs(ind.currentMinusDi-ind.currentPlusDi) / tmp)
		} else {
			result = 0.0
		}

		ind.UpdateIndicatorWithNewValue(result, streamBarIndex)
	}

	return &ind, err
}

// A Directional Movement Index Indicator (Dx)
type Dx struct {
	*DxWithoutStorage

	// public variables
	Data []float64
}

// NewDx creates a Directional Movement Index Indicator (Dx) for online usage
func NewDx(timePeriod int) (indicator *Dx, err error) {

	ind := Dx{}
	ind.DxWithoutStorage, err = NewDxWithoutStorage(timePeriod,
		func(dataItem float64, streamBarIndex int) {
			ind.Data = append(ind.Data, dataItem)
		})

	return &ind, err
}

// NewDefaultDx creates a Directional Movement Index (Dx) for online usage with default parameters
//	- timePeriod: 14
func NewDefaultDx() (indicator *Dx, err error) {
	timePeriod := 14
	return NewDx(timePeriod)
}

// NewDxWithSrcLen creates a Directional Movement Index (Dx) for offline usage
func NewDxWithSrcLen(sourceLength uint, timePeriod int) (indicator *Dx, err error) {
	ind, err := NewDx(timePeriod)

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDefaultDxWithSrcLen creates a Directional Movement Index (Dx) for offline usage with default parameters
func NewDefaultDxWithSrcLen(sourceLength uint) (indicator *Dx, err error) {
	ind, err := NewDefaultDx()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Data = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDxForStream creates a Directional Movement Index (Dx) for online usage with a source data stream
func NewDxForStream(priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Dx, err error) {
	ind, err := NewDx(timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultDxForStream creates a Directional Movement Index (Dx) for online usage with a source data stream
func NewDefaultDxForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Dx, err error) {
	ind, err := NewDefaultDx()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDxForStreamWithSrcLen creates a Directional Movement Index (Dx) for offline usage with a source data stream
func NewDxForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Dx, err error) {
	ind, err := NewDxWithSrcLen(sourceLength, timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultDxForStreamWithSrcLen creates a Directional Movement Index (Dx) for offline usage with a source data stream
func NewDefaultDxForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Dx, err error) {
	ind, err := NewDefaultDxWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *DxWithoutStorage) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	ind.minusDI.ReceiveDOHLCVTick(tickData, streamBarIndex)
	ind.plusDI.ReceiveDOHLCVTick(tickData, streamBarIndex)
}
