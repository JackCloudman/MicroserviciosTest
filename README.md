# MicroserviciosTest
Prueba de microservicios usando docker y golang
## Requisitos
- Golang >=1.13
- Docker (opcional)
- Mux == v1.7.3
## Instalación (normal)
Clona el repo

    $ git clone https://github.com/JackCloudman/MicroserviciosTest.git
    $ cd MicroserviciosTest
Instala requerimientos

    $ go get github.com/gorilla/mux v1.7.3
### Ejecutar

    $ go run main.go
  Si lo prefieres puedes compilarlo de la siguiente forma:


    $ go build
    $ ./MicroserviciosTest #Ejecutamos el microservicio
## Instalación (Docker)
    $ git clone https://github.com/JackCloudman/MicroserviciosTest.git
    $ cd MicroserviciosTest
    $ docker build -t jackcloudman/microserviciotest .
### Ejecutar
    $ docker run --publish 8090:8090 jackcloudman/microserviciotest
## Notas
- Se puede compilar el proyecto y crear una imagen de docker que solo contenga este archivo binario.
- Hay mas información en la carpeta doc
- La estructura del proyecto se puede mejorar
