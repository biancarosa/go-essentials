# Compilando Código Go

## Introdução

Go é uma linguagem compilada, similar a C. O processo de compilação segue o seguinte fluxo:

```
CÓDIGO -> COMPILADOR DO CÓDIGO -> LINGUAGEM DE MÁQUINA -> LINKING (LINKAGEM) -> EXECUTÁVEL (PROGRAMA)
```

Exemplo: `main.go -> compilador do Go -> executável main`

### Características Importantes

- Go é uma linguagem multi-plataforma
- Linguagens de programação podem ser interpretadas ou compiladas
- Linguagens compiladas geralmente são mais performáticas, mas isso não é uma regra absoluta

### Comandos Básicos

- `go run main.go`: Compila e executa o arquivo main.go
- `go build`: Compila o código e gera um executável para a plataforma atual

## Compilação Cruzada em Go

Go oferece poderosas capacidades de compilação cruzada, permitindo que você compile código para diferentes sistemas operacionais e arquiteturas a partir de uma única máquina. Isso é particularmente útil para desenvolvimento e distribuição de software multiplataforma.

### Conceitos Chave

1. **GOOS (Go Operating System)**: Esta variável de ambiente especifica o sistema operacional alvo para o qual você está compilando.
   - Valores comuns: `linux`, `windows`, `darwin` (para macOS)

2. **GOARCH (Go Architecture)**: Esta variável define a arquitetura do processador alvo.
   - Valores comuns: `amd64` (para 64-bit x86), `386` (para 32-bit x86), `arm` (para ARM)

### Como Funciona

Quando você compila código Go, o compilador gera um binário específico para a combinação GOOS/GOARCH alvo. Por padrão, Go compila para o sistema operacional e arquitetura da máquina onde está sendo executado.

Para compilar para uma plataforma diferente, você pode definir as variáveis GOOS e GOARCH antes de executar o comando `go build`. Por exemplo:

```
GOOS=windows GOARCH=amd64 go build
```

Este comando compilará seu código para Windows 64-bit, independentemente do sistema em que você está trabalhando.

### Exemplos Práticos

1. Compilando para Linux 64-bit a partir de qualquer sistema:
   ```
   GOOS=linux GOARCH=amd64 go build
   ```

2. Compilando para macOS (Darwin) ARM64 (por exemplo, para M1 Macs):
   ```
   GOOS=darwin GOARCH=arm64 go build
   ```

3. Compilando para Windows 32-bit:
   ```
   GOOS=windows GOARCH=386 go build
   ```

### Vantagens da Compilação Cruzada

1. **Desenvolvimento Multiplataforma**: Desenvolva em uma plataforma e compile facilmente para outras.
2. **CI/CD Simplificado**: Gere binários para múltiplas plataformas em um único pipeline de CI/CD.
3. **Distribuição Facilitada**: Crie versões do seu software para diferentes sistemas sem necessidade de múltiplas máquinas.

### Considerações Importantes

- Nem todas as bibliotecas suportam compilação cruzada. Bibliotecas que dependem de código C específico da plataforma podem causar problemas.
- Testes em plataformas alvo são cruciais, pois alguns comportamentos podem variar entre sistemas operacionais.
- Para projetos complexos, pode ser necessário configurar ferramentas de compilação cruzada adicionais, especialmente ao lidar com dependências nativas.

## Casos de Uso Úteis

A compilação em Go oferece diversas vantagens práticas. Aqui estão alguns casos de uso úteis que demonstram o poder e a flexibilidade da compilação em Go:

1. **Distribuição de Aplicativos Standalone**
   - Cenário: Você desenvolveu uma ferramenta CLI em Go.
   - Uso: Compile o aplicativo para diferentes sistemas operacionais e arquiteturas.
   - Benefício: Usuários podem executar seu aplicativo sem instalar Go ou dependências adicionais.
   - Exemplo:
     ```
     GOOS=linux GOARCH=amd64 go build -o myapp-linux-amd64
     GOOS=darwin GOARCH=arm64 go build -o myapp-mac-arm64
     GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64.exe
     ```

