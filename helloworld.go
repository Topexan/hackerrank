package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    fmt.Println("Hello, World.")
    fmt.Println(input)
}