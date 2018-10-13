package main

import "os"

//遅延実行はリソースの解放に便利

func main() {

    //ファイルのオープン
    file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)

    //オープンに失敗した時は終了
    if err != nil {
        os.Exit(1)
    }

    //ファイルのクローズを遅延指定
    defer file.Close()

    //ファイルへテキスト出力
    file.WriteString("あいうえお")

}