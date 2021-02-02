package doggit_coordinates

import "github.com/JustTalDevelops/doggit"

// BuildLibrary is called from Doggit.
func BuildLibrary(_ *doggit.Service) *doggit.Library {
	library := doggit.NewLibrary("Doggit Coordinates", "A coordinates system for Doggit", "JustTal")

	library.AddHandler(DoggitHandler{})

	return library
}