package main

type abstractRoutes interface {
	Create()
}

type Kong struct {
	host   string
	user   string
	pass   string
	routes abstractRoutes
}

type Routes struct {
	name string
}

func New(n string) Routes {
	return Routes{
		name: n,
	}
}

func (r Routes) Create() {
	println("create route kong")
}

func main() {
	var k Kong
	// k.routes = Routes{name: "route-kong1"}
	k.routes = New("route-kong1")

	k.host = "localhost"
	k.user = "guest"
	k.pass = "12345"
	k.routes.Create()

	var i interface{}

	i = "jeffotoni"
	println(i.(string))
	i = 10
	println(i.(int))

	switch i.(type) {
	case int:
		println("ok int")
	}
}
