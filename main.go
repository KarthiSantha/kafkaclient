package main 

import (
    "fmt"
    "io/ioutil"
    "kafkaclient/Producer"
)
// "kafkaclient/reader"
func main(){

    fmt.Printf("Main class starting")
    //go reader.Reader()
    file, _ := ioutil.ReadFile("test.json")
    fmt.Printf(" data from file is \n %v \n",string(file))
    Producer.Producer()

}