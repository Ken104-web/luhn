package api

import (
	"encoding/json"
	"luhn/pattern"
    "net/http"
    "github.com/Ken104-web/luhn/tree/master/luhn2"

)

func CardNum(w http.ResponseWriter, r *http.Request) {
    var card pattern.CheckCard
    
    err := json.NewDecoder(r.Body).Decode(&card)
    // err handling
    if err != nil {
        http.Error(w, "Not the wanted JSON body", http.StatusBadRequest)
        return
    }
    // cn validation
    checkValid, cardType := luhn2.IsCnValid(card.Num)
    // to hold json resp like javascript
    resp := make(map[string]string)
    // set basic content type in our case(json)
    w.Header().Add("Content-Type", "application/json")
    // return appropiate message based on the input
    if checkValid != nil{
        w.WriteHeader(http.StatusBadRequest)
        resp["Message"] = "not valid card"
    } else {
        resp["Message"] = "valid card"
        resp["cardType"] = cardType
        w.WriteHeader(http.StatusOK)
    }
    // convert to json and Request
    jresp, _:= json.Marshal(resp)
    w.Write(jresp)


}
