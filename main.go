package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// Define el esquema con un campo "hello"
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String, // El tipo de datos que se devolverá es String
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// El resolver devuelve el mensaje "Hello, world!"
				return "Hello, world!", nil
			},
		},
	}

	// Crea el objeto de la consulta raíz (RootQuery)
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	// Configura el esquema
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Error al crear el esquema: %v", err)
	}

	// Crea un manejador GraphQL en la raíz ("/")
	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true, // Imprime respuestas en formato bonito
		GraphiQL: true, // Habilita GraphiQL para consultas interactivas
	})

	// Configura la ruta raíz ("/") para manejar las consultas GraphQL
	http.Handle("/", gqlHandler)

	// Inicia el servidor HTTP en el puerto 80
	log.Println("Servidor GraphQL en ejecución en http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
