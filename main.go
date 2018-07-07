package main

import (
	"archive/zip"
	logger "github.com/cihub/seelog"
	"gopkg.in/urfave/cli.v1"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	swaggerVersion = "2"
	swaggerlink    = "https://github.com/beego/swagger/archive/v" + swaggerVersion + ".zip"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "router_path",
			Value: "./main.go",
			Usage: "router define path",
		},
		cli.StringFlag{
			Name:  "output_path",
			Value: "./",
			Usage: "output swagger path",
		},
		cli.StringFlag{
			Name:  "downdoc",
			Value: "false",
			Usage: "Enable auto-download of the swagger file if it does not exist",
		},
	}

	app.Action = func(c *cli.Context) {
		routerPath := c.String("router_path")
		outputPath := c.String("output_path")
		downdoc := c.String("downdoc")
		currpath, _ := os.Getwd()

		if downdoc == "true" {
			if _, err := os.Stat(path.Join(currpath, "swagger", "index.html")); err != nil {
				if os.IsNotExist(err) {
					downloadFromURL(swaggerlink, "swagger.zip")
					unzipAndDelete("swagger.zip")
				}
			}
		}

		generateDocs(routerPath, outputPath)
		logger.Info("'swagger.json' 'swagger.yml' already gen!")
	}

	app.Run(os.Args)
}

func downloadFromURL(url, fileName string) {
	var down bool
	if fd, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		down = true
	} else if fd.Size() == int64(0) {
		down = true
	} else {
		logger.Infof("'%s' already exists", fileName)
		return
	}
	if down {
		logger.Infof("Downloading '%s' to '%s'...", url, fileName)
		output, err := os.Create(fileName)
		if err != nil {
			logger.Errorf("Error while creating '%s': %s", fileName, err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			logger.Errorf("Error while downloading '%s': %s", url, err)
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			logger.Errorf("Error while downloading '%s': %s", url, err)
			return
		}
		logger.Infof("%d bytes downloaded!", n)
	}
}

func unzipAndDelete(src string) error {
	logger.Infof("Unzipping '%s'...", src)
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	rp := strings.NewReplacer("swagger-"+swaggerVersion, "swagger")
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fname := rp.Replace(f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fname, f.Mode())
		} else {
			f, err := os.OpenFile(
				fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	logger.Infof("Done! Deleting '%s'...", src)
	return os.RemoveAll(src)
}
