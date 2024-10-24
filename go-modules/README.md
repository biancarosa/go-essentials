# github.com/biancarosa/go-essentials/go-modules

## Quando precisamos usar?
Precisamos usar Go Modules quando estamos lidando com projetos maiores, compostos por mais de um arquivo Go. Em um projeto simples com apenas um arquivo, o uso de módulos não é estritamente necessário, mas assim que o projeto cresce e começa a envolver múltiplos arquivos, pacotes ou até mesmo dependências de terceiros, o uso de módulos é altamente recomendado. Eles ajudam a organizar e gerenciar o código de forma eficiente, principalmente em projetos maiores e colaborativos.

## Por que usar?

- Publicação de Projetos: Go Modules facilitam a publicação de projetos. Ao criar um projeto modularizado, ele pode ser facilmente importado por outros projetos, funcionando como uma biblioteca reutilizável. Isso é importante quando você quer compartilhar seu código com a comunidade ou com outros projetos internos.

- Compilação Simplificada: Usando módulos, você pode compilar seu projeto de forma simples sem precisar especificar manualmente todos os arquivos Go envolvidos na linha de comando. Por exemplo, ao usar o comando go build, o Go Modules vai identificar automaticamente os arquivos relevantes, em vez de ter que compilar cada arquivo manualmente, como seria necessário em um comando como go build main.go print.go. Isso reduz a complexidade e a possibilidade de erro na hora da compilação.

- Gerenciamento de Dependências: Uma das maiores vantagens dos módulos é o gerenciamento automático de dependências. O Go Modules facilita a inclusão e atualização de bibliotecas externas, garantindo que o seu projeto sempre tenha as versões corretas das dependências, sem necessidade de hacks manuais.

## O que são Go Modules?
Os Go Modules são um agrupamento de pacotes e arquivos Go, que são identificados por um nome (normalmente, o nome do repositório onde o código reside) e suas dependências declaradas em um arquivo go.mod. Eles foram introduzidos na versão 1.11 do Go e representam uma solução oficial para o gerenciamento de dependências e versões.

## Componentes principais

- go.mod: Este arquivo define o nome do módulo e as dependências externas do projeto. Quando você inicializa um módulo usando go mod init, este arquivo é criado automaticamente e nele ficam registradas todas as dependências do projeto, com as versões correspondentes.

- go.sum: O arquivo go.sum armazena checksums (somas de verificação) das dependências, garantindo que as versões exatas sejam baixadas e utilizadas. Ele serve como uma camada extra de segurança para garantir que você não esteja baixando uma versão modificada ou maliciosa de uma dependência.

## Bibliotecas
Quando exportamos módulos para uso por outras pessoas ou aplicações, esses módulos podem ser considerados bibliotecas. Assim como Go tem uma biblioteca padrão (com pacotes como fmt, io, etc.), você pode criar bibliotecas não padrão, que podem ser importadas por outros desenvolvedores.

Essas bibliotecas não padrão são amplamente usadas na comunidade Go e contribuem para a vasta coleção de pacotes disponíveis através de repositórios como o pkg.go.dev.

## Histórico e Evolução
Antes da introdução dos módulos na versão 1.11, o gerenciamento de dependências em Go era feito manualmente através do GOPATH. O comando go get era usado para baixar pacotes diretamente para o GOPATH, mas isso criava diversos problemas, como a falta de controle de versões. Muitas soluções de gerenciamento de dependências foram criadas pela comunidade, como dep, glide, entre outros. No entanto, a partir da versão 1.11, o Go Modules foi introduzido como a solução oficial, e na versão 1.13 ele se tornou o padrão.

Agora, com Go Modules, o gerenciamento de dependências é muito mais eficiente e seguro, além de ser integrado diretamente à linguagem.

## Links úteis

Referência do go.mod
Documentação oficial do Go Modules
Blog post oficial sobre o uso de Go Modules
