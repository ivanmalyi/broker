package technodom

import (
	"github.com/ivanmalyi/broker/internal/app/repository"
)

type Technodom struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Technodom {
	return &Technodom{repo}
}

func (technodom *Technodom) Auth() {

}

func (technodom *Technodom) CreateApp() string {
	return `{"status":"Success","code":"200","redirectUrl":"https://alfabank.kz","requestUuid":"3244-412mf-dt23g-4352g","message":"Заявка успешно создана"}`
}
