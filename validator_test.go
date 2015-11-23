package input

import (
    "fmt"
    "testing"
    "net/url"
)

func TestRequired(t *testing.T) {
    data := url.Values{
        "123": []string{"123"},
    }
    v := Validator{}
    v.Required(&data, "123")
    v.Required(&data, "asdb")
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}

func TestMaxLen(t *testing.T) {
    data := url.Values{
        "123": []string{"123123"},
    }
    v := Validator{}
    v.Required(&data, "123")
    v.MaxLen(data.Get("123"), 7)

    v.Required(&data, "123")
    v.MaxLen(data.Get("123"), 4)

    v.Required(&data, "abc")
    v.MaxLen(data.Get("abc"), 3)
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}

func TestInt(t *testing.T) {
    data := url.Values {
        "a": []string{"123123"},
    }

    v := Validator{}
    v.Required(&data, "a")
    dv := v.Int(&data, "a", 0)
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
    int1 := v.Int(&data, "tt", 5)
    int2 := v.Int(&data, "t", 10)
    fmt.Println(int1)
    fmt.Println(int2)
    if v.Err != nil {
        fmt.Println(v.Err)
    }
}
