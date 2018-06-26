package macd

import (
	"fmt"
	"indicator/ema"
)


type Macd struct {
	short  *ema.Ema
	long   *ema.Ema
	signal *ema.Ema
	diff   float64
	dea    float64
	macd   float64
}

func NewMacd(short, long, signal int32) *Macd {
	return &Macd{short:ema.NewEma(short), long:ema.NewEma(long), signal:ema.NewEma(signal)}
}

func (this *Macd)Update(price float64) (float64, float64, float64) {
	s := this.short.Update(price)
	l := this.long.Update(price)
	this.diff = s - l
	this.dea = this.signal.Update(this.diff)
	this.macd = 2.0 * (this.diff - this.dea)

	return this.diff, this.dea, this.macd
}

func (this *Macd)Clone() *Macd {
	return &Macd{short:this.short.Clone(), long:this.long.Clone(), signal:this.signal.Clone(), diff:this.diff, dea:this.dea, macd:this.macd}

}

func main() {
	macd := NewMacd(12, 26, 9)
	var price float64 = 35.66
	f1, f2, f3 := macd.Update(price)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)

}