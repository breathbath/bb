package observers

type Observer interface {
	AcceptResult(result ObservingResult)
}
