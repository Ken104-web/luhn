package luhn2

import (
	"errors"
	"regexp"
	"strconv"

)

// card type constants

const (
    MASTERCARD = "MASTERCARD"
    VISA = "VISA"
    DISCOVER = "DISCOVER"
    UNSUPPORTED = "UNSUPPORTED"
)

// returns regex patterns for supported card networks

func SupportedCardNetworks() map[string]string{
    return map[string]string{
        MASTERCARD: "^5[1 2 3 4 5]\\d{14}$",
        VISA: "^4(?:\\d{12}|\\d{15})$",
        DISCOVER: "^6011\\d{12}$",
    }
}

// func return (true, card type) on success else return (false, "")

func IsCnValid(cn string) (bool, string){
    cardType, err := getCardType(cn)
    if err != nil{
        return false, ""
    }
    rawNum, err := getIntArrayFromStr(cn)
    if err != nil{
        return false, ""
    }
    return Valid(rawNum), cardType
    
}
// func returns a cn type on success and an err on faliure

func getCardType(cn string) (string, error) {
    networkName, same := UNSUPPORTED, false
    for n, p := range SupportedCardNetworks(){
        exp := regexp.MustCompile(p)
        if exp.MatchString(cn){
            networkName, same = n, true
            break
        }
    }
    if same{
        return networkName, nil
    } else{
        return "", errors.New(UNSUPPORTED)
    }
}



// check if valid or not based on luhn algorithm
func Valid(number int) bool {
    return  (number%10+checksum(number/10))%10 == 0
}
// checks the lat digit after computing 0-9 in checkSum
func Validateluhn(number int) int {
    checkNum := checksum(number)
    if checkNum == 0 {
        return 0
    }
    return 10 - checkNum   
}
func checksum(number int) int {
    var luhn2 int
    debugPrintf("working")
    for i:= 0; number > 0; i++ {
        rem := number % 10

        if i%2 == 0 { // meaning even
            rem = rem * 2
            if rem > 9 { // add individual digits after operation
               rem = rem%10 + rem/10
            }
        }
        luhn2 += rem
        number = number / 10
    }
    return luhn2 % 10
}

// checks & converts string to int
// Take the cn num as str
// remove the last digit
// convert remainig digits to int
// then return it
func getIntArrayFromStr(s string) (int, error){
    debugPrintf("working")
    res := s[:len(s)-1]
    // convertion
    item, err := strconv.Atoi(res)
    if err != nil {
        return 0, err
    }
    return item, nil 
}

// func getIntArrayFromStr(s string) int {
//     res := make(int, len(s))
//     for i, u := range s{
//         res[i] = int(u - '0')
//     }
//     return res
// }
