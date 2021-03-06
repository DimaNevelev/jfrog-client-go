package utils

import (
	"github.com/jfrog/jfrog-client-go/utils/io/fileutils"
	"github.com/jfrog/jfrog-client-go/artifactory/buildinfo"
)

type FileHashes struct {
	Sha256 string `json:"sha256,omitempty"`
	Sha1   string `json:"sha1,omitempty"`
	Md5    string `json:"md5,omitempty"`
}

type FileInfo struct {
	*FileHashes
	LocalPath       string `json:"localPath,omitempty"`
	ArtifactoryPath string `json:"artifactoryPath,omitempty"`
}

func (fileInfo *FileInfo) ToBuildArtifacts() buildinfo.Artifact {
	artifact := buildinfo.Artifact{Checksum: &buildinfo.Checksum{}}
	artifact.Sha1 = fileInfo.Sha1
	artifact.Md5 = fileInfo.Md5
	filename, _ := fileutils.GetFileAndDirFromPath(fileInfo.LocalPath)
	artifact.Name = filename
	return artifact
}
