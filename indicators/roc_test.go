package indicators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jaybutera/gotrade"
	"github.com/jaybutera/gotrade/indicators"
)

var _ = Describe("when creating an rocwithoutstorage", func() {
	var (
		indicator      *indicators.RocWithoutStorage
		indicatorError error
	)

	Context("and the indicator was not given a value available action", func() {
		BeforeEach(func() {
			indicator, indicatorError = indicators.NewRocWithoutStorage(4, nil)
		})

		It("the indicator should not be created and return the appropriate error message", func() {
			Expect(indicator).To(BeNil())
			Expect(indicatorError).To(Equal(indicators.ErrValueAvailableActionIsNil))
		})
	})

	Context("and the indicator was given a timePeriod below the minimum", func() {
		BeforeEach(func() {
			indicator, indicatorError = indicators.NewRocWithoutStorage(0, fakeFloatValAvailable)
		})

		It("the indicator should not be created and return the appropriate error message", func() {
			Expect(indicator).To(BeNil())
		})
	})

	Context("and the indicator was given a timePeriod above the maximum", func() {
		BeforeEach(func() {
			indicator, indicatorError = indicators.NewRocWithoutStorage(indicators.MaximumLookbackPeriod+1, fakeFloatValAvailable)
		})

		It("the indicator should not be created and return the appropriate error message", func() {
			Expect(indicator).To(BeNil())
		})
	})
})

var _ = Describe("when calculating an rate of change (roc) with DOHLCV source data", func() {
	var (
		period         int = 7
		indicator      *indicators.Roc
		inputs         IndicatorWithFloatBoundsSharedSpecInputs
		stream         *fakeDOHLCVStreamSubscriber
		indicatorError error
	)

	Context("given the indicator is created via the standard constructor", func() {
		BeforeEach(func() {
			indicator, _ = indicators.NewRoc(period, gotrade.UseClosePrice)

			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has received less ticks than the lookback period", func() {

			BeforeEach(func() {
				for i := 0; i < indicator.GetLookbackPeriod(); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedFewerTicksThanItsLookbackPeriod(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has received ticks equal to the lookback period", func() {

			BeforeEach(func() {
				for i := 0; i <= indicator.GetLookbackPeriod(); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedTicksEqualToItsLookbackPeriod(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)
		})

		Context("and the indicator has received more ticks than the lookback period", func() {

			BeforeEach(func() {
				for i := range sourceDOHLCVData {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedMoreTicksThanItsLookbackPeriod(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)
		})
	})

	Context("given the indicator is created via the standard constructor with a nil data selection func", func() {
		BeforeEach(func() {
			indicator, indicatorError = indicators.NewRoc(period, nil)
		})

		It("the indicator should not be created and return the appropriate error message", func() {
			Expect(indicator).To(BeNil())
			Expect(indicatorError).To(Equal(indicators.ErrDOHLCVDataSelectFuncIsNil))
		})
	})

	Context("given the indicator is created via the constructor with defaulted parameters", func() {
		BeforeEach(func() {
			indicator, _ = indicators.NewDefaultRoc()
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)
		})
	})

	Context("given the indicator is created via the constructor with fixed source length", func() {
		BeforeEach(func() {
			indicator, _ = indicators.NewRocWithSrcLen(uint(len(sourceDOHLCVData)), 4, gotrade.UseClosePrice)
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		It("should have pre-allocated storge for the output data", func() {
			Expect(cap(indicator.Data)).To(Equal(len(sourceDOHLCVData) - indicator.GetLookbackPeriod()))
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)

			It("no new storage capcity should have been allocated", func() {
				Expect(len(indicator.Data)).To(Equal(cap(indicator.Data)))
			})
		})
	})

	Context("given the indicator is created via the constructor with defaulted parameters and fixed source length", func() {
		BeforeEach(func() {
			indicator, _ = indicators.NewDefaultRocWithSrcLen(uint(len(sourceDOHLCVData)))
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		It("should have pre-allocated storge for the output data", func() {
			Expect(cap(indicator.Data)).To(Equal(len(sourceDOHLCVData) - indicator.GetLookbackPeriod()))
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)

			It("no new storage capcity should have been allocated", func() {
				Expect(len(indicator.Data)).To(Equal(cap(indicator.Data)))
			})
		})
	})

	Context("given the indicator is created via the constructor for use with a price stream", func() {
		BeforeEach(func() {
			stream = newFakeDOHLCVStreamSubscriber()
			indicator, _ = indicators.NewRocForStream(stream, 4, gotrade.UseClosePrice)
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		It("should have requested to be attached to the stream", func() {
			Expect(stream.lastCallToAddTickSubscriptionArg).To(Equal(indicator))
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)
		})
	})

	Context("given the indicator is created via the constructor for use with a price stream with defaulted parameters", func() {
		BeforeEach(func() {
			stream = newFakeDOHLCVStreamSubscriber()
			indicator, _ = indicators.NewDefaultRocForStream(stream)
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		It("should have requested to be attached to the stream", func() {
			Expect(stream.lastCallToAddTickSubscriptionArg).To(Equal(indicator))
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)
		})
	})

	Context("given the indicator is created via the constructor for use with a price stream with fixed source length", func() {
		BeforeEach(func() {
			stream = newFakeDOHLCVStreamSubscriber()
			indicator, _ = indicators.NewRocForStreamWithSrcLen(uint(len(sourceDOHLCVData)), stream, 4, gotrade.UseClosePrice)
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		It("should have pre-allocated storge for the output data", func() {
			Expect(cap(indicator.Data)).To(Equal(len(sourceDOHLCVData) - indicator.GetLookbackPeriod()))
		})

		It("should have requested to be attached to the stream", func() {
			Expect(stream.lastCallToAddTickSubscriptionArg).To(Equal(indicator))
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)

			It("no new storage capcity should have been allocated", func() {
				Expect(len(indicator.Data)).To(Equal(cap(indicator.Data)))
			})
		})
	})

	Context("given the indicator is created via the constructor for use with a price stream with fixed source length with defaulted parmeters", func() {
		BeforeEach(func() {
			stream = newFakeDOHLCVStreamSubscriber()
			indicator, _ = indicators.NewDefaultRocForStreamWithSrcLen(uint(len(sourceDOHLCVData)), stream)
			inputs = NewIndicatorWithFloatBoundsSharedSpecInputs(indicator, len(sourceDOHLCVData), indicator,
				func() float64 {
					return GetFloatDataMax(indicator.Data)
				},
				func() float64 {
					return GetFloatDataMin(indicator.Data)
				})
		})

		It("should have pre-allocated storge for the output data", func() {
			Expect(cap(indicator.Data)).To(Equal(len(sourceDOHLCVData) - indicator.GetLookbackPeriod()))
		})

		It("should have requested to be attached to the stream", func() {
			Expect(stream.lastCallToAddTickSubscriptionArg).To(Equal(indicator))
		})

		Context("and the indicator has not yet received any ticks", func() {
			ShouldBeAnInitialisedIndicator(&inputs)

			ShouldNotHaveAnyFloatBoundsSetYet(&inputs)
		})

		Context("and the indicator has recieved all of its ticks", func() {
			BeforeEach(func() {
				for i := 0; i < len(sourceDOHLCVData); i++ {
					indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
				}
			})

			ShouldBeAnIndicatorThatHasReceivedAllOfItsTicks(&inputs)

			ShouldHaveFloatBoundsSetToMinMaxOfResults(&inputs)

			It("no new storage capcity should have been allocated", func() {
				Expect(len(indicator.Data)).To(Equal(cap(indicator.Data)))
			})
		})
	})
})
