package controllers

import (
    "cars_notify/connection"
    "cars_notify/models"
    "encoding/json"
    "log"

    "github.com/gorilla/websocket"
)

func SetupConsumer(ws *websocket.Conn) {
    rabbit := connection.NewRabbitMQ() 
    // Tratamiento de un mensaje
  msgs := rabbit.GetMessages()

  go func() {
    for d := range msgs {
        var status models.Message
        err := json.Unmarshal(d.Body, &status)
        if err != nil {
            log.Printf("Error al decodificar el mensaje: %s", err)
            continue
        }
        payload, _ := json.Marshal(status)

        if err := ws.WriteMessage(websocket.TextMessage, payload); err != nil {
          log.Println("Error enviando mensaje:", err)
          ws.Close()
          break
        }

        log.Printf(" [x] Recibido: Accion sobre la reta: %s", status.Action)
    }
}()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  select {}
}