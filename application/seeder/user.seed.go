package seeder

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
	"github.com/bxcodec/faker"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (ms UserSeeder) Run(db *gorm.DB) {
	var (
		userRepository repository.UserRepository = repository.NewUserRepository(db)
	)

	group := &sync.WaitGroup{}

	type UserFaker struct {
		Email    string `faker:"email"`
		UserName string `faker:"username"`
		Sentence string `faker:"sentence"`
	}
	go func() {
		group.Add(10)
		defer group.Done()
		for i := 0; i < 10; i++ {
			fake := UserFaker{}
			err := faker.FakeData(&fake)
			if err != nil {
				fmt.Println(err)
			}
			pass := md5.Sum([]byte("12345"))
			token := sha256.Sum256([]byte(fake.Email))
			userRepository.Insert(model.User{
				Name:     fake.UserName,
				Email:    fake.Email,
				Password: fmt.Sprintf("%x", pass),
				Token:    fmt.Sprintf("%x", token),
				Alamat: []*model.Alamat{
					{
						ProvinceId:  1,
						CityId:      3,
						Description: fake.Sentence,
					},
				},
			})
		}
		group.Wait()
	}()

}
