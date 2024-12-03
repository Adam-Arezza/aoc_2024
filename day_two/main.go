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

    //fmt.Printf("Current report: \n")
    //fmt.Println(report)
    for i := range report{
        if i == 0 {
            continue
        }

        if report[i] > report[i-1]{
     //       fmt.Printf("%d is greater than %d, setting increasing to true\n", report[i], report[i-1])
            isIncreasing = true
        }

        if report[i] < report[i-1]{
      //      fmt.Printf("%d is less than %d, setting decreasing to true\n", report[i], report[i-1])
            isDecreasing = true
        }

        if isIncreasing && isDecreasing{
            isDecreasing = false
            isIncreasing = false
            break
        }

        //fmt.Println("checking if increment/decrement is within range 1 to 3")
        if math.Abs(float64(report[i]) - float64(report[i-1])) > 0 && math.Abs(float64(report[i]) - float64(report[i-1])) <= 3{
            increments = true
        }else{
            increments = false
            break
        }

       // fmt.Printf("Increment within range: %v\n", increments)
    }

    if (isDecreasing || isIncreasing) && isIncreasing != isDecreasing && increments{
        isSafe = 1
    }

    //fmt.Printf("Is the report safe: %d\n", isSafe)
    return isSafe
}

func main(){
    var totalSafe int
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

    for r := range reports{
        totalSafe += checkCriteria(reports[r])
    }

    fmt.Printf("The total number of safe reports is: %d", totalSafe)
}
