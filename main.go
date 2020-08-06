package main 

import (
    "fmt"
    "kafkaclient/reader"
    "kafkaclient/Producer"
)

func main(){

    fmt.Printf("Main class starting")
    go reader.Reader()
	Producer.Producer()

}