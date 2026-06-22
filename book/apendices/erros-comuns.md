# Erros comuns — Go Essentials

## `declared and not used`
Você criou uma variável e não usou. Em Go isso é erro de compilação.

## `imported and not used`
Importou pacote sem usar. Remova o import ou use o pacote.

## `cannot use X as Y`
Tipos diferentes sem conversão explícita. Converta (`int(x)`, `float64(x)`, etc.).

## `undefined: X`
Nome não existe no escopo atual. Pode ser typo, falta de import, ou item não exportado.

## Erros de package/import path
- Caminho de módulo em `go.mod` não bate com imports.
- Pacote com nome diferente do diretório esperado.

## `index out of range`
Acesso fora do limite de slice/array. Valide `len(...)` antes.

## `assignment to entry in nil map`
Map não foi inicializado. Use `make(map[T]U)` antes de atribuir.

## `missing go.mod`
Você rodou comando de módulo fora de um projeto inicializado.
Use `go mod init <caminho-do-modulo>`.
