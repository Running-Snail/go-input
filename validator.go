package input

import (
    "fmt"
    "errors"
    "net/url"
    "strconv"
    "encoding/json"
    "unicode/utf8"
    "gopkg.in/mgo.v2/bson"
)

type Validator struct {
    Err error
}

func (v *Validator) Required(data *url.Values, key string, errMsg string) string {
    dStrings, exists := (*data)[key]
    if v.Err != nil {
        return data.Get(key)
    }
    if !exists {
        v.Err = errors.New(errMsg)
        return ""
    }
    return dStrings[0]
}

func (v *Validator) Int(data *url.Values, key string, preset int, errMsg string) int {
    if v.Err != nil {
        return preset
    }
    dStrings, exists := (*data)[key]
    if !exists {
        return preset
    }
    d, err := strconv.Atoi(dStrings[0])
    if err != nil {
        v.Err = errors.New(errMsg)
        return preset
    }
    return d
}

func (v *Validator) Float(data *url.Values, key string, preset float64, errMsg string) float64 {
    if v.Err != nil {
        return preset
    }
    dStrings, exists := (*data)[key]
    if !exists {
        return preset
    }
    d, err := strconv.ParseFloat(dStrings[0], 64)
    if err != nil {
        v.Err = errors.New(errMsg)
        return 0.0
    }
    return d
}

func (v *Validator) MaxLen(data string, max int, errMsg string) string {
    if v.Err != nil {
        return data
    }
    if utf8.RuneCountInString(data) > max {
        v.Err = errors.New(errMsg)
    }
    return data
}

func (v *Validator) MinLen(data string, min int, errMsg string) string {
    if v.Err != nil {
        return data
    }
    if utf8.RuneCountInString(data) < min {
        v.Err = errors.New(errMsg)
    }
    return data
}

func (v *Validator) MaxInt(data int, max int, errMsg string) int {
    if v.Err != nil {
        return data
    }
    if data > max {
        v.Err = errors.New(errMsg)
    }
    return data
}

func (v *Validator) MinInt(data int, min int, errMsg string) int {
    if v.Err != nil {
        return data
    }
    if data < min {
        v.Err = errors.New(errMsg)
    }
    return data
}

func (v *Validator) ObjectId(data string, errMsg string) {
    if v.Err != nil {
        return
    }
    if !bson.IsObjectIdHex(data) {
        v.Err = errors.New(errMsg)
    }
}

func (v *Validator) JSONString(data string, errMsg string) {
    if v.Err != nil {
        return
    }
    var r interface{}
    if err := json.Unmarshal([]byte(data), r); err != nil {
        fmt.Println(err)
        v.Err = errors.New(errMsg)
    }
}

func (v *Validator) Reset() {
    v.Err = nil
}
