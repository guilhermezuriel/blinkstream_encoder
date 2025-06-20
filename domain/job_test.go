package domain_test

import (
	"encoder/domain"
	"testing"
	"time"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJobValidation(t *testing.T){
	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()


	job, err := domain.NewJob("path", "pending", video)
	
	require.NotNil(t, job)
	require.Nil(t, err)
}

func TestJobValidateIfBucketIsEmpty(t *testing.T){ 
	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("", "pending", video)

	require.Nil(t, job)
	require.Error(t, err)
}

func TestJobValidateIfStatusIsEmpty(t *testing.T){
	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "", video)
	
	require.Nil(t, job)
	require.Error(t, err)
}
