package command

type CreateUserCommand struct {
	// TODO: Implement idempotency key. Что это и как это есть?
	Name  string
	Email string
}
