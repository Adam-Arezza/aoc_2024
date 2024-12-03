package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
    "math"
)

func checkCriteria(report []int) int{
    isSafe := 0
    increments := false
    isIncreasing := false
    isDecreasing := false
    for i := range report{
        if i == 0 {
            continue
        }

        if report[i] > report[i-1]{
            isIncreasing = true
        }

        if report[i] < report[i-1]{
            isDecreasing = true
        }

        if isIncreasing && isDecreasing{
            isDecreasing = false
            isIncreasing = false
            break
        }

        if math.Abs(float64(report[i]) - float64(report[i-1])) > 0 && math.Abs(float64(report[i]) - float64(report[i-1])) <= 3{
            increments = true
        }else{
            increments = false
            break
        }

    }

    if (isDecreasing || isIncreasing) && isIncreasing != isDecreasing && increments{
        isSafe = 1
    }
    return isSafe
}


func main(){
    var totalSafe int
    var totalSafeWithProblemDamper int
    currentDir, err := os.Getwd()
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    dataFile, err := os.Open(path.Join(currentDir,"/data.txt"))
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    fileScanner := bufio.NewScanner(dataFile)
    fileScanner.Split(bufio.ScanLines)
    var reports [][]int
    for fileScanner.Scan(){
        var reportNumbers []int
        var report = strings.Split(fileScanner.Text(), " ")
        for i := range report{
            num, err := strconv.Atoi(report[i])
            if err != nil{
                fmt.Printf("Error converting: %s to integer", report[i])
            }else{
                reportNumbers = append(reportNumbers, num)
            }
        }
        reports = append(reports, reportNumbers)
    }

    //checks each report to see if it's safe
    for r := range reports{
        if checkCriteria(reports[r]) == 1{
            totalSafe += 1
        }else{
            //apply problem damper
            fmt.Printf("Original report: %d\n", reports[r])
            for j := range len(reports[r]){
                var test []int
                if j == 0{
                    test = append([]int{}, reports[r][j+1:]...)
                }else{
                    start := append([]int{}, reports[r][:j]...)
                    end := append([]int{}, reports[r][j+1:]...)
                    test = append(start,end...)
                }

                fmt.Printf("Testing: %d\n", test)
                if checkCriteria(test) == 1{
                    totalSafeWithProblemDamper += 1
                    break
                }
            }
        }
    }

    fmt.Printf("The total number of safe reports is: %d\n", totalSafe)
    fmt.Printf("The total number of safe reports using the problem damper is: %d\n", totalSafeWithProblemDamper + totalSafe)

}
