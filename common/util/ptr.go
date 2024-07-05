package util

func StringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
