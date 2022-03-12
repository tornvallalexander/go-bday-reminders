package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// special   = "\"!#€%&/()=?`^*-.,;:<>§°' "
	// swedish   = "åÅäÄöÖ"
	digits = "0123456789"
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

func RandomNumber(n int, s string) int {
	var sb strings.Builder
	k := len(s)

	for i := 0; i < n; i++ {
		c := s[rand.Intn(k)]
		sb.WriteByte(c)
	}

	num, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatal("Could not convert string to int:", err)
	}

	return num
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

// RandomPnr returns a randomly generated personal number (12 digits)
func RandomPnr() int64 {
	return int64(RandomNumber(12, digits))
}

// RandomPhoneNumber returns a randomly generated phone number (10 digits)
func RandomPhoneNumber() int64 {
	return int64(RandomNumber(10, digits))
}

// RandomUserName returns a randomly generated username (11 characters)
func RandomUserName() string {
	return RandomString(11, lowercase+uppercase+digits)
}

// RandomHashedPassword returns a randomly generated and hashed password (15 characters)
func RandomHashedPassword() string {
	return RandomString(15, lowercase+uppercase+digits)
}

// RandomEmail returns a randomly generated email (12 characters before "@")
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(12, lowercase+uppercase+digits))
}
