// @generated by CFF

package example

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.uber.org/cff"
	"github.com/uber-go/tally"
	"go.uber.org/zap"
)

// Request TODO
type Request struct {
	LDAPGroup string
}

// Response TODO
type Response struct {
	MessageIDs []string
}

type fooHandler struct {
	mgr    *ManagerRepository
	users  *UserRepository
	ses    *SESClient
	scope  tally.Scope
	logger *zap.Logger
}

func (h *fooHandler) HandleFoo(ctx context.Context, req *Request) (*Response, error) {
	var res *Response
	err := func(ctx context.Context, scope tally.Scope,

		logger *zap.Logger, v1 *Request) (err error) {
		var _ = (cff.FlowOption)(nil)

		var emitter cff.Emitter

		emitter = cff.DefaultEmitter(scope)

		var flowEmitterReplace sync.Once
		var _ = &flowEmitterReplace
		flowEmitter := emitter.FlowInit(
			&cff.FlowInfo{
				Flow:   "HandleFoo",
				File:   "go.uber.org/cff/examples/magic.go",
				Line:   34,
				Column: 9,
			},
		)
		startTime := time.Now()
		defer func() { flowEmitter.FlowDone(ctx, time.Since(startTime)) }()
		type task struct {
			name        string
			taskEmitter cff.TaskEmitter
			ran         bool
		}

		tasks := [][]*task{
			{
				{

					ran: false,
				},
			},
			{
				{

					ran: false,
				},
				{
					name: "FormSendEmailRequest",
					taskEmitter: emitter.TaskInit(
						&cff.TaskInfo{
							Task:   "FormSendEmailRequest",
							File:   "go.uber.org/cff/examples/magic.go",
							Line:   62,
							Column: 4,
						},
						&cff.FlowInfo{
							Flow: "HandleFoo",

							File:   "go.uber.org/cff/examples/magic.go",
							Line:   34,
							Column: 9,
						},
					),
					ran: false,
				},
			},
			{
				{
					name: "FormSendEmailRequest",
					taskEmitter: emitter.TaskInit(
						&cff.TaskInfo{
							Task:   "FormSendEmailRequest",
							File:   "go.uber.org/cff/examples/magic.go",
							Line:   67,
							Column: 4,
						},
						&cff.FlowInfo{
							Flow: "HandleFoo",

							File:   "go.uber.org/cff/examples/magic.go",
							Line:   34,
							Column: 9,
						},
					),
					ran: false,
				},
			},
			{
				{

					ran: false,
				},
			},
			{
				{

					ran: false,
				},
			},
		}
		defer func() {
			for _, sched := range tasks {
				for _, task := range sched {
					if task.name == "" || task.ran {
						continue
					}
					task.taskEmitter.TaskSkipped(ctx, err)
					if err == nil {
						logger.Debug("task skipped", zap.String("task", task.name))
					} else {
						logger.Debug("task skipped", zap.String("task", task.name), zap.Error(err))
					}
				}
			}
			if err != nil {
				flowEmitter.FlowSkipped(ctx, err)
				logger.Debug("taskflow skipped", zap.String("flow", "HandleFoo"), zap.Error(err))
			}

		}()

		if ctx.Err() != nil {
			return ctx.Err()
		}
		var (
			once0 sync.Once
		)

		var v2 *GetManagerRequest
		var v3 *ListUsersRequest

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

			v2, v3 = func(req *Request) (*GetManagerRequest, *ListUsersRequest) {
				return &GetManagerRequest{
						LDAPGroup: req.LDAPGroup,
					}, &ListUsersRequest{
						LDAPGroup: req.LDAPGroup,
					}
			}(v1)
			tasks[0][0].ran = true

		}()

		if err != nil {
			flowEmitter.FlowError(ctx, err)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = &once0
			_ = &v2
			_ = &v3
		)

		if ctx.Err() != nil {
			return ctx.Err()
		}
		var (
			wg1 sync.WaitGroup

			once1 sync.Once
		)

		wg1.Add(2)

		var v4 *GetManagerResponse
		var err1 error
		go func() {
			defer wg1.Done()

			defer func() {
				recovered := recover()
				if recovered != nil {

					once1.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)

						err = recoveredErr
					})

				}
			}()

			v4, err1 = h.mgr.Get(v2)
			tasks[1][0].ran = true
			if err1 != nil {

				once1.Do(func() {
					err = err1
				})
			}

		}()

		var v5 *ListUsersResponse
		var err4 error
		go func() {
			defer wg1.Done()
			taskEmitter := tasks[1][1].taskEmitter
			startTime := time.Now()
			defer func() { taskEmitter.TaskDone(ctx, time.Since(startTime)) }()
			defer func() {
				recovered := recover()
				if recovered != nil {

					taskEmitter.TaskPanic(ctx, recovered)
					taskEmitter.TaskRecovered(ctx, recovered)
					recoveredErr, ok := recovered.(error)
					if ok {
						logger.Error("task panic recovered",
							zap.String("task", "FormSendEmailRequest"),
							zap.Stack("stack"),
							zap.Error(recoveredErr))
					} else {
						logger.Error("task panic recovered",
							zap.String("task", "FormSendEmailRequest"),
							zap.Stack("stack"),
							zap.Any("recoveredValue", recovered))
					}
					v5, err4 = &ListUsersResponse{}, nil

				}
			}()

			v5, err4 = h.users.List(v3)
			tasks[1][1].ran = true
			if err4 != nil {
				taskEmitter.TaskError(ctx, err4)
				taskEmitter.TaskRecovered(ctx, err4)
				logger.Error("task error recovered",
					zap.String("task", "FormSendEmailRequest"),
					zap.Error(err4),
				)

				v5, err4 = &ListUsersResponse{}, nil
			} else {
				taskEmitter.TaskSuccess(ctx)
				logger.Debug("task succeeded", zap.String("task", "FormSendEmailRequest"))
			}

		}()

		wg1.Wait()
		if err != nil {
			flowEmitter.FlowError(ctx, err)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = &once1
			_ = &v4
			_ = &v5
		)

		if ctx.Err() != nil {
			return ctx.Err()
		}
		var (
			once2 sync.Once
		)

		var v6 []*SendEmailRequest

		func() {
			taskEmitter := tasks[2][0].taskEmitter
			startTime := time.Now()
			defer func() { taskEmitter.TaskDone(ctx, time.Since(startTime)) }()
			defer func() {
				recovered := recover()
				if recovered != nil {

					once2.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)
						taskEmitter.TaskPanic(ctx, recovered)
						logger.Error("task panic",
							zap.String("task", "FormSendEmailRequest"),
							zap.Stack("stack"),
							zap.Error(recoveredErr))
						err = recoveredErr
					})

				}
			}()

			if func(req *GetManagerRequest) bool {
				return req.LDAPGroup != "everyone"
			}(v2) {

				v6 = func(mgr *GetManagerResponse, users *ListUsersResponse) []*SendEmailRequest {
					var reqs []*SendEmailRequest
					for _, u := range users.Emails {
						reqs = append(reqs, &SendEmailRequest{Address: u})
					}
					return reqs
				}(v4, v5)
				tasks[2][0].ran = true

				taskEmitter.TaskSuccess(ctx)
				logger.Debug("task succeeded", zap.String("task", "FormSendEmailRequest"))

			}

		}()

		if err != nil {
			flowEmitter.FlowError(ctx, err)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = &once2
			_ = &v6
		)

		if ctx.Err() != nil {
			return ctx.Err()
		}
		var (
			once3 sync.Once
		)

		var v7 []*SendEmailResponse
		var err2 error
		func() {

			defer func() {
				recovered := recover()
				if recovered != nil {

					once3.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)

						err = recoveredErr
					})

				}
			}()

			v7, err2 = h.ses.BatchSendEmail(v6)
			tasks[3][0].ran = true
			if err2 != nil {

				once3.Do(func() {
					err = err2
				})
			}

		}()

		if err != nil {
			flowEmitter.FlowError(ctx, err)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = &once3
			_ = &v7
		)

		if ctx.Err() != nil {
			return ctx.Err()
		}
		var (
			once4 sync.Once
		)

		var v8 *Response

		func() {

			defer func() {
				recovered := recover()
				if recovered != nil {

					once4.Do(func() {
						recoveredErr := fmt.Errorf("task panic: %v", recovered)

						err = recoveredErr
					})

				}
			}()

			v8 = func(responses []*SendEmailResponse) *Response {
				var r Response
				for _, res := range responses {
					r.MessageIDs = append(r.MessageIDs, res.MessageID)
				}
				return &r
			}(v7)
			tasks[4][0].ran = true

		}()

		if err != nil {
			flowEmitter.FlowError(ctx, err)
			return err
		}

		// Prevent variable unused errors.
		var (
			_ = &once4
			_ = &v8
		)

		*(&res) = v8

		if err != nil {
			flowEmitter.FlowError(ctx, err)
		} else {
			flowEmitter.FlowSuccess(ctx)
			logger.Debug("taskflow succeeded", zap.String("flow", "HandleFoo"))
		}

		return err
	}(ctx, h.scope, h.logger, req)
	return res, err
}

