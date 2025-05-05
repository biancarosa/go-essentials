# Básico de Git

Git é um sistema de controle de versão distribuído que permite rastrear mudanças em arquivos ao longo do tempo. É essencial para desenvolvimento de software colaborativo.

## Por que usar Git?

- **Histórico completo**: Mantém registro de todas as alterações feitas no código
- **Trabalho colaborativo**: Permite que múltiplas pessoas trabalhem no mesmo projeto simultaneamente
- **Gerenciamento de versões**: Facilita a criação e manutenção de diferentes versões do software
- **Backup**: Fornece uma cópia de segurança do código em repositórios remotos

## O Git na prática de desenvolvimento

O Git se tornou uma ferramenta essencial na carreira de qualquer desenvolvedor moderno. Embora não seja a única forma de versionamento de código, é certamente a mais popular e amplamente adotada.

Pense no Git como o "Dropbox ou Google Drive do código", mas com recursos avançados para controle de versão e colaboração. Ele permite que equipes trabalhem na mesma base de código sem precisar trocar arquivos por e-mail, mensagens ou outros meios menos eficientes.

### Git vs. Plataformas de hospedagem

É importante entender a diferença entre o Git (a ferramenta de controle de versão) e as plataformas que o utilizam:

- **Git**: O sistema de controle de versão em si
- **GitHub**: Uma plataforma web que utiliza Git e adiciona recursos sociais e de colaboração
- **GitLab**: Outra plataforma similar ao GitHub
- **Bitbucket**: Mais uma alternativa para hospedar repositórios Git

Todas essas plataformas são construídas sobre o protocolo Git, mas oferecem interfaces e recursos adicionais distintos.

### Abordagens de uso do Git

Existem diferentes níveis de uso do Git:
1. **Via linha de comando**: Usando diretamente os comandos Git no terminal
2. **Via interfaces gráficas**: Usando programas que facilitam o uso do Git com recursos visuais
3. **Abordagem híbrida**: Combinando comandos de terminal com interfaces visuais

Neste guia, adotaremos uma abordagem híbrida, que oferece o melhor custo-benefício entre o que você precisa aprender (comandos) e a visualização das mudanças no seu código.

## Links importantes

- [Sincronizando a sua conta do Github com o VS Code](https://code.visualstudio.com/docs/editor/settings-sync)
- [Crie uma conta no Github (se não possuir)](https://docs.github.com/en/get-started/start-your-journey/creating-an-account-on-github)

## Configuração Inicial

Antes de começar a usar o Git, configure seu nome de usuário e email:

```bash
git config --global user.name "Seu Nome"
git config --global user.email "seu.email@exemplo.com"
```

## Comandos Básicos

### Inicializando um Repositório

```bash
git init
```

Cria um novo repositório Git vazio no diretório atual.

### Verificando o Status

```bash
git status
```

Mostra o estado atual do repositório, incluindo arquivos modificados, adicionados ou removidos.

### Adicionando Arquivos ao Stage

```bash
git add arquivo.go     # Adiciona um arquivo específico
git add .              # Adiciona todos os arquivos modificados
```

Prepara os arquivos para serem commitados.

### Commitando Alterações

```bash
git commit -m "Mensagem descrevendo as alterações"
```

Salva as alterações no histórico do repositório com uma mensagem descritiva.

### Visualizando o Histórico

```bash
git log                # Histórico completo
git log --oneline      # Histórico resumido, uma linha por commit
```

Mostra o histórico de commits do repositório.

## Trabalhando com Repositórios Remotos

### Conectando a um Repositório Remoto

```bash
git remote add origin https://github.com/usuario/repositorio.git
```

Associa seu repositório local a um repositório remoto.

### Enviando Alterações para o Repositório Remoto

```bash
git push -u origin main    # Na primeira vez
git push                   # Nas próximas vezes
```

Envia seus commits locais para o repositório remoto.

### Recebendo Alterações do Repositório Remoto

```bash
git pull
```

Baixa e integra as alterações do repositório remoto.

### Clonando um Repositório Existente

```bash
git clone https://github.com/usuario/repositorio.git
```

Cria uma cópia local de um repositório remoto existente.

## Branches

### Criando uma Nova Branch

```bash
git branch nome-da-branch          # Apenas cria a branch
git checkout -b nome-da-branch     # Cria e muda para a nova branch
```

Branches permitem trabalhar em funcionalidades isoladamente.

### Mudando entre Branches

```bash
git checkout nome-da-branch
```

Alterna para uma branch específica.

### Mesclando Branches

```bash
git checkout main          # Primeiro, mude para a branch de destino
git merge nome-da-branch   # Mescla a branch especificada na branch atual
```

Combina as alterações de uma branch com outra.

## Ferramentas Visuais para Git

Além da linha de comando, você pode usar ferramentas visuais para facilitar o uso do Git:

- **GitHub Desktop**: Interface gráfica simples para GitHub
- **GitKraken**: Cliente Git com interface visual poderosa
- **VSCode**: O próprio VSCode tem integração com Git

## Boas Práticas

1. **Commits pequenos e frequentes**: Faça commits de alterações relacionadas e com mensagens claras
2. **Pull antes de Push**: Sempre atualize seu repositório local antes de enviar alterações
3. **Branches para funcionalidades**: Crie branches separadas para cada funcionalidade ou correção
4. **Revisão de código**: Utilize pull requests para revisar código antes de mesclar

## Escolhendo seu fluxo de trabalho

Como em qualquer ferramenta, é importante que você encontre seu próprio fluxo de trabalho com o Git. Absorva as informações de diferentes fontes, teste diferentes abordagens e escolha o que funciona melhor para você e sua equipe.

Para iniciantes, recomendamos a abordagem híbrida: aprender os comandos básicos do Git para entender os conceitos, mas utilizar interfaces visuais quando elas facilitarem o trabalho.

## Recursos para Aprender Mais

- [Git Documentation](https://git-scm.com/doc)
- [GitHub Guides](https://guides.github.com/)
- [Learn Git Branching](https://learngitbranching.js.org/)
- [Pro Git Book](https://git-scm.com/book/en/v2) 