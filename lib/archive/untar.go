package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func Untargz(archive, out string) {
	file, err := os.Open(archive)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		path := filepath.Join(out, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				panic(err)
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if err := os.MkdirAll(dir, 0755); err != nil {
				panic(err)
			}

			file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}

			if _, err := io.Copy(file, tarReader); err != nil {
				file.Close()
				panic(err)
			}
			file.Close()
		}
	}
}

func Untargz2(gzipPath, dest string) error {
	f, err := os.Open(gzipPath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzReader, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dest, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(target, os.ModePerm)
		case tar.TypeReg:
			os.MkdirAll(filepath.Dir(target), os.ModePerm)
			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			io.Copy(outFile, tarReader)
			outFile.Close()
		}
	}
	return nil
}
