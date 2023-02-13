package importTest

func (DB Database) Auth(u User) bool {
	for i := 0; i < len(DB.Users); i++ {
		if DB.Users[i].Name == u.Name && DB.Users[i].Login == u.Login && DB.Users[i].Password == u.Password {
			return true
		}
	}
	return false
}
