package bot

import (
//    "github.com/Heribio/termChat/internal/cli"
    "github.com/Goscord/goscord/goscord"
    "github.com/Goscord/goscord/goscord/discord"
    "github.com/Goscord/goscord/goscord/gateway"
    "github.com/Goscord/goscord/goscord/gateway/event"
    "io"
    "os"
    "fmt"
    "encoding/json"
    "net/http"
)

var messages []discordMessage

var client *gateway.Session

func Run() {
    go RunBot()
    go RunServer()

    select {}
}

func RunBot() {
    jsonFile, err := os.Open("../../SECRETS/secrets.json")
    if err != nil {
        fmt.Println(err)
        fmt.Println("Create a file ./SECRETS/secrets.json with the following structure: {\"WEBHOOK_URL\": YOUR_WEBHOOK_URL HERE\"}")
        os.Exit(1)
    }
    byteValue, err := io.ReadAll(jsonFile)

    var result Secrets

    json.Unmarshal(byteValue, &result)

    client := goscord.New(&gateway.Options{
            Token: result.BOT_TOKEN,
            Intents: gateway.IntentsAll,
    })


    client.On(event.EventMessageCreate, func(msg *discord.Message) {

        messages = append(messages, discordMessage{Content: msg.Content, Username: msg.Author.Username})
        fmt.Println(msg.Username + ": " + msg.Content)
    })
    
    client.Login()
}

func RunServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        jsonResponse, err := json.Marshal(messages)
        if err != nil {
            http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
    })
    http.ListenAndServe(":8080", nil)
}

/*
../../internal/bot/bot.go:48:1: syntax error: non-declaration statement outside function body
../../internal/bot/bot.go:48:67: method has multiple receivers
../../internal/bot/bot.go:50:2: syntax error: unexpected ) after top level declaration
*/
