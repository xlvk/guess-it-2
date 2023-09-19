package main

import (
	"bufio"
	"fmt"
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
	fileScanner := bufio.NewScanner(os.Stdin)
	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			fmt.Printf("%T \n %v", ha, ha)
		}
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
		// r := (limit*xySum - xSum*ySum) / math.Sqrt((limit*xSqSum-(xSum*xSum))*(limit*ySqSum-(ySum*ySum)))

		// b := math.Round(((XYNum-((sum*xsum)/float64(len(arr))))/(XXsum-((xsum*xsum)/float64(len(arr)))))*1000000) / 1000000
		// a := math.Round((((sum*XXsum)-(xsum*XYNum))/((float64(len(arr))*XXsum)-(xsum*xsum)))*1000000) / 1000000
		y := 0.0
		rrr := 0.0
		if len(arr) == 1 {
			y = float64(arr[0])
			rrr = 10.0
			// y = float64(len(arr)) + float64(149)
		} else {
			y = b*(float64(len(arr)+1)) + a
			rrr = float64(arr[len(arr)-1]) - float64(arr[len(arr)-2])
			if (float64(arr[len(arr)-1]) - float64(arr[len(arr)-2])) < 0 {
				rrr = float64(arr[len(arr)-1]) - float64(arr[len(arr)-2])
			} else {
				rrr = float64(arr[len(arr)-2]) - float64(arr[len(arr)-1]) 
			}
		}
		y = y + (float64(len(arr)+1))
		low := int(math.Round(y - rrr))
		up := int(math.Round(y + rrr))
		fmt.Printf("%d %d\n", low, up)
	}
}


// for i := 0; i < len(arr); i++ {
		// 	xsum = xsum + float64(i)
		// 	XXsum = XXsum + (float64(i) * float64(i))
		// 	if arr[i] < 100 || arr[i] > 13000 {
		// 		// arr = append(arr, (int(math.Round(sum / float64(len(arr))))))
		// 		sum = sum + float64((int(math.Round(sum / float64(i)))))
		// 		XYNum = XYNum + (float64((int(math.Round(sum / float64(i+1))))) * float64(i))
		// 		YYsum = YYsum + (float64((int(math.Round(sum / float64(i+1))))) * float64((int(math.Round(sum / float64(i+1))))))
		// 	} else {
		// 		// arr = append(arr, arr[i])
		// 		sum = sum + float64(arr[i])
		// 		XYNum = XYNum + (float64(arr[i]) * float64(i))
		// 		YYsum = YYsum + (float64(arr[i]) * float64(arr[i]))
		// 	}
		// }