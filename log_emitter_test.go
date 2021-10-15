package cff

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogFlowEmitter_IncludesFlowName(t *testing.T) {
	core, observed := observer.New(zapcore.DebugLevel)

	em := LogEmitter(zap.New(core)).FlowInit(&FlowInfo{Name: "myflow"})
	em.FlowSuccess(context.Background())
	em.FlowError(context.Background(), errors.New("foo"))

	for _, logEntry := range observed.TakeAll() {
		fields := logEntry.ContextMap()
		assert.Equalf(t, "myflow", fields["flow"],
			"flow name expected in %#v", fields)
	}
}
func TestLogFloWEmitter_ErrorLevelChange(t *testing.T) {
	core, observed := observer.New(zapcore.DebugLevel)

	LogEmitter(
		zap.New(core),
		LogErrors(zapcore.WarnLevel),
	).FlowInit(&FlowInfo{Name: "myflow"}).
		FlowError(context.Background(), errors.New("great sadness"))

	logs := observed.TakeAll()
	require.Len(t, logs, 1)
	logs[0].Level = zapcore.WarnLevel
	assert.Equal(t, "great sadness", logs[0].ContextMap()["error"])
}

func TestLogTaskEmitter(t *testing.T) {
	ctx := context.Background()
	core, observed := observer.New(zapcore.DebugLevel)
	emitter := LogEmitter(zap.New(core))
	tem := emitter.TaskInit(&TaskInfo{Name: "mytask"}, &FlowInfo{Name: "myflow"})

	t.Run("includes task and flow name", func(t *testing.T) {
		tem.TaskSuccess(ctx)
		tem.TaskErrorRecovered(ctx, errors.New("great sadness"))

		for _, logEntry := range observed.TakeAll() {
			fields := logEntry.ContextMap()
			assert.Equalf(t, "mytask", fields["task"],
				"task name expected in %#v", fields)
			assert.Equalf(t, "myflow", fields["flow"],
				"flow name expected in %#v", fields)
		}
	})

	t.Run("panic with value", func(t *testing.T) {
		tem.TaskPanic(ctx, "foo")
		logs := observed.TakeAll()
		require.Len(t, logs, 1)
		assert.Equal(t, "foo", fmt.Sprint(logs[0].ContextMap()["panic-value"]))
	})

	t.Run("panic with error", func(t *testing.T) {
		tem.TaskPanic(ctx, errors.New("great sadness"))
		logs := observed.TakeAll()
		require.Len(t, logs, 1)
		assert.Equal(t, "great sadness", fmt.Sprint(logs[0].ContextMap()["panic-value"]))
		assert.Equal(t, "great sadness", fmt.Sprint(logs[0].ContextMap()["error"]))
	})

	t.Run("panic recovered with value", func(t *testing.T) {
		tem.TaskPanicRecovered(ctx, "foo")
		logs := observed.TakeAll()
		require.Len(t, logs, 1)
		assert.Equal(t, "foo", fmt.Sprint(logs[0].ContextMap()["panic-value"]))
	})

	t.Run("panic recovered with error", func(t *testing.T) {
		tem.TaskPanicRecovered(ctx, errors.New("great sadness"))
		logs := observed.TakeAll()
		require.Len(t, logs, 1)
		assert.Equal(t, "great sadness", fmt.Sprint(logs[0].ContextMap()["panic-value"]))
		assert.Equal(t, "great sadness", fmt.Sprint(logs[0].ContextMap()["error"]))
	})
}

func TestLogTaskEmitter_CustomizeLevels(t *testing.T) {
	core, observed := observer.New(zapcore.DebugLevel)
	ctx := context.Background()

	em := LogEmitter(
		zap.New(core),
		LogErrors(zapcore.WarnLevel),
		LogPanics(zapcore.InfoLevel),
		LogRecovers(zapcore.DebugLevel),
	).TaskInit(&TaskInfo{Name: "mytask"}, &FlowInfo{Name: "myflow"})

	t.Run("error level", func(t *testing.T) {
		em.TaskError(ctx, errors.New("great sadness"))

		logs := observed.TakeAll()
		require.Len(t, logs, 1)
		logs[0].Level = zapcore.WarnLevel
		assert.Equal(t, "great sadness", logs[0].ContextMap()["error"])
	})

	t.Run("panic level", func(t *testing.T) {
		em.TaskPanic(ctx, "something went wrong")

		logs := observed.TakeAll()
		require.Len(t, logs, 1)
		logs[0].Level = zapcore.InfoLevel
		assert.Equal(t, "something went wrong", logs[0].ContextMap()["panic-value"])
	})

	t.Run("recover level", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			em.TaskErrorRecovered(ctx, errors.New("great sadness"))

			logs := observed.TakeAll()
			require.Len(t, logs, 1)
			logs[0].Level = zapcore.DebugLevel
			assert.Equal(t, "great sadness", logs[0].ContextMap()["error"])
		})

		t.Run("panic", func(t *testing.T) {
			em.TaskPanicRecovered(ctx, "something went wrong")

			logs := observed.TakeAll()
			require.Len(t, logs, 1)
			logs[0].Level = zapcore.DebugLevel
			assert.Equal(t, "something went wrong", logs[0].ContextMap()["panic-value"])
		})
	})
}

func TestLogTaskEmitter_EmitScheduler(t *testing.T) {
	core, observed := observer.New(zapcore.DebugLevel)

	t.Run("log a scheduler emission", func(t *testing.T) {
		em := LogEmitter(
			zap.New(core),
		).SchedulerInit(&SchedulerInfo{Name: "myflow", Directive: FlowDirective})

		em.EmitScheduler(SchedulerState{})
		logs := observed.TakeAll()
		require.Len(t, logs, 1)

		assert.Contains(t, "scheduler state", logs[0].Message)

		v, ok := logs[0].ContextMap()["flow"]
		require.True(t, ok)
		assert.Equal(t, "myflow", v)
	})

	t.Run("empty flow is skipped", func(t *testing.T) {
		em := LogEmitter(
			zap.New(core),
		).SchedulerInit(&SchedulerInfo{})

		em.EmitScheduler(SchedulerState{})
		logs := observed.TakeAll()
		require.Len(t, logs, 1)

		_, ok := logs[0].ContextMap()["flow"]
		assert.False(t, ok)
	})
}
