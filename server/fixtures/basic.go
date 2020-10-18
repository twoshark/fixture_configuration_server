package fixtures

//BasicFixtures ...
type BasicFixtures []BasicFixture

//BasicFixture is a 5 Channel Lighting Fixture
type BasicFixture struct {
	Name      string `json:"name"`
	Opacity   uint8  `json:"opacity"`
	Animation uint8  `json:"animation"`
	Option    uint8  `json:"option"`
	Speed     uint8  `json:"speed"`
	Strobe    uint8  `json:"strobe"`
}

//Add ...
func (b BasicFixtures) Add(newFixture BasicFixture) {
	b = append(b, newFixture)
}

//Find ...
func (b BasicFixtures) Find(name string) (int, BasicFixture) {
	for index, fixture := range b {
		if fixture.Name == name {
			return index, fixture
		}
	}
	return -1, BasicFixture{}
}

//Update ...
func (b BasicFixtures) Update(index int, update BasicFixture) {
	b[index] = update
}

//Delete ...
func (b BasicFixtures) Delete(i int, preserveOrder bool) {
	copy(b[i:], b[i+1:])         // Shift a[i+1:] left one index.
	b[len(b)-1] = BasicFixture{} // Erase last element (write zero value).
	b = b[:len(b)-1]             // Truncate slice.

}
