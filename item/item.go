package item

import "github.com/Kurtz1993/gql-example/gql/model"

var items = map[string][]*model.Item{
	"1": {
		{
			ID:   "1",
			Name: "Sword",
		},
		{
			ID:   "2",
			Name: "Katana",
		},
		{
			ID:   "3",
			Name: "Bow",
		},
	},
}

func QueryItem(playerId string) []*model.Item {
	return items[playerId]
}
