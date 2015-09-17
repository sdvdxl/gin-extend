package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/sdvdxl/gin-extend/sources/util/log"
)

var (
	SessionStore sessions.CookieStore
)

func init() {
	log.Logger.Info("init session ...")

	SessionStore = sessions.NewCookieStore([]byte("_IlikeUkeTravel.com_oyear"))

	log.Logger.Info("session inited")
}
