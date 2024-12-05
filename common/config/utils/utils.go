package utils

import (
	commonutils "colegio/server/common/utils"

	"github.com/pkg/errors"
)

func SwitchOnStage[T any](stage commonutils.Stage,
	prodValueFunc func() T,
	localValueFunc func() T,
) T {
	switch stage {
	case commonutils.Prod:
		return prodValueFunc()
	case commonutils.Local:
		return localValueFunc()
	default:
		panic(errors.Errorf("Not supported stage: %v", stage))
	}
}
