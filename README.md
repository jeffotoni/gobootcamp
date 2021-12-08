## Go BootCamp

Este material tem o objetivo de contribuir para o aprendizado da linguagem de programação Go😍 colaborando com a comunidade como uma fonte de aprendizado na linguagem Go. Um material que aborda o que precisamos saber para iniciarmos a programação na linguagem Go 😍.

O conteúdo visa o nível básico do aluno muitos exemplos práticos foram feitos com riqueza de detalhes para tornar a vida mais fácil de quem está iniciando na linguagem.

Se você conhece pouco e quase nada a programação não será problema, todo o manual foi feito para o nível inicial ao avançado.

Espero que gostem e possam servir de base para aprendizado e ajudar vários Gophers possíveis.

Existem milhares de referências hoje em relação ao Golang, vamos começar do início e não podíamos deixar de falar do [Golang Tour](https://go.dev/tour/welcome/1), [Play Golang](https://go.dev/play) ou [Play Go Space](https://goplay.space/) são formas online de brincarmos com a linguagem Go, lindo não é ? 😊

Criamos esta página para ajuda-lo a encontrar com mais facilidade alguns links que acreditamos ser essenciais para o aprendizado da linguagem Go: [referencias Go](https://github.com/jeffotoni/gobootcamp/tree/main/references).

Todo manual foi baseado nestas referências apresetados acima, ele encontra-se aqui [gobootcamp](https://gobootcamp.jeffotoni.com/).

## Instalar com Docker

Você pode instalar gobootcampmanual com docker.

```bash
$  docker run --rm --name gobootcampmanual -it \
-p 8080:8080 jeffotoni/gobootcampmanual:latest
```

## Instalar o manual localmente

Para instalar o manual e executar localmente basta rodar o script abaixo. Seu $GOPATH tem que está configurado.

_Observação_ 
_É necessário $GOPATH está configurado._

```bash
$ sh -c "$(wget https://raw.githubusercontent.com/jeffotoni/gobootcamp/main/install/v1/install.sh -O -)"
```

## Executar o manual localmente

Você poderá instalar o manual em sua máquina local, vamos fazer clone do projeto e executá-lo localmente. 

_Observação_ 
_É necessário o Go instalado na máquina._

```bash
$ git clone https://github.com/jeffotoni/gobootcamp
$ cd gomanual
$ go run main.go
Run Server: http://localhost:8181
```
Agora basta acessar o link para acessar o manual localmente, desta forma você consegue alterar o manual seja para colaborar enviando um PR (Pull request) com melhorias ou novos temas como fazer um fork para seu uso pessoal 😊.

- ![gobootcamp](img/gobootcamp1.jpg?raw=true "gobootcamp")

