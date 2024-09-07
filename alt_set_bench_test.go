package swiss

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
)

type testMapType map[uint32]int

func prepareDataUint32(initialSetSize, finalSetSize, searchListSize int, minimalHitRatio float32) (resultSet *Set[uint32], resultMap testMapType, searchElements []uint32) {
	resultSet = NewSet[uint32](uint32(initialSetSize))
	resultMap = make(testMapType, initialSetSize)
	for n := 0; n < finalSetSize; n++ {
		element := rand.Uint32()
		resultSet.Add(element)
		resultMap[element] = 1
	}
	nrOfElemToCopy := int(minimalHitRatio * float32(searchListSize))
	tempList := make([]uint32, 0, searchListSize)
	countCopied := 0
	for countCopied < nrOfElemToCopy {
		resultSet.Iter(func(e uint32) (stop bool) {
			tempList = append(tempList, e)
			countCopied++
			return countCopied >= nrOfElemToCopy
		})
	}
	for n := countCopied; n < searchListSize; n++ {
		element := rand.Uint32()
		tempList = append(tempList, element)
	}
	perm := rand.Perm(searchListSize)
	searchElements = make([]uint32, searchListSize)
	for i, idx := range perm {
		searchElements[i] = tempList[idx]
	}
	return
}

func Algorithm1(resultSet *Set[uint32], resultMap testMapType, searchElements []uint32) {
	x := uint64(0)
	for _, e := range searchElements {
		if resultSet.Contains(e) {
			x += uint64(e)
		}
	}
	resultSet.Clear()
	for k := range resultMap {
		delete(resultMap, k)
	}
}

func Algorithm2(resultSet *Set[uint32], resultMap testMapType, searchElements []uint32) {
	x := uint64(0)
	for _, e := range searchElements {
		_, b := resultMap[e]
		if b {
			x += uint64(e)
		}
	}
	resultSet.Clear()
	for k := range resultMap {
		delete(resultMap, k)
	}
}

func BenchmarkAlgorithm1(b *testing.B) {
	resultSet, resultMap, searchElements := prepareDataUint32(10, 5000, 8000, 0.3)
	for i := 0; i < b.N; i++ {
		Algorithm1(resultSet, resultMap, searchElements)
	}
}

func BenchmarkAlgorithm2(b *testing.B) {
	resultSet, resultMap, searchElements := prepareDataUint32(10, 5000, 8000, 0.3)
	for i := 0; i < b.N; i++ {
		Algorithm2(resultSet, resultMap, searchElements)
	}
}

func main() {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	resultSet, resultMap, searchElements := prepareDataUint32(10, 20, 30, 0.3)
	println("Set: %v", resultSet)
	println("Map: %v", resultMap)
	println("Search: %v", searchElements)

	Algorithm1(resultSet, resultMap, searchElements) // oder Algorithm2()
}

func myGenStringData(size, count int) (result []string) {
	src := rand.New(rand.NewSource(int64(size * count)))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r := make([]rune, size*count)
	for i := range r {
		r[i] = letters[src.Intn(len(letters))]
	}
	result = make([]string, count)
	for i := range result {
		result[i] = string(r[:size])
		r = r[size:]
	}
	return
}

func myGenUint32Data(count int) (result []uint32) {
	result = make([]uint32, count)
	for i := range result {
		result[i] = rand.Uint32()
	}
	return
}

func myGenerateInt64Data(n int) (data []int64) {
	data = make([]int64, n)
	for i := range data {
		data[i] = rand.Int63()
	}
	return
}
