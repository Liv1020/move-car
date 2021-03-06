package components

// Config Config
type Config struct {
	Mode string `required:"true"`
	Port int    `required:"true"`
	DB   struct {
		Host        string `required:"true"`
		Port        int    `required:"true"`
		User        string `required:"true"`
		Password    string `required:"true"`
		Database    string `required:"true"`
		Charset     string `default:"utf8mb4"`
		MaxIdle     int    `default:"20"`
		MaxOpen     int    `default:"100"`
		MaxLifetime int    `default:"25200"`
	}
	Wechat struct {
		AppID          string `required:"true"`
		AppSecret      string `required:"true"`
		TemplateID     string `required:"true"`
		ConfirmUrl     string `required:"true"`
		Token          string `required:"true"`
		EncodingAesKey string `required:"true"`
	}
	Aliyun struct {
		AccessKeyId     string `required:"true"`
		AccessKeySecret string `required:"true"`
		Vms             struct {
			CalledShowNumber string `required:"true"`
			TtsCode          string `required:"true"`
		}
		Sms struct {
			SignName     string `required:"true"`
			TemplateCode string `required:"true"`
		}
	}
}
