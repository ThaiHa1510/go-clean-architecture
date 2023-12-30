package listeners

type Observer interface {
    Excute[T any](payload T)
}