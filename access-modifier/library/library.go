package library

const secondInMinute int = 60
const minutesInHour int = 60
const HourInADay int = 24

func daysToHour(day int) int {
	return day * HourInADay
}

func DaysToMinutes(day int) int {
	return day * HourInADay * minutesInHour
}
