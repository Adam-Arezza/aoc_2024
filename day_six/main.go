package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"slices"
)

type StartingPosition struct {
    x int
    y int
    direction rune
}

func main(){
    data := readFile("/test.txt")
    directions := []rune{'<','v','>','^'}
   // for n := range data{
   //     fmt.Println(string(data[n]))
   // }
    startingPosition := getStartPosition(data, directions)
    stepsTotal := walk(data, startingPosition)
    fmt.Printf("The total steps are: %d", len(stepsTotal))
}

func walk(matrix [][]rune, start StartingPosition)[][]int{
    var steps [][]int
    stop := '#'
    tile := '.'
    x := start.x
    y := start.y
    currentDir := start.direction
    steps = append(steps, []int{y,x})
    matrix[y][x] = '.'

    fmt.Printf("Starting walk at location: %d, %d\n", y,x)
    for y <= len(matrix) && x >= 0 && x <= len(matrix[0]) && y >= 0{
        nextX := x
        nextY := y
        if currentDir == '>'{
            nextX = x + 1
        }else if currentDir == '<'{
            nextX = x - 1
        }else if currentDir == '^'{
            nextY = y - 1
        }else{
            nextY = y + 1
        }

        if nextX > len(matrix[0])-1 || nextX < 0 || nextY > len(matrix)-1 || nextY < 0{
            break
        }

        switch matrix[nextY][nextX]{
            case tile:
                x = nextX
                y = nextY
                duplicate := false
                for s := range steps{
                    if equal(steps[s], []int{y,x}){
                        duplicate = true
                    }
                }

                if !duplicate{
                    steps = append(steps, []int{y,x})
                }

                //fmt.Printf("Moving to next space: %d, %d\n", y,x)
            case stop:
                currentDir = turn(currentDir)
            default:
        }
    }
    return steps
}

func equal(a,b []int)bool{
    if len(a) != len(b){
        return false
    }
    for i, v := range a{
        if v != b[i]{
            return false
        }
    }
    return true
}

func turn(direction rune)rune{
    var newDirection rune
    switch direction{
        case '^':
            newDirection = '>'
        case '>':
            newDirection = 'v'
        case 'v':
            newDirection = '<'
        case '<':
            newDirection = '^'
        default:
    }
    return newDirection
}

func getStartPosition(matrix [][]rune, dirs []rune)StartingPosition{
    var start StartingPosition
    for i := range len(matrix){
        for n := range matrix[i]{
            if slices.Contains(dirs, matrix[i][n]){
                start.y = i
                start.x = n
                start.direction = matrix[i][n]
            }
        }
    }
    return start
}

func readFile(filename string)[][]rune{
    var result [][]rune
    currentDir, err := os.Getwd()
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    dataFile, err := os.Open(path.Join(currentDir,filename))
    if err != nil{
        fmt.Printf("error: %s", err.Error())
    }
    fileScanner := bufio.NewScanner(dataFile)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan(){
        chars := []rune(fileScanner.Text())
        result = append(result, chars)
    }
    return result
}
