package tmpl

type optionName int

const (
	optionLayoutName optionName = iota
)

// Optioner interface
type Optioner interface {
	getName() optionName
	getValue() string
}

// OptionLayoutName struct
type OptionLayoutName struct {
	Name string
}

func (o OptionLayoutName) getName() optionName {
	return optionLayoutName
}

func (o OptionLayoutName) getValue() string {
	return o.Name
}
