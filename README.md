# InfraForge

InfraForge é uma plataforma de monitoramento de serviços e infraestrutura, construída em Go com foco em confiabilidade, observabilidade e arquitetura moderna.

## 🧠 Visão do projeto

O objetivo do InfraForge é simular ferramentas utilizadas em ambientes reais de DevOps/SRE, como:

- Monitoramento contínuo de serviços
- Verificação automática de saúde (health check)
- Coleta de métricas básicas (status e tempo de resposta)
- Persistência de dados para análise
- Base para sistemas de alerta e observabilidade

---

## ⚙️ Funcionalidades

- Cadastro de serviços (URL, ambiente)
- Monitoramento automático em background (worker)
- Verificação de status (online/offline)
- Medição de tempo de resposta
- Registro de status code HTTP
- Listagem de serviços monitorados
- Remoção de serviços
- Health check da própria API
- Controle do worker via variáveis de ambiente
- Graceful shutdown do servidor

---

## 🏗️ Arquitetura

O projeto segue uma estrutura modular:

```txt
cmd/
internal/
  ├── config/
  ├── database/
  ├── health/
  ├── services/
  ├── worker/
  └── server/
migrations/

```
🔹 Backend
- Go (Golang)
- Gin (HTTP server)
- PostgreSQL (persistência)
- pgxpool (pool de conexões)

🔹 Infra
- Docker + Docker Compose
- Migrations SQL versionadas
- Configuração via .env

🔄 Worker (monitoramento)

O InfraForge possui um worker que roda em background utilizando goroutines:

- Executa checks periódicos
- Faz requisições HTTP para os serviços cadastrados
- Calcula tempo de resposta
- Atualiza status no banco de dados
- Controlado via context (shutdown seguro)

🧪 Endpoints
Health check:
- GET /health

Criar serviço:
- POST /services

{
  "name": "InfraForge API",
  "url": "http://localhost:8080/health",
  "environment": "dev"
}

Listar serviços:
- GET /services

Remover serviço:
- DELETE /services/:id

⚙️ Configuração

Crie um .env baseado em:

- PORT=8080
- APP_ENV=development
- DATABASE_URL=postgres://postgres:postgres@localhost:5433/infraforge?sslmode=disable
- MONITOR_ENABLED=true
- MONITOR_INTERVAL_SECONDS=30

🐳 Rodando com Docker
- docker compose up -d

▶️ Rodando a API
- go run ./cmd/api

🗄️ Migrations

As migrations estão na pasta:

migrations/

Para executar:

Get-Content migrations/arquivo.sql | docker exec -i infraforge-db psql -U postgres -d infraforge

📈 Próximos passos:

- Dashboard (React)
- Sistema de alertas
- Métricas com Prometheus
- Logs estruturados
- Deploy em cloud

👨‍💻 Autor

Caique Brandão
Desenvolvedor Full Stack, especialista em arquitetura e sistemas escaláveis.
