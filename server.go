package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/albrow/gqlgen-todos/graph"
	"github.com/albrow/gqlgen-todos/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(playgroundPage)
	}))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var playgroundPage = []byte(`
<!DOCTYPE html>
<html>

<head>
  <meta charset=utf-8/>
  <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
  <title>GraphQL Playground</title>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react@1.7.23/build/static/css/index.css" />
  <link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react@1.7.23/build/favicon.png" />
  <script src="//cdn.jsdelivr.net/npm/graphql-playground-react@1.7.23/build/static/js/middleware.js"></script>
</head>

<body>
  <div id="root">
    <style>
      body {
        background-color: rgb(23, 42, 58);
        font-family: Open Sans, sans-serif;
        height: 90vh;
      }

      #root {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
      }

      .loading {
        font-size: 32px;
        font-weight: 200;
        color: rgba(255, 255, 255, .6);
        margin-left: 20px;
      }

      img {
        width: 78px;
        height: 78px;
      }

      .title {
        font-weight: 400;
	  }
	  
	  .graphiql-wrapper>div:first-child>div:first-child>div:nth-child(2) {
		height: 100% !important;
	  }
	  .graphiql-wrapper>div:first-child>div:first-child>div:nth-child(2)>div:nth-child(2) {
		height: 100% !important;
	  }
    </style>
    <img src='//cdn.jsdelivr.net/npm/graphql-playground-react/build/logo.png' alt=''>
    <div class="loading"> Loading
      <span class="title">GraphQL Playground</span>
    </div>
  </div>
  <script>
		window.addEventListener('load', function (event) {
			const root = document.getElementById('root');
			root.classList.add('playgroundIn');
			const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
			GraphQLPlayground.init(root, {
				endpoint: location.protocol + '//' + location.host + '/query',
				subscriptionsEndpoint: wsProto + '//' + location.host + '/query',
			shareEnabled: true,
				settings: {
					'request.credentials': 'same-origin'
				}
			})
		})
  </script>
</body>

</html>
`)
