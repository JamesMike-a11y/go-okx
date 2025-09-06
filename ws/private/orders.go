package private

import (
	"encoding/json"

	"github.com/JamesMike-a11y/go-okx/common"
	"github.com/JamesMike-a11y/go-okx/rest/api/trade"
	"github.com/JamesMike-a11y/go-okx/ws"
)

type HandlerOrders func(EventOrders)

type EventOrders struct {
	Arg  ws.Args `json:"arg"`
	Data []Order `json:"data"`
}

type Order struct {
	trade.Order
}

// default subscribe
func SubscribeOrders(args *ws.Args, auth common.Auth, handler HandlerOrders, handlerError ws.HandlerError) error {
	args.Channel = "orders"

	h := func(message []byte) {
		var event EventOrders
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPrivate(auth).Subscribe(args, h, handlerError)
}
