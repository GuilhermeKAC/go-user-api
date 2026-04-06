# go-user-api

API REST para gerenciamento de usuários construída com Go puro, sem frameworks externos. Utiliza PostgreSQL como banco de dados e segue uma arquitetura em camadas com separação entre handlers, repositórios e modelos.

---

## Tecnologias

- **Go 1.22+** — linguagem principal
- **PostgreSQL** — banco de dados relacional
- **`database/sql`** — acesso ao banco sem ORM
- **`lib/pq`** — driver PostgreSQL para Go
- **`google/uuid`** — geração de IDs únicos

---

## Estrutura do projeto

```
go-user-api/
├── cmd/
│   ├── api/
│   │   └── main.go           # Entrypoint da API
│   └── create_table.go       # Script para criação da tabela
├── internal/
│   ├── config/
│   │   └── database.go       # Conexão com o banco de dados
│   ├── database/
│   │   └── migrations.go     # Criação das tabelas
│   ├── handlers/
│   │   └── user_handler.go   # Handlers HTTP
│   ├── models/
│   │   └── user.go           # Modelo de domínio
│   └── repository/
│       └── user_repository.go # Camada de acesso a dados
├── docs/
│   └── openapi.yaml          # Documentação OpenAPI 3.0
├── go.mod
└── go.sum
```

---

## Pré-requisitos

- Go 1.22+
- PostgreSQL rodando localmente

---

## Configuração

### 1. Banco de dados

Crie o banco de dados no PostgreSQL:

```sql
CREATE DATABASE go_user_api;
```

A conexão está configurada em `internal/config/database.go`:

```
user=postgres dbname=go_user_api sslmode=disable password=root
```

Ajuste as credenciais conforme seu ambiente.

### 2. Criar a tabela

Execute o script de migração:

```bash
go run cmd/create_table.go
```

Isso cria a tabela `users` com a seguinte estrutura:

```sql
CREATE TABLE IF NOT EXISTS users (
    id       VARCHAR(36)  PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    email    VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);
```

---

## Executando a API

```bash
go run cmd/api/main.go
```

O servidor sobe em `http://localhost:8080`.

---

## Endpoints

| Método   | Rota          | Descrição              |
|----------|---------------|------------------------|
| `POST`   | `/users`      | Criar novo usuário     |
| `GET`    | `/users`      | Listar todos os usuários |
| `GET`    | `/users/{id}` | Buscar usuário por ID  |
| `PUT`    | `/users/{id}` | Atualizar usuário      |
| `DELETE` | `/users/{id}` | Deletar usuário        |

> A documentação completa com schemas e exemplos está em [`docs/openapi.yaml`](docs/openapi.yaml).

---

## Exemplos de uso

### Criar usuário

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva", "email": "joao@example.com", "password": "senha123"}'
```

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "João Silva",
  "email": "joao@example.com"
}
```

### Listar usuários

```bash
curl http://localhost:8080/users
```

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "João Silva",
    "email": "joao@example.com"
  }
]
```

### Buscar por ID

```bash
curl http://localhost:8080/users/550e8400-e29b-41d4-a716-446655440000
```

### Atualizar usuário

```bash
curl -X PUT http://localhost:8080/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{"name": "João Atualizado", "email": "joao.novo@example.com"}'
```

### Deletar usuário

```bash
curl -X DELETE http://localhost:8080/users/550e8400-e29b-41d4-a716-446655440000
```

Retorna `204 No Content` em caso de sucesso.

---

## Modelo de dados

### Request body (criação/atualização)

| Campo      | Tipo   | Obrigatório | Descrição                          |
|------------|--------|-------------|------------------------------------|
| `name`     | string | sim         | Nome do usuário                    |
| `email`    | string | sim         | E-mail único                       |
| `password` | string | sim         | Senha (nunca retornada nas respostas) |

### Response

| Campo   | Tipo   | Descrição               |
|---------|--------|-------------------------|
| `id`    | string | UUID gerado pelo servidor |
| `name`  | string | Nome do usuário         |
| `email` | string | E-mail do usuário       |

> O campo `password` nunca é exposto nas respostas da API.
