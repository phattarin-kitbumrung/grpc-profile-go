package services

import (
	context "context"
	"fmt"
	"io"
	"time"
)

type ProfileService interface {
	Profile(profile ...map[string]interface{}) error
}

type profileService struct {
	profileClient ProfileClient
}

func NewProfileService(profileClient ProfileClient) ProfileService {
	return profileService{profileClient}
}

func (base profileService) Profile(profile ...map[string]interface{}) error {
	stream, err := base.profileClient.Profile(context.Background())
	if err != nil {
		return err
	}

	go func() {
		for _, data := range profile {
			req := ProfileRequest{
				Profile: &MyProfile{
					Introduction: data["profile"].(map[string]interface{})["introduction"].(string),
					AboutMe: &AboutMe{
						Fullname: data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["fullname"].(string),
						Address:  data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["address"].(string),
						Email:    data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["email"].(string),
						Mobile:   data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["mobile"].(string),
						Nickname: data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["nickname"].(string),
						Birthday: data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["birthday"].(string),
						Religion: data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["religion"].(string),
						Github:   data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["github"].(string),
						Linkedin: data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["linkedin"].(string),
						Medium:   data["profile"].(map[string]interface{})["aboutMe"].(map[string]interface{})["medium"].(string),
					},
					Educations:      data["profile"].(map[string]interface{})["educations"].([]string),
					Certifications:  data["profile"].(map[string]interface{})["certifications"].([]string),
					Workexperiences: data["profile"].(map[string]interface{})["workexperiences"].([]string),
					Skills:          data["profile"].(map[string]interface{})["skills"].([]string),
				},
				Filename: data["filename"].(string),
			}
			stream.Send(&req)
			fmt.Printf("Request : %v\n", req.Profile)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	done := make(chan bool)
	errs := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
			}
			fmt.Printf("Response: %v\n", res.Filename)
		}
		done <- true
	}()

	if <-done {
		return nil
	}

	return <-errs
}
