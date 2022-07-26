package todoist

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

// NOTE: Do not use generics at this time to support versions below go1.18
// func addOptionalValueToMap[T any](m map[string]interface{}, k string, v *T) {
// 	if v != nil {
// 		m[k] = *v
// 	}
// }
