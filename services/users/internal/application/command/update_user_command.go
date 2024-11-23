package command

// TODO нужно ли здесь ставить все что можно? Думаю нет. Надо мб сделать для каждого метода свой инпут.
// Будет как в gprc с реквестами... надо ли оно... ps. мб чето с дженериками придумать.
type UpdateUserCommand struct {
	Id       string
	Name     string
	Email    string
	Password string
}
