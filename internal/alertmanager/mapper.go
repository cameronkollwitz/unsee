package alertmanager

import (
	"github.com/cameronkollwitz/unsee/internal/mapper"
	"github.com/cameronkollwitz/unsee/internal/mapper/v04"
	"github.com/cameronkollwitz/unsee/internal/mapper/v05"
	"github.com/cameronkollwitz/unsee/internal/mapper/v061"
	"github.com/cameronkollwitz/unsee/internal/mapper/v062"
)

// initialize all mappers
func init() {
	mapper.RegisterAlertMapper(v04.AlertMapper{})
	mapper.RegisterAlertMapper(v05.AlertMapper{})
	mapper.RegisterAlertMapper(v061.AlertMapper{})
	mapper.RegisterAlertMapper(v062.AlertMapper{})
	mapper.RegisterSilenceMapper(v04.SilenceMapper{})
	mapper.RegisterSilenceMapper(v05.SilenceMapper{})
}
