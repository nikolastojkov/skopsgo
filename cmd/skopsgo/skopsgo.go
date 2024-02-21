package main

import (
  "github.com/nikolastojkov/skopsgo/internal/config"
  "github.com/nikolastojkov/skopsgo/web/app"
)

func main() {
  config.LoadEnvConfig()
  app.LoadApp()
}
