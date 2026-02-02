package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ClientResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição (server down ou timeout): %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	var quotation ClientResponse
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		log.Fatalf("Erro ao fazer parse do JSON: %v", err)
	}

	content := fmt.Sprintf("Dólar: %s", quotation.Bid)

	err = os.WriteFile("cotacao.txt", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Erro ao salvar arquivo: %v", err)
	}

	fmt.Println("Sucesso! Cotação salva em 'cotacao.txt'")
}
