package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// special   = "\"!#€%&/()=?`^*-.,;:<>§°' "
	// swedish   = "åÅäÄöÖ"
	// digits    = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // random int between min and max
}

func RandomString(n int, s string) string {
	var sb strings.Builder
	k := len(s)

	for i := 0; i < n; i++ {
		c := s[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomFullName() string {
	return fmt.Sprintf(
		"%s%s %s%s",
		RandomString(1, uppercase),
		RandomString(5, lowercase),
		RandomString(1, uppercase),
		RandomString(7, lowercase),
	)
}

func RandomDate() time.Time {
	today := time.Now().UTC()
	futureBirthday := today.AddDate(1, 0, 0)
	return futureBirthday
}
