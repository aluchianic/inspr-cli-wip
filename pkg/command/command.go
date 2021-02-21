package command

var (
	varInitFncs []func()
	cmdInitFncs []func()
)

// RegisterCommandVar is used to register with inspr the initialization function
// for the command variable.
// Something must be returned to use the `var _ = ` trick.
func RegisterCommandVar(c func()) bool {
	varInitFncs = append(varInitFncs, c)

	return true
}

// RegisterCommandInit is used to register with inspr the initialization function
// for the command flags.
// Something must be returned to use the `var _ = ` trick.
func RegisterCommandInit(c func()) bool {
	cmdInitFncs = append(cmdInitFncs, c)
	return true
}

// Main starts the inspr cli
func Setup() {
	// Setup all command variables.
	for _, v := range varInitFncs {
		v()
	}

	// Call all command inits.
	for _, f := range cmdInitFncs {
		f()
	}
}
