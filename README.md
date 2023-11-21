# movieBlog
Aplicação de um mini blog de filmes em que o usuário pode criar uma lista de filmes com seus títulos e diretores,
para a prática de Golang + HTMX.
- Feito em Go 1.21.4
- Postgres 16.0.1 de DB 
- HTMX no front

Para rodar, basta criar um arquivo .env na raiz do projeto, onde os dados são o seguinte:

- PGHOST= URL em que o banco em Postgres está rodando. Se você está rodando o banco localmente, seria "localhost"
- PGPORT= Porta em que o Postgres está rodando.
- PGUSER= Usuário responsável pelo Postgres.
- PGPASSWORD= Senha do usuário do Postgres.
- PGDATABASE= Nome do banco que você criou para essa aplicação.
- SECRET_KEY= Não é usado na aplicação, mas seria onde colocaríamos a chave de codificação de coisas como o Bcrypt
- PORT= Porta em que a aplicação rodará no browser.

Depois de tudo criado, basta rodar `go run cmd/main.go` em um terminal aberto na raiz do repositório.
