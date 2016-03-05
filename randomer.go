package randomer

// Randomer - Main structure to perform the random string creation
type Randomer struct {
	settings *Settings
}

// NewRandomer - Create a new Randomer
func NewRandomer(settings *Settings) *Randomer {
	var r = new(Randomer)
	r.settings = settings

	return r
}

// GetSettings - Get specified settings
func (r Randomer) GetSettings() *Settings {
	return r.settings
}

// GetRandomString - Gets a random string
func (r *Randomer) GetRandomString() string {
	// actually return some data!
	return "aa"
}

// Settings - Setting struct for Randomer
type Settings struct {
	characters string
}

// SetCharacters - Set the characters
func (s *Settings) SetCharacters(characters string) {
	s.characters = characters
}

// GetCharacters - Get the characters
func (s Settings) GetCharacters() string {
	return s.characters
}
