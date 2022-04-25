package validate

import (
	"github.com/kingofzihua/learn-go/protoc/validate/proto"
	"testing"
)

func TestProtoValidate(t *testing.T) {
	p := new(proto.Person)

	var err error
	err = p.Validate()

	IdError := err.(proto.PersonValidationError)
	t.Error(IdError)
	p.Id = 1000

	err = p.Validate()
	emailError := err.(proto.PersonValidationError)
	t.Error(emailError)

	p.Email = "kingofzihua@outlook.com"
	p.Name = "kingofzihua"
	p.Mobile = "15020866740"
	location := &proto.Person_Location{Lat: 35, Lng: 180}
	err = location.Validate()
	if err != nil {
		t.Error(err)
	} else {
		p.Home = location
	}

	err = p.Validate()
	if err != nil {
		t.Error(err)
	}
}
