package tests

import (
	"encoding/json"
	"github.com/WalterPaes/ManUtdApi/pkg/domains/player"
	"net/http"
	"reflect"
	"testing"
)

type ServiceMock struct {}

func (svc ServiceMock) GetPlayer(_ string) ([]byte, error, int)  {
	return []byte(`{"player":[{"strPlayer":"Daniel James","dateBorn":"1997-11-10","strNationality":"United Kingdom","strPosition":"Winger","strNumber":"21","strWage":"£2,184,000 (£42,000 a week)"}]}`), nil, http.StatusOK
}

func TestSearchPlayer(t *testing.T) {
	data := player.Player{
		Data: player.Data{
			{
				StrPlayer: "Daniel James",
				DateBorn: "1997-11-10",
				StrNationality: "United Kingdom",
				StrPosition: "Winger",
				StrNumber: "21",
				StrWage: "£2,184,000 (£42,000 a week)",
			},
		},
	}
	svc := &ServiceMock{}
		
	got, _ := player.New(svc).Search("Daniel James")
	want, _ := json.Marshal(data)

	if !reflect.DeepEqual(got, string(want)) {
		t.Errorf("Want: %v; Got %v", string(want), got)
	}
}