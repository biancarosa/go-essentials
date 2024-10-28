# Idéias de Projeto

Para aplicar o que você aprendeu no curso, seguem algumas idéias de projeto:

## JWT Encoder/Decode

Crie uma aplicação que faça o encoding/decoding de JWT tokens: https://jwt.io usando HMACSHA256.

- `go run . decode eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o secret`
- `go run . encode '{"sub": "1234567890","name": "John Doe","iat": 1516239022}' secret`

## Calculadora

A calculadora é um excelente projeto para manipular números e argumentos de linha de comando.

Se você quiser continuar brincando com esse projeto, seguem algumas idéias:

Garantir que seu programa funciona e retorna valores corretos para as entradas:

- `go run . =100/10`
- `go run . =100*1.5`
- `go run . =100/0`
- `go run . =100*0`
- `go run . =5+10+10+5+200`
- `go run . =(5+10+10)*5+200`

(E por aí vai :D)

# Implementações

Fique à vontade pra abrir um PR pra esse repo com o link da sua implementação:

- [Calculadora - @biancarosa](https://github.com/biancarosa/go-essentials/tree/main/code-examples/calc)


