package notifications

import (
	"github.com/xlab/android-go/app"
	"github.com/xlab/android-go/notification"
)

type AndroidNotifier struct {
	name string
	ctx  app.Context
	mngr notification.Manager
}

func (n AndroidNotifier) NewNotifier(name *string) *AndroidNotifier {
	n.ctx = app.Context()
	n.mngr = notification.NewManager(n.ctx, *name)
	return &n
}

func (n *AndroidNotifier) Notify(title *string, text *string, autocancel bool) {
	builder := notification.NewBuilder(n.ctx, n.mngr.ChannelId()).
		SetSmallIcon("android.R.Drawable.ic_notification").
		SetContentTitle(title).
		SetContentText(text).
		SetAutoCancel(autocancel)

	n.mngr.Notify(11111111, builder.Build())
}