// ManagerRepository TODO
type ManagerRepository struct{}

// GetManagerRequest TODO
type GetManagerRequest struct {
	LDAPGroup string
}

// GetManagerResponse TODO
type GetManagerResponse struct {
	Email string
}

// Get TODO
func (*ManagerRepository) Get(req *GetManagerRequest) (*GetManagerResponse, error) {
	return &GetManagerResponse{Email: "boss@example.com"}, nil
}

// UserRepository TODO
type UserRepository struct{}

// ListUsersRequest TODO
type ListUsersRequest struct {
	LDAPGroup string
}

// ListUsersResponse TODO
type ListUsersResponse struct {
	Emails []string
}

// List TODO
func (*UserRepository) List(req *ListUsersRequest) (*ListUsersResponse, error) {
	return &ListUsersResponse{
		Emails: []string{"a@example.com", "b@example.com"},
	}, nil
}

// SESClient TODO
type SESClient struct{}

// SendEmailRequest TODO
type SendEmailRequest struct {
	Address string
}

// SendEmailResponse TODO
type SendEmailResponse struct {
	MessageID string
}

// BatchSendEmail TODO
func (*SESClient) BatchSendEmail(req []*SendEmailRequest) ([]*SendEmailResponse, error) {
	res := make([]*SendEmailResponse, len(req))
	for i := range req {
		res[i] = &SendEmailResponse{MessageID: strconv.Itoa(i)}
	}
	return res, nil
}
