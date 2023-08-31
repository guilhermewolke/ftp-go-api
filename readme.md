# API de Livros

## O que é?

Um exemplo de API desenvolvido utilizando Golang, MySQL e Redis como Caching.

Trata-se de uma API que expõe um endpoint de listagem de autores, e de livros vinculados a estes autores, consultados
à partir de um banco de dados MySQL.

Entre a camada do MySQL e do endpoint que realiza esta consulta, existe uma camada de Caching em Redis, que é consultado
primeiramente. Se o registro existir no Redis, a API entrega na requisição a informação recuperada deste serviço.

Caso a informação não seja encontrada no Redis, a consulta é realizada no MySQL e um cache é criado no Redis, expirando 10 segundos após esta criação.  

## Como Executar?

1) Baixe o repositório em sua estação;
2) Com o serviço Docker ativo, execute o seguinte comando em seu terminal:

   > docker-compose up

    Este comando iniciará os serviços MySQL, Redis, e executará o "build" presente no diretório raíz (/app dentro do container).
3) Acesse em seu navegador o seguinte endereço:

   > http://localhost:8080/

4) Caso este seja seu primeiro acesso, acesse o endpoint abaixo, para popular as tabelas:


      > http://localhost:8080/setup


    
## Endpoints disponíveis

      > /autores

Lista todos os autores cadastrados no sistema

      > /autores/{id}

Exibe os dados de um autor, baseado no ID deste autor. O ID é um valor numérico.

      > /autores/{id}/livros

Lista de livros vinculados à este autor, baseado no ID do autor. O ID é um valor numérico

      > /livros

Lista todos os livros cadastrados no sistema

      > /livros/{id}

Exibe os dados de um livro e de seu autor, baseado no ID do livro. O ID é um valor numérico.
