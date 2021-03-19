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
    data_str := string(data[:])
    data_str = strings.Replace(data_str, "\n", "", -1)
    fmt.Println("INPUT DATA:", data_str)
}
