package main

import (
	"client/services"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:50051", "gRPC server host")
	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		creds = insecure.NewCredentials()
	}

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	profileClient := services.NewProfileClient(cc)
	profileService := services.NewProfileService(profileClient)

	myProfile1 := map[string]interface{}{
		"profile": map[string]interface{}{
			"introduction": "This is minizymint profile1",
			"aboutMe": map[string]interface{}{
				"fullname": "phattarin kitbumrung",
				"address":  "xxx",
				"email":    "xxx",
				"mobile":   "xxx",
				"nickname": "mint",
				"birthday": "xxx",
				"religion": "xxx",
				"github":   "xxx",
				"linkedin": "xxx",
				"medium":   "xxx",
			},
			"educations": []string{"cs22"},
			"certifications": []string{
				"microsoft",
			},
			"workexperiences": []string{
				"w1",
				"w2",
			},
			"skills": []string{
				"s1",
				"s2",
			},
		},
		"filename": "test1.png",
	}
	myProfile2 := map[string]interface{}{
		"profile": map[string]interface{}{
			"introduction": "This is minizymint profile2",
			"aboutMe": map[string]interface{}{
				"fullname": "phattarin kitbumrung",
				"address":  "xxx",
				"email":    "xxx",
				"mobile":   "xxx",
				"nickname": "mint",
				"birthday": "xxx",
				"religion": "xxx",
				"github":   "xxx",
				"linkedin": "xxx",
				"medium":   "xxx",
			},
			"educations": []string{"cs22"},
			"certifications": []string{
				"microsoft",
			},
			"workexperiences": []string{
				"w1",
				"w2",
			},
			"skills": []string{
				"s1",
				"s2",
			},
		},
		"filename": "test2.png",
	}
	myProfile3 := map[string]interface{}{
		"profile": map[string]interface{}{
			"introduction": "This is minizymint profile3",
			"aboutMe": map[string]interface{}{
				"fullname": "phattarin kitbumrung",
				"address":  "xxx",
				"email":    "xxx",
				"mobile":   "xxx",
				"nickname": "mint",
				"birthday": "xxx",
				"religion": "xxx",
				"github":   "xxx",
				"linkedin": "xxx",
				"medium":   "xxx",
			},
			"educations": []string{"cs22"},
			"certifications": []string{
				"microsoft",
			},
			"workexperiences": []string{
				"w1",
				"w2",
			},
			"skills": []string{
				"s1",
				"s2",
			},
		},
		"filename": "test3.png",
	}
	err = profileService.Profile(myProfile1, myProfile2, myProfile3)

	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}
}
