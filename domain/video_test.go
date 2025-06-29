package domain_test

import (
	"encoder/domain"
	"testing"
	"time"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoValidation(t *testing.T){
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.ResourceID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}

func TestValidateIfVideoIsEmpty(t *testing.T){ 
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAnUUID(t *testing.T){
	video := domain.NewVideo()
	video.ID = "123"
	video.ResourceID = "123"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}	