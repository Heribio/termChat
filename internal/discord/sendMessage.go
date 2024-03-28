package discord

import (
    "fmt"
    "os"
    "encoding/json"
    "io"
    "github.com/Heribio/termChat/pkg/discordWebhook"
)

type Secrets struct {
    WEBHOOK_URL string `json:"WEBHOOK_URL"`
}


func SendMessage(message discordWebhook.Message) {
    jsonFile, err := os.Open("../../SECRETS/secrets.json")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    byteValue, err := io.ReadAll(jsonFile)

    var result Secrets

    json.Unmarshal(byteValue, &result)
    discordWebhook.SendMessage(message, result.WEBHOOK_URL)
}
