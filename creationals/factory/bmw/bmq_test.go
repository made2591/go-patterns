package bmw

import (
	"strings"
	"testing"
)

func TestCreateBMWSerieOne(t *testing.T) {

	vehicle := createBMWSerieOne()

	if vehicle == nil {
		t.Errorf("CreateBMWSerieOne was incorrect, got: nil")
	}

}

func TestCreateBMWSerieTwo(t *testing.T) {

	vehicle := createBMWSerieTwo()

	if vehicle == nil {
		t.Errorf("CreateBMWSerieTwo was incorrect, got: nil")
	}

}

func TestGetModel(t *testing.T) {

	serieOne := createBMWSerieOne()
	serieOneModel := serieOne.GetModel().(string)

	if strings.Compare(serieOneModel, SerieOneModel) != 0 {
		t.Errorf("GetModel was incorrect, got %s, wanted %s", serieOneModel, SerieOneModel)
	}

	serieTwo := createBMWSerieTwo()
	serieTwoModel := serieTwo.GetModel().(string)

	if strings.Compare(serieTwoModel, SerieTwoModel) != 0 {
		t.Errorf("GetModel was incorrect, got %s, wanted %s", serieTwoModel, SerieTwoModel)
	}

}

func TestSetModel(t *testing.T) {

	serieOne := createBMWSerieOne()
	serieOne.SetModel("Serie 1 100CV")
	serieOneModel := serieOne.GetModel().(string)

	if strings.Compare(serieOneModel, "Serie 1 100CV") != 0 {
		t.Errorf("SetModel was incorrect, got %s, wanted %s", serieOneModel, "Serie 1 100CV")
	}

	serieTwo := createBMWSerieTwo()
	serieTwo.SetModel("Serie 2 200CV")
	serieTwoModel := serieTwo.GetModel().(string)

	if strings.Compare(serieTwoModel, "Serie 2 200CV") != 0 {
		t.Errorf("SetModel was incorrect, got %s, wanted %s", serieTwo, "Serie 2 200CV")
	}

}
