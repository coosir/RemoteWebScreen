package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kbinani/screenshot"
)

var CaptureScreenquality int = 80

func SimulateDesktopHDMessage(conn *websocket.Conn, msg []byte) {
	var message map[string]interface{}
	if err := json.Unmarshal(msg, &message); err != nil {
		//log.Printf("if err := json.Unmarshal(msg, &message); err != nil Error: %v", err)
		return
	}

	switch messageType := message["type"].(string); messageType {
	case "updateSettings":
		quality := message["quality"].(float64)
		CaptureScreenquality = int(quality)
	case "9":
		currentScreen++
		if currentScreen >= screenshot.NumActiveDisplays() {
			currentScreen = 0
		}
	case "10":
		conn.Close()
	}
}
