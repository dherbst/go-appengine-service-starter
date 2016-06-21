package starter

import (
	"os"
	"strconv"
	"time"
)

// Config holds the configuration necessary for this service.
type Config struct {
	GitBranch string
	GitHash   string
	GitTime   string
	BuildTime time.Time
}

// NewConfig creates a Config object and reads environment variables for values.
func NewConfig() Config {
	c := Config{}
	c.GitTime = os.Getenv("GIT_TIME")
	c.GitBranch = os.Getenv("GIT_BRANCH")
	c.GitHash = os.Getenv("GIT_HASH")
	c.BuildTime = convertTime(c.GitTime)

	return c
}

// convertTime converts the seconds from epoch into a golang time struct.  If there is an error
// computing the time, a time for date 1/1/1972 is returned.
func convertTime(timeString string) time.Time {
	if timeString == "" {
		return time.Unix(0, 0)
	}

	i, err := strconv.ParseInt(timeString, 10, 64)
	if err != nil {
		return time.Unix(63072000, 0)
	}

	return time.Unix(i, 0)
}
