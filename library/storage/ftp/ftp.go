package ftp

import (
	"context"
	"mime/multipart"

	"esb-test/library/logger"

	"github.com/jlaffaye/ftp"
)

type FTPConfig struct {
	Host            string
	Username        string
	Password        string
	BaseDomainImage string
	RootDirectory   string
}

type FileTransferProtocol struct {
	cfg FTPConfig
}

type FTPClient interface {
	UploadToFTP(ctx context.Context, fileHeader *multipart.FileHeader) (fileDomain string, err error)
	DeleteFileFTP(ctx context.Context, fileName string) (err error)
}

func InitFTP(cfg FTPConfig) FTPClient {
	client := &FileTransferProtocol{}
	client.cfg = cfg
	return client
}

func (f *FileTransferProtocol) UploadToFTP(ctx context.Context, fileHeader *multipart.FileHeader) (fileDomain string, err error) {
	// Open file
	file, err := fileHeader.Open()
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to open file: %v", err)
		return
	}
	defer file.Close()

	// Dial to ftp server
	c, err := ftp.Dial(f.cfg.Host)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to dial ftp server: %v", err)
		return
	}
	defer c.Quit()

	// Login to ftp server
	err = c.Login(f.cfg.Username, f.cfg.Password)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to login to ftp server: %v", err)
		return
	}

	// Change Directory to root directory
	err = c.ChangeDir(f.cfg.RootDirectory)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to change directory: %v", err)
		return
	}

	// Upload the file to the FTP server
	err = c.Stor(fileHeader.Filename, file)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to upload file to ftp server: %v", err)
		return
	}

	// Set file domain
	fileDomain = f.cfg.BaseDomainImage + fileHeader.Filename
	return
}

func (f *FileTransferProtocol) DeleteFileFTP(ctx context.Context, pathToFile string) (err error) {
	// Dial to ftp server
	c, err := ftp.Dial(f.cfg.Host)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to dial ftp server: %v", err)
		return
	}
	defer c.Quit()

	// Login to ftp server
	err = c.Login(f.cfg.Username, f.cfg.Password)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to login to ftp server: %v", err)
		return
	}

	// delete the file to the FTP server
	err = c.Delete(pathToFile)
	if err != nil {
		logger.GetLogger(ctx).Errorf("failed to delete file in path: %v , err : %v", pathToFile, err)
		return
	}

	return
}
