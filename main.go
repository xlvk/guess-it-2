package main

import (
	"bufio"
	"fmt"
	"log"
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
	// readFile, err := os.Open("data.txt")
	// defer readFile.Close()

	// if err != nil {
	// 	log.Fatalf("failed to open file: %s", err)
	// }

	fileScanner := bufio.NewScanner(os.Stdin)
	// fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			fmt.Printf("%T \n %v", ha, ha)
		}
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
		xsum = xsum + float64(len(arr)-1)
		XXsum = XXsum + (float64(len(arr)-1) * float64(len(arr)-1))
	}
	b := math.Round(((XYNum-((sum*xsum)/float64(len(arr))))/(XXsum-((xsum*xsum)/float64(len(arr)))))*1000000) / 1000000
	a := math.Round((((sum*XXsum)-(xsum*XYNum))/((float64(len(arr))*XXsum)-(xsum*xsum)))*1000000) / 1000000
	// PCC := (math.Round((((float64(len(arr))*XYNum)-(sum*(xsum)))/math.Sqrt(((float64(len(arr))*XXsum)-(xsum*xsum))*((float64(len(arr))*YYsum)-(sum*sum))))*10000000000) / 10000000000)
	// fmt.Printf("Linear Regression Line: y = %.6f",b)
	// fmt.Printf("x + %.6f\n", a)
	// fmt.Printf("Pearson Correlation Coefficient: %.10f\n", PCC)
	y := b*(float64(len(arr)))+a
	low := int(y - 2 )
	up := int(y + 2)
	fmt.Printf("%d %d\n", low, up)
}
