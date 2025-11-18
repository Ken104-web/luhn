package luhn2




func Validateluhn(number int) int {
    checkNum := checksum(number)
    if checkNum == 0 {
        return 0
    }
    return 10 - checkNum
    
}

// check if valid or not based on luhn algorithm
func Valid(number int) bool {
    return  (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
    var luhn2 int

    for _, i:= 0, number > 0; i++ {
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
