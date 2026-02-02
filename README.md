# Desafio Go - Client-Server API (Cotação do Dólar)

Este projeto é a solução para o desafio de implementação de APIs em Go, focado no uso de **Context**, **Webserver HTTP**, **JSON**, **Manipulação de Arquivos** e **Banco de Dados (SQLite)**.

## 📋 Sobre o Desafio

O objetivo é criar dois sistemas (`client.go` e `server.go`) que trocam informações sobre a cotação do dólar, respeitando regras estritas de timeout (tempo limite) para cada operação.

### Requisitos Funcionais

1.  **server.go**:
    * Deve consumir a API externa `https://economia.awesomeapi.com.br/json/last/USD-BRL`.
    * **Timeout API:** Máximo de **200ms** para receber a resposta externa.
    * **Persistência:** Deve salvar cada cotação recebida em um banco de dados SQLite (`quotation.db`).
    * **Timeout DB:** Máximo de **10ms** para persistir os dados.
    * **Endpoint:** Disponibilizar os dados na rota `/cotacao` na porta `:8080`.

2.  **client.go**:
    * Deve realizar uma requisição HTTP ao `server.go`.
    * **Timeout Client:** Máximo de **300ms** para receber a resposta do servidor.
    * **Arquivo:** Deve salvar apenas o valor do câmbio (campo `bid`) em um arquivo `cotacao.txt` no formato `Dólar: {valor}`.

3.  **Geral**:
    * O sistema deve gerar logs de erro caso os tempos de execução (timeouts) sejam excedidos.

## 🚀 Como Executar

### Pré-requisitos
* Go instalado (versão 1.18+)
* GCC instalado (necessário para o driver do SQLite - `go-sqlite3`)

### Passo 1: Clone o repositório e baixe as dependências

```bash
# Clone o projeto
git clone <seu-link-do-github-aqui>
cd <nome-da-pasta>

# Baixe as dependências (Driver SQLite)
go mod tidy
```
### Passo 2: Execute o Servidor

Abra um terminal e rode:
```Bash

go run server.go

O servidor iniciará na porta 8080 e criará o banco de dados cotacoes.db automaticamente.
```
### Passo 3: Execute o Cliente

Abra um segundo terminal e rode:
```Bash

go run client.go
```
### ✅ Resultado Esperado

No terminal do cliente, você verá a mensagem:
```Bash
Sucesso! Cotação salva em 'cotacao.txt'.
```
Um arquivo cotacao.txt será criado na raiz com o conteúdo:

```Bash
    Dólar: 5.2568

    O banco de dados cotacoes.db terá o registro histórico da cotação.
```
### 🛠 Tecnologias Utilizadas

Golang (Standard Library: net/http, context, encoding/json, io, os)

SQLite3 (Persistência de dados)

Context (Gerenciamento de timeouts e cancelamento de requisições)