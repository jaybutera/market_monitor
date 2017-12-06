package indicators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jaybutera/gotrade"
	"github.com/jaybutera/gotrade/indicators"
)

var _ = Describe("when executing the gotrade simple moving average with a years data and known output", func() {
	var (
		sma             *indicators.Sma
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("sma_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			sma, err = indicators.NewSma(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(sma)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(sma.Data)).To(Equal(len(priceStream.Data) - sma.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the simple moving average for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", sma.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade exponential moving average with a years data and known output", func() {
	var (
		ema             *indicators.Ema
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("ema_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			ema, err = indicators.NewEma(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ema)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(ema.Data)).To(Equal(len(priceStream.Data) - ema.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the exponential moving average for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ema.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade weighted moving average with a years data and known output", func() {
	var (
		wma             *indicators.Wma
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("wma_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			wma, err = indicators.NewWma(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(wma)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(wma.Data)).To(Equal(len(priceStream.Data) - wma.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the weighted moving average for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", wma.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade double exponential moving average with a years data and known output", func() {
	var (
		dema            *indicators.Dema
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("dema_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			dema, err = indicators.NewDema(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(dema)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(dema.Data)).To(Equal(len(priceStream.Data) - (dema.GetLookbackPeriod())))
		})

		It("it should have correctly calculated the double exponential moving average for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", dema.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade triple exponential moving average with a years data and known output", func() {
	var (
		tema            *indicators.Tema
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("tema_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			tema, err = indicators.NewTema(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(tema)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(tema.Data)).To(Equal(len(priceStream.Data) - (tema.GetLookbackPeriod())))
		})

		It("it should have correctly calculated the triple exponential moving average for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", tema.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade variance with a years data and known output", func() {
	var (
		variance        *indicators.Var
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("variance_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			variance, err = indicators.NewVar(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(variance)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(variance.Data)).To(Equal(len(priceStream.Data) - variance.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the variance for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", variance.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade standard deviation with a years data and known output", func() {
	var (
		stdDev          *indicators.StdDev
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("stddev_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			stdDev, err = indicators.NewStdDev(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(stdDev)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(stdDev.Data)).To(Equal(len(priceStream.Data) - stdDev.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the standard deviation for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", stdDev.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade bollinger bands with a years data and known output", func() {
	var (
		bb              *indicators.BollingerBands
		period          int
		expectedResults []BollingerBand
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVBollingerPriceDataFromFile("bb_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			bb, err = indicators.NewBollingerBands(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(bb)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(bb.Length()).To(Equal(len(priceStream.Data) - bb.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the bollinger upper, middle and lower bands for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].U()).To(BeNumerically("~", bb.UpperBand[k], 0.01))
				Expect(expectedResults[k].M()).To(BeNumerically("~", bb.MiddleBand[k], 0.01))
				Expect(expectedResults[k].L()).To(BeNumerically("~", bb.LowerBand[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade macd with a years data and known output", func() {
	var (
		macd            *indicators.Macd
		expectedResults []MacdData
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVMacdPriceDataFromFile("macd_12_26_9_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback periods of 12, 26, 9", func() {

		BeforeEach(func() {
			macd, err = indicators.NewMacd(12, 26, 9, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(macd)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(macd.Length()).To(Equal(len(priceStream.Data) - macd.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the macd, signal and histogram for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].M()).To(BeNumerically("~", macd.Macd[k], 0.01))
				Expect(expectedResults[k].S()).To(BeNumerically("~", macd.Signal[k], 0.01))
				Expect(expectedResults[k].H()).To(BeNumerically("~", macd.Histogram[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade aroon with a years data and known output", func() {
	var (
		aroon           *indicators.Aroon
		expectedResults []AroonData
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVAroonPriceDataFromFile("aroon_25_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback periods of 25", func() {

		BeforeEach(func() {
			aroon, err = indicators.NewAroon(25)
			priceStream.AddTickSubscription(aroon)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(aroon.Length()).To(Equal(len(priceStream.Data) - aroon.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the aroon up for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].U()).To(BeNumerically("~", aroon.Up[k], 0.01))
			}
		})

		It("it should have correctly calculated the aroon down for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].D()).To(BeNumerically("~", aroon.Down[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade aroonosc with a years data and known output", func() {
	var (
		aroon           *indicators.AroonOsc
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("aroonosc_25_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback periods of 25", func() {

		BeforeEach(func() {
			aroon, err = indicators.NewAroonOsc(25)
			priceStream.AddTickSubscription(aroon)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(aroon.Length()).To(Equal(len(priceStream.Data) - aroon.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the aroon oscillator for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", aroon.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade truerange with a years data and known output", func() {
	var (
		trueRange       *indicators.TrueRange
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("truerange_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using an implicit lookback period of 1", func() {

		BeforeEach(func() {
			trueRange, err = indicators.NewTrueRange()
			priceStream.AddTickSubscription(trueRange)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookback period", func() {
			Expect(trueRange.Length()).To(Equal(len(priceStream.Data) - trueRange.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the truerangefor each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", trueRange.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade average truerange with a years data and known output", func() {
	var (
		avgTrueRange    *indicators.Atr
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("atr_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using an implicit lookback period of 14", func() {

		BeforeEach(func() {
			avgTrueRange, err = indicators.NewAtr(14)
			priceStream.AddTickSubscription(avgTrueRange)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookback period", func() {
			Expect(avgTrueRange.Length()).To(Equal(len(priceStream.Data) - avgTrueRange.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the truerangefor each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", avgTrueRange.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade accumulation distribution line with a years data and known output", func() {
	var (
		adl             *indicators.Adl
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("adl_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using no lookback", func() {

		BeforeEach(func() {
			adl, err = indicators.NewAdl()
			priceStream.AddTickSubscription(adl)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(adl.Length()).To(Equal(len(priceStream.Data)))
		})

		It("it should have correctly calculated the truerangefor each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", adl.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade chaikin oscilator with a years data and known output", func() {
	var (
		chaikinOsc      *indicators.ChaikinOsc
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
		fastPeriod      int
		slowPeriod      int
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("chaikinosc_3_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
		fastPeriod = 3
		slowPeriod = 10
	})

	Describe("using no a fast Time Period of 3 and a slow Time Period of 10", func() {

		BeforeEach(func() {
			chaikinOsc, err = indicators.NewChaikinOsc(3, 10)
			priceStream.AddTickSubscription(chaikinOsc)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length - the lookback Period", func() {
			Expect(chaikinOsc.Length()).To(Equal(len(priceStream.Data) - chaikinOsc.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the chaikin oscillator for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", chaikinOsc.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade on balance volume indicator with a years data and known output", func() {
	var (
		obv             *indicators.Obv
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("obv_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using no lookback period", func() {

		BeforeEach(func() {
			obv, err = indicators.NewObv()
			priceStream.AddTickSubscription(obv)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(obv.Length()).To(Equal(len(priceStream.Data)))
		})

		It("it should have correctly calculated the on balance volume for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", obv.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade average price indicator with a years data and known output", func() {
	var (
		avgPrice        *indicators.AvgPrice
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("avgprice_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using no lookback period", func() {

		BeforeEach(func() {
			avgPrice, err = indicators.NewAvgPrice()
			priceStream.AddTickSubscription(avgPrice)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(avgPrice.Length()).To(Equal(len(priceStream.Data)))
		})

		It("it should have correctly calculated the avg price for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", avgPrice.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade median price indicator with a years data and known output", func() {
	var (
		medPrice        *indicators.MedPrice
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("medprice_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using no lookback period", func() {

		BeforeEach(func() {
			medPrice, err = indicators.NewMedPrice()
			priceStream.AddTickSubscription(medPrice)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(medPrice.Length()).To(Equal(len(priceStream.Data)))
		})

		It("it should have correctly calculated the median price for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", medPrice.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade typical price indicator with a years data and known output", func() {
	var (
		typPrice        *indicators.TypPrice
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("typprice_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using no lookback period", func() {

		BeforeEach(func() {
			typPrice, err = indicators.NewTypPrice()
			priceStream.AddTickSubscription(typPrice)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(typPrice.Length()).To(Equal(len(priceStream.Data)))
		})

		It("it should have correctly calculated the typical price for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", typPrice.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade plus directional movement indicator (1) with a years data and known output", func() {
	var (
		plusDM          *indicators.PlusDm
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("plusdm_1_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 1", func() {

		BeforeEach(func() {
			plusDM, err = indicators.NewPlusDm(1)
			priceStream.AddTickSubscription(plusDM)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(plusDM.Length()).To(Equal(len(priceStream.Data) - plusDM.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the plus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", plusDM.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade plus directional movement indicator (14) with a years data and known output", func() {
	var (
		plusDM          *indicators.PlusDm
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("plusdm_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			plusDM, err = indicators.NewPlusDm(14)
			priceStream.AddTickSubscription(plusDM)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(plusDM.Length()).To(Equal(len(priceStream.Data) - plusDM.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the plus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", plusDM.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade minus directional movement indicator (1) with a years data and known output", func() {
	var (
		minusDM         *indicators.MinusDm
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("minusdm_1_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 1", func() {

		BeforeEach(func() {
			minusDM, err = indicators.NewMinusDm(1)
			priceStream.AddTickSubscription(minusDM)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(minusDM.Length()).To(Equal(len(priceStream.Data) - minusDM.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the minus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", minusDM.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade minus directional movement indicator (14) with a years data and known output", func() {
	var (
		minusDM         *indicators.MinusDm
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("minusdm_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			minusDM, err = indicators.NewMinusDm(14)
			priceStream.AddTickSubscription(minusDM)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(minusDM.Length()).To(Equal(len(priceStream.Data) - minusDM.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the minus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", minusDM.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade plus directional indicator (1) with a years data and known output", func() {
	var (
		plusDI          *indicators.PlusDi
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("plusdi_1_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 1", func() {

		BeforeEach(func() {
			plusDI, err = indicators.NewPlusDi(1)
			priceStream.AddTickSubscription(plusDI)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(plusDI.Length()).To(Equal(len(priceStream.Data) - plusDI.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the plus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", plusDI.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade plus directional indicator (14) with a years data and known output", func() {
	var (
		plusDI          *indicators.PlusDi
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("plusdi_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			plusDI, err = indicators.NewPlusDi(14)
			priceStream.AddTickSubscription(plusDI)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(plusDI.Length()).To(Equal(len(priceStream.Data) - plusDI.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the plus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", plusDI.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade minus directional indicator (1) with a years data and known output", func() {
	var (
		minusDI         *indicators.MinusDi
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("minusdi_1_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 1", func() {

		BeforeEach(func() {
			minusDI, err = indicators.NewMinusDi(1)
			priceStream.AddTickSubscription(minusDI)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(minusDI.Length()).To(Equal(len(priceStream.Data) - minusDI.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the minus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", minusDI.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade minus directional indicator (14) with a years data and known output", func() {
	var (
		minusDI         *indicators.MinusDi
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("minusdi_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			minusDI, err = indicators.NewMinusDi(14)
			priceStream.AddTickSubscription(minusDI)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(minusDI.Length()).To(Equal(len(priceStream.Data) - minusDI.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the minus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", minusDI.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade directional movement indicator (14) with a years data and known output", func() {
	var (
		dx              *indicators.Dx
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("dx_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			dx, err = indicators.NewDx(14)
			priceStream.AddTickSubscription(dx)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(dx.Length()).To(Equal(len(priceStream.Data) - dx.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the minus directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", dx.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade average directional movement indicator (14) with a years data and known output", func() {
	var (
		adx             *indicators.Adx
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("adx_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			adx, err = indicators.NewAdx(14)
			priceStream.AddTickSubscription(adx)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(adx.Length()).To(Equal(len(priceStream.Data) - adx.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the average directional movement for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", adx.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade average directional movement rating (14) with a years data and known output", func() {
	var (
		adxr            *indicators.Adxr
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("adxr_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			adxr, err = indicators.NewAdxr(14)
			priceStream.AddTickSubscription(adxr)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(adxr.Length()).To(Equal(len(priceStream.Data) - adxr.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the average directional rating for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", adxr.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade relative strength index with a years data and known output", func() {
	var (
		rsi             *indicators.Rsi
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("rsi_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 14", func() {

		BeforeEach(func() {
			period = 14
			rsi, err = indicators.NewRsi(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(rsi)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(rsi.Data)).To(Equal(len(priceStream.Data) - rsi.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the standard deviation for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", rsi.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade momentum with a years data and known output", func() {
	var (
		ind             *indicators.Mom
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("mom_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			ind, err = indicators.NewMom(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(ind.Data)).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the momentum for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade rate of change with a years data and known output", func() {
	var (
		ind             *indicators.Roc
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("roc_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			ind, err = indicators.NewRoc(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(ind.Data)).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the standard deviation for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade rate of change percentage with a years data and known output", func() {
	var (
		ind             *indicators.RocP
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("rocp_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			ind, err = indicators.NewRocP(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(ind.Data)).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the standard deviation for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade rate of change ratio with a years data and known output", func() {
	var (
		ind             *indicators.RocR
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("rocr_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			ind, err = indicators.NewRocR(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(ind.Data)).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the standard deviation for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade rate of change ratio 100 scale with a years data and known output", func() {
	var (
		ind             *indicators.RocR100
		period          int
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("rocr100_10_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback period of 10", func() {

		BeforeEach(func() {
			period = 10
			ind, err = indicators.NewRocR100(period, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(len(ind.Data)).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the standard deviation for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.1))
			}
		})
	})
})

var _ = Describe("when executing the gotrade money flow index (14) with a years data and known output", func() {
	var (
		ind             *indicators.Mfi
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("mfi_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewMfi(14)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the money flow index for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade parabolic stop and reverse (Sar) with a years data and known output", func() {
	var (
		ind             *indicators.Sar
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("sar_002_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using an accelleration factor of 0.02", func() {

		BeforeEach(func() {
			ind, err = indicators.NewSar(0.02, 0.20)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the Sar for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade linearregression (LinReg) with a years data and known output", func() {
	var (
		ind             *indicators.LinReg
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("linear_regression_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewLinReg(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the linear regression for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade linearregression slope (LinRegSlp) with a years data and known output", func() {
	var (
		ind             *indicators.LinRegSlp
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("linear_regression_slope_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewLinRegSlp(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the linear regression slope for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade linearregression intercept (LinRegInt) with a years data and known output", func() {
	var (
		ind             *indicators.LinRegInt
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("linear_regression_intercept_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewLinRegInt(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the linear regression intercept for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade linearregression angle (LinRegAng) with a years data and known output", func() {
	var (
		ind             *indicators.LinRegAng
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("linear_regression_angle_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewLinRegAng(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the linear regression angle for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade time series forecast (Tsf) with a years data and known output", func() {
	var (
		ind             *indicators.Tsf
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("tsf_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewTsf(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the times series forecast for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade kaufman adaptive moving average (Kama) with a years data and known output", func() {
	var (
		ind             *indicators.Kama
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("kama_30_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 30", func() {

		BeforeEach(func() {
			ind, err = indicators.NewKama(30, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the kama for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade triangular moving average (Trima) with a years data and known output", func() {
	var (
		ind             *indicators.Trima
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("trima_30_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 30", func() {

		BeforeEach(func() {
			ind, err = indicators.NewTrima(30, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the trima for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade williams percent r (WillR) with a years data and known output", func() {
	var (
		ind             *indicators.WillR
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("willr_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewWillR(14)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the willr for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade highest high value (Hhv) with a years data and known output", func() {
	var (
		ind             *indicators.Hhv
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("hhv_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewHhv(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the hhv for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade lowest low value (Llv) with a years data and known output", func() {
	var (
		ind             *indicators.Llv
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("llv_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewLlv(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the llv for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade highest high bars (HhvBars) with a years data and known output", func() {
	var (
		ind             *indicators.HhvBars
		expectedResults []int64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVIntPriceDataFromFile("hhvbars_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewHhvBars(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the hhv for each item in the result", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(Equal(ind.Data[k]))
			}
		})
	})
})

var _ = Describe("when executing the gotrade lowest low bars (LlvBars) with a years data and known output", func() {
	var (
		ind             *indicators.LlvBars
		expectedResults []int64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVIntPriceDataFromFile("llvbars_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewLlvBars(14, gotrade.UseClosePrice)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the llv for each item in the result", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(Equal(ind.Data[k]))
			}
		})
	})
})

var _ = Describe("when executing the gotrade stochastic oscillator with a years data and known output", func() {
	var (
		stoch           *indicators.StochOsc
		expectedResults []StochData
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVStochPriceDataFromFile("stoch_5_3_3_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback periods of 5,3,3", func() {

		BeforeEach(func() {
			stoch, err = indicators.NewStochOsc(5, 3, 3)
			priceStream.AddTickSubscription(stoch)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(stoch.Length()).To(Equal(len(priceStream.Data) - stoch.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the stoch slowk for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].K()).To(BeNumerically("~", stoch.SlowK[k], 0.01))
			}
		})

		It("it should have correctly calculated the stoch slowd for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].D()).To(BeNumerically("~", stoch.SlowD[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade stochastic rsi oscillator with a years data and known output", func() {
	var (
		stoch           *indicators.StochRsi
		expectedResults []StochData
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVStochPriceDataFromFile("stochrsi_14_5_3_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a lookback periods of 14,5,3", func() {

		BeforeEach(func() {
			stoch, err = indicators.NewStochRsi(14, 5, 3)
			priceStream.AddTickSubscription(stoch)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length less the lookbackperiod", func() {
			Expect(stoch.Length()).To(Equal(len(priceStream.Data) - stoch.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the stoch rsi fastk for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].K()).To(BeNumerically("~", stoch.SlowK[k], 0.01))
			}
		})

		It("it should have correctly calculated the stoch rsi fastd for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k].D()).To(BeNumerically("~", stoch.SlowD[k], 0.01))
			}
		})
	})
})

var _ = Describe("when executing the gotrade commodity channel index (Cci) with a years data and known output", func() {
	var (
		ind             *indicators.Cci
		expectedResults []float64
		err             error
		priceStream     *gotrade.InterDayDOHLCVStream
	)

	BeforeEach(func() {
		// load the expected results data
		expectedResults, _ = LoadCSVPriceDataFromFile("cci_14_expectedresult.data")
		priceStream = gotrade.NewDailyDOHLCVStream()
	})

	Describe("using a time period of 14", func() {

		BeforeEach(func() {
			ind, err = indicators.NewCci(14)
			priceStream.AddTickSubscription(ind)
			csvFeed.FillDOHLCVStream(priceStream)
		})

		It("the result set should have a length equal to the source data length", func() {
			Expect(ind.Length()).To(Equal(len(priceStream.Data) - ind.GetLookbackPeriod()))
		})

		It("it should have correctly calculated the cci for each item in the result set accurate to two decimal places", func() {
			for k := range expectedResults {
				Expect(expectedResults[k]).To(BeNumerically("~", ind.Data[k], 0.01))
			}
		})
	})
})
