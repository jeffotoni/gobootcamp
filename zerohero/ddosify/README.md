#ddosify 

Esta ferramenta faz um test de stress em nossa api. Pode [entrar aqui](https://github.com/ddosify/ddosify) para todos os detalhes inclusive a instalação.

Eu criei um arquivo de configuração para testar nossa API ZeroHero.

Você pode executar pelo config que criei ou executar manualmente.

### Usuando Config json
```bash
$ ddosify -config ddosify/config.config.json
```

### Run ddosify
```bash
$ ddosify -n 10000 -m GET -l linear -h 'Content-Type: application/json' -d 20 -t http://localhost:8080/api/hulk
```
saída:
```bash
  Initializing...
🔥 Engine fired.

🛑 CTRL+C to gracefully stop.
✔️  Successful Run: 350    100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00779s
✔️  Successful Run: 725    100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00838s
✔️  Successful Run: 1100   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00840s
✔️  Successful Run: 1475   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00843s
✔️  Successful Run: 1850   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00855s
✔️  Successful Run: 2225   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00921s
✔️  Successful Run: 2600   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00961s
✔️  Successful Run: 2975   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00950s
✔️  Successful Run: 3350   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00962s
✔️  Successful Run: 3725   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00953s
✔️  Successful Run: 4100   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00946s
✔️  Successful Run: 4475   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00900s
✔️  Successful Run: 4850   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00876s
✔️  Successful Run: 5000   100%       ❌ Failed Run: 0        0%       ⏱️  Avg. Duration: 0.00873s

RESULT
-------------------------------------
Step 1
-------------------------------------
Success Count:    5000  (100%)
Failed Count:     0     (0%)

Durations (Avg):
  DNS                  :0.0003s
  Connection           :0.0003s
  Request Write        :0.0002s
  Server Processing    :0.0039s
  Response Read        :0.0002s
  Total                :0.0048s

Status Code (Message) :Count
  201 (Created)    :5000

Step 2
-------------------------------------
Success Count:    5000  (100%)
Failed Count:     0     (0%)

Durations (Avg):
  DNS                  :0.0001s
  Connection           :0.0003s
  Request Write        :0.0001s
  Server Processing    :0.0033s
  Response Read        :0.0001s
  Total                :0.0039s

Status Code (Message) :Count
  200 (OK)    :5000

```

