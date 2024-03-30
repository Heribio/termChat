package discord

import (
    "fmt"
    "os"
    "encoding/json"
    "io"
    "github.com/Heribio/termChat/pkg/discordWebhook"
)

func SendMessage(message discordWebhook.Message) {
    jsonFile, err := os.Open("../../SECRETS/secrets.json")
    if err != nil {
        fmt.Println("Error opening JSON file:", err)
        fmt.Println("Create a file ./SECRETS/secrets.json with the following structure: {\"WEBHOOK_URL\": YOUR_WEBHOOK_URL HERE\"}")
        os.Exit(1)
    }
    defer jsonFile.Close() // Close the file when done reading

    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        os.Exit(1)
    }

    var result Secrets
    err = json.Unmarshal(byteValue, &result)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        os.Exit(1)
    }
    discordWebhook.SendMessage(message, result.WEBHOOK_URL)
}
