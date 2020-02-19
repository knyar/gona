package gona

// Location is a struct for storing the id and name of a location
type Location struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
}

// GetLocations public method on Client to get a list of locations
func (c *Client) GetLocations() ([]Location, error) {

	var locationMap map[string]Location
	var locationList []Location

	if err := c.get("cloud/locations", &locationMap); err != nil {
		return nil, err
	}

	for _, loc := range locationMap {
		locationList = append(locationList, loc)
	}

	return locationList, nil
}
