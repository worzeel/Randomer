package randomer

// Randomer - Main structure to perform the random string creation
type Randomer struct {
	setting Setter
}

// NewRandomer - Create a new Randomer
func NewRandomer(setting Setter) *Randomer {
	var r = new(Randomer)
	r.setting = setting

	return r
}

// GetSettings - Get specified settings
func (r Randomer) GetSettings() Setter {
	return r.setting
}

// GetRandomString - Gets a random string
func (r *Randomer) GetRandomString() string {
	// actually return some data!
	return "aabbccddee"
}
