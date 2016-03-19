package randomer

// Settings - Setting struct for Randomer
type Settings struct {
	characters string
}

// NewSettings - Create a new Settings
func NewSettings() *Settings {
	return new(Settings)
}

// SetCharacters - Set the characters
func (s *Settings) SetCharacters(characters string) {
	s.characters = characters
}

// GetCharacters - Get the characters
func (s Settings) GetCharacters() string {
	return s.characters
}
