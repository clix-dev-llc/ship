package ship

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/replicatedhq/ship/pkg/util/warnings"
)

func (s *Ship) UnforkAndMaybeExit(ctx context.Context) error {
	if err := s.Unfork(ctx); err != nil {
		s.ExitWithError(err)
		return err
	}
	return nil
}

func (s *Ship) Unfork(ctx context.Context) error {
	debug := level.Debug(log.With(s.Logger, "method", "unfork"))
	ctx, cancelFunc := context.WithCancel(ctx)
	defer s.Shutdown(cancelFunc)

	existingState, _ := s.State.CachedState()
	if !existingState.IsEmpty() {
		debug.Log("event", "existing.state")

		if s.Viper.GetString("state-from") != "file" {
			debug.Log("event", "existing.state", "state-from", "not file")
			return warnings.WarnCannotRemoveState
		}

		if err := s.promptToRemoveState(); err != nil {
			debug.Log("event", "state.remove.prompt.fail")
			return err
		}
	}

	s.State.UpdateVersion()

	upstream := s.Viper.GetString("upstream")
	if upstream == "" {
		return errors.New("No upstream provided")
	}

	p := s.Resolver.NewContentProcessor()
	maybeVersionedUpstream, err := p.MaybeResolveVersionedUpstream(ctx, upstream, existingState)
	if err != nil {
		return errors.Wrap(err, "create versioned upstream release")
	}

	fork := s.Viper.GetString("fork")
	if fork == "" {
		return errors.New("No fork provided")
	}

	maybeVersionedFork, err := p.MaybeResolveVersionedUpstream(ctx, fork, existingState)
	if err != nil {
		return errors.Wrap(err, "create versioned fork release")
	}

	release, err := s.Resolver.ResolveUnforkRelease(ctx, maybeVersionedUpstream, maybeVersionedFork)
	if err != nil {
		return errors.Wrap(err, "resolve release")
	}

	release.Spec.Lifecycle = s.IDPatcher.EnsureAllStepsHaveUniqueIDs(release.Spec.Lifecycle)

	if err := s.execute(ctx, release, nil); err != nil {
		return errors.Wrap(err, "execute")
	}

	if err := s.State.CommitState(); err != nil {
		return errors.Wrap(err, "commit state")
	}

	return nil
}
