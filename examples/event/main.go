package main

import (
	"context"
	"fmt"

	"github.com/xmcontinue/sui-go-sdk/constant"
	"github.com/xmcontinue/sui-go-sdk/models"
	"github.com/xmcontinue/sui-go-sdk/sui"
	"github.com/xmcontinue/sui-go-sdk/utils"
)

var ctx = context.Background()
var cli = sui.NewSuiClient(constant.BvMainnetEndpoint)

func main() {
	SuiGetEvents()
	SuiXQueryEvents()
}

func SuiGetEvents() {
	rsp, err := cli.SuiGetEvents(ctx, models.SuiGetEventsRequest{
		Digest: "AxwPoyvpPRcfyuURg6vuc2wDgrMh8BJkT1rAtc1dfU5p",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiXQueryEvents() {
	rsp, err := cli.SuiXQueryEvents(ctx, models.SuiXQueryEventsRequest{
		SuiEventFilter: models.EventFilterByMoveEventType{
			MoveEventType: "0x3::validator::StakingRequestEvent",
		},
		Limit:           5,
		DescendingOrder: false,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
