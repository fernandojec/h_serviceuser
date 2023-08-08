package users

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/fernandojec/h_serviceuser/domain/notifications"
	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	//"github.com/fernandojec/assignment-2/pkg/smtp"
)

type userService struct {
	repo_users         repo_users
	grpc_notifications notifications.MessagingServiceClient
}

type repo_users interface {
	saveRegistrasiUsers_services
}

type saveRegistrasiUsers_services interface {
	getbyUserID(userID string) (sverifyUsers verifyUsers, err error)
	saveRegistrasiUsers(data tableUsersModel) (user_id string, err error)
}

func NewServiceUsers(repo_users repo_users, grpc_notifications notifications.MessagingServiceClient) userService {
	return userService{repo_users: repo_users, grpc_notifications: grpc_notifications}
}

func (s userService) saveRegistrasi(ctx context.Context, req tableUsersModel, baseVerifyEmail string) (err error) {
	log.Println("test 1")
	stableUsersModel, errCheck := s.repo_users.getbyUserID(req.UserID)

	log.Println("test 2")
	if errCheck != nil && errCheck != sql.ErrNoRows {
		err = errCheck
		return
	}
	if stableUsersModel.UserID != "" {
		return errors.New("user id exists")
	}

	log.Println("test 3")

	_, err = s.repo_users.saveRegistrasiUsers(req.ConvertToUsers())
	if err != nil {
		return err
	}
	emailReq := notifications.EmailRequest{
		From:    "ricky.lai@jec.co.id",
		To:      []string{req.Email},
		Subject: "Registrasi Berhasil",
		Body:    "Selamat, anda telah terdaftar di aplikasi kami. Dengan user id:" + req.UserID,
	}
	emailResp, errSend := s.grpc_notifications.SendEmail(ctx, &emailReq)
	if errSend != nil {
		loghelper.Errorf(ctx, "Error Send Email:%+v", emailResp)
		err = errors.New("email not sended")
	}
	return
}
