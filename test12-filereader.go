package main
import (
    "fmt"
    "strings"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("[FILE]")
    if err != nil {
        fmt.Println("File input ERROR:", err)
        return
    }
    data_del := string(data[:])
    data_del = strings.Replace(data_del, "\n", "", -1)
    fmt.Println("INPUT DATA:", data_del)
}
