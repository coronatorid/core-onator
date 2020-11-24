package command_test

import (
	"testing"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/coronatorid/core-onator/provider/infrastructure/command"
	"github.com/stretchr/testify/assert"
)

func TestWhatsappLogin(t *testing.T) {

	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		panic(err)
	}

	whatsappLoginCMD := command.NewWhatsappLogin(wac)

	t.Run("Use", func(t *testing.T) {
		assert.Equal(t, "whatsapp:login", whatsappLoginCMD.Use())
	})

	t.Run("Example", func(t *testing.T) {
		assert.Equal(t, "whatsapp:login", whatsappLoginCMD.Example())
	})

	t.Run("Short", func(t *testing.T) {
		assert.Equal(t, "Get whatsapp session for enviroment variables", whatsappLoginCMD.Short())
	})
}
