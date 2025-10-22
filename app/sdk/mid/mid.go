// Package mid asdas das dsa dsq.
package mid

import (
	"errors"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func isError(e web.Encoder) error {
	err, isError := e.(error)
	if isError {
		var appErr *errs.Error
		if errors.As(err, &appErr) && appErr.Code == errs.None {
			return nil
		}
		return err
	}
	return nil
}
