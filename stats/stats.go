package stats

import (
	"math"

	"github.com/ren-/osu/api"
)

type Statistics struct {
	HighestPP               float64
	LowestPP                float64
	AveragePP               float64
	StandardDeviation       float64
	MedianAbsoluteDeviation float64
}

func Calculate(scores []api.TopScore) *Statistics {
	stats := new(Statistics)
	stats.HighestPP = scores[0].PP
	stats.LowestPP = scores[0].PP
	var sum float64
	var pp []float64
	for _, element := range scores {
		pp = append(pp, element.PP)

		if element.PP > stats.HighestPP {
			stats.HighestPP = element.PP
		}
		if element.PP < stats.LowestPP {
			stats.LowestPP = element.PP
		}

		sum += element.PP

	}

	stats.AveragePP = sum / float64(len(scores))
	stats.StandardDeviation = stdDev(pp, stats.AveragePP)
	stats.MedianAbsoluteDeviation = medianAbsoluteDeviation(pp)

	return stats

}
func stdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(float64(number-mean), 2)
	}
	variance := total / float64(len(numbers)-1)
	return math.Sqrt(variance)
}
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func absoluteDeviations(numbers []float64, median float64) []float64 {
	var result []float64
	for _, element := range numbers {
		result = append(result, math.Abs(element-median))
	}

	return result
}

func medianAbsoluteDeviation(numbers []float64) float64 {
	ppMedian := median(numbers)
	absoluteDeviations := absoluteDeviations(numbers, ppMedian)
	//fmt.Println(absoluteDeviations)

	return median(absoluteDeviations)

}
