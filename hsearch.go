package hsearch

import (
	"errors"
)


// Find needle's position in haystack.
// Search progresses linearly.
func LinearInt(haystack []int, needle int) (int, error) {
	if len(haystack) < 1 {
		return -1, errors.New("Invalid haystack size")
	}

	for i, v := range haystack {
		if needle == v {
			return i, nil
		}
	}

	return -1, errors.New("Not found")
}
