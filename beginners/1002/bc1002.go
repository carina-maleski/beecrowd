package bc1002

import (
	"fmt"
	"log"
)

var R, n, A float64
 
func main() {
    
    n := 3.14159
    
    _, err := fmt.Scanf("%v", &R)
    if err != nil {
        log.Fatal(err)
    }
    
    A = n * (R*R)
    
    fmt.Printf("A=%.4f\n", A)
}