# M23 Desafio - Gemini Demo 

![Gemini Banner](https://cdn.prod.website-files.com/630d4d1c4a462569dd189855/6584a9558d6e7530491b6ce1_jn_g9l0wjUIo8WuTO03lpTEkc0kq9qQplZghgziZaUrBxStW8sTXZQ96Lo-dQrya7PW4US6ZMiisw8WZeSUt2XrEkESjtiQCEHsCs5F8UFDwjUrDKy954tcbD9FyugnK8RU4wPo-YAwfpZaGZv1O3ME.jpeg)

![Made with Go](https://img.shields.io/badge/Made%20with-Go-blue)
![Using VS Code](https://img.shields.io/badge/Using-VS%20Code-blue)

Este projeto demonstra como configurar e usar o Google Gemini com Go. Siga as etapas abaixo para clonar o repositório, configurar sua chave de API e executar o projeto.

## Pré-requisitos

- Ter o [Go](https://golang.org/dl/) instalado no seu computador.
- Ter o [Git](https://git-scm.com/downloads) instalado no seu computador.
- Ter uma chave de API do Google Gemini.

## Passo a passo

### 1. Clonar o Repositório

Primeiro, clone este repositório:

```sh
git clone https://github.com/seu-usuario/gemini_project.git
cd gemini_project
```
### 2. Configurar a Chave de API

Obtenha uma chave de API do Google Gemini seguindo as instruções na documentação do Google Gemini.

### 3. Renomear Arquivo de Variáveis de Ambiente

```sh
mv config/.env.sample config/.env.development.local
```

### 4. Adicionar a Chave de API

```sh
GEMINI_API_KEY=your_api_key_here
```

### 5. Adicionar a Chave de API

```sh
cd cmd/app
go run .
```