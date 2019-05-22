package manipulador

import "html/template"

//ModeloOla armazena o modelo de pagina ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModelsLocal armazena o modelo de pagina local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
