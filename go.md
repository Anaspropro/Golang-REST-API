
Go dependency installation

go get -u github.com/gin-gonic/gin
go get -u github.com/jackc/pgx/v5
go get -u github.com/jackc/pgx/v5/pgxpool
go get -u github.com/golang-jwt/jwt/v5
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/joho/godotenv
go install github.com/cosmtrek/air@latest
