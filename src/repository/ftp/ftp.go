package ftp

import (
	"context"
	"mime/multipart"

	"esb-test/library/storage/ftp"
)

type FtpRepository struct {
	ftpClient ftp.FTPClient
}

func InitFTPRepository(ctx context.Context, ftp ftp.FTPClient) (*FtpRepository, error) {
	return &FtpRepository{
		ftpClient: ftp,
	}, nil
}

func (f *FtpRepository) UploadToFTP(ctx context.Context, fileHeader *multipart.FileHeader) (fileDomain string, err error) {
	return f.ftpClient.UploadToFTP(ctx, fileHeader)
}

func (f *FtpRepository) DeleteFileFTP(ctx context.Context, pathToFile string) (err error) {
	return f.ftpClient.DeleteFileFTP(ctx, pathToFile)
}
