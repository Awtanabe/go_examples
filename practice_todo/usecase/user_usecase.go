package usecase

import (
	"os"
	"practice_todo/models"
	"practice_todo/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


type IUserUsecase interface {
	SignUp(user models.User) (models.UserResponse,error)
	Login(user models.User)( string, error)
	GetUser(userId uint) (models.UserResponse2, error)
}

type userUsecase struct {
	// userユースケースはリポジトリを持っておくことでdb操作が可能になる
	ur repository.IUserRepogitory
}

func NewUserUsecase(ur repository.IUserRepogitory) IUserUsecase{
	return &userUsecase{ur}
}

// userはこんな感じで渡される
// user := models.User{}
// c.Bind(&user)
func (uu *userUsecase)SignUp(user models.User) (models.UserResponse, error) {
	//パスワードをhashで暗号化する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return models.UserResponse{},err
	}

	// パラメーターを利用して
	newUser := models.User{Email: user.Email, Password: string(hash)}

	if err := uu.ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}

	resUser := models.UserResponse {
		ID: newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user models.User)(string, error){
	//clientからくるemailがdbに存在するか確認する
	storedUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storedUser,user.Email); err != nil{
		return "", err
	}
	//パスワードの一致確認
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	//JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": int(storedUser.ID),
		"exp": time.Now().Add(time.Hour * 12).Unix(), //有効期限
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		return "",err
	}
	return tokenString, nil
}

func (uu *userUsecase) GetUser(userId uint) (models.UserResponse2, error) {
	user := models.User{}
	if err := uu.ur.GetUser(userId, &user); err != nil {
		return models.UserResponse2{}, err
	}
	return models.UserResponse2{
		ID: user.ID,
		Email: user.Email,
		Todos: user.Todos,
	}, nil
}

// if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(todos).Error; err != nil{
