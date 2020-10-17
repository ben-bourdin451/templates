package src

// Event maps the received payload into the Lambda
type Event struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}
