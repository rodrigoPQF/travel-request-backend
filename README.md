# Teste Técnico - Desenvolvedor Golang Onfly

## Introdução

Este é um projeto técnico desenvolvido para o processo seletivo da Onfly. A aplicação está escrita em Go (Golang) e utiliza Docker para a configuração de ambientes, facilitando a execução e o gerenciamento do banco de dados Postgres em contêineres. Além disso, a aplicação suporta desenvolvimento interativo e inclui testes, documentação swagger.

### Arquitetura e Expansão

Este projeto foi desenvolvido com uma arquitetura baseada em microserviços, permitindo fácil expansão e adição de novos serviços conforme as necessidades do negócio.

### Algumas técnologias utilizadas

- Fiber
- Docker
- Mockgen
- Swagger
- Postgres
- Gorm

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas no sistema:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Dev Container](https://code.visualstudio.com/docs/remote/containers) (opcional para desenvolvimento em containers no VSCode)

## Configuração do Ambiente

### 1. Clone o repositório:

```shell
git clone <URL_DO_REPOSITORIO>
cd <PASTA_DO_PROJETO>
```

### 2. Configure as variáveis de ambiente:

- Copie o arquivo .env.example para um novo arquivo .env:

```shell
cp .env.example .env
```

- Preencha as variáveis de ambiente necessárias no arquivo .env conforme sua configuração.

### 3. O banco de dados Postgres já está configurado no Docker, então não é necessário configurá-lo separadamente.

## Executando o Serviço Localmente

Para rodar o projeto, siga um dos métodos abaixo:

### Método 1: Executar com Docker Compose

Inicie o serviço com Docker Compose:

```shell
docker compose up
```

### Método 2: Executar Interativamente

Para executar de forma interativa dentro do contêiner, use:

```shell
docker-compose run --rm app bash
make start-dev
```

## Gerar Executável

Para compilar e gerar o executável do projeto, rode:

```bash
docker-compose run --rm app bash
make runbuild
```

## Executando os Testes

Para rodar os testes interativamente com o uso de mocks, utilize:

```bash
docker-compose run --rm app bash
make test
```

Este projeto utiliza o pacote Testify para facilitar a criação de mocks e verificação de comportamentos, o que simplifica os testes unitários.

# Documentação da API com Swagger

A documentação da API foi gerada usando o Swagger e está disponível em uma interface interativa para facilitar o teste e visualização dos endpoints.

- Após iniciar a aplicação, acesse a documentação no navegador em:

```bash
http://localhost:5432/swagger/index.html
```

# Utilização com Dev Container (Opcional)

Caso prefira utilizar um Dev Container, esta aplicação está preparada para funcionar no ambiente de contêineres do VSCode. Basta abrir o projeto em um ambiente configurado para Dev Containers, e o VSCode irá configurar automaticamente o ambiente de desenvolvimento.
