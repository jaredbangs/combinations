package combinations

// For each combination of [numberOfValuesInEachSet] elements
// call the function [processCombination] passing a list of m integers in 0-[maxValue]
// with repetitions
func ForEachCombinationWithRepetitions(maxValue, numberOfValuesInEachSet int, processCombination func([]int)) {

	indices := make([]int, numberOfValuesInEachSet)

	processCombination(indices)

	for {
		var i int
		for i = numberOfValuesInEachSet - 1; i >= 0; i-- {
			if indices[i] != maxValue-1 {
				break
			}
		}
		if i < 0 {
			break
		}

		indices_i := indices[i]
		for k := i; k < numberOfValuesInEachSet; k++ {
			indices[k] = indices_i + 1
		}
		processCombination(indices)
	}
}

// For each combination of [numberOfValuesInEachSet] elements
// call the function [processCombination] passing a list of m integers in 0-[maxValue]
// without repetitions
func ForEachCombination(maxValue, numberOfValuesInEachSet int, processCombination func([]int)) {

	s := make([]int, numberOfValuesInEachSet)
	last := numberOfValuesInEachSet - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < maxValue; j++ {
			s[i] = j
			if i == last {
				processCombination(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
