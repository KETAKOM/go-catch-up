//参考サイト
//http://cuto.unirita.co.jp/gostudy/post/standard-library-json/

package main

import (
    "encoding/json"
    "fmt"
)

type Country struct {
    Name string              `json:"name"`
    Prefectures []Prefecture `json:"prefectures"`
}

type Prefecture struct {
    Name string    `json:"name"`
    Capital string `json:"capital"`
    Population int `json:"population"`
}

func main() {
    jsonStr := `
{
    "name": "日本",
    "prefectures": [
        {
            "name": "東京都",
            "capital": "東京",
            "population": 13482040
        },
        {
             "name": "埼玉県",
             "capital": "さいたま市",
             "population": 7249287
        },
        {
             "name": "神奈川県",
             "capital": "横浜市",
             "population": 9116252
        }
    ]
}
`
    jsonBytes := ([]byte)(jsonStr)
    data := new(Country)

    err := json.Unmarshal(jsonBytes, data);

    if err != nil {
        fmt.Println("JSON Unmarshal error: ", err)
        return
    }

    fmt.Println(data)
}