package main

import "os/exec"

func main() {
  err := exec.Command("open", "https://www.google.co.jp").Start()
  if err != nil {
    panic(err)
  }
}

