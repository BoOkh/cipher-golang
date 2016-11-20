package main

import (
    "fmt"
    "math"
)

const A, S, N int = 13, 16, 33

func main() {
    ver := Vertex{}

    ver.inputLength()

    ver.x = make([]int, ver.m)
    ver.E = make([]int, ver.m)
    ver.R = make([]float64, ver.m)

    ver.inputData()
	
    ver.EncryptData()
    ver.getEncryptedData()

    ver.decryptKey()
	
    ver.decryptData()
    ver.getDecryptedData() 
}

type Vertex struct {
    m int
    x []int // Input slice  
    E []int // Encrypted data
    R []float64 // Decrypted data
    as, dKey int
}

func (ver *Vertex) inputLength() {
    fmt.Println("Enter the length of name and lastname: ")
    fmt.Scanln(&ver.m)
}

func (ver Vertex) inputData() {
    fmt.Println("Enter the sequence number of each letter in the alphabet: ");	
    for i := 0; i < ver.m; i++ {
        fmt.Printf("Letter #%d: ", i)
        fmt.Scanln(&ver.x[i])
    }
}

func (ver Vertex) EncryptData() {
    for i := 0; i < ver.m; i++ {
        ver.E[i] = (A * ver.x[i] + S) % N
    }
}

func (ver Vertex) getEncryptedData() {
    fmt.Printf("Encrypted text:\n");
    for i := 0; i < ver.m; i++ {
        fmt.Printf("E[%d] = %d\n", i, ver.E[i])
    }
}

func (ver *Vertex) decryptKey() {
    var counter int = 0

    for {
        if ((counter * N + 1) % A) == 0 {
            ver.as = (counter * N + 1) / A;
            break    
        }
        counter++
    }

    ver.dKey = (-ver.as * S) % N

    fmt.Printf("Decrypt key value: %d\n", ver.dKey)
}


func (ver Vertex) decryptData() {
    for i := 0; i < ver.m; i++ {
        ver.R[i] = math.Abs(float64((ver.as * ver.E[i] + ver.dKey) % N))
    }
}

func (ver Vertex) getDecryptedData() {
    fmt.Printf("Decrypted text:\n");
    for i := 0; i < ver.m; i++ {
        fmt.Printf("R[%d] = %d\n", i, int(ver.R[i]))
    }
}
