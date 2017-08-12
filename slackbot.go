package main

import (
  "fmt"
  "log"
  "math/rand"
  "time"
  "os"
  "strings"
  "github.com/nlopes/slack"
)

var API_TOKEN string = "SET-YOUR_API_TOKEN"

func main() {
  api := slack.New(API_TOKEN)
  logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
  slack.SetLogger(logger)
  api.SetDebug(true)

  rtm := api.NewRTM()
  go rtm.ManageConnection()

  for msg := range rtm.IncomingEvents {
    fmt.Print("[*] Event Received: ")
    switch ev := msg.Data.(type) {
    case *slack.HelloEvent:
      // ignore Hello

    case *slack.ConnectedEvent:
      //fmt.Println("Infos: ", ev.Info)
      fmt.Println("Connection counter: ", ev.ConnectionCount)

    case *slack.MessageEvent:
      fmt.Printf("Message %v\n", ev)
      info := rtm.GetInfo()
      prefix := fmt.Sprintf("<@%s>", info.User.ID)

      if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
        respond(rtm, ev, prefix)
        //rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy?", ev.Channel))
      }

    case *slack.PresenceChangeEvent:
      fmt.Printf("Presence Change: %v\n", ev)

    case *slack.LatencyReport:
      fmt.Printf("Current Latency: %v\n", ev)

    case *slack.RTMError:
      fmt.Printf("Error: %s\n", ev)

    case *slack.InvalidAuthEvent:
      fmt.Printf("Invalid Credentials")
      return

    default:
      // ignore other events...
      // fmt.Printf("[-] Unexpected: %v\n", msg.Data)
    }
  }
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
  text := msg.Text
  text = strings.TrimPrefix(text, prefix)
  text = strings.TrimSpace(text)
  text = strings.ToLower(text)

  if text == "" {
    greetings := []string{"what's up buddy?", "What is your purpose?"}
    rand.Seed(time.Now().UnixNano())
    rtm.SendMessage(rtm.NewOutgoingMessage(greetings[rand.Intn(len(greetings))], msg.Channel))
  }

  if strings.Contains(strings.ToLower(text), "fortune") {
    var response string = fortune()
    rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
  }
}

func fortune() (string){
  var result string
  rand.Seed(time.Now().UnixNano())

  switch rand.Intn(10) {
  case 0:
    result = "Very good luck!!!"

  case 1:
    result = "Bad luck ..."

  case 2,3:
    result = "Good luck!!"

  case 4,5,6:
    result = "Fair luck!"

  default:
    // case 7,8,9
    result = "A little luck."
  }

  return result
}
