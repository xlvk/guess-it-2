package main

import (
	"bufio"
	"fmt"
	// "log"
	"math"
	"os"
	"strconv"
)

func main() {
	var arr []int
	sum := 0.0
	xsum := 0.0
	XYNum := 0.0
	XXsum := 0.0
	YYsum := 0.0
	index := 0.0
	// file := ""
	fileScanner := bufio.NewScanner(os.Stdin)
	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			fmt.Printf("%T \n %v", ha, ha)
		}
		index++
		xsum = xsum + float64(len(arr))
		XXsum = XXsum + (float64(len(arr)) * float64(len(arr)))
		if ha < 100 || ha > 200 {
			arr = append(arr, (int(math.Round(sum / float64(len(arr))))))
			sum = sum + float64((int(math.Round(sum / float64(len(arr))))))
			XYNum = XYNum + (float64((int(math.Round(sum / float64(len(arr)))))) * float64(len(arr)-1))
			YYsum = YYsum + (float64((int(math.Round(sum / float64(len(arr)))))) * float64((int(math.Round(sum / float64(len(arr)))))))
		} else {
			arr = append(arr, ha)
			sum = sum + float64(ha)
			XYNum = XYNum + (float64(ha) * float64(len(arr)-1))
			YYsum = YYsum + (float64(ha) * float64(ha))
		}

		a := ((sum * XXsum) - (xsum * XYNum)) / (float64(len(arr))*XXsum - (xsum * xsum))
		b := ((float64(len(arr)) * XYNum) - (xsum * sum)) / (float64(len(arr))*XXsum - (xsum * xsum))

		// a := ((sum * XXsum) - (xsum * XYNum))*(float64(len(arr))* - (xsum * sum))
		// bb := ((float64(len(arr)) * XYNum) - (xsum * sum)) / (float64(len(arr))*XXsum - (xsum * xsum))

		// b := ((XYNum-((sum*xsum)/float64(len(arr))))/(XXsum-((xsum*xsum)/float64(len(arr)))))
		// a := (((sum*XXsum)-(xsum*XYNum))/((float64(len(arr))*XXsum)-(xsum*xsum)))
		calVal := b*index + a + index
		// y := 0.0
		// rrr := 0
		if len(arr) == 1 {
			calVal = float64(arr[0])
		}
		// 	y = float64(arr[0])
		// rrr = 10
		// } else {
		// 	y = b*(float64(len(arr)+1)) + a
		// 	rrr = arr[len(arr)-1] - arr[len(arr)-2]
		// 	if (arr[len(arr)-1] - arr[len(arr)-2]) > 0 {
		// 		rrr = arr[len(arr)-1] - arr[len(arr)-2]
		// 	} else {
		// 		rrr = arr[len(arr)-2] - arr[len(arr)-1]
		// 	}
		// }

		// y = y + (float64(len(arr) + 1))
		low := int(math.Round(calVal - 5))
		up := int(math.Round(calVal + 6))

		fmt.Printf("%d %d\n", low, up)
	}
}
