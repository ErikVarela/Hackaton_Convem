# AI Log - Incident Hub

## Interação 1 - Setup Inicial do Projeto
- **Objetivo:** Criar a estrutura base da aplicação Go com banco SQLite e servidor HTTP.
- **Contexto:** Prompt inicial para arquitetura simples sem necessidade de CGO.
- **Instrução:** Gerar `go.mod`, `main.go` com driver `modernc.org/sqlite` e endpoint `/health`.
- **Resultado:** Código gerado e dependências baixadas com sucesso.
- **Validação:** Servidor iniciado na porta `:8080` e endpoint `/health` retornando status 200 OK.
- **Decisão:** Avançar para a modelagem do banco de dados e tabela de incidentes.

## Interação 2 - Modelagem de Banco e Dados de Seed
- **Objetivo:** Criar tabelas no SQLite (`incidents` e `incident_history`) e carregar os dados iniciais de exemplo.
- **Contexto:** Especificação dos 3 incidentes iniciais e regras de persistência.
- **Instrução:** Implementar função `InitDB` em `db.go` para criação das tabelas e execução de seed idempotente.
- **Resultado:** Banco configurado com suporte ao histórico de transições e 3 registros criados.
- **Validação:** Verificação dos logs no terminal confirmando a criação das tabelas e inserção dos dados.
- **Decisão:** Desenvolver as regras do domínio (validação de transição de status para severidade `Critical`).

## Interação 3 - Regras de Negócio e Testes Automatizados
- **Objetivo:** Criar as estruturas de dados e implementar a regra de validação que impede transição direta de `Open` para `Resolved` em incidentes `Critical`.
- **Contexto:** Regra de negócio obrigatória especificada na seção 7 do desafio.
- **Instrução:** Criar `models.go` com a validação e `models_test.go` com casos de testes unitários cobrindo cenários válidos e inválidos.
- **Resultado:** Código da regra e suíte de testes unitários criados.
- **Validação:** Execução do comando `go test -v ./...` com 100% de aprovação nos testes.
- **Decisão:** Criar a camada de handlers/serviços HTTP e as páginas HTML para exibição e alteração de status.

## Interação 4 - Handlers HTTP e Endpoints da API
- **Objetivo:** Implementar os handlers para consulta do Dashboard, listagem com filtros, criação, detalhes com histórico e atualização de status.
- **Contexto:** Especificações de requisição e resposta do desafio para gerenciamento de incidentes.
- **Instrução:** Criar `handlers.go` contendo as rotas `/api/*` e registrá-las no `main.go`.
- **Resultado:** Endpoints REST funcionais cobrindo todos os requisitos de backend.
- **Validação:** Servidor executado via `go run .` com endpoints prontos para integração.
- **Decisão:** Construir a interface visual HTML (Dashboard e formulários de gerenciamento).

## Interação 5 - Interface Gráfica Web (Frontend HTML/CSS/JS)
- **Objetivo:** Criar a interface de usuário para exibição do Dashboard, filtros de lista, formulário de cadastro, visualização de histórico e alteração de status.
- **Contexto:** Necessidade de disponibilizar uma aplicação funcional para teste pelo avaliador.
- **Instrução:** Criar `static/index.html` e configurar o `http.FileServer` em `main.go` para a rota raiz `/`.
- **Resultado:** Interface responsiva e funcional integrada com todos os endpoints da API.
- **Validação:** Acesso a `http://localhost:8080` no navegador, testando a transição bloqueada de `Critical` (`Open` -> `Resolved`), verificação da mensagem de alerta na tela e atualização do histórico.
- **Decisão:** Finalizar a documentação obrigatória do repositório (`README.md` e `FINAL_REPORT.md`).