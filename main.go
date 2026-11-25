package main

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

type WeatherCondition int

const (
	Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
	Rain
	HeavyRain
	Snow
)

const (
	minPrice = 99.0
	maxPrice = 20000
)

type TripParameters struct {
	Distance, Duration float64
}

func CalculateBasePrice(t TripParameters) float64 {
	return t.Distance*pricePerKm + t.Duration*pricePerMinute
}

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	k := 1.0
	switch weather.Condition {
	case 1:
		k += 0.125
	case 2:
		k += 0.2
	case 3:
		k += 0.15
	}
	if weather.WindSpeed > 15 {
		k += 0.1
	}
	return k
}

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	}
	if price > maxPrice {
		return maxPrice
	}
	return price
}
