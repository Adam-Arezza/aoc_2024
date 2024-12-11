package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
    "sort"
    "reflect"
)

func main(){
    rules, updates := readFile("/data.txt")
//PART 1
    invalidUpdates := checkValidUpdate(rules, updates)
    fmt.Println("The following updates are valid:")
    var result int
    for n := range updates{
        if !slices.Contains(invalidUpdates, n){
            //fmt.Println(updates[n])
            length := len(updates[n])
            middle := length / 2
            result += updates[n][middle]
        }
    }

    fmt.Printf("The total is: %d\n", result)
    fmt.Printf("The total invalid is: %d\n", len(invalidUpdates))
    fmt.Println(invalidUpdates)

//PART 2
    //var correctedUpdateTotal int
    ruleMap := make(map[int][]int)
    for r := range rules{
        key := rules[r][0]
        _, ok := ruleMap[key]
        if ok{
            ruleMap[key] = append(ruleMap[key], rules[r][1])
        }else{
            ruleMap[key] = []int{rules[r][1]}
        }
    }

    first,second := alternativeAns(ruleMap,updates)
    fmt.Println(first,second)
}


///This solution is not mine...///
//////////////////found solution online///////////////////////////
func customLess(ruleMap map[int][]int, update []int, i, j int) bool {
	if _, ok := ruleMap[update[i]]; ok {
		for _, char := range ruleMap[update[i]] {
			if char == update[j] {
				return true
			}
		}
	}
	return false
}

func alternativeAns(rulesMap map[int][]int, updates [][]int) (int, int) {
	sum, fixedSum := 0, 0
	for _, update := range updates {

		sortedUpdate := make([]int, len(update))
		copy(sortedUpdate, update)

		sort.Slice(sortedUpdate, func(i, j int) bool {
			return customLess(rulesMap, sortedUpdate, i, j)
		})

		if reflect.DeepEqual(update, sortedUpdate) {
			sum += update[len(update)/2]
		} else {
			fixedSum += sortedUpdate[len(update)/2]
		}
	}
	return sum, fixedSum
}
///////////////////found solution online//////////////////////////


func checkValidUpdate(rules [][]int, updates [][]int )[]int{
    var invalidUpdates []int
    for i := range rules{
            for j := range updates{
                if slices.Contains(updates[j], rules[i][0]){
                    updateIdx := slices.Index(updates[j],rules[i][0])
                    if slices.Contains(updates[j], rules[i][1]){
                        updateIdxNext := slices.Index(updates[j], rules[i][1])
                        if updateIdx > updateIdxNext{
                            if !slices.Contains(invalidUpdates,j){
                                invalidUpdates = append(invalidUpdates,j)
                            }
                        }
                    }

                }else{
                    continue
                }
        }
    }
    //fmt.Printf("Invalid updates: %d\n", invalidUpdates)
    return invalidUpdates
}

func readFile(file string)([][]int,[][]int){
    var rules [][]int
    var updates [][]int
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
        if len(fileScanner.Text()) == 5{
            var numOne int
            var numTwo int
            line := fileScanner.Text()
            numStrs := strings.Split(line, "|")
            numOne, err := strconv.Atoi(numStrs[0])
            if err != nil{
                fmt.Printf("Error in conversion: %s", err.Error())
            }
            numTwo, err = strconv.Atoi(numStrs[1])           
            if err != nil{
                fmt.Printf("Error in conversion: %s", err.Error())
            }
            newRule := []int{numOne,numTwo}
            rules = append(rules, newRule)
        }else if len(fileScanner.Text()) == 0{
            continue
        }else{
            var newUpdate []int
            line := fileScanner.Text()
            newLine := strings.Split(line, ",")
            for i := range newLine{
                num, err := strconv.Atoi(newLine[i])
                if err != nil{
                    fmt.Printf("Error in conversion: %s", err.Error())
                }
                newUpdate = append(newUpdate, num)
            }
            updates = append(updates, newUpdate)
        }
    }
    return rules, updates
}
