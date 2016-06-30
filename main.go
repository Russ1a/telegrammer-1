package main

import (
  "flag"
  "fmt"
  "net/http"
  "net/url"
  "os"
)

func main() {
  args := os.Args[1:]

  var token string
  var userid string
  var message string

  cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
  cmdFlags.StringVar(&token, "t", "", "Telegram API Token")
  cmdFlags.StringVar(&userid, "userid", "general", "Userid")
  cmdFlags.StringVar(&message, "m", "", "Message")

  if err := cmdFlags.Parse(args); err != nil {
    fmt.Printf("All is wrong")
    os.Exit(1)
  }
  userid, _ = url.QueryUnescape(userid)
  message, _ = url.QueryUnescape(message)
  resp, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token), url.Values{"text": {message}, "chat_id": {userid}})

  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", resp.Body)
}
