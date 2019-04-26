package auth

import (
	"fmt"
	"log"
	"os"
	"testing"
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
		jwtExp     string
		fail       bool
	}{
		{
			"should pass with correct database values without basic auth credentials",
			"localhost:27017", "auth", "users", "", "",
			"mySuperSecret", "10", false,
		},
		{
			"should pass with correct database values and basic auth credentials",
			"localhost:27017", "auth", "users", "admin",
			"", "mySuperSecret", "10", false,
		},
		{
			"should fail without url",
			"", "auth", "users", "admin",
			"", "mySuperSecret", "10", true,
		},
		{
			"should fail without dbName",
			"localhost:27017", "", "users", "admin",
			"", "mySuperSecret", "10", true,
		},
		{
			"should fail without dbCollection",
			"localhost:27017", "auth", "", "admin",
			"awesomePassword", "mySuperSecret", "10", true,
		},
		{
			"should pass with correct database values and jwt values",
			"localhost:27017", "auth", "users", "admin",
			"awesomePassword", "mySuperSecret", "10", false,
		},
		{
			"should fail with correct database values and NaN jwt exp",
			"localhost:27017", "auth", "users", "admin",
			"awesomePassword", "mySuperSecret", "asd", true,
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

		if c.fail {
			if c.fail && err == nil {
				t.Errorf("Expected an error test: %s but error did not occur", c.name)
			}
			continue
		}

		assertEqual(c.name, c.url, e.DbUrl, t)
		assertEqual(c.name, c.dbName, e.DbName, t)
		assertEqual(c.name, c.collection, e.DbCollection, t)
		assertEqual(c.name, c.username, e.DbUsername, t)
		assertEqual(c.name, c.password, e.DbPassword, t)
		assertEqual(c.name, c.jwtSecret, e.JwtSecret, t)
		assertEqual(c.name, c.jwtExp, fmt.Sprintf("%d", e.JwtExpire), t)

	}
}

func assertEqual(testName, val1, val2 string, t *testing.T) {
	if val1 != val2 {
		t.Errorf("ERROR-TEST: %s expected value %s got %s", testName, val1, val2)
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
