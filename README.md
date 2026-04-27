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
