package balance

type Balance interface {
	DoBalance(instances []*Instance, key ...string)(inst *Instance, err error)
}
