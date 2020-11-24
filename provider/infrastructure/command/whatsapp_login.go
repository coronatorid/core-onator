package command

import (
	"fmt"
	"os"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
)

// WhatsappLogin return cli to get whatsapp session
type WhatsappLogin struct {
	wac *whatsapp.Conn
}

// NewWhatsappLogin return cli to get whatsapp session
func NewWhatsappLogin(wac *whatsapp.Conn) *WhatsappLogin {
	return &WhatsappLogin{wac: wac}
}

// Use of the command
func (w *WhatsappLogin) Use() string {
	return "whatsapp:login"
}

// Example of the command
func (w *WhatsappLogin) Example() string {
	return "whatsapp:login"
}

// Short description about the command
func (w *WhatsappLogin) Short() string {
	return "Get whatsapp session for enviroment variables"
}

// Run the command with the args given by the caller
func (w *WhatsappLogin) Run(args []string) {
	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	session, err := w.wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
		return
	}

	fmt.Println("login successful!")
	fmt.Println("client_id:", session.ClientId)
	fmt.Println("client_token:", session.ClientToken)
	fmt.Println("server_token:", session.ServerToken)
	fmt.Println("enc_key:", fmt.Sprintf("%x", session.EncKey))
	fmt.Println("mac_key:", fmt.Sprintf("%x", session.MacKey))
	fmt.Println("wid:", session.Wid)
}
