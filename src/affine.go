package main

import (
	"fmt"
	"math"
)


func main() {
    var a, s, n int = 13, 16, 33
    var m, as, dKey int
    var counter int = 0

    fmt.Println("Enter the length of name and lastname: ")
    fmt.Scanln(&m)
    
    var x = make([]int, m) // Input slice  
    var E = make([]int, m) // Encrypted data
    var R = make([]float64, m) // Decrypted data
    
    fmt.Println("Enter the sequence number of each letter in the alphabet: ");	
    for i := 0; i < m; i++ {
        fmt.Printf("Letter #%d: ", i)
	fmt.Scanln(&x[i])
    }

    for i := 0; i < m; i++ {
	E[i] = (a * x[i] + s) % n
    }

    fmt.Printf("Encrypted text:\n");
    for i:=0; i < m; i++ {
	fmt.Printf("E[%d] = %d\n", i, E[i])
    }	

    for {
        if ((counter * n + 1) % a) == 0 {
	    as = (counter * n + 1) / a;
	    break    
	}
 	counter++
    }

    dKey = (-as * s) % n

    fmt.Printf("Decrypt key value: %d\n", dKey)

    for i:=0; i < m; i++ {
        R[i] = math.Abs(float64((as * E[i] + dKey) % n))
    }

    fmt.Printf("Decrypted text:\n");
    for i:=0; i < m; i++ {
	fmt.Printf("R[%d] = %d\n", i, int(R[i]))
    }

}
