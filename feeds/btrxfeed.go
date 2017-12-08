package feeds

import (
   "github.com/jaybutera/gotrade"
   "github.com/toorop/go-bittrex"
)

type BtrxHistFeed struct {
   market string
   interval string
   ticks []DOHLCVDataItem
}

type BtrxLiveFeed struct {
   bittrex *Bittrex
   market string
   interval string
}

// Market format ex. "BTC-USD"
// Interval can be -> ["oneMin", "fiveMin", "thirtyMin", "hour", "day"]
NewBtrxHistFeed(market string, interval string, bittrex *Bittrex) (*BtrxHistFeed) {
   ticks, err := bittrex.GetTicks(market, interval)
   if err != nil {
      return nil, err
   }

   return &BtrxHistFeed{market, interval, ticks}, nil
}

NewBtrxLiveFeed(market string, interval string, bittrex *Bittrex) *BtrxLiveFeed {
   return &BtrxLiveFeed{bittrex, market, interval}
}
