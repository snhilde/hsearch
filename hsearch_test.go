package hsearch

import (
	"testing"
	"math/rand"
	"time"
)


// Build two slices of random numbers: one with our target value and one without.
// Send each slice and target value through LinearInt().
func TestLinearInt(t *testing.T) {
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

	// Test the first slice.
	i, err := LinearInt(listA, target)
	if err != nil {
		t.Log("Expected to find target")
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

	// Test the second slice.
	i, err = LinearInt(listB, target)
	if i != -1 {
		t.Error("Found unexpected target in list B")
	} else if err == nil {
		t.Error("Unexpectedly passed Test 2")
	} else if err.Error() != "Not found" {
		t.Error("Received unexpected error:")
		t.Error(err)
	}
}

// Build two slices of sorted random numbers: one with our target value and one without.
// Send each slice and target value through BinaryInt().
func TestBinaryInt(t *testing.T) {
	var listA []int
	var listB []int

	// Get maximum int.
	maxInt := int((1 << 31) - 1)

	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	target := random.Int()
	length := 10000

	// Create our first slice, which will include our target value.
	for i := 1; i < length; i++ {
		section := float64(i) / float64(length)
		section *= float64(maxInt)
		v := random.Int31n(int32(section))
		listA = append(listA, int(v))
	}
	index := int(random.Int31n(int32(length)))
	listA[index] = target

	// Test the first slice.
	i, err := BinaryInt(listA, target)
	if err != nil {
		t.Log("Expected to find target")
		t.Error(err)
	} else if i != index {
		t.Error("Incorrect index")
		t.Log("Expected index =", index)
		t.Log("Returned index =", i)
	}

	// Create our second slice, which will not include our target value.
	for i := 1; i < length; i++ {
		v := target
		for v == target {
			section := float64(i) / float64(length)
			section *= float64(maxInt)
			v = int(random.Int31n(int32(section)))
		}
		listB = append(listB, v)
	}

	// Test the second slice.
	i, err = BinaryInt(listB, target)
	if i != -1 {
		t.Error("Found unexpected target in list B")
	} else if err == nil {
		t.Error("Unexpectedly passed Test 2")
	} else if err.Error() != "Not found" {
		t.Error("Received unexpected error:")
		t.Error(err)
	}
}
