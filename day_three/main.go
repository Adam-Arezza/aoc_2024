package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

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
    var total int
    for fileScanner.Scan(){
        str := fileScanner.Text()
        str = strings.TrimSpace(str)
        r,_ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)        
        //r,_ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
        matches := r.FindAllStringSubmatch(str, -1)
        //fmt.Println(len(matches))
        // at a "don't", get the index, at the next "do" get the index
        // remove the section of items between the indicies
        // continue until the slice only has the valid operations
        // find total as before
        var skipOperation bool = false
        for match := range matches{
            fmt.Println(matches[match])
            if matches[match][0] == `don't()`{
               skipOperation = true
            }

            if matches[match][0] == `do()`{
                skipOperation = false
                continue
            }

            if skipOperation{
                continue
            }else{
                var num1 int
                var num2 int
                newStr := strings.ReplaceAll(matches[match][0], "mul", "")
                newStr = strings.ReplaceAll(newStr, "(","")
                newStr = strings.ReplaceAll(newStr, ")","")
                numStrings := strings.Split(newStr, ",")
                num1, err = strconv.Atoi(numStrings[0]) 
                if err != nil{
                    fmt.Printf("Error converting first number: %s\n", numStrings[0])
                    fmt.Println(err.Error())
                }
                num2, err = strconv.Atoi(numStrings[1])
                if err != nil{
                    fmt.Printf("Error converting second number: %s\n", numStrings[1])
                    fmt.Println(err.Error())
                }
                total += num1 * num2
            }
        }
    }
    fmt.Printf("The total is: %d", total)
}
