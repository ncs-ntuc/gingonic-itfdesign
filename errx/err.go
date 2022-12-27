package errx

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ErrType int

const (
	MarshalErr = iota + 10
	CacheQryErr
)

type CartErr struct {
	Ty          ErrType // identifies the type of the err
	Internal    error   // cascading internal error
	Trace       string  // location from which the err has emanated
	Diagnosis   string  // tips to solve this erro
	UserMessage string  // message that gets exposed to the client
}

// so that we can cotinue sending this as error interface from deep within the functions
func (crterr *CartErr) Error() string {
	return fmt.Sprintf("%d,%s-%s", crterr.Ty, crterr.Internal, crterr.Trace)
}
func (crterr *CartErr) UMsg() string {
	return crterr.UserMessage
}
func (crterr *CartErr) Typ() ErrType {
	return crterr.Ty
}
func Throw(ty ErrType, internal error, trace, diag, umsg string) *CartErr {
	return &CartErr{
		Ty:          ty,
		Internal:    internal,
		Trace:       trace,
		Diagnosis:   diag,
		UserMessage: umsg,
	}
}

// DigestErr : one place where all the errors get logged and retrofitted into the gin context
// this is vital since we want a uniform logging - downstream analysis
func DigestErr(c *gin.Context, err error) {
	if err != nil {
		// potential error
		crtEr := err.(*CartErr)
		switch crtEr.Ty {
		case MarshalErr, CacheQryErr:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"err": crtEr.UMsg(),
			})
		}
		log.WithFields(log.Fields{
			"trace":    crtEr.Trace,
			"internal": crtEr.Internal.Error(),
		}).Error("Error in one or more server operations")
	}
}
