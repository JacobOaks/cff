// +build !cff
// @generated

package panic

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync"

	"github.com/uber-go/tally"
)

type panicker struct {
	scope  tally.Scope
	logger *zap.Logger
}

func (p *panicker) FlowPanicsParallel() error {
	var b bool

	err := func(ctx context.Context, scope tally.Scope,
		logger *zap.Logger) (err error) {
		flowTags := map[string]string{"name": "PanicParallel"}
		flowTagsMutex := new(sync.Mutex)
		if ctx.Err() != nil {
			s0t0Tags := map[string]string{"name": "T1"}
			scope.Tagged(s0t0Tags).Counter("task.skipped").Inc(1)
			logger.Debug("task skipped",
				zap.String("name", "T1"),
				zap.Error(ctx.Err()),
			)
			scope.Tagged(flowTags).Counter("taskflow.skipped").Inc(1)
			logger.Debug("taskflow skipped", zap.String("name", "PanicParallel"))
			return ctx.Err()
		}
		var (
			wg0 sync.WaitGroup

			once0 sync.Once
		)

		wg0.Add(2)

		var v1 string

		go func() {
			defer wg0.Done()
			tags := map[string]string{"name": "T1"}
			timer := scope.Tagged(tags).Timer("task.timing").Start()
			defer timer.Stop()
			defer func() {
				recovered := recover()
				if recovered != nil {
					once0.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)
						scope.Tagged(map[string]string{"name": "T1"}).Counter("task.panic").Inc(1)
						logger.Error("task panic",
							zap.String("name", "T1"),
							zap.Stack("stack"),
							zap.Error(recoveredErr))
						err = recoveredErr
					})
				}
			}()

			v1 = func() string {
				panic("panic")
				return ""
			}()

		}()

		var v2 int64

		go func() {
			defer wg0.Done()

			defer func() {
				recovered := recover()
				if recovered != nil {
					once0.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)

						err = recoveredErr
					})
				}
			}()

			v2 = func() int64 {
				return 0
			}()

		}()

		wg0.Wait()
		if err != nil {
			scope.Tagged(flowTags).Counter("taskflow.error").Inc(1)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = flowTagsMutex

			_ = &once0
			_ = &v1
			_ = &v2
		)

		if ctx.Err() != nil {
			scope.Tagged(flowTags).Counter("taskflow.skipped").Inc(1)
			logger.Debug("taskflow skipped", zap.String("name", "PanicParallel"))
			return ctx.Err()
		}
		var (
			once1 sync.Once
		)

		var v3 bool

		func() {

			defer func() {
				recovered := recover()
				if recovered != nil {
					once1.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)

						err = recoveredErr
					})
				}
			}()

			v3 = func(string, int64) bool {
				return true
			}(v1, v2)

		}()

		if err != nil {
			scope.Tagged(flowTags).Counter("taskflow.error").Inc(1)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = flowTagsMutex

			_ = &once1
			_ = &v3
		)

		*(&b) = v3

		if err != nil {
			scope.Tagged(flowTags).Counter("taskflow.error").Inc(1)
		} else {
			scope.Tagged(flowTags).Counter("taskflow.success").Inc(1)
			logger.Debug("taskflow succeeded", zap.String("name", "PanicParallel"))
		}

		return err
	}(context.Background(), p.scope, p.logger)

	return err
}

func (p *panicker) FlowPanicsSerial() error {
	var r string

	err := func(ctx context.Context, scope tally.Scope,
		logger *zap.Logger) (err error) {
		flowTags := map[string]string{"name": "FlowPanicsSerial"}
		flowTagsMutex := new(sync.Mutex)
		if ctx.Err() != nil {
			s0t0Tags := map[string]string{"name": "T1"}
			scope.Tagged(s0t0Tags).Counter("task.skipped").Inc(1)
			logger.Debug("task skipped",
				zap.String("name", "T1"),
				zap.Error(ctx.Err()),
			)
			scope.Tagged(flowTags).Counter("taskflow.skipped").Inc(1)
			logger.Debug("taskflow skipped", zap.String("name", "FlowPanicsSerial"))
			return ctx.Err()
		}
		var (
			once0 sync.Once
		)

		var v1 string

		func() {
			tags := map[string]string{"name": "T1"}
			timer := scope.Tagged(tags).Timer("task.timing").Start()
			defer timer.Stop()
			defer func() {
				recovered := recover()
				if recovered != nil {
					once0.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)
						scope.Tagged(map[string]string{"name": "T1"}).Counter("task.panic").Inc(1)
						logger.Error("task panic",
							zap.String("name", "T1"),
							zap.Stack("stack"),
							zap.Error(recoveredErr))
						err = recoveredErr
					})
				}
			}()

			v1 = func() string {
				panic("panic")
				return ""
			}()

		}()

		if err != nil {
			scope.Tagged(flowTags).Counter("taskflow.error").Inc(1)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = flowTagsMutex

			_ = &once0
			_ = &v1
		)

		*(&r) = v1

		if err != nil {
			scope.Tagged(flowTags).Counter("taskflow.error").Inc(1)
		} else {
			scope.Tagged(flowTags).Counter("taskflow.success").Inc(1)
			logger.Debug("taskflow succeeded", zap.String("name", "FlowPanicsSerial"))
		}

		return err
	}(context.Background(), p.scope, p.logger)

	return err
}