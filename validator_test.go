package input

import (
    "fmt"
    "testing"
    "net/url"
    "encoding/json"
)

func TestRequired(t *testing.T) {
    data := url.Values{
        "123": []string{"123"},
    }
    v := Validator{}
    v.Required(&data, "123", "123:missing")
    v.Required(&data, "asdb", "asdb:missing")
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}

func TestMaxLen(t *testing.T) {
    data := url.Values{
        "123": []string{"123123"},
    }
    v := Validator{}
    v.Required(&data, "123", "123:missing")
    v.MaxLen(data.Get("123"), 7, "123:max_len_error")

    v.Required(&data, "123", "123:missing")
    v.MaxLen(data.Get("123"), 4, "123:max_len_error")

    v.Required(&data, "abc", "abs")
    v.MaxLen(data.Get("abc"), 3, "abc:max_len_error")
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}

func TestInt(t *testing.T) {
    data := url.Values {
        "a": []string{"123123"},
    }

    v := Validator{}
    v.Required(&data, "a", "a:missing")
    dv := v.Int(&data, "a", 0, "a:invalid")
    fmt.Println(dv)
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}

func TestOptionalInt(t *testing.T) {
    data := url.Values {
        "int1": []string{"123"},
    }

    v := Validator{}
    int1 := v.Int(&data, "tt", 5, "tt:invalid")
    int2 := v.Int(&data, "t", 10, "t:invalid")
    fmt.Println(int1)
    fmt.Println(int2)
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}

func TestJSONString(t *testing.T) {
    v := Validator{}
    v.JSONString("ddsf", "not_json_error")
    if v.Err == nil {
        t.Error("json check failed")
    }
}

func TestJSONStrin(t *testing.T) {
    data := map[string]interface{} {
        "b": "a",
        "a": 123,
    }
    jsonString, err := json.Marshal(data)
    if err != nil {
        t.Error(err)
    }
    v := Validator{}
    v.JSONString(string(jsonString), "json:invalid")
    if v.Err != nil {
        t.Error(v.Err)
    }
}
