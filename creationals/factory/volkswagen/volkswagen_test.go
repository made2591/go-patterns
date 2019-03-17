package volkswagen

import (
	"strings"
	"testing"
)

func TestCreateVolkswagenPolo(t *testing.T) {

	vehicle := createVolkswagenPolo()

	if vehicle == nil {
		t.Errorf("CreateVolkswagenPolo was incorrect, got: nil")
	}

}

func TestCreateVolkswagenGolf(t *testing.T) {

	vehicle := createVolkswagenGolf()

	if vehicle == nil {
		t.Errorf("CreateVolkswagenGolf was incorrect, got: nil")
	}

}

func TestGetModel(t *testing.T) {

	polo := createVolkswagenPolo()
	poloModel := polo.GetModel().(string)

	if strings.Compare(poloModel, PoloModel) != 0 {
		t.Errorf("GetModel was incorrect, got %s, wanted %s", poloModel, PoloModel)
	}

	golf := createVolkswagenGolf()
	golfModel := golf.GetModel().(string)

	if strings.Compare(golfModel, GolfModel) != 0 {
		t.Errorf("GetModel was incorrect, got %s, wanted %s", golfModel, GolfModel)
	}

}

func TestSetModel(t *testing.T) {

	polo := createVolkswagenPolo()
	polo.SetModel("Polo GTI")
	PoloModel := polo.GetModel().(string)

	if strings.Compare(PoloModel, "Polo GTI") != 0 {
		t.Errorf("SetModel was incorrect, got %s, wanted %s", PoloModel, "Polo GTI")
	}

	golf := createVolkswagenGolf()
	golf.SetModel("Golf GTI")
	GolfModel := golf.GetModel().(string)

	if strings.Compare(GolfModel, "Golf GTI") != 0 {
		t.Errorf("SetModel was incorrect, got %s, wanted %s", golf, "Golf GTI")
	}

}
