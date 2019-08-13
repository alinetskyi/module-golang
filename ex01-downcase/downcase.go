package main
import "fmt"

func lowercase(s string) string{
    out := []rune(s)
    i:=0
    for ; i < len(out); i++ {
        if out[i] <= 97 {
  	    out[i] = out[i] + 32
	} 
    }
    return string(out)
}

func main() {
    fmt.Println(lowercase("HelloWorld"))
}
