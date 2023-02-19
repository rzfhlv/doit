package repository

var (
	RegisterQuery = `INSERT INTO users 
		(name, email, username, password, created_at) 
		VALUES ($1, $2, $3, $4, $5) RETURNING *`
	LoginQuery = `SELECT * FROM users WHERE username = $1`
)
