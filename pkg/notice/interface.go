package notice

type Notifier interface {
	SendNotify(msg string)
	Getenable() bool
}
