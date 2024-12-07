package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
    "regexp"
)

func main() {
    matrix := readFile("/data.txt")
    for row := range matrix{
        fmt.Println(string(matrix[row]))
    }
//    PART 1
//    xmasCount := 0
//    checkList := [][]int{{-1,-1}, {-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1},}
//    target := []byte{'X','M','A','S'}
//
//    for i := range len(matrix){
//        for j := range len(matrix[i]){
//            if matrix[i][j] == target[0]{
//               //check directions
//               for check := range checkList{
//                   direction := checkList[check]
//                   if checkDirection(direction, matrix, i, j, target, 1){
//                       xmasCount += 1
//                   }else{
//                       continue
//                   }
//               }
//            }        
//        }
//    }
//    fmt.Printf("The total xmas count is: %d", xmasCount)


//  PART 2
    masCount := 0
    r, _ := regexp.Compile("MAS|SAM")

    for i := range len(matrix){
        for j := range len(matrix[i]){
            if matrix[i][j] == 'A'{
                test := []byte{matrix[i + -1][j + -1], matrix[i][j], matrix[i + 1][j + 1]}
                if len(r.FindAllSubmatch(test, -1)) > 0{
                    test = []byte{matrix[i + -1][j + 1], matrix[i][j], matrix[i + 1][j + -1]}
                    if len(r.FindAllSubmatch(test, -1)) > 0{
                        masCount += 1
                    }
                }
            }
        }
    }
    fmt.Printf("The total X-mas count is: %d", masCount)

}

func readFile(file string)[][]byte{
    var result [][]byte
    currentDir, err := os.Getwd()
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    dataFile, err := os.Open(path.Join(currentDir,file))
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    fileScanner := bufio.NewScanner(dataFile)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan(){
        row := append([]byte{'Z'}, []byte(fileScanner.Text())...)
        row = append(row, 'Z')
        //fmt.Println(row)
        result = append(result, row)
        //fmt.Println(string(fileScanner.Text()[0]))
    }
    padRow := make([]byte, len(result[0]))
    for i := range padRow{
        padRow[i] = 'Z'
    }
    result = append([][]byte{padRow},result...)
    result = append(result, padRow)
    return result
}

func checkDirection(direction []int, mat [][]byte, i int, j int, target []byte, targetIdx int)bool{
    directionRow := i + direction[0]
    directionCol := j + direction[1]
    if mat[directionRow][directionCol] == target[targetIdx] && targetIdx < len(target)-1{
        targetIdx += 1
        return checkDirection(direction, mat, directionRow, directionCol, target, targetIdx)
    }

    if mat[directionRow][directionCol] == target[targetIdx] && targetIdx == len(target) -1{
        return true
    }

    return false
}
