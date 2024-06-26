package subscription

type SignupNewsletterDto struct {
	Email string `json:"email" validate:"required,email,max=50"`
}
