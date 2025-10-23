package dbtest

import (
	"time"

	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/business/domain/userbus/extensions/userotel"
	"github.com/ardanlabs/service/business/domain/userbus/stores/usercache"
	"github.com/ardanlabs/service/business/domain/userbus/stores/userdb"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	User userbus.ExtBusiness
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	userStorage := usercache.NewStore(log, userdb.NewStore(log, db), time.Hour)

	userBus := userbus.NewBusiness(log, userStorage, userotel.NewExtension())

	return BusDomain{
		User: userBus,
	}
}
