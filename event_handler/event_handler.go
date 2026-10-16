package eventhandler

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	e_uuid uuid.UUID
	name string

	datetime time.DateTime
	duration time.Time
	roles []QuantizedRole
	edit_lock bool
}
