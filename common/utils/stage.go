package utils

type Stage string

const (
	Local Stage = "local"
	Prod  Stage = "prod"
)

func (s Stage) String() string {
	return string(s)
}

func GetStage() Stage {
	stageString := GetEnvVar("ENV", false)
	switch stageString {
	case Prod.String():
		return Prod
	case Local.String():
		return Local
	default:
		return Local
	}
}
