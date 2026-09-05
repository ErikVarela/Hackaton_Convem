# 🚨 Incident Hub

Aplicação web para registro, acompanhamento e gestão do ciclo de vida de incidentes operacionais.

---

## 🛠️ Pré-requisitos
- **Go** (versão 1.22 ou superior)

---

## 📥 Instalação

Clone o repositório e acesse a pasta do projeto:
```bash
git clone [https://github.com/ErikVarela/Hackaton_Convem.git](https://github.com/ErikVarela/Hackaton_Convem.git)
cd Hackaton_Convem

Sincronize e baixe as dependências:

go mod tidy

🚀 Execução
Inicie o servidor da aplicação:

go run .

Acesse no navegador:

http://localhost:8080

📊 Dados Iniciais (Seed)
A aplicação popula o banco SQLite (incidents.db) automaticamente na primeira inicialização com os 3 incidentes de exemplo exigidos:

Payment API instability — Critical | Open

Reconciliation delay — High | In Progress

Incorrect customer notification — Medium | Resolved

💡 Nota: Para resetar o banco de dados para o estado inicial, basta excluir o arquivo incidents.db na raiz e reiniciar o servidor com go run ..

🧪 Execução dos Testes
Para rodar a suíte de testes unitários das regras de negócio:

go test -v ./...

🏗️ Arquitetura
Backend: Go (net/http) com arquitetura simples baseada em handlers REST.

Persistência: SQLite (modernc.org/sqlite — driver nativo em Go sem dependência de CGO).

Frontend: HTML5, CSS3 e JavaScript puro servidos estaticamente.

Regras Críticas: Validação estrita impedindo transição direta de Open ➔ Resolved em incidentes Critical (exige passagem intermediária por In Progress) com registro em tabela de histórico.

---

### Passo para Atualizar o Git

Substitua o conteúdo do seu arquivo **`README.md`** local por esse código e execute os comandos abaixo no terminal:

```powershell
git add README.md
git commit -m "docs: reestiliza layout visual do README"
git push origin main
git rev-parse HEAD

(Lembre-se de utilizar git push origin master se a sua branch for master).