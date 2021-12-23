package birdwatcher

//import "fmt"
// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
     var sum = 0
	 for _, value := range birdsPerDay {
        sum += value
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var start int = 7* (week - 1)
    var end  int  = start + 7
    return TotalBirdCount(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i+=2 {
       birdsPerDay[i] += 1
    }
    return birdsPerDay
}
