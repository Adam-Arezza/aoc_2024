package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

//get the smallest value in column 1 and column 2, determine the distance between them
//sum all the distances

func findMin(distances []int) (int,int){
    minDist := 1000000
    idx := 0
    for i := range len(distances){
        if distances[i] < int(minDist){
            minDist = distances[i]
            idx = i
        }else{
            continue
        }
    }
    return idx,minDist
}

func getTotalDicstance(left []int, right []int) int{
    var total int
    for range left{
        leftIdx, minLeft := findMin(left)
        rightIdx, minRight := findMin(right)
        left = append(left[:leftIdx], left[leftIdx+1:]...)
        right = append(right[:rightIdx], right[rightIdx+1:]...)
        dist := minLeft - minRight
        total += int(math.Abs(float64(dist)))
    }
    return total
}

func getSimilarity(left []int, right []int) int{
    var similarity int
    for i := range len(left){
        countInRight := 0
        for j := range len(right){
            if left[i] == right[j]{
                countInRight ++
            }
        }
        similarity += left[i] * countInRight
    }
    return similarity
}

func main(){
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
    var leftSide []int
    var rightSide []int

    for fileScanner.Scan(){
        line := strings.Split(fileScanner.Text(), "   ")
        left,err := strconv.Atoi(strings.TrimSpace(line[0]))
        if err != nil{
            fmt.Printf("error converting string to int: %s", line[0])
            fmt.Printf(err.Error())
            return
        }
        right,err := strconv.Atoi(strings.TrimSpace(line[1]))
        if err != nil{
            fmt.Printf("error converting string to int: %s", line[1])
            fmt.Printf(err.Error())
            return
        }
        leftSide = append(leftSide, left)
        rightSide = append(rightSide, right)
    }
    
    fmt.Printf("The similarity score: %d\n", getSimilarity(leftSide, rightSide))
    fmt.Print("The total distance: ")
    fmt.Println(getTotalDicstance(leftSide, rightSide))
    
}
