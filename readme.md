# PixelPress

Uma ferramenta de linha de comando (CLI) com uma interface interativa para processar imagens em massa. Com ela, você pode facilmente converter formatos ou comprimir múltiplas imagens para um tamanho de arquivo específico diretamente do seu terminal.

## Visualização

https://github.com/user-attachments/assets/f3c457a3-4d25-4dbb-afbe-6a869526f629

## Recursos Principais

-   **Detecção Automática de Imagens**: Encontra todos os arquivos `.jpg`, `.jpeg`, `.png` e `.webp` no diretório atual.
-   **Seleção Múltipla**: Use a barra de espaço para selecionar quantos arquivos desejar processar de uma vez.
-   **Conversão de Formato**: Converta os arquivos selecionados para `jpg`, `png` ou `webp`. Os novos arquivos são salvos com o sufixo `_converted`.
-   **Compressão Inteligente**: Defina um tamanho máximo em Megabytes (MB) e a ferramenta tentará comprimir as imagens para que fiquem abaixo desse limite.
    Os arquivos comprimidos são salvos com o prefixo `compressed_`.

## Pré-requisitos

Para que a ferramenta funcione corretamente, você precisa ter instalados em seu sistema:

**FFmpeg**: A dependência mais importante, usada para todo o processamento de imagem. - Você precisa instalar o FFmpeg e garantir que o executável esteja acessível através do `PATH` do seu sistema. Você pode verificar se ele está instalado corretamente abrindo um terminal e digitando `ffmpeg -version`.

## Instalação

### Opção 1: Baixar o Binário Pré-compilado (Recomendado)

1.  Acesse a [página de **Releases** do projeto no GitHub](https://github.com/AndreXime/pixel-press/releases).
2.  Procure pela versão mais recente e baixe o arquivo compatível com o seu sistema operacional.
3.  **(Para macOS e Linux)** Abra o terminal, navegue até a pasta onde está o arquivo e dê permissão de execução para o binário:
    ```bash
    chmod +x <nome-do-executavel>
    ```
4.  Pronto\! Você já pode mover o executável para um diretório com imagens ou configure o executavel para ser global no sistema adicionando no `PATH` _(Linux/macOS)_.

### Opção 2: Compilar do Código-Fonte

Se você tem o **Go** instalado e prefere compilar o projeto manualmente:

1.  Clone o repositório:

    ```bash
    git clone https://github.com/AndreXime/pixel-press.git
    cd pixel-press
    ```

2.  Compile o projeto. O Go irá baixar as dependências e criar o arquivo executável:

    ```bash
    go build .
    ```
