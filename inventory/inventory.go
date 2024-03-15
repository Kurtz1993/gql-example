package inventory

import "github.com/Kurtz1993/gql-example/gql/model"

func QueryInventory(playerId string) *model.Inventory {
	return &model.Inventory{
		PlayerID: playerId,
	}
}
