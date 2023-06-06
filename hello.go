import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)
creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{}, // Возможность брать ключи из переменных окружения
			&credentials.SharedCredentialsProvider{
				Filename: path.Join(os.Getenv("HOME"), ".aws/credentials") // Путь до credentials
				// Profile:  "<profile>",
			},
		},
	)
	sess := session.Must(session.NewSession(&aws.Config{
		Region:                         aws.String("ru-msk"), // Указываем регион
		Endpoint:                       aws.String("hb.vkcs.cloud"), // Указываем endpoint
		Credentials:                    creds, // Указываем credentials
		S3ForcePathStyle:               aws.Bool(true),
	}))
client:= s3.New(sess) // Создаем s3 клиент
