package domain

// List defines the pepocampaigns list structure
type List struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
	Subscribers uint64 `json:"subscribers"`
}
