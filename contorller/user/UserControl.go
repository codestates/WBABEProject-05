package user

var instance *userControl

type userControl struct {
}

func GetUserControl() *userControl {
	if instance != nil {
		return instance
	}
	instance = &userControl{}
	return instance
}

func (u *userControl) FindUser() {

}
func (u *userControl) UpdateUser() {

}
func (u *userControl) DeleteUser() {

}
func (u *userControl) InsertUser() {

}
