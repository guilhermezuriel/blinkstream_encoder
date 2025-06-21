package domain

import (
	"time"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Job struct {
	ID string  `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string  `json:"output_bucket_path" valid:"notnull"`
	Status string  `json:"status" valid:"notnull"`
	Video *Video  `json:"video" valid:"-"`
	VideoID string `json:"-" valid:"-" gorm:"type:uuid;column:video_id;notnull"`
	Error string  `json:"error" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func(job *Job) prepare() error {
	job.ID = uuid.New().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	return nil
}

func init(){
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status: status,
		Video: video,
	}
	job.prepare()

	err := job.Validate()
	if err != nil{
		return nil, err
	}
	return &job, nil
}

func (j *Job) Validate() error {	
	_, err := govalidator.ValidateStruct(j)
	if err != nil{
		return err
	}
	return nil
}
