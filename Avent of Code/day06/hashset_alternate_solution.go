//roughly 1.8x slower than that just loops through string window, but more concise
package main

func findMarker2(signal string) int {

	for i := 14;i < len(signal);i++{
		hashSet := map[byte]int{} //second data type doesn't really matter
		for j := 0;j < 14;j++{
			if _, ok := hashSet[signal[i - 14 + j]]; ok{
				break
			}
			hashSet[signal[i - 14 + j]] = j
		}
		if len(hashSet) == 14{
			return i
		}
	}
	panic("Could not find marker")
}
