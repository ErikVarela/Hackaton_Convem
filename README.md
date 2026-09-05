# Incident Hub

Aplicação web para registro, acompanhamento e gestão do ciclo de vida de incidentes operacionais.

## Pré-requisitos
- Go (versão 1.22 ou superior)

## Instalação
Clone o repositório e baixe as dependências:
```bash
git clone [https://github.com/ErikVarela/Hackaton_Convem.git](https://github.com/ErikVarela/Hackaton_Convem.git)
cd Hackaton_Convem
go mod tidy

Execução
Inicie a aplicação executando: go run .

Acesse a aplicação no navegador em: http://localhost:8080

Dados Iniciais (Seed)
A aplicação popula o banco SQLite (incidents.db) automaticamente na primeira inicialização com os 3 incidentes de exemplo exigidos:

Payment API instability (Critical - Open)

Reconciliation delay (High - In Progress)

Incorrect customer notification (Medium - Resolved)

Para resetar os dados, basta deletar o arquivo incidents.db na raiz e reiniciar o servidor com go run ..

Execução dos Testes
Execute a suíte de testes unitários automatizados da regra de negócio: go test -v ./...

Arquitetura
Backend: Go (net/http) com arquitetura simples baseada em handlers REST.

Persistência: SQLite (modernc.org/sqlite - driver nativo em Go sem dependência de CGO).

Frontend: HTML5, CSS3 e JavaScript puro servidos estaticamente.

Regras Críticas: Validação estrita impedindo transição direta de Open -> Resolved em incidentes Critical (exige passagem intermediária por In Progress) com registro completo em tabela de histórico.