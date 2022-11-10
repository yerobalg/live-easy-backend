package null

func StringFrom(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func Int64From(i int64) *int64 {
	if i == int64(0) {
		return nil
	}
	return &i
}

func Float64From(f float64) *float64 {
	if f == float64(0) {
		return nil
	}
	return &f
}

func BoolFrom(b bool) *bool {
	if b == false {
		return nil
	}
	return &b
}
