# Pratique — Tipos compostos

## Revisão rápida

1. Qual é a diferença entre array e slice?
2. Por que slices aparecem muito mais no dia a dia do que arrays?
3. O que `len` retorna em um slice?
4. O que `cap` retorna em um slice?
5. O que `append` faz quando o slice não tem capacidade suficiente?
6. Como inicializar um map antes de adicionar valores?

## Leia o código

```go
contacts := []string{"Bia", "Ana"}
contacts = append(contacts, "Joao")

fmt.Println(len(contacts))
fmt.Println(contacts[2])
```

1. Qual é o tamanho final do slice?
2. Qual valor aparece em `contacts[2]`?
3. O que aconteceria se o código tentasse acessar `contacts[3]`?

## Maps sem susto

```go
phones := map[string]string{
    "Bia": "11999990000",
}

phone, ok := phones["Ana"]
```

1. Qual valor vai para `phone`?
2. Qual valor vai para `ok`?
3. Por que o padrão `valor, ok` é importante em maps?

## Pensando na agenda

1. Quando faz sentido retornar contatos em `[]Contact`?
2. Por que `map[string]Contact` é adequado enquanto o UUID ainda é representado como string?
3. Por que `Contact.ID`, e não `Email`, deve ser a chave do map?
4. Como evitar `assignment to entry in nil map`?
5. O que muda na Aula 6 quando a chave passa para `uuid.UUID`?
