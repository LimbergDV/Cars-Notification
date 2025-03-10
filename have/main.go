package main

import (
	"encoding/json"
	"have/entities"
	"have/src"
	"log"

	"github.com/joho/godotenv"
)

func main() {
  // Cargar las variables de entorno
  godotenv.Load()
  rabbit := src.NewRabbitMQ()
  
  // Tratamiento de un mensaje
  msgs := rabbit.GetMessages()
  var forever chan struct{}

  go func() {
    for d := range msgs {
        var notify entities.Notify
        err := json.Unmarshal(d.Body, &notify)
        if err != nil {
            log.Printf("Error al decodificar el mensaje: %s", err)
            continue
        }
        log.Printf(" [x] Renta recibida: id_customer=%d, return_date=%s", notify.Id_customer, notify.Return_date)
        
        if notify.Return_date != "null" {
          src.FetchAPI(notify, "rent")
        } else {
		      src.FetchAPI(notify, "return")
        } 
    }
}()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}

