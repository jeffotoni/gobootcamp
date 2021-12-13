# k6.io

k6 é uma ferramenta de teste de carga de código aberto desenvolvida pela linguagem Go 😍. O k6 vem com recursos, sobre os quais você pode aprender tudo na documentação. Os principais recursos incluem:

- Ferramenta CLI com APIs amigáveis ​​ao desenvolvedor.
- Scripting em JavaScript ES2015 / ES6 - com suporte para módulos locais e remotos
- Verificações e limites - para teste de carga orientado a metas

O k6 criou sua própria lib javascript para comportar como nodejs, então ao construir os scripts irá usar a linguagem javascript porém com libs disponibilizada pela k6.io.

Posso usar npm e suas libs para criação dos meus scripts ?
Sim pode, importar módulos npm ou bibliotecas, você pode [agrupar módulos npm com webpack](https://k6.io/docs/using-k6/modules/#bundling-node-modules) e importá-los em seus testes.

### Github k6.io
[github k6.io](https://github.com/k6io/k6)

### Instalar k6.io

Existe várias formas de instalação, e por ser feito em Go tudo fica mais fácil basta instalar seu binário em sua máquina.

Aqui está o link com todas possibilidades de instalação:
[Install k6.io](https://k6.io/docs/getting-started/installation/)

### Instalar com docker

Vamos mostrar a instalação usando Docker desta forma não irá precisar instalar nadinha na sua máquina.

```bash
$ docker pull loadimpact/k6
```

Agora vamos executar nossa massa de testes. O detalhe abaixo é que como está em container a rede que irá executar é outra, então passei parametro para nosso Script para buscar nosso domain ou hosname da api e o volume para carregar nosso json.
```bash
$ docker run -v $(pwd):/data \
-e DOMAIN=http://192.168.0.70:8080 \
-i loadimpact/k6 run - <script.js
```
Nosso script já deixamos pré-prontos, fazendo chamada do POST que envia um json e do Get buscando a informação e do nosso famigerado ping 😍.

#### Executando k6

```bash
$ k6 run myscript.js
```
Ou se estiver usando **Docker:**
```bash
$ docker run -i loadimpact/k6 run - <myscript.js
```
ou

```bash
$  k6 run -e DOMAIN=http://localhost:8080 --vus 100 --duration 20s script.js

```

Saída:
```bash

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: script.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 50s max duration (incl. graceful stop):
           * default: 100 looping VUs for 20s (gracefulStop: 30s)


running (20.1s), 000/100 VUs, 1000 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  20s

     data_received..................: 1.6 MB 81 kB/s
     data_sent......................: 236 kB 12 kB/s
     http_req_blocked...............: avg=915.78µs min=909ns   med=5.2µs    max=28.05ms p(90)=9.32µs   p(95)=188.61µs
     http_req_connecting............: avg=888.55µs min=0s      med=0s       max=27.65ms p(90)=0s       p(95)=2.91µs  
     http_req_duration..............: avg=2.56ms   min=46.94µs med=1.02ms   max=27.01ms p(90)=6.83ms   p(95)=9.41ms  
       { expected_response:true }...: avg=2.56ms   min=46.94µs med=1.02ms   max=27.01ms p(90)=6.83ms   p(95)=9.41ms  
     http_req_failed................: 0.00%  ✓ 0         ✗ 2000 
     http_req_receiving.............: avg=55.03µs  min=6.34µs  med=36.93µs  max=3.45ms  p(90)=92.03µs  p(95)=112.41µs
     http_req_sending...............: avg=118.26µs min=3.6µs   med=26.56µs  max=25.99ms p(90)=180.94µs p(95)=447.84µs
     http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s      p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=2.38ms   min=31.4µs  med=916.17µs max=21.12ms p(90)=6.53ms   p(95)=9.19ms  
     http_reqs......................: 2000   99.477445/s
     iteration_duration.............: avg=2s       min=2s      med=2s       max=2.03s   p(90)=2.02s    p(95)=2.03s   
     iterations.....................: 1000   49.738723/s
     vus............................: 100    min=100     max=100
     vus_max........................: 100    min=100     max=100

```