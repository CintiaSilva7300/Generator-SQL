# Gerador de SQL a partir de Esquema Go

Este é um projeto simples que demonstra como gerar declarações SQL para criar tabelas a partir de um esquema Go definido em uma estrutura de dados.

## Funcionalidades

- Gera declarações SQL para criar tabelas a partir de um esquema Go.
- Suporta tipos de dados fixos e variáveis.

## Pré-requisitos

Certifique-se de ter Go instalado em sua máquina.

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu_usuario/gerador-sql-go.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd gerador-sql-go
   ```

## Uso

1. Defina o esquema Go no arquivo `schema.go`. Exemplo:

   ```go
   package schema

   type Person struct {
       ID         int    `json:"id"`
       FirstName  string `size:"fixed" maxlen:"50" json:"first_name"`
       LastName   string `size:"fixed" maxlen:"50" json:"last_name"`
       Age        int    `size:"fixed" maxlen:"3" json:"age"`
       // Adicione mais campos conforme necessário
   }
   ```

2. Execute o programa principal para gerar a declaração SQL:

   ```bash
   go run main.go
   ```

3. A declaração SQL será exibida no terminal.

## Exemplo de Saída

```JAVA
CREATE TABLE Person (
ID VARCHAR(255),
FirstName VARCHAR(50) NOT NULL,
LastName VARCHAR(50) NOT NULL,
Age VARCHAR(3) NOT NULL
)
```

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas ou enviar solicitações de pull.
