package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
    currentDir, err := os.Getwd()
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    dataFile, err := os.Open(path.Join(currentDir,"/test.txt"))
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    fileScanner := bufio.NewScanner(dataFile)
    fileScanner.Split(bufio.ScanLines)
    var charMatrix [][]string
    for fileScanner.Scan(){
        row := strings.Split(fileScanner.Text(), "")
        row = append([]string{"Z"}, row...)
        row = append(row, "Z")
        //fmt.Println(row)
        charMatrix = append(charMatrix, row)
        //fmt.Println(string(fileScanner.Text()[0]))
    }

    //add buffer to matrix
    padRow := make([]string, len(charMatrix[0]))
    for i := range padRow{
        padRow[i] = string("Z")
    }

    charMatrix = append([][]string{padRow},charMatrix...)
    charMatrix = append(charMatrix, padRow)
    for row := range charMatrix{
        fmt.Println(charMatrix[row])
    }

    //fmt.Print("Final matrix: \n")
    xmasCount := 0
    for i := range len(charMatrix){
        for j := range len(charMatrix[i]){
            if charMatrix[i][j] == string("X"){
                if checkAdjacentCells(i,j,charMatrix, string("M")){
                    xmasCount += 1
                }
            }else{
                continue
            }
        }
    }
    fmt.Printf("The total xmas count is: %d", xmasCount)
}

func checkAdjacentCells(i,j int, mat [][]string, target string) bool{
    checkList := [][]int{
                         {-1,-1}, 
                         {-1,0},
                         {-1,1},
                         {0,-1},
                         {0,1},
                         {1,-1},
                         {1,0},
                         {1,1},
                        }
    fmt.Printf("Searching around X at position: %d, %d\n", i, j)
    for coords := range checkList{
        fmt.Printf("looking around X, currently checking: %d\n", checkList[coords])
        checkRow := i + checkList[coords][0]
        checkCol := j + checkList[coords][1]
        fmt.Printf("found: %s\n", mat[checkRow][checkCol])
        if mat[checkRow][checkCol] == target{
            fmt.Println("Found M, looking for A now...")
            target = "A"
            if checkDirection(mat, checkRow, checkCol, target, checkList[coords]){
                fmt.Println("Found XMAS")
                return true 
            }else{
                target = "M"
                continue
            }
        }
    }
    return false
}

func checkDirection(mat [][]string, i int, j int, target string, direction []int) bool{
    directionRow := i + direction[0]
    directionCol := j + direction[1]
    if mat[directionRow][directionCol] == target{
        if target == "S"{
            fmt.Println("Found S")
            return true
        }

        if target == "A"{
            fmt.Println("Found A, looking for S now...")
            target = "S"
            return checkDirection(mat, directionRow, directionCol, target, direction)
        }
    }else{
        fmt.Printf("Couldn't find target: %s\n", target)
    }
    return false
}
