package main

import  "fmt"
import  "time"

func main()  {


  ch1 := make(chan string) //channel 1
  ch2:= make(chan string)// channel 2
  go func(){          //goroutine 1
    for  {
      t1 := time.Now().Format("2006-01-02 15:04:05")
    ch1 <- "Time: "+t1 + " Process1 :   Odd"
    time.Sleep(time.Second*1) 
  }
  }()
  go func(){      //goroutine 2
    for  {
      t1 := time.Now().Format("2006-01-02 15:04:05")
    ch2 <- "Time: "+t1 + " Process2 :   Even"
    time.Sleep(time.Second*1)
  }
  }()
  for i := 0; i < 5; i++{ //Iterate 5 times
    fmt.Println(<- ch1)
      fmt.Println(<- ch2)
  }
}
