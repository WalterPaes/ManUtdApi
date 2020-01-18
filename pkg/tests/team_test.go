package tests

import (
	"encoding/json"
	"github.com/WalterPaes/ManUtdApi/pkg/domains/player"
	"github.com/WalterPaes/ManUtdApi/pkg/domains/team"
	"net/http"
	"reflect"
	"testing"
)

type ServiceMock struct{}

func (_ ServiceMock) GetPlayer(_ string) ([]byte, error, int) {
	return []byte(`{"player":[{"strPlayer":"Daniel James","dateBorn":"1997-11-10","strNationality":"United Kingdom","strPosition":"Winger","strNumber":"21","strWage":"£2,184,000 (£42,000 a week)"}]}`), nil, http.StatusOK
}

func (_ ServiceMock) GetAllPlayers() ([]byte, error, int) {
	return []byte(`{"player":[{"strPlayer":"Daniel James","dateBorn":"1997-11-10","strNationality":"United Kingdom","strPosition":"Winger","strNumber":"21","strWage":"£2,184,000 (£42,000 a week)"}]}`), nil, http.StatusOK
}

func TestSearch(t *testing.T) {
	data := team.Team{
		Player: []player.Player{
			{
				StrPlayer:      "Daniel James",
				DateBorn:       "1997-11-10",
				StrNationality: "United Kingdom",
				StrPosition:    "Winger",
				StrNumber:      "21",
				StrWage:        "£2,184,000 (£42,000 a week)",
			},
		},
	}
	svc := &ServiceMock{}

	t.Run("Search One Player", func(t *testing.T) {
		got, _ := team.New(svc).GetPlayer("Daniel James")
		want, _ := json.Marshal(data)

		if !reflect.DeepEqual(got, string(want)) {
			t.Errorf("Want: %v; Got %v", string(want), got)
		}
	})

	t.Run("Search All Players", func(t *testing.T) {
		got, _ := team.New(svc).GetPlayer()
		want, _ := json.Marshal(data)

		if !reflect.DeepEqual(got, string(want)) {
			t.Errorf("Want: %v; Got %v", string(want), got)
		}
	})
}
