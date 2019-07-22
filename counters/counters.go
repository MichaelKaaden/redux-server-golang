// Package counters handles HTTP requests regarding counters.
package counters

import (
	"fmt"
	"sort"
)

type Counter struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

type Counters []Counter

// New returns an initial Counters structure.
func New() *Counters {
	return &Counters{}
}

// GetCounter returns the counter for the given index. If
// necessary, it creates a new one.
func GetCounter(cp *Counters, i int) Counter {
	for _, c := range *cp {
		if c.Index == i {
			return c
		}
	}

	newCounter := Counter{Index: i, Value: 0}
	*cp = append(*cp, newCounter)

	sort.Slice(*cp, func(i, j int) bool {
		return (*cp)[i].Index < (*cp)[j].Index
	})

	return newCounter
}

// SetCounter sets a counter with given index to the given
// value and returns this counter. If the counter does not yet exist,
// it returns an error.
func SetCounter(cp *Counters, i, v int) (Counter, error) {
	for idx := range *cp {
		// operate on the counter itself, not on a copy
		if (*cp)[idx].Index == i {
			(*cp)[idx].Value = v
			return (*cp)[idx], nil
		}
	}

	return Counter{}, fmt.Errorf("counter with index %d not found", i)
}

// Increment increments a counter by a given value.
func Increment(cp *Counters, i, by int) (Counter, error) {
	c := GetCounter(cp, i)
	c, err := SetCounter(cp, i, c.Value+by)
	if err != nil {
		return Counter{}, err
	}

	return c, nil
}

// Decrement decrements a counter by a given value.
func Decrement(cp *Counters, i, by int) (Counter, error) {
	c := GetCounter(cp, i)
	c, err := SetCounter(cp, i, c.Value-by)
	if err != nil {
		return Counter{}, err
	}

	return c, nil
}
