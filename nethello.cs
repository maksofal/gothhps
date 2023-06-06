AmazonS3Config configsS3 = new AmazonS3Config {
    ServiceURL = "https://s3.yandexcloud.net"
};

AmazonS3Client s3client = new AmazonS3Client(configsS3);
