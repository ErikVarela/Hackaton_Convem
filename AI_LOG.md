# AI Log - Incident Hub

## Interação 1 - Setup Inicial e Estrutura do Projeto
- **Objetivo:** Inicializar o módulo Go, configurar o servidor HTTP nativo (`net/http`) e o driver SQLite sem CGO (`modernc.org/sqlite`).
- **Contexto:** Definição da stack tecnológica sem frameworks adicionais.
- **Instrução:** Criar `main.go`, `go.mod` e testar a rota `/health`.
- **Resultado:** Servidor operando com resposta 200 OK no endpoint de health check.
- **Validação:** Requisição para `http://localhost:8080/health`.

## Interação 2 - Modelagem de Banco e Dados de Seed
- **Objetivo:** Criar tabelas no SQLite (`incidents` e `incident_history`) e carregar os dados iniciais de exemplo.
- **Contexto:** Especificação dos 3 incidentes iniciais e regras de persistência.
- **Instrução:** Implementar função `InitDB` em `db.go` para criação das tabelas e execução de seed idempotente.
- **Resultado:** Banco configurado com suporte ao histórico de transições e 3 registros criados.
- **Validação:** Verificação dos logs no terminal confirmando a criação das tabelas e inserção dos dados.

## Interação 3 - Regras de Negócio e Testes Automatizados
- **Objetivo:** Criar as estruturas de dados e implementar a regra de validação que impede transição direta de `Open` para `Resolved` em incidentes `Critical`.
- **Contexto:** Regra de negócio obrigatória especificada no desafio.
- **Instrução:** Criar `models.go` com a validação e `models_test.go` com casos de testes unitários.
- **Resultado:** Código da regra e suíte de testes unitários criados.
- **Validação:** Execução do comando `go test -v ./...` com 100% de aprovação.

## Interação 4 - Handlers HTTP e Endpoints da API
- **Objetivo:** Implementar os handlers para consulta do Dashboard, listagem com filtros, criação, detalhes com histórico e atualização de status.
- **Contexto:** Especificações de requisição e resposta do desafio para gerenciamento de incidentes.
- **Instrução:** Criar `handlers.go` contendo as rotas `/api/*` e registrá-las no `main.go`.
- **Resultado:** Endpoints REST funcionais cobrindo todos os requisitos de backend.
- **Validação:** Servidor executado via `go run .` com endpoints operacionais.

## Interação 5 - Interface Gráfica Web (Frontend HTML/CSS/JS)
- **Objetivo:** Criar a interface de usuário para exibição do Dashboard, filtros de lista, formulário de cadastro, visualização de histórico e alteração de status.
- **Contexto:** Necessidade de disponibilizar uma aplicação funcional para teste pelo avaliador.
- **Instrução:** Criar `static/index.html` e configurar o `http.FileServer` em `main.go` para a rota raiz `/`.
- **Resultado:** Interface responsiva e funcional integrada com todos os endpoints da API.
- **Validação:** Acesso a `http://localhost:8080` no navegador com validação visual e operacional.

## Interação 6 - Change Request #1 (Comentários e Timeline Cronológica)
- **Objetivo:** Implementar sistema de comentários em incidentes e unificar histórico de status e comentários em uma timeline cronológica.
- **Contexto:** Recebimento do Change Request #1 exigindo comentários obrigatórios e visão unificada da atividade.
- **Instrução:** Atualizar `db.go` (tabela `incident_comments`), `models.go`, `models_test.go`, `handlers.go` (`POST /api/incidents/comment`), `main.go` e `static/index.html`.
- **Resultado:** Comentários persistidos e timeline integrada exibindo alterações de status e comentários em ordem cronológica.
- **Validação:** Execução com sucesso da suíte de testes unitários (`go test -v ./...`) e validação via frontend.
