// From http://www.sarathlakshman.com/2016/06/15/pitfall-of-golang-scheduler
// Upnderstanding the go scheduler and the go concurrency.

package main
import "fmt"
import "time"
import "runtime"

func main() {
    var x int
    threads := runtime.GOMAXPROCS(0)-1
    for i := 0; i < threads; i++ {
        go func() {
            for { x++ }
        }()
    }
    time.Sleep(time.Second)
    fmt.Println("x =", x)
}
