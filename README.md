構造体名取得ライブラリ
===
ソースコードを解析し、構造体名飲みを取得するライブラリです。

```go
package main

import (
    "github.com/ochipin/classes"
)

func main() {
    // Newでディレクトリ名を渡す
    // Classlist() 、ディレクトリ内にある*.goソースコードを解析し、構造名のみを抜き出す
    classlist, err := classes.New("src/structs").Classlist() // ([]string, error) の値が返却される
    if err != nil {
        panic(err)
    }

    // []string 型に格納された構造体一覧を表示
    for _, name := range classlist {
        fmt.Println(name) // App, Base, List, ...
    }
}
```