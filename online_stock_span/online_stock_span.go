package online_stock_span

type Pair struct {
	price int
	index int
}

type StockSpanner struct {
	s []Pair
	n int
}

func Constructor() StockSpanner {
	return StockSpanner{make([]Pair, 0), 0}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.s) > 0 && this.s[len(this.s)-1].price <= price {
		this.s = this.s[:len(this.s)-1]
	}

	this.n++
	days := this.n
	if len(this.s) > 0 {
		days = this.n - this.s[len(this.s)-1].index
	}

	this.s = append(this.s, Pair{price, this.n})
	return days
}
