package bot 

type Secrets struct {
    WEBHOOK_URL string `json:"WEBHOOK_URL"`
    BOT_TOKEN string `json:"BOT_TOKEN"`
    BOT_ID string `json:"BOT_ID"`
}

type discordMessage struct {
    Content string `json:"content"`
    Username string `json:"username"`
}
