package luhn2

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T){
    validNum := []int{
        68431579246,
        3719928465537204,
        8274503916824457,
        9364578220135794,
        5527801946381276,
        8146792058431168,
        43298567124039,
    }
    for _, i := range validNum{
        if !Valid(i){
        t.Errorf("%v you should have a valid", i)
    }
    if Validateluhn(i/10) != i%10{
        t.Errorf("%v's check number should be %v, but got %v instead", i, i%10, Validateluhn(i/10))
    }

}
}

func GetValidCn(t *testing.T){
    // check if card num is valid/not
    exp := map[string]bool{
        "378282246310005": true,
        "6011000990139424": true,
        "9937373": false,
        "4222222222222": true,
        "371449635398431": true,
    }
    actualResult := make(map[string]bool)
    for cn := range exp{
        r, _ := IsCnValid(cn)
        actualResult[cn] = r
}
if !reflect.DeepEqual(actualResult, exp) {
    t.Errorf("Not working!\nActual result should be %v\n", actualResult)
    }
}
