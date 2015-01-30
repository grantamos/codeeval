package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        values := strings.Split(scanner.Text(), " ")
        firstDivisor, _ := strconv.Atoi(values[0])
        secondDivisor, _ := strconv.Atoi(values[1])
        n, _ := strconv.Atoi(values[2])
        fizzbuzz(firstDivisor, secondDivisor, n)
    }
}

func fizzbuzz(firstDivisor int, secondDivisor int, n int) {
    for i := 1; i <= n; i++ {
        output := ""
        if i % firstDivisor == 0 {
            output += "F"
        }
        if i % secondDivisor == 0 {
            output += "B"
        }
        if output == "" {
            output += strconv.Itoa(i)
        }

        fmt.Print(output + " ")
    }

    fmt.Println()
}