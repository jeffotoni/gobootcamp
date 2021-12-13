# API zerohero 🐿️ 

Um simples exemplo de uma API zerohero inspirada na api [superheroapi](https://superheroapi.com/).
Este serviço está bem no inicio e será usado no Bootcamp e Hand-on no DevOpsFest.

Fique a vontade em comentar e colaborar abra um [issue](https://github.com/jeffotoni/gzerohero/issues) e deixe seus comentários.

ZeroHero é uma api que carrega uma base com os supereróis e permite pesquisas por biography, powerstats, connections , image, work e appearance.

legal não é 😍?. O objetivo é entender a construção de uma API Go usando somente a strand library. 

Na pasta json temos alguns supererois para que possamos brincar em nossa base.

Exemplo de Json do Odin, eles são encontrados no diretório "json/"
```bash
{
   "response":"success",
   "id":"498",
   "name":"Odin",
   "powerstats":{
      "intelligence":"100",
      "strength":"83",
      "speed":"67",
      "durability":"95",
      "power":"100",
      "combat":"90"
   },
   "biography":{
      "full-name":"Odin Borson",
      "alter-egos":"No alter egos found.",
      "aliases":[
         "All-Father,Sky-Father",
         "Atum-Re",
         "Woden",
         "Wotan",
         "Oden",
         "Orrin",
         "Harbard",
         "King of Asgard"
      ],
      "place-of-birth":"Asgard",
      "first-appearance":"Journey into Mystery #85",
      "publisher":"Marvel Comics",
      "alignment":"good"
   },
   "appearance":{
      "gender":"Male",
      "race":"God / Eternal",
      "height":[
         "6'9",
         "206 cm"
      ],
      "weight":[
         "650 lb",
         "293 kg"
      ],
      "eye-color":"Blue",
      "hair-color":"White"
   },
   "work":{
      "occupation":"Deity, Monarch of Asgard, Asgardian God of the Sky, Wind, Wisdom, Crafts, Time, and the Dead, Warrior",
      "base":"City of Asgard, Asgard"
   },
   "connections":{
      "group-affiliation":"Asgardians, Council of Godheads",
      "relatives":"Buri (Tiwaz) (paternal grandfather), Bolthorn (maternal grandfather), Bor Burison (father, deceased), Bestla (mother), Mimir Burison (paternal uncle), Njord (paternal uncle), Vili, Ve, Cul (brothers), Frigga (wife), Freyr (father-in-law), Gullveig (sister-in-law), Thor (son by Jord), Vidar (son by Grid), Balder (son by Frigga), Tyr, Hermod (allegedly sons by Frigga), Angela (daughter by Frigga), Laussa (daughter by Frigga and Surtur), Loki (foster son), Hoder (nephew), Skadi (niece)"
   },
   "image":{
      "url":"https://www.superherodb.com/pictures2/portraits/10/100/10388.jpg"
   }
}

```

#### Docker Compose

Podemos usar o docker-compose para subir o mongodb e a api zerohero, e também pdoerá rodar local.
Geralmente comento as linhas do serviço zerohero para debugar e testar tudo antes localmente o serviço e deixe o mongo no docker-compose.

```bash
$ docker-compose up -d 
```

#### Go run

Faça clone do zerohero e rode localmente a api

```bash
$ git clone https://github.com/jeffotoni/gzerohero.git
$ cd gzerohero
$ go run .
```
saída:
```bash
2021/12/13 00:45:35 Running on http://0.0.0.0:8080 (Press CTRL+C to quit)
```

Agora vamos testar nossa API 🦾

#### GET ping

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/ping

```

saída:
```bash
HTTP/1.1 200 OK
Date: Mon, 13 Dec 2021 03:57:20 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8

pong😍

```

#### POST api

```bash
$ curl -i -XPOST \
-H "Content-Type:application/json" \
http://localhost:8080/api -d @json/hulk.json

```
saída:
```bash
$ HTTP/2 201
```

#### PUT api/<heroi>

```bash
$ curl -i -XPUT \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk -d @json/hulk.json

```
saída:
```bash
$ HTTP/2 200
```

#### DELETE api/<heroi>

```bash
$ curl -i -XDELETE \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk

```
saída:
```bash
$ HTTP/2 204
```

#### GET api

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:48:39 GMT
Content-Length: 1396

{"response":"success","id":"332","uuid":"06508da1-d3a3-40f4-9db2-64b9e66df3de","name":"hulk","powerstats":{"intelligence":"88","strength":"100","speed":"63","durability":"100","power":"98","combat":"85"},"biography":{"full-name":"Bruce Banner","alter-egos":"No alter egos found.","aliases":["Annihilator","Captain Universe","Joe Fixit","Mr. Fixit","Mechano","Professor","Jade Jaws","Golly Green Giant"],"place-of-birth":"Dayton, Ohio","first-appearance":"Incredible Hulk #1 (May, 1962)","publisher":"Marvel Comics","alignment":"good"},"appearance":{"gender":"Male","race":"Human / Radiation","height":["8'0","244 cm"],"weight":["1400 lb","630 kg"],"eye-color":"Green","hair-color":"Green"},"work":{"occupation":"Nuclear physicist, Agent of S.H.I.E.L.D.","base":"(Banner) Hulkbuster Base, New Mexico, (Hulk) mobile, but prefers New Mexico"},"connections":{"group-affiliation":"Defenders, former leader of the new Hulkbusters, member of the Avengers, Pantheon, Titans Three, the Order, Hulkbusters of Counter-Earth-Franklin, alternate Fantastic Four","relatives":"Betty Ross Talbot Banner (wife), Brian Banner (father, apparently deceased), Rebecca Banner (mother, deceased), Morris Walters (uncle), Elaine Banner Walters (aunt, deceased), Jennifer Walters (She-Hulk, cousin), Thaddeus E. 'Thunderbolt' Ross (father"},"image":{"url":"https://www.superherodb.com/pictures2/portraits/10/100/83.jpg"}}
```

#### GET api/<heroi>/<caracteristica>

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk/work
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:50:34 GMT
Content-Length: 142

{"occupation":"Nuclear physicist, Agent of S.H.I.E.L.D.","base":"(Banner) Hulkbuster Base, New Mexico, (Hulk) mobile, but prefers New Mexico"}
```


```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk/image
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:50:34 GMT
Content-Length: 142

{"url":"https://www.superherodb.com/pictures2/portraits/10/100/83.jpg"}
```

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk/powerstats
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:50:34 GMT
Content-Length: 142

{"intelligence":"88","strength":"100","speed":"63","durability":"100","power":"98","combat":"85"}
```

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk/biography
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:50:34 GMT
Content-Length: 142

{"full-name":"Bruce Banner","alter-egos":"No alter egos found.","aliases":["Annihilator","Captain Universe","Joe Fixit","Mr. Fixit","Mechano","Professor","Jade Jaws","Golly Green Giant"],"place-of-birth":"Dayton, Ohio","first-appearance":"Incredible Hulk #1 (May, 1962)","publisher":"Marvel Comics","alignment":"good"}
```

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk/appearance
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:50:34 GMT
Content-Length: 142

{"gender":"Male","race":"Human / Radiation","height":["8'0","244 cm"],"weight":["1400 lb","630 kg"],"eye-color":"Green","hair-color":"Green"}
```

```bash
$ curl -i -XGET \
-H "Content-Type:application/json" \
http://localhost:8080/api/hulk/Connections
```

saída:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 03:50:34 GMT
Content-Length: 142

{"group-affiliation":"Defenders, former leader of the new Hulkbusters, member of the Avengers, Pantheon, Titans Three, the Order, Hulkbusters of Counter-Earth-Franklin, alternate Fantastic Four","relatives":"Betty Ross Talbot Banner (wife), Brian Banner (father, apparently deceased), Rebecca Banner (mother, deceased), Morris Walters (uncle), Elaine Banner Walters (aunt, deceased), Jennifer Walters (She-Hulk, cousin), Thaddeus E. 'Thunderbolt' Ross (father"}
```