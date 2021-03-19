package main
import (
    "fmt"
    "strings"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("[FILE]")
    if err != nil {
        fmt.Println("File reader ERROR:", err)
        return
    }
    data_del := string(data[:])
    data_del = strings.Replace(data_del, "\n", "", -1)
    //data_del = strings.TrimSpace(data_del)
    fmt.Println("READ IN DATA:")
    fmt.Println(data_del)
}
