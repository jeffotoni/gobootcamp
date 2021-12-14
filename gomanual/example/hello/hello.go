
//// PRIMEIRO MOMENTO 
/// PACKAGE
package main

/// SEGUNDO MOMENTO
/// IMPORT

import "net/http"
import  "github.com/jeffotoni/gconcat" 

func main() {
    
    conc := gconcat.ConcatStr("jeffotoni", "golang","devops")
    conc2 := "jeffotoni"+ "golang" + "devops"
    println(conc)
    println (conc2)

    mux := http.NewServeMux()
    http.ListenAndServe(":8080", mux)
}
