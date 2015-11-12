# combinations

## ForEachCombination
<pre>func ForEachCombination(maxValue, numberOfValuesInEachSet int, processCombination func([]int))</pre>

## ForEachCombinationWithRepetitions
<pre>func ForEachCombinationWithRepetitions(maxValue, numberOfValuesInEachSet int, processCombination func([]int))</pre>

For each combination of [numberOfValuesInEachSet] elements, call the function [processCombination] passing a list of [maxValue] integers in 0-[maxValue] with or without repetitions
