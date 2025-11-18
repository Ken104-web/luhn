package main

import (
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
    var sum int64 = 0
    parity := len(ss) % 2

    for i, v := range ss {
        item, err := strconv.Atoi(string(v))

        if err != nil {
         return false        
     }
    if int64(i)%2 != int64(parity) {
        sum += int64(item)
    } else if item > 4 {
        sum += int64(2*item - 9)
    } else {
        sum += int64(2* item)
    }
    }
    return sum%10 == 0
}



