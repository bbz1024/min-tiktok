package config

import "fmt"

type UserInfoStructure struct {
	SignatureUrl string
	AvatarUrl    string
	BackImageUrl string
}
type MysqlStructure struct {
	DataSource string
}
type QiNiuStructure struct {
	AccessKey   string
	SecretKey   string
	VideoDomain string
	Bucket      string
}
type RabbitMQStructure struct {
	Host  string
	Port  int
	User  string
	Pass  string
	VHost string
}

func (r *RabbitMQStructure) Dns() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		r.User,
		r.Pass,
		r.Host,
		r.Port,
		r.VHost,
	)
}

type AlibabaNslStructure struct {
	AccessKeyId     string
	AccessKeySecret string
	AppKey          string
}
type GptStructure struct {
	ApiKey  string
	ModelID string
}
type GorseStructure struct {
	GorseAddr   string
	GorseApikey string
}
