# Planning Document - Incident Hub

## 1. Entendimento do Problema
O objetivo é desenvolver uma aplicação web funcional chamada **Incident Hub** para apoiar pequenas equipes de operações no registro, acompanhamento e gerenciamento do ciclo de vida de incidentes operacionais. A aplicação resolverá a falta de visibilidade atual centralizando o status dos incidentes, severidade, responsáveis, histórico de mudanças e um dashboard resumido.

## 2. Escopo

### Obrigatório (MVP)
- **Gestão de Incidentes:**
  - Criação de incidentes com os campos obrigatórios (título, descrição, severidade, responsável).
  - Status inicial padrão `Open` com timestamp automático.
- **Listagem e Detalhes:**
  - Visualização de lista e detalhes completos de cada incidente.
  - Filtros por `status` e por `severidade`.
- **Regras de Negócio e Transições de Status:**
  - Validação estrita: Incidentes com severidade `Critical` **não podem** transitar diretamente de `Open` para `Resolved` (devem obrigatoriamente passar por `In Progress`).
  - Bloqueio e feedback amigável ao usuário em transições inválidas.
- **Histórico e Dashboard:**
  - Registro e persistência do histórico com data/hora para todas as mudanças de status.
  - Dashboard resumido em tempo real com contadores: incidentes abertos, incidentes críticos não resolvidos e incidentes resolvidos.
- **Persistência e Seed Inicial:**
  - Armazenamento persistente em banco de dados SQLite.
  - Carregamento inicial automático/mecanismo fácil para disponibilizar os 3 incidentes de exemplo especificados.
- **Testes e Qualidade:**
  - Testes unitários/de integração cobrindo a regra de transição de status e validações críticas.

### Desejável
- Interface web limpa e responsiva utilizando HTML/CSS (ou HTMX/Tailwind via CDN).
- Botão/mecanismo na interface para resetar/recarregar os dados iniciais de exemplo (Seed).

### Fora de Escopo
- Autenticação e controle de permissões por usuário.
- Multi-tenancy ou gestão de múltiplas organizações.
- Notificações por e-mail/webhooks.

## 3. Decisões Técnicas
- **Linguagem / Backend:** Go (Golang) — Escolhido pela alta performance, simplicidade, compilação rápida e suporte nativo robusto a servidores HTTP.
- **Frontend:** HTML5 + CSS / Templates HTML nativos de Go (ou HTMX para interações sem reflash pesado) — Garante simplicidade, sem necessidade de pipelines complexos de build frontend (como npm/webpack).
- **Persistência:** Banco de dados SQLite — Ideal para o escopo, garante persistência local real dos dados sem exigir setup de serviços externos pesados por parte dos avaliadores.
- **Estratégia de Testes:** Testes automatizados com a biblioteca padrão `testing` de Go cobrindo a camada de regras de negócio (transição de status e validações).

## 4. Decomposição de Tarefas
1. **Setup Inicial:** Estrutura do projeto Go, conexão com SQLite e criação do arquivo `START.md` e `PLAN.md`.
2. **Modelagem & Banco:** Definição das tabelas de `incidents` e `incident_history`, e script de carga dos dados iniciais (seed).
3. **Regra de Negócio:** Implementação da lógica do domínio e validação de transição de status de incidentes `Critical` com testes automatizados.
4. **Endpoints & Handlers HTTP:** Criação das rotas (Criar, Listar/Filtrar, Detalhar, Atualizar Status, Dashboard).
5. **Interface do Usuário (UI):** Construção das páginas HTML/Visões para listagem, dashboard e histórico.
6. **Validar & Refinar:** Verificação do fluxo completo, tratamento de erros e atualização dos documentos (`AI_LOG.md`, `README.md`, `FINAL_REPORT.md`).

## 5. Critérios de Aceite
- [ ] Incidentes `Critical` em status `Open` **não** aceitam transição direta para `Resolved` e retornam mensagem de erro clara.
- [ ] O histórico registra cada transição de status com data/hora e permanece salvo no SQLite após reiniciar a aplicação.
- [ ] O dashboard exibe a contagem correta dos 3 indicadores principais.
- [ ] A aplicação inicia já com os 3 incidentes de exemplo carregados.
- [ ] Todos os testes automatizados executam e passam com sucesso via `go test`.

## 6. Riscos Identificados
- **Tempo para alinhamento com a IA:** Dificuldade em instruir a IA para refatorações específicas mantendo a regra de 100% linguagem natural.
- **Gargalos de UI:** Gastar tempo excessivo ajustando estilos CSS em vez das regras de negócio.
- *Mitigação:* Focar primeiro na funcionalidade, regra de negócio e persistência; utilizar estilos CSS simples e funcionais.

## 7. Estratégia de Utilização de IA
- Utilização de prompts estruturados por camadas (Domínio/Regra → Persistência → Handlers HTTP → UI).
- Todo o código será gerado via instruções em linguagem natural (prompting).
- Validação e execução contínua de testes a cada alteração para garantir que novas funcionalidades não quebrem o que já está funcionando.