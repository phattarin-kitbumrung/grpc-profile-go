package services

import (
	"image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type profileServer struct {
}

func NewProfileServer() ProfileServer {
	return profileServer{}
}

func (profileServer) mustEmbedUnimplementedProfileServer() {}

func (profileServer) Profile(stream Profile_ProfileServer) error {
	os.RemoveAll("images")
	os.Mkdir("images", os.ModePerm)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		qrCode, _ := qr.Encode(req.Profile.AboutMe.Fullname, qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, 200, 200)
		path := filepath.Join("images", req.Filename)
		file, _ := os.Create(path)
		defer file.Close()
		png.Encode(file, qrCode)

		res := ProfileResponse{
			Profile:  req.Profile,
			Filename: req.Filename,
		}
		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}

	return nil
}
