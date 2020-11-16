# showboat
Portfolio page creator

![Go](https://github.com/thealamu/showboat/workflows/Go/badge.svg)

You build amazing stuff a lot, this app helps you create a portfolio page to showcase them.

## Docker Deploy :rocket:
To deploy on docker you can easily clone and build the image.
Two environment variables are required: 
 - FRONTEND - which is a http link to a deployment of the [frontend app](https://github.com/thealamu/showboat-ui) 
 - HMACSECRET - a secret password used to sign [jwts](jwt.io)

### Build Image
```shell
git clone github.com/thealamu/showboat
cd showboat
docker build -t showboat .
```

### Start the app
```shell
docker run -e FRONTEND={path_to_your_frontend} -e HMACSECRET={any_secret_password} -dp 8080:8080 showboat
```
