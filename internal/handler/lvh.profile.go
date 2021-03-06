package handler

import (
	"context"
	"html/template"

	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/domain"

	"github.com/jfyne/live"
)

const (
	// events
	eventUserProfileUpdatePage = "challenge-list-update-page"
	// params
	paramUserProfilePage = "page"
	// params value
)

type (
	ProfileInstance struct {
		*CommonInstance
		Summary    *domain.SystemSummary
		Page       int
		MaxPage    int
		TimeLayout string
	}
)

func (ins *ProfileInstance) withError(err error) *ProfileInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *ProfileInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (ins *ProfileInstance) NextPage() int {
	if ins.Page >= ins.MaxPage {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *ProfileInstance) PrevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (h *Handler) NewProfileInstance(s live.Socket) *ProfileInstance {
	m, ok := s.Assigns().(*ProfileInstance)
	if !ok {
		return &ProfileInstance{
			CommonInstance: h.NewCommon(s, viewProfile),
			Page:           1,
			TimeLayout:     h.app.Cfg.App.DefaultTimeLayout,
		}
	}

	return m
}

func (h *Handler) Profile() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.profile.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewProfileInstance // NB: make sure constructor is correct
		// SAFE TO COPY
		lvh.HandleEvent(eventCloseAuthModals, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseAuthModals()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLogoutModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLogoutModal()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLoginModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLoginModal()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseError, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseError()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseMessage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseMessage()
			return instance, nil
		})

		lvh.HandleEvent(eventToggleDark, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
			instance := constructor(s)
			if instance.User != nil {
				instance.User.UseDarkTheme = !instance.User.UseDarkTheme
			}
			instance.User, err = h.app.UpdateUser(ctx, instance.User)
			if err != nil {
				return instance.withError(err), nil
			}
			return instance, nil
		})

		// update locale logic
		lvh.HandleParams(func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.SetLocale(p.String(paramLocale))
			err := instance.updateForLocale(ctx, s, h)
			if err != nil {
				return nil, err
			}
			return instance, nil
		})

		lvh.HandleError(func(ctx context.Context, err error) {
			h.HandleError(ctx, err)
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewProfileInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil || instance.UserID == uuid.Nil {
			s.Redirect(h.url404())
			return nil, nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventUserProfileUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramUserProfilePage)
		instance := h.NewProfileInstance(s)
		instance.Page = page

		return instance, nil
	})

	return lvh
}
