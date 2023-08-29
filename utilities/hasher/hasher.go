package hasher

import "golang.org/x/crypto/bcrypt"

type HashPassword interface {
	HashedPassword(password string) (hashed string, err error)
	VerifyPassword(hashed, password string) (err error)
}

type HasherPassword struct{}

func (h *HasherPassword) HashedPassword(password string) (hashed string, err error) {
	hashedPassword, err := h.Hash(password)
	if err != nil {
		return
	}
	hashed = string(hashedPassword)
	return
}

func (h *HasherPassword) VerifyPassword(hashed, password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func (h *HasherPassword) Hash(password string) (hashed []byte, err error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
