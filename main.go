package main

import (
	"gocheat/gui"
	"log"
	"os"
	"path"
)

func main() {
  gui := gui.New()

  gui.Run()
}

func init() {
  home, err := os.UserHomeDir()
  confDir := path.Join(home, ".gocheat")
  if err != nil {
    log.Fatalln(err)
  }
  if _, err := os.Stat(confDir); err != nil {
    err := os.Mkdir(confDir, 0777)
    if err != nil {
      log.Fatalln(err)
    }
  }
}

