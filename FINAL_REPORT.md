```markdown
# Final Report - Incident Hub

### 1. O que foi entregue?
Dashboard de indicadores, listagem de incidentes com filtros por status e severidade, criação de incidentes, alteração de status com validação de regra de negócio, histórico de auditoria e persistência SQLite com seed inicial.

### 2. O que não foi entregue?
Nenhuma funcionalidade obrigatória ficou pendente.

### 3. O que você deliberadamente decidiu não fazer?
Autenticação de usuários, níveis de acesso, relatórios complexos e pipelines de build de frontend para manter o projeto enxuto e confiável no tempo disponível.

### 4. Quais foram as três principais decisões técnicas?
1. **Uso de Go nativo (`net/http`)**: Reduziu dependências externas e garantiu inicialização instantânea.
2. **SQLite sem CGO (`modernc.org/sqlite`)**: Evitou problemas de compilação CGO mantendo reprodutibilidade total.
3. **Frontend em HTML/JS estático**: Eliminou necessidade de `npm`/`webpack` facilitando a execução pelos avaliadores.

### 5. Qual foi o maior erro produzido pela IA durante o desenvolvimento?
Divergência na importação inicial dos pacotes Go quando executado fora da pasta correta do módulo.

### 6. Como você identificou esse erro?
Através das mensagens de erro apresentadas no terminal durante o comando `go mod tidy`.

### 7. Como você corrigiu e validou a correção?
Ajustando o caminho do diretório no terminal para a raiz do repositório antes da execução das ferramentas do Go.

### 8. Houve alguma regressão?
Não foram observadas regressões durante o desenvolvimento.

### 9. Em qual parte houve mais retrabalho?
No alinhamento do schema do banco entre a tabela principal e a tabela de histórico de transições.

### 10. Cite uma situação em que você rejeitou ou alterou uma abordagem sugerida pela IA.
Ajuste da estrutura de rotas para servir a pasta estática `/static` diretamente na rota raiz `/` sem frameworks web adicionais.

### 11. Qual parte da aplicação você considera menos confiável?
O formulário de criação no frontend por utilizar validação simples via JavaScript.

### 12. Se tivesse mais duas horas, quais seriam suas três prioridades?
1. Adicionar paginação e busca textual nos incidentes.
2. Implementar gráficos visuais no dashboard.
3. Adicionar mais testes de integração para os handlers HTTP.

### 13. Como você avalia sua estratégia inicial?
Muito eficiente. Priorizar a regra de negócio e os testes automatizados garantiu segurança antes do desenvolvimento da interface.

### 14. Aproximadamente quantas interações relevantes com IA foram necessárias?
5 interações principais bem delimitadas e documentadas no `AI_LOG.md`.

### 15. Quais ferramentas de IA foram utilizadas?
Gemini.