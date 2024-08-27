# Documentação do Projeto

## Visão Geral

Este projeto é uma aplicação de backend em Go que gerencia um catálogo de produtos. A arquitetura é baseada em Clean Architecture, utilizando pacotes para isolar as responsabilidades e promover uma estrutura organizada e modular.

## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

```
.
│
├── cmd/
│   └── main.go                    # Ponto de entrada da aplicação
│
├── internal/
│   ├── domain/
│   │   ├── product.go             # Definições das entidades e interfaces de repositório
│   │   └── product_repository.go  # Interface do repositório
│   │
│   ├── infrastructure/
│   │   ├── db/
│   │   │   ├── db.go              # Função ConnectDB e configurações do banco de dados
│   │   │   └── product_repository_db.go # Implementação concreta do repositório
│   │   └── config.go              # Carregamento de variáveis de ambiente e configurações
│   │
│   └── use_cases/
│       ├── get_all_products.go     # Caso de uso para obter todos os produtos
│       ├── get_by_id_product.go    # Caso de uso para obter um produto por ID
│       ├── save_product.go         # Caso de uso para salvar um novo produto
│       ├── update_product.go       # Caso de uso para atualizar um produto existente
│       └── delete_product.go       # Caso de uso para excluir um produto
│
├── interfaces/
│   └── api/
│       └── controllers/
│           └── product_controller.go # Controladores HTTP
│
├── tests/
│
├── migrations/
│
├── go.mod
└── go.sum
```
