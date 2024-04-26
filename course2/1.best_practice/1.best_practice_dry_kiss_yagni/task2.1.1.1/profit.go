package main

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)            //  устанавливает продукт для статистики прибыли.
	GetAverageProfit() float64        //  возвращает среднюю прибыль.
	GetAverageProfitPercent() float64 //  возвращает средний процент прибыли.
	GetCurrentProfit() float64        //  возвращает текущую прибыль.
	GetDifferenceProfit() float64     //  возвращает разницу между текущей ценой продукта и средней ценой продажи.
	GetAllData() []float64            //  возвращает все данные о прибыли в виде среза чисел.
	Average(prices []float64) float64 //  вычисляет среднее значение из среза чисел.
	Sum(prices []float64) float64     //  вычисляет сумму чисел в срезе.
}

type StatisticProfit struct {
	product *Product
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

func (s *StatisticProfit) GetAverageProfit() float64 {
	return s.Average(s.product.Sells) - s.Average(s.product.Buys)
}
func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	return s.GetAverageProfit() / s.Average(s.product.Buys) * 100
}
func (s *StatisticProfit) GetCurrentProfit() float64 {
	return s.Sum(s.product.Sells) - s.Sum(s.product.Buys)
}
func (s *StatisticProfit) GetDifferenceProfit() float64 {
	return s.product.CurrentPrice - s.Average(s.product.Sells)
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	return s.Sum(prices) / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	var res float64
	for _, xs := range prices {
		res += xs
	}
	return res
}

func (s *StatisticProfit) getAllData() []float64 {
	res := make([]float64, 0, 4)
	res = append(res, s.GetAverageProfit())
	res = append(res, s.GetAverageProfitPercent())
	res = append(res, s.GetCurrentProfit())
	res = append(res, s.GetDifferenceProfit())
	return res
}
