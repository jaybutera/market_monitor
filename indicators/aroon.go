package indicators

import (
	"container/list"
	"errors"
	"github.com/jaybutera/gotrade"
	"math"
)

// An Aroon (Aroon), no storage, for use in other indicators
type AroonWithoutStorage struct {
	*baseIndicatorWithFloatBoundsAroon

	// private variables
	periodCounter     int
	periodHighHistory *list.List
	periodLowHistory  *list.List
	aroonFactor       float64
	timePeriod        int
}

// NewAroonWithoutStorage creates an Aroon (Aroon) without storage
func NewAroonWithoutStorage(timePeriod int, valueAvailableAction ValueAvailableActionAroon) (indicator *AroonWithoutStorage, err error) {

	// an indicator without storage MUST have a value available action
	if valueAvailableAction == nil {
		return nil, ErrValueAvailableActionIsNil
	}

	// the minimum timeperiod for an Aroon indicator is 2
	if timePeriod < 2 {
		return nil, errors.New("timePeriod is less than the minimum (2)")
	}

	// check the maximum timeperiod
	if timePeriod > MaximumLookbackPeriod {
		return nil, errors.New("timePeriod is greater than the maximum (100000)")
	}

	lookback := timePeriod
	ind := AroonWithoutStorage{
		baseIndicatorWithFloatBoundsAroon: newBaseIndicatorWithFloatBoundsAroon(lookback, valueAvailableAction),
		periodCounter:                     (timePeriod + 1) * -1,
		periodHighHistory:                 list.New(),
		periodLowHistory:                  list.New(),
		aroonFactor:                       100.0 / float64(timePeriod),
	}

	return &ind, nil
}

// An Aroon (Aroon)
type Aroon struct {
	*AroonWithoutStorage

	// public variables
	Up   []float64
	Down []float64
}

// NewAroon creates an Aroon (Aroon) for online usage
func NewAroon(timePeriod int) (indicator *Aroon, err error) {
	ind := Aroon{}
	ind.AroonWithoutStorage, err = NewAroonWithoutStorage(timePeriod,
		func(dataItemAroonUp float64, dataItemAroonDown float64, streamBarIndex int) {
			ind.Up = append(ind.Up, dataItemAroonUp)
			ind.Down = append(ind.Down, dataItemAroonDown)
		})
	return &ind, err
}

// NewDefaultAroon creates an Aroon (Aroon) for online usage with default parameters
//	- timePeriod: 14
func NewDefaultAroon() (indicator *Aroon, err error) {
	timePeriod := 14
	return NewAroon(timePeriod)
}

// NewAroonWithSrcLen creates an Aroon (Aroon) for offline usage
func NewAroonWithSrcLen(sourceLength uint, timePeriod int) (indicator *Aroon, err error) {
	ind, err := NewAroon(timePeriod)

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Up = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
		ind.Down = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewDefaultAroonWithSrcLen creates an Aroon (Aroon) for offline usage with default parameters
func NewDefaultAroonWithSrcLen(sourceLength uint) (indicator *Aroon, err error) {
	ind, err := NewDefaultAroon()

	// only initialise the storage if there is enough source data to require it
	if sourceLength-uint(ind.GetLookbackPeriod()) > 1 {
		ind.Up = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
		ind.Down = make([]float64, 0, sourceLength-uint(ind.GetLookbackPeriod()))
	}

	return ind, err
}

// NewAroonForStream creates an Aroon (Aroon) for online usage with a source data stream
func NewAroonForStream(priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Aroon, err error) {
	ind, err := NewAroon(timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultAroonForStream creates an Aroon (Aroon) for online usage with a source data stream
func NewDefaultAroonForStream(priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Aroon, err error) {
	ind, err := NewDefaultAroon()
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewAroonForStreamWithSrcLen creates an Aroon (Aroon) for online usage with a source data stream
func NewAroonForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber, timePeriod int) (indicator *Aroon, err error) {
	ind, err := NewAroonWithSrcLen(sourceLength, timePeriod)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// NewDefaultAroonForStreamWithSrcLen creates an Aroon (Aroon) for online usage with a source data stream
func NewDefaultAroonForStreamWithSrcLen(sourceLength uint, priceStream gotrade.DOHLCVStreamSubscriber) (indicator *Aroon, err error) {
	ind, err := NewDefaultAroonWithSrcLen(sourceLength)
	priceStream.AddTickSubscription(ind)
	return ind, err
}

// ReceiveDOHLCVTick consumes a source data DOHLCV price tick
func (ind *AroonWithoutStorage) ReceiveDOHLCVTick(tickData gotrade.DOHLCV, streamBarIndex int) {
	ind.periodCounter += 1
	ind.periodHighHistory.PushBack(tickData.H())
	ind.periodLowHistory.PushBack(tickData.L())

	if ind.periodHighHistory.Len() > (1 + ind.GetLookbackPeriod()) {
		var first = ind.periodHighHistory.Front()
		ind.periodHighHistory.Remove(first)
		first = ind.periodLowHistory.Front()
		ind.periodLowHistory.Remove(first)
	}

	if ind.periodCounter >= 0 {
		var aroonUp float64
		var aroonDwn float64

		var highValue float64 = math.SmallestNonzeroFloat64
		var highIdx int = -1
		var i int = (1 + ind.GetLookbackPeriod())
		for e := ind.periodHighHistory.Front(); e != nil; e = e.Next() {
			i--
			var value float64 = e.Value.(float64)
			if highValue <= value {
				highValue = value
				highIdx = i
			}
		}
		var daysSinceHigh = highIdx

		var lowValue float64 = math.MaxFloat64
		var lowIdx int = -1
		i = (1 + ind.GetLookbackPeriod())
		for e := ind.periodLowHistory.Front(); e != nil; e = e.Next() {
			i--
			var value float64 = e.Value.(float64)
			if lowValue >= value {
				lowValue = value
				lowIdx = i
			}

		}
		var daysSinceLow = lowIdx

		aroonUp = ind.aroonFactor * float64(ind.GetLookbackPeriod()-daysSinceHigh)
		aroonDwn = ind.aroonFactor * float64(ind.GetLookbackPeriod()-daysSinceLow)

		ind.UpdateIndicatorWithNewValue(aroonUp, aroonDwn, streamBarIndex)
	}
}
