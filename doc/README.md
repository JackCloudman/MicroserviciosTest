## Introducción
Hoy en día existen distintas formas de crear microservicios, en este repo se explicara como se hizo el microservicio usando Golang y Docker.

## Estructura
Para empezar el endpoint que se creará simulara una persona haciendo una orden para comprar por ejemplo, un rollo de sushi.

La estructura de nuestra API se puede representar usando la notación RAML.

    #%RAML 1.0
	title: "Microservicio"
	version: 1.0

	/api/sps/helloworld/v1:
	  description: Endpoint que recibe y devuelve las entradas
	  /ordenar:
	    post:
	      description: |
	        Endpoint para ordenar una comida
	      body:
	        application/json:
	          example: |
	            {"id_cliente": 1,"id_comida": 2,"time": "2018-09-22T12:42:31Z"}
	      responses:
	        200:
	          body:
	            application/json:
	              example: |
	                {"message": "Orden realizada con exito!"}
	        404:
	          body:
	            application/json:
	              example: |
	                {"message": "No se ha encontrado el usuario"}
Podemos ver que nuestra ruta principal es /api/sps/helloworld/v1 pero la que hara la orden se encuentra en /api/sps/helloworld/v1/ordenar. Recibiremos un json el cual contendra la siguiente estructura:

    {
	    "id_cliente": 1234
	    "id_comida": 1234
	    "time": "2018-09-22T12:42:31Z"
	}
Como salida podemos obtener un status 200,400 o 404 y además un mensaje alusivo, por ejemplo, en status 200:

    {"message": "Orden realizada con exito!"}

   En el caso de status 400

    {"message": "Error en la estructura"}

   En el caso de status 404

    {"message": "No se ha encontrado el usuario"}

## Ejemplos con cURL
Ordenar un platillo

     $ curl -d '{"id_cliente": 1,"id_comida": 2,"time": "2018-09-22T12:42:31Z"}' -H "Content-Type: application/json" -X POST http://localhost:8090/api/sps/helloworld/v1/ordenar
     # Respuesta
     {"message":"Orden realizada con exito!"}
Usuario o platillo no encontrado

     $ curl -d '{"id_cliente": 3,"id_comida": 2,"time": "2018-09-22T12:42:31Z"}' -H "Content-Type: application/json" -X POST http://localhost:8090/api/sps/helloworld/v1/ordenar
    # Respuesta
    {"message":"Usuario no encontrado"}

## Extras
Puedes clonar el repo y ver el archivo autogenarado (DocumentacionAPI.html) a partir de la especificación RAML.
