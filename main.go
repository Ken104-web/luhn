package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// create the func as a card number.
// receive the num as a string as they have always have spaces

func validate(cn string) bool {
    // removes white spaces
    var ss string

    for _, r := range cn{
        if !unicode.IsSpace(r) {
            ss += string(r)
        }
    }
     // setup 
    var sum int64 = 0
    parity := len(ss) % 2
    
    cardNumWithoutChecksum := ss[:len(ss)-1]
    // converts each string to int
    for i, v := range cardNumWithoutChecksum {
        item, err := strconv.Atoi(string(v))

        if err != nil {
            fmt.Println(err)
            return false        
        }
    // apply formula
    if int64(i)%2 != int64(parity) {
        sum += int64(item)
    } else if item > 4 {
        sum += int64(2*item - 9)
    } else {
        sum += int64(2* item)
    }
    }
    checkDigit, err := strconv.Atoi(ss[len(ss)-1:])

    if err != nil {
        fmt.Println(err)
        return false
    }
     SumMod := sum % 10

     if SumMod == int64(0) {
         return SumMod == int64(checkDigit)
     }
     return int64(10)-SumMod == int64(checkDigit)

}
func main(){
    isValid := validate("4000056655665556")
    fmt.Println(isValid)
}

