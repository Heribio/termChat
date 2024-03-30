package bot

import (
//    "github.com/Heribio/termChat/internal/discord"
    "github.com/Goscord/goscord/goscord"
    "github.com/Goscord/goscord/goscord/discord"
    "github.com/Goscord/goscord/goscord/gateway"
    "github.com/Goscord/goscord/goscord/gateway/event"
    "io"
    "os"
    "fmt"
    "encoding/json"
)

var client *gateway.Session

func Run() {
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
        //TODO send message to cli
    })
    
    client.Login()

    select {}
}
