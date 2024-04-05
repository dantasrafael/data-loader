# data-loader

## O que esse código faz?

> Executar sql statements num banco de dados usando as go routines

## Configurações

> Pasta dos scripts SQL (./sql)

> Exemplo de Env File (.env):
- DB_HOST=localhost
- DB_NAME=data_loader_test
- DB_PORT=5432
- DB_USERNAME=postgres
- DB_PASSWORD=postgres
- DB_MAX_OPEN_CONNS=20 (opcional - padrão é 10)
- DB_MAX_IDLE_CONNS=10 (opcional - padrão é 3)
