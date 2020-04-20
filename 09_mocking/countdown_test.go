package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

type CountdownSpy struct {
	Sleeper *SpySleeper
	Buffer  *bytes.Buffer
	Calls   []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *CountdownSpy) Sleep() {
	s.Sleeper.Sleep()
	s.Calls = append(s.Calls, "sleep")
}

func (s *CountdownSpy) Write(b []byte) (int, error) {
	s.Buffer.Write(b)
	s.Calls = append(s.Calls, "write")
	return 0, nil
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := SpySleeper{}
	calls := make([]string, 8)
	spy := CountdownSpy{&sleeper, &buffer, calls}
	Countdown(&spy, &spy)

	expectedCalls := []string{
		sleep, write, sleep, write,
		sleep, write, sleep, write,
	}
	if reflect.DeepEqual(expectedCalls, calls) {
		t.Errorf("expected calls %v got %v", expectedCalls, calls)
	}

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	sleepCalls := sleeper.Calls
	if sleepCalls != 3 {
		t.Errorf("expected sleeper to be called 3 times")
	}
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("expected sleep duration of %v got %v", sleepTime, spyTime.durationSlept)
	}
}
