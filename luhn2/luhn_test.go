package luhn2


import "testing"

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
        t.Errorf("bro's check number should be %v, but got v% instead",i,i%10, Validateluhn(i/10))
    }

}
}