2. **Microserviços em Contêineres**
   - Cenário: Desenvolvimento de microserviços para uma arquitetura em nuvem.
   - Uso: Compile aplicativos Go para contêineres leves.
   - Benefício: Imagens de contêiner menores e mais eficientes.
   - Exemplo Dockerfile:
     ```dockerfile
     FROM golang:1.22 as builder
     WORKDIR /app
     COPY . .
     RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

     FROM alpine:latest  
     RUN apk --no-cache add ca-certificates
     WORKDIR /root/
     COPY --from=builder /app/main .
     CMD ["./main"]
     ```

3. **Desenvolvimento e Testes Multiplataforma**
   - Cenário: Você está desenvolvendo um aplicativo que deve rodar em Windows, macOS e Linux.
   - Uso: Compile e teste em todas as plataformas a partir de uma única máquina de desenvolvimento.
   - Benefício: Identificação precoce de problemas específicos de plataforma.
   - Exemplo de script de teste:
     ```bash
     #!/bin/bash
     echo "Compilando e testando para múltiplas plataformas"

     platforms=("windows/amd64" "darwin/amd64" "linux/amd64")

     for platform in "${platforms[@]}"
     do
         platform_split=(${platform//\// })
         GOOS=${platform_split[0]}
         GOARCH=${platform_split[1]}
         output_name='myapp-'$GOOS'-'$GOARCH
         if [ $GOOS = "windows" ]; then
             output_name+='.exe'
         fi

         env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
         if [ $? -ne 0 ]; then
             echo 'Falha na compilação para '${GOOS}' '${GOARCH}
             exit 1
         fi
     done

     echo "Compilação concluída para todas as plataformas"
     ```

4. **Implantação em Sistemas Embarcados**
   - Cenário: Desenvolvimento de software para dispositivos IoT com recursos limitados.
   - Uso: Compile Go para arquiteturas ARM ou MIPS com flags de otimização.
   - Benefício: Software eficiente e de baixo consumo para dispositivos com recursos limitados.
   - Exemplo:
     ```
     GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o iot-app
     ```

5. **Distribuição de Plugins ou Extensões**
   - Cenário: Você está criando um sistema que suporta plugins desenvolvidos pela comunidade.
   - Uso: Os desenvolvedores podem compilar seus plugins em Go para diferentes plataformas.
   - Benefício: Fácil distribuição e integração de plugins em várias plataformas.
   - Exemplo (para um desenvolvedor de plugin):
     ```
     go build -buildmode=plugin -o myplugin.so myplugin.go
     ```

Estes casos de uso demonstram como a compilação em Go pode ser aproveitada em diversos cenários, desde o desenvolvimento de aplicativos desktop até sistemas embarcados e arquiteturas de microserviços, oferecendo flexibilidade, eficiência e facilidade de distribuição.

## Vantagens da Compilação em Go vs C

Para muitos casos de uso:
- O binário Go é autocontido, incluindo a maioria das dependências necessárias
- Código compilado em C frequentemente requer que bibliotecas dinâmicas (.so, .dll) estejam presentes no sistema alvo

## Flags de Compilação

- Permitem customizar o binário (executável) gerado
- Podem ajudar a reduzir o tamanho do executável
- Comumente usadas em Dockerfiles para gerar binários menores para ambientes de produção

Exemplo de flag comum:
```
go build -ldflags="-s -w"
```
Esta flag remove informações de debug e tabela de símbolos, reduzindo o tamanho do binário.

## Conclusão

A compilação em Go, especialmente com suas capacidades de compilação cruzada, é uma das características que tornam Go uma excelente escolha para desenvolvimento de software multiplataforma. Ela simplifica significativamente o processo de criação e distribuição de aplicativos para diferentes ambientes, oferecendo flexibilidade e eficiência no desenvolvimento de software.