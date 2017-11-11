package gigasecond

// import path for the time package from the standard library
import "time"
import "math"

const testVersion = 4

func AddGigasecond(dob time.Time) time.Time {
	duration := time.Duration(math.Pow10(18))
	return dob.Add(duration)
}
