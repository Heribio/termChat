package main

import (
   "github.com/Heribio/termChat/internal/cli" 
    "github.com/Heribio/termChat/internal/bot"
)

func main() {
    go bot.Run()
    cli.Run()
}
