package auth

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetEnv(t *testing.T) {
	// cases
	cases := []struct {
		name       string
		url        string
		dbName     string
		collection string
		username   string
		password   string
		jwtSecret  string
		eJwtSecret string
		jwtExp     string
		eJwtExp    time.Duration
		fail       bool
	}{
		{
			"should pass with correct database values without basic auth credentials",
			"localhost:27017", "auth", "users", "", "",
			"mySuperSecret", "mySuperSecret", "10", time.Duration(10), false,
		},
		{
			"should pass with correct database values and basic auth credentials",
			"localhost:27017", "auth", "users", "admin", "awesomePassword",
			"mySuperSecret", "mySuperSecret", "10", time.Duration(10), false,
		},
		{
			"should fail without url",
			"", "auth", "users", "admin", "awesomePassword",
			"mySuperSecret", "mySuperSecret", "10", time.Duration(10), true,
		},
		{
			"should fail without dbName",
			"localhost:27017", "", "users", "admin", "awesomePassword",
			"mySuperSecret", "mySuperSecret", "10", time.Duration(10), true,
		},
		{
			"should fail without dbCollection",
			"localhost:27017", "auth", "", "admin", "awesomePassword",
			"mySuperSecret", "mySuperSecret", "10", time.Duration(10), true,
		},
		{
			"jwt options should equal to defaults when didn't given",
			"localhost:27017", "auth", "", "admin", "awesomePassword",
			"", "secret-string", "", time.Duration(5), false,
		},
	}

	for _, c := range cases {
		os.Clearenv()
		setEnv("DB_URL", c.url)
		setEnv("DB_NAME", c.dbName)
		setEnv("DB_COLLECTION", c.collection)
		setEnv("DB_USERNAME", c.username)
		setEnv("DB_PASSWORD", c.password)
		setEnv("JWT_SECRET", c.jwtSecret)
		setEnv("JWT_EXPIRE", c.jwtExp)

		e, err := GetEnv()

		if c.fail && err == nil {
			t.Error("Expected an error but error did not occur")
		}

		assertEqual(c.url, e.DbUrl, t)
		assertEqual(c.dbName, e.DbName, t)
		assertEqual(c.collection, e.DbCollection, t)
		assertEqual(c.username, e.DbUsername, t)
		assertEqual(c.password, e.DbPassword, t)
		assertEqual(c.eJwtSecret, e.JwtSecret, t)
		assertEqual(fmt.Sprintf("%d", c.eJwtExp), fmt.Sprintf("%d", e.JwtExpire), t)
	}
}

func assertEqual(val1, val2 string, t *testing.T) {
	if val1 != val2 {
		t.Errorf("Error expected value %s got %s", val1, val2)
	}
}

func setEnv(key, value string) {
	if value != "" {
		err := os.Setenv(key, value)
		if err != nil {
			log.Fatal(err)
		}
	}
}
