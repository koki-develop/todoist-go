package todoist

import (
	"sort"
	"strconv"
)

// Returns a string as a pointer.
func String(s string) *string { return &s }

// Returns a int as a pointer.
func Int(i int) *int { return &i }

// Returns a bool as a pointer.
func Bool(b bool) *bool { return &b }

func addOptionalStringToMap(m map[string]interface{}, k string, v *string) {
	if v != nil {
		m[k] = *v
	}
}

func addOptionalIntToMap(m map[string]interface{}, k string, v *int) {
	if v != nil {
		m[k] = *v
	}
}

func addOptionalBoolToMap(m map[string]interface{}, k string, v *bool) {
	if v != nil {
		m[k] = *v
	}
}

func addOptionalStringToStringMap(m map[string]string, k string, v *string) {
	if v != nil {
		m[k] = *v
	}
}

func addOptionalIntToStringMap(m map[string]string, k string, v *int) {
	if v != nil {
		m[k] = strconv.Itoa(*v)
	}
}

func intsToStrings(is []int) []string {
	ss := []string{}
	for _, i := range is {
		ss = append(ss, strconv.Itoa(i))
	}
	return ss
}

func getKeysFromStringMap(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getSortedKeysFromStringMap(m map[string]string) []string {
	keys := getKeysFromStringMap(m)
	sort.Strings(keys)
	return keys
}
