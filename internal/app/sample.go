package app

import (
	"context"
	"fmt"

	"go.uber.org/multierr"

	"github.com/tinygodsdev/orrery/internal/domain"
)

func (a *App) initSamples() error {
	samples := []*domain.Sample{
		{
			Code:     domain.SampleAllCode,
			Criteria: domain.SampleCriteria{},
		},
		{
			Code: domain.SampleAllNonSuspiciousCode,
			Criteria: domain.SampleCriteria{
				NotSuspicious: true,
			},
		},
	}

	// create samples for all locales
	for _, s := range samples {
		for _, l := range domain.Locales() {
			crit := s.Criteria
			crit.Locale = l
			samples = append(samples, &domain.Sample{
				Code:     fmt.Sprintf("%s-%s", s.Code, l),
				Criteria: crit,
			})
		}
	}

	var errs error
	for _, s := range samples {
		smp, err := a.repo.CreateOrUpdateSample(context.TODO(), s)
		a.log.Info("created sample", fmt.Sprintf("%+v", smp))
		if err != nil {
			a.log.Error("init samples failed", err)
			if a.IsDev() {
				return err
			} else {
				errs = multierr.Append(errs, err)
				continue
			}
		}
	}

	return errs
}
