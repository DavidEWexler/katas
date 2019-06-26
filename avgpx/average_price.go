package avgpx

// AveragePrice Calculates the average price for a list of trades
func AveragePrice(trades ...Trade) float64 {
	var shares int
	var value float64
	for _, trade := range trades {
		shares += trade.Shares
		value += float64(trade.Shares) * trade.Price
	}
	if shares > 0 {
		return value / float64(shares)
	}
	return 0.0
}
