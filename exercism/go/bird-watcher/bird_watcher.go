package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0
	for count := 0; count < len(birdsPerDay); count++ {
		birdCount += birdsPerDay[count]
	}
	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdCount := 0
	for day := (week-1) * 7; day < week*7; day++ {
		birdCount += birdsPerDay[day]
	}
	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day := 0; day < len(birdsPerDay)-1; day+=2 {
		birdsPerDay[day] +=1
	}
	return birdsPerDay
}
