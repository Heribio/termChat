package discordWebhook

import (
    "net/http"
    "fmt"
    "os"
    "bytes"
    "encoding/json"
)

func SendMessage(content Message, webhookUrl string) {
    payload, err := json.Marshal(content)
    if err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
    if err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }

    req, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
    fmt.Println(req.Status)
}
