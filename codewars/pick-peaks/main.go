package main

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	last := array[0]
	delta := 0
	for i := 1; i < len(array); i++ {
		v := array[i]
		dif := v - last
		delta = dif
	}
	return PosPeaks{Pos: []int{delta}, Peaks: []int{delta}}
}
