package eventhandler

type Role int

type QuantizedRole struct {
	role Role
	number uint
}

const (
	PetLivingCleaning Role = iota
	NonMedicalAnimalCare
	MedicalAnimalCare
	EventSupport
	EventSetup
	EventTearDown
)
