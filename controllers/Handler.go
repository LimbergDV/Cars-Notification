package controllers

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

type WsHandler struct {
    upgrader websocket.Upgrader
}

func NewWsHandler(upgrader websocket.Upgrader) *WsHandler{
    return &WsHandler{upgrader: upgrader}
}

// Handler para el endpoint WebSocket
func (wsh *WsHandler) WsHandler(w http.ResponseWriter, r *http.Request, ) {
    // Realizamos el upgrade de la conexión HTTP a WebSocket.
    ws, err := wsh.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error en Upgrade:", err)
        return
    }
    // Asegurarse de cerrar la conexión cuando termine.
    defer ws.Close()

    log.Println("Cliente conectado")

    go SetupConsumer(ws)

    select {}
}