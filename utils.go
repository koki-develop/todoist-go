package todoist

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// Returns a string as a pointer.
func String(s string) *string { return &s }

// Returns a int as a pointer.
func Int(i int) *int { return &i }

// Returns ints as a pointer.
func Ints(is ...int) *[]int { return &is }

// Returns a bool as a pointer.
func Bool(b bool) *bool { return &b }

func intsToStrings(is []int) []string {
	ss := []string{}
	for _, i := range is {
		ss = append(ss, strconv.Itoa(i))
	}
	return ss
}

func toMap(obj interface{}, dest map[string]interface{}) error {
	if obj == nil {
		return nil
	}

	var m map[string]interface{}
	cfg := &mapstructure.DecoderConfig{TagName: "json", Result: &m}
	dec, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return err
	}

	if err := dec.Decode(obj); err != nil {
		return err
	}

	for k, v := range m {
		dest[k] = v
	}

	return nil
}
