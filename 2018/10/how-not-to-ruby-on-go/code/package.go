package dao
import (
"errors"
"fmt"
"time"

"github.com/astaxie/beego/validation"
"github.com/gtforge/order_locations_service/dto"
"github.com/gtforge/order_locations_service/models"
"github.com/gtforge/services_common_go/gett-storages"
	"github.com/satori/go.uuid"
	"github.com/gtforge/rides_api/models/location"
)
// START OMIT
package dao

func Find(uuid uuid.UUID) User, error {}

func FindLocations(uuid uuid.UUID) []location.Location, error { ... }
.
.
.

type ldao struct{}

func (t *ldao) Find(uuid uuid.UUID) []location.Location { ... }
// END OMIT
