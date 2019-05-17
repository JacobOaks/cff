// +build !cff
// @generated

package fallbackwith

import (
	"context"
	"fmt"
	"sync"
)

// Serial executes a flow that fails with the given error, if any and recovers
// with the given string.
func Serial(e error, r string) (string, error) {
	var s string
	err := func(ctx context.Context) (err error) {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		var (
			once0 sync.Once
		)

		var v1 string
		var err0 error
		func() {

			defer func() {
				recovered := recover()
				if recovered != nil {
					once0.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)

						err = recoveredErr
					})
				}
			}()

			v1, err0 = func() (string, error) {
				return "foo", e
			}()
			if err0 != nil {

				v1, err0 = r, nil
			}

		}()

		if err != nil {

			return err
		}

		// Prevent variable unused errors.
		var (
			_ = &once0
			_ = &v1
		)

		*(&s) = v1

		return err
	}(context.Background())
	return s, err
}