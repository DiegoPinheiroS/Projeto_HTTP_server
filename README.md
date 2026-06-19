# Projeto Korp - Automação de Stack de Monitorização com Ansible e Docker

Este projeto implementa uma automação completa do provisionamento e configuração de uma infraestrutura de monitorização utilizando **Ansible** como ferramenta de Infraestrutura como Código (IaC) e **Docker / Docker Compose** para a orquestração dos serviços de monitorização (**Prometheus** e **Grafana**), para monitoramente de uma api simples desenvolvida em Golang.

## Arquitetura do Projeto

O projeto foi desenhado seguindo a arquitetura padrão de mercado para gerenciamento de configuração:

* **Control Node (Nó de Controle):** Notebook executando Linux, responsável por armazenar os playbooks e disparar as automações via SSH, além de servir como base para o desenvolvimento.
* **Managed Node (Nó Gerenciado):** Máquina Virtual (VirtualBox) executando Ubuntu Server, configurada em modo *Bridge*, que recebe a stack de monitorização.

A comunicação entre as máquinas é feita de forma segura através de par de chaves criptográficas SSH, eliminando a necessidade de autenticação interativa por senha durante o deploy.

---

## Tecnologias Utilizadas

* **Ansible (v2.x+):** Automatização do deploy, configuração do sistema operativo e garantia de idempotência.
* **Docker & Docker Compose (Plugin V2):** Conteinerização e isolamento das aplicações de monitorização.
* **Prometheus:** Coleta e armazenamento de métricas de séries temporais da infraestrutura.
* **Grafana:** Construção de dashboards interativos para visualização de dados e métricas.
* **VirtualBox:** Para testes do deploy com um instalação limpa do ubuntu.

---

## 📂 Estrutura de Diretórios

```text
├── ansible/
│   ├── inventory.ini       # Definição dos hosts alvo e variáveis de conexão
│   └── playbook.yml        # Instruções e tarefas automatizadas de provisionamento
├── prometheus/
│   └── prometheus.yml      # Configuração de targets e retenção do Prometheus
├── grafana/                # Diretório de persistência de dados e dashboards do Grafana
├── docker-compose.yml      # Definição dos serviços (containers) da stack
└── README.md               # Documentação do projeto
