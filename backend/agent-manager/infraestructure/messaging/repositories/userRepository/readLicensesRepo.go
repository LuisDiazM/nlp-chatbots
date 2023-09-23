package userRepository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/repositories"

	"github.com/LuisDiazM/agent-manager/infraestructure/messaging"
)

type UserLicensesMessagingRepository struct {
	Nats *messaging.NatsImp
}

func NewUserLicenseMessagingRepository(nats *messaging.NatsImp) repositories.LicensesRepoGateway {
	return &UserLicensesMessagingRepository{Nats: nats}
}

func (repository *UserLicensesMessagingRepository) GetLicensesByUser(userId string) *[]entities.License {
	var getLicenseRequest RequestGetLicenses = RequestGetLicenses{UserId: userId}
	data, err := json.Marshal(getLicenseRequest)
	if err != nil {
		log.Println(err)
		return nil
	}
	msg, err := repository.Nats.Conn.Request(queryGetLicenses, data, timeout)
	if err != nil {
		log.Println(err)
		return nil
	}
	var licenses []entities.License
	err = json.Unmarshal(msg.Data, &licenses)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &licenses
}

func (repository *UserLicensesMessagingRepository) CreateLicenseByUser(userId string, licenseType string) error {
	requestData := RequestCreateLicense{UserId: userId, LicenseType: licenseType}
	data, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = repository.Nats.Conn.Publish(eventCreateLicense, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repository *UserLicensesMessagingRepository) GetLastLicensesByUser(userId string) *entities.License {
	var getLicenseRequest RequestGetLicenses = RequestGetLicenses{UserId: userId}
	data, err := json.Marshal(getLicenseRequest)
	if err != nil {
		log.Println(err)
		return nil
	}
	msg, err := repository.Nats.Conn.Request(queryGetLastLicense, data, timeout)
	if err != nil {
		log.Println(err)
		return nil
	}
	var license entities.License
	err = json.Unmarshal(msg.Data, &license)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &license
}

func (repository *UserLicensesMessagingRepository) GetLastLicensesUsage(licenseId string) *entities.LicensesUsage {
	var getLicenseRequest RequestGetLicenseUsage = RequestGetLicenseUsage{LicenseId: licenseId}
	data, err := json.Marshal(getLicenseRequest)
	if err != nil {
		log.Println(err)
		return nil
	}
	msg, err := repository.Nats.Conn.Request(queryGetLicenseUsage, data, timeout)
	if err != nil {
		log.Println(err)
		return nil
	}
	var license entities.LicensesUsage
	err = json.Unmarshal(msg.Data, &license)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &license
}
