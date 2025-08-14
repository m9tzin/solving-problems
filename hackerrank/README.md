# HackerRank Solutions

Este diretório contém soluções para problemas do HackerRank em Go.

## Como executar os algoritmos

### Bubble Sort
```bash
go run -tags bubble bubble-sort.go
```

### Simple Array Sum
```bash
go run -tags sum simple-array-sum.go
```

## Estrutura do Projeto

```
hackerrank/
├── go.mod              # Módulo Go
├── bubble-sort.go      # Algoritmo Bubble Sort (build tag: bubble)
├── simple-array-sum.go # Algoritmo Simple Array Sum (build tag: sum)
└── README.md          # Este arquivo
```

## Explicação dos Build Tags

Os build tags (`//go:build`) permitem que múltiplos arquivos com função `main` coexistam no mesmo diretório:

- `bubble-sort.go` usa o tag `bubble`
- `simple-array-sum.go` usa o tag `sum`

Isso permite executar cada algoritmo independentemente sem conflitos de função `main`.

## Resolução de Problemas

Se você encontrar o erro "No packages found", certifique-se de que:
1. O arquivo `go.mod` existe no diretório
2. Você está executando os comandos com as tags corretas
3. O Go está configurado corretamente no seu ambiente
