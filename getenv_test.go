package getenv

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestGetStringEnv(t *testing.T) {
	testCases := []struct {
		desc                 string
		key                  string
		want                 string
		shouldBeCreatedFirst bool
	}{
		{
			desc:                 "ENV string when exists",
			key:                  "TEST_KEY",
			want:                 "test",
			shouldBeCreatedFirst: true,
		},
		{
			desc:                 "ENV string when not exists",
			key:                  "TEST_KEY_THAT_DOES_NOT_EXISTS",
			want:                 "",
			shouldBeCreatedFirst: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.shouldBeCreatedFirst {
				os.Setenv(tC.key, tC.want)
				defer os.Unsetenv(tC.key)
			}
			got := GetStringEnv(tC.key, tC.want)
			if tC.want != got {
				t.Errorf("want %s but got %s\n", tC.want, got)
			}
		})
	}
}

func TestGetIntEnv(t *testing.T) {
	testCases := []struct {
		desc                 string
		key                  string
		want                 int
		shouldBeCreatedFirst bool
	}{
		{
			desc:                 "ENV int when exists",
			key:                  "TEST_KEY",
			want:                 100,
			shouldBeCreatedFirst: true,
		},
		{
			desc:                 "ENV int when not exists",
			key:                  "TEST_KEY_THAT_DOES_NOT_EXISTS",
			want:                 0,
			shouldBeCreatedFirst: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.shouldBeCreatedFirst {
				os.Setenv(tC.key, strconv.Itoa(tC.want))
				defer os.Unsetenv(tC.key)
			}
			got := GetIntEnv(tC.key, tC.want)
			if tC.want != got {
				t.Errorf("want %d but got %d\n", tC.want, got)
			}
		})
	}
}

func TestGetBoolEnv(t *testing.T) {
	testCases := []struct {
		desc                 string
		key                  string
		want                 bool
		shouldBeCreatedFirst bool
	}{
		{
			desc:                 "ENV bool when exists",
			key:                  "TEST_KEY",
			want:                 true,
			shouldBeCreatedFirst: true,
		},
		{
			desc:                 "ENV bool when not exists",
			key:                  "TEST_KEY_THAT_DOES_NOT_EXISTS",
			want:                 false,
			shouldBeCreatedFirst: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.shouldBeCreatedFirst {
				os.Setenv(tC.key, strconv.FormatBool(tC.want))
				defer os.Unsetenv(tC.key)
			}
			got := GetBoolEnv(tC.key, tC.want)
			if tC.want != got {
				t.Errorf("want %v but got %v\n", tC.want, got)
			}
		})
	}
}

func TestGetDurationEnv(t *testing.T) {
	testCases := []struct {
		desc                 string
		key                  string
		want                 time.Duration
		shouldBeCreatedFirst bool
	}{
		{
			desc:                 "ENV time.Duration when exists",
			key:                  "TEST_KEY",
			want:                 1 * time.Second,
			shouldBeCreatedFirst: true,
		},
		{
			desc:                 "ENV time.Duration when not exists",
			key:                  "TEST_KEY_THAT_DOES_NOT_EXISTS",
			want:                 0,
			shouldBeCreatedFirst: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.shouldBeCreatedFirst {
				os.Setenv(tC.key, strconv.Itoa(int(tC.want.Seconds())))
				defer os.Unsetenv(tC.key)
			}
			got := GetDurationEnv(tC.key, tC.want)
			if tC.want != got {
				t.Errorf("want %v but got %v\n", tC.want, got)
			}
		})
	}
}
