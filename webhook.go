package main

import (
	"fmt"
	"net/http"
	"bytes"
)

func main() {
	url := "https://chat.googleapis.com/v1/spaces/AAAARyc6yA8/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=Yxf1NvbziWxyRpgsVfvo-yf08lgjaAX-33IhU-ECRJs%3D" // substitua pelo URL do seu webhook

	jsonStr := []byte(`{ "text": "Teste para envio de mensagem para BOT! \nExemplo da mensagem: Erro DP2040" }`) // substitua pelo corpo da mensagem que deseja enviar

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
