package utils

func ShortHash(c string) string {
	if len(c) > 7 {
		return c[:7]
	}
	return c
}