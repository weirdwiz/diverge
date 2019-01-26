package data

import (
	"time"
)

// User struct for the user
type User struct {
	ID        int
	UUID      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Session struct for a session
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

//LeaderBoardRow one row of the leaderboard
type LeaderBoardRow struct {
	Rank     int
	Username string
	Level    string
}

// CreateSession : Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned ID into the Session struct
	err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// Session : Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.ID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// GetQuestion : returns current question (on the level) the user is at
func (user *User) GetQuestion() (question string, err error) {
	level, err := user.GetLevel()
	if err != nil {
		return
	}
	err = Db.QueryRow("SELECT question FROM question_table WHERE id = $1", level-1).Scan(question)
	return
}

// GetLevel : returns the current level of the user
func (user *User) GetLevel() (level int, err error) {
	err = Db.QueryRow("SELECT level FROM leaderboard WHERE user_id = $1", user.ID).Scan(level)
	return
}

// Check : Check if session is valID in the database
func (session *Session) Check() (valID bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.UUID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	if err != nil {
		valID = false
		return
	}
	if session.ID != 0 {
		valID = true
	}
	return
}

// DeleteByUUID : Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.UUID)
	return
}

// User : Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid,  username, email, created_at FROM users WHERE id = $1", session.UserID).
		Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.CreatedAt)
	return
}

// SessionDeleteAll : Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert ID, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "insert into users (uuid, username, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned ID into the User struct
	err = stmt.QueryRow(createUUID(), user.Username, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.ID, &user.UUID, &user.CreatedAt)
	if err != nil {
		return
	}
	l := LeaderBoardRow{}
	err = Db.QueryRow("INSERT INTO leaderboard (username, level, solve_time) values ($1,$2,$3) returning username, level", user.Username, 0, time.Now()).Scan(&l.Username, &l.Level)
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	statement := "delete from users whereid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update users set  username = $2, email = $3 whereid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Username, user.Email)
	return
}

// UserDeleteAll : Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Users : Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid,  username, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// UserByEmail : Get a single user given the username
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, username, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// UserByUUID : Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid,  username, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// GetLeaderBoard : returns the current leaderboard
func GetLeaderBoard() (leaderbaord []*LeaderBoardRow, err error) {
	rows, err := Db.Query("SELECT username, level FROM leaderboard ORDER BY level DESC, solve_time ASC")
	rank := 1
	if err != nil {
		return
	}
	for rows.Next() {
		l := LeaderBoardRow{}
		l.Rank = rank
		if err = rows.Scan(&l.Username, &l.Level); err != nil {
			return
		}
		leaderbaord = append(leaderbaord, &l)
		rank++
	}
	rows.Close()
	return
}
