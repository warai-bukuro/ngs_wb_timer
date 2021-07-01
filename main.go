package main

import (
  "fmt"
  "os"
  "os/signal"
  //"sync"
  "time"

  "golang.org/x/sys/windows"
)

var (
  moduser32 = windows.NewLazyDLL("user32.dll")
  procGetAsyncKeyState = moduser32.NewProc("GetAsyncKeyState")
)

func main() {
  fmt.Println("Start")
  countdown()
}

func countdown(){
  t := time.NewTicker(1 * time.Second)
  defer t.Stop()

  reset := make(chan int, 1)
  reset <- 0
  go func (){
    for {
      // If you want to change the value of the key, please refer to the URL.
      // https://docs.microsoft.com/ja-jp/windows/win32/inputdev/virtual-key-codes?redirectedfrom=MSDN
      button01, _, _ := procGetAsyncKeyState.Call(uintptr(0x05))
      button02, _, _ := procGetAsyncKeyState.Call(uintptr(0x06))
      if button01 == 0 || button02 == 0 {
        reset <- 0
        continue
      }
      fmt.Printf("\r!!")
      reset <- 1
    }
  }()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  go func (){
    select {
      case <-c:
        fmt.Printf("\n")
        t.Stop()
        close(reset)
        close(c)
        os.Exit(130)
        return
    }
  }()

  for i := 15; true; i-- {
    select {
    case <-t.C:
      fmt.Printf("\r%2d",i)
    }
    if <- reset != 0 || i < 1 {
      i = 14
    }
  }
}

