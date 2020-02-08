package hsearch

import (
	"testing"
	"math/rand"
	"time"
	"sort"
)


// Build two slices of random numbers: one with our target value and one without.
// Send each slice and target value through the provided hsearc function.
func testSearch(t *testing.T, searchFunc func([]int, int)(int, error), iters int, sorted bool) {
	for i := 0; i < iters; i++ {
		var listA []int
		var listB []int

		seed   := time.Now().UnixNano()
		source := rand.NewSource(seed)
		random := rand.New(source)

		target := random.Int()
		length := 10000

		// Create our first slice, which will include our target value.
		for i := 0; i < length; i++ {
			v := random.Int()
			listA = append(listA, v)
		}
		index := int(random.Int31n(int32(length)))
		listA[index] = target
		if sorted {
			// Sort the list and find our new target index.
			sort.Ints(listA)
			index = sort.SearchInts(listA, target)
		}

		// Test the first slice.
		i, err := searchFunc(listA, target)
		if err != nil {
			t.Log("Expected to find target at index", index)
			t.Error(err)
		} else if i != index {
			t.Error("Incorrect index")
			t.Log("Expected index =", index)
			t.Log("Returned index =", i)
		}

		// Create our second slice, which will not include our target value.
		for i := 0; i < length; i++ {
			v := target
			for v == target {
				v = random.Int()
			}
			listB = append(listB, v)
		}
		if sorted {
			sort.Ints(listB)
		}

		// Test the second slice.
		i, err = searchFunc(listB, target)
		if i != -1 {
			t.Error("Found unexpected target in list B")
		} else if err == nil {
			t.Error("Unexpectedly passed Test 2")
		} else if err.Error() != "Not found" {
			t.Error("Received unexpected error:")
			t.Error(err)
		}
	}
}

func TestLinearInt(t *testing.T) {
	testSearch(t, LinearInt, 100, false)
}

func TestBinaryInt(t *testing.T) {
	testSearch(t, BinaryInt, 100, true)
}
