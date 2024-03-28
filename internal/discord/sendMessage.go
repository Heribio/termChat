package discord

import (
    "fmt"
    "os"
    "encoding/json"
    "io"
    "github.com/Heribio/termChat/pkg/discordWebhook"
)

func SendMessage(message string) {
    jsonFile, err := os.Open("../../SECRETS/secrets.json")
    if err != nil {
        fmt.Println(err)
    }

    byteValue, err := io.ReadAll(jsonFile)

    type Secrets struct {
        WEBHOOK_URL string `json:"WEBHOOK_URL"`
    }

    var result Secrets

    json.Unmarshal(byteValue, &result)
    discordWebhook.SendMessage(message, result.WEBHOOK_URL)
}
