package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sort"
)

func main() {
	var arr []float64
	ysum := 0.0
	xsum := 0.0
	XYNum := 0.0
	XXsum := 0.0
	YYsum := 0.0
	index := 0.0
	avr := 0.0
	sd := 0.0
	fileScanner := bufio.NewScanner(os.Stdin)
	for fileScanner.Scan() {
		num, e := strconv.ParseFloat(fileScanner.Text(), 64)
		if e != nil {
			fmt.Printf("%T \n %v", num, num)
		}
		xsum = xsum + index
		XXsum = XXsum + (index * index)
		if num < 100 || num > 200 {
			arr = append(arr, ysum/(index+1))
			ysum = ysum + ysum/(index+1)
			XYNum = XYNum + (ysum/(index+1))*index
			YYsum = YYsum + (ysum/(index+1))*(ysum/(index+1))
		} else {
			arr = append(arr, num)
			ysum = ysum + num
			XYNum = XYNum + (num * index)
			YYsum = YYsum + (num * num)
		}
		index++
		// a := ((ysum * XXsum) - (xsum * XYNum)) / (index*XXsum - (xsum * xsum))
		// b := ((index * XYNum) - (xsum * ysum)) / (index*XXsum - (xsum * xsum))
		b := ((XYNum - ((ysum * xsum) / index)) / (XXsum - ((xsum * xsum) / index)))
		a := (((ysum * XXsum) - (xsum * XYNum)) / ((index * XXsum) - (xsum * xsum)))
		// PCC := (((index*XYNum)-(ysum*(xsum)))/math.Sqrt(((index*XXsum)-(xsum*xsum))*((index*YYsum)-(ysum*ysum))))
		rrr := 0.0
		
		calVal1 := b*index + a + index
		if index <= 3 {
			calVal1 = arr[int(index)-1]
		} else {
			// if (arr[index-1] - arr[index-2]) > 0 {
			rrr = arr[int(index)-1] - arr[int(index)-2]
		// } else {
		// 	rrr = arr[index-2] - arr[index-1]
		// }
		}
		calVal := num + rrr
		// used := calVal-PCC
		// used = calVal1 - used
		// used = used + PCC
		// oooo := calVal1 + used
		// if index <= 3 {
		// 	oooo = calVal1
		// }
		sort.Float64s(arr)
		avr = math.Round(ysum / float64(len(arr)))
		for j := 0; j < len(arr); j++ {
			sd += math.Pow((float64(arr[j]) - avr), 2)
		}
		sd = math.Round(math.Sqrt(sd / float64(len(arr))))
		y := avr
		// low := int(y - sd )
		// up := int(y + sd)
		used := (calVal+calVal1)/2
		low := int(math.Round(used-sd))
		up := int(math.Round(used+sd))

		fmt.Printf("%d %d\n", low, up)
	}
}
