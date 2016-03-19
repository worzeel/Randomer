package randomer

// Setter - Interface for settings informaiton
type Setter interface {
	SetCharacters(characters string)
	GetCharacters() string
}

// Setting - Setting struct for Randomer
type Setting struct {
	characters string
}

// NewSetting - Create a new Settings
func NewSetting() *Setting {
	return new(Setting)
}

// SetCharacters - Set the characters
func (s *Setting) SetCharacters(characters string) {
	s.characters = characters
}

// GetCharacters - Get the characters
func (s Setting) GetCharacters() string {
	return s.characters
}
