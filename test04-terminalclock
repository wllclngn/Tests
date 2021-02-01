package main

import "fmt"
import "os"
import "os/exec"
import "runtime"
import "time"

// Map for storing clear funcs
var clear map[string]func()

func init() {
    clear = make(map[string]func()) 
    clear["linux"] = func() { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func callClear() {
    // runtime.GOOS -> linux, windows, darwin etc.
    value, ok := clear[runtime.GOOS]
    if ok {
        value()
    } else {
        panic("Done pooped the bed...")
    }
}

func main() {
    for true {
        time2 := time.Now()
        callClear()
        fmt.Println("")
        fmt.Println(time2.Format("2006-01-02 15:04:05.000000000"))
        // fmt.Println(time2)
        // time.Sleep(time.Second)
        time.Sleep(10 * time.Millisecond)
        // time.Sleep(time.Millisecond)
        // time.Sleep(time.Nanosecond)
    }
}
