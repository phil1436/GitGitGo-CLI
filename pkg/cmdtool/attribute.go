package cmdtool

// This represents a attribute that is bind to a command for example: npm install <attribute>
type Attribute struct {
	Name        string
	Description string
}

// Create a new attribute
func NewAttribute(name string, description string) *Attribute {
	return &Attribute{
		Name:        name,
		Description: description,
	}
}
