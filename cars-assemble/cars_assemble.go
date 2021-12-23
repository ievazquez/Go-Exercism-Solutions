package cars

const carsProducedPerHour = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed <= 0 {
        return float64(0.0)
    } else if speed <= 4 {
        return float64(1.0)
    } else if speed <= 8 {
        return float64(0.9)
    } else if speed <= 10 {
        return float64(0.77)
    }
    return float64(0.0);
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * carsProducedPerHour * float64(speed) 
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60) 
    //return int(  (CalculateProductionRatePerHour(1) / 60 ) * float64(speed) )
    //return int(3.3 * float64(speed)) 
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
    var productionPerHour float64 = CalculateProductionRatePerHour(speed)
	if productionPerHour > limit {
        return limit
    }
     return productionPerHour
}
