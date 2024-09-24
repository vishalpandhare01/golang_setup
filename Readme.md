### start golang project
- go mod init github.com/githubUserName/gitHubProjectName
- go get github.com/gofiber/fiber/v2
- go get github.com/gin-gonic/gin
- go get -u gorm.io/gorm
- github.com/spf13/viper
- go get github.com/google/uuid
- go get github.com/lib/pq
- go get github.com/joho/godotenv
- go get github.com/adhtanjung/go_rest_api

### import golang private repository by ssh
write command in terminal:-
- GOPRIVATE=github.com/xyz
- export GOPRIVATE
- go env -w GOPRIVATE=github.com/xyz
- go env GOPRIVATE


- git config --global .url."enter here ssh url".insteadOf "https://github.com/cowhite/reponame"
- cat ~/.gitconfig
