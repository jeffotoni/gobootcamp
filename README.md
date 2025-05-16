## Go Bootcamp
This material aims to contribute to the learning of the Go programming language😍 by collaborating with the community as a source of learning in the Go language. A material that covers what we need to know to start programming in the Go language 😍.
The content is aimed at the basic level of the student, many practical examples were made with rich details to make life easier for those who are starting in the language.
If you know little or almost nothing about programming, it won't be a problem, the entire manual was made for beginners to advanced levels.
I hope you like it and that it can serve as a basis for learning and help many possible Gophers.

<span style="display: inline-flex; align-items: center;">
  <img src="img/youtube.png?raw=true" alt="youtube" title="youtube" width="5%" style="margin-right: 10px;" />
  Here is the live stream with over 7 hours of content: <a href="https://www.youtube.com/watch?v=XVE3hHW7Wvs">ZeroHero Bootcamp</a> ❤️
</span>

There are thousands of references today regarding Golang, let's start from the beginning and we couldn't fail to mention the [Golang Tour](https://go.dev/tour/welcome/1), [Play Golang](https://go.dev/play) or [Play Go Space](https://goplay.space/) are online ways to play with the Go language, isn't it beautiful? 😊

We created this page to help you find some links that we believe are essential for learning the Go language more easily:
* **[Go references](https://github.com/jeffotoni/gobootcamp/tree/main/references)**

We created a Go roadmap to make it easier to have a macro view when learning Go.
* **[Go roadmap](roadmap/goroadmap.png)**

The entire manual was based on the references presented above, and can be found here:
* **[gobootcamp](https://gobootcamp.jeffotoni.com/)**

We created this manual page
* **[Go manual](https://github.com/jeffotoni/gobootcamp/tree/main/gomanual)**

We created our frontend for our rEST Zerohero API
* **[Zerohero Front](https://zerohero.web.s3apis.com)**

Our Zerohero backend standard library repo
* **[Zerohero backend source](https://github.com/jeffotoni/gzerohero)**

Our Zerohero front repo
* **[Zerohero front source](https://github.com/jeffotoni/gzerohero.web)**
- ![zerohero front](img/zerohero-front.png?raw=true "zerohero front")

## Install with Docker
You can install gobootcampmanual with docker.

```bash
$  docker run --rm --name gobootcampmanual -it \
-p 8080:8080 jeffotoni/gobootcampmanual:latest
```

## Install the manual locally

To install the manual and run it locally, simply run the script below. Your $GOPATH must be configured.

_Note_
_$GOPATH must be configured._

```bash
$ sh -c "$(wget https://raw.githubusercontent.com/jeffotoni/gobootcamp/main/install/v1/install.sh -O -)"
```

## Run the manual locally
You can install the manual on your local machine, we will clone the project and run it locally.

_Note_
_Go must be installed on the machine._

```bash
$ git clone https://github.com/jeffotoni/gobootcamp
$ cd gomanual
$ go run .
Run Server: http://localhost:8181
```
Now just access the link to access the manual locally, this way you can change the manual either to collaborate by sending a PR (Pull request) with improvements or new topics or to make a fork for your personal use 😊.

- ![gobootcamp](img/gobootcamp1.jpg?raw=true "gobootcamp")

