package environments

import "github.com/joho/godotenv"

func Set() {
	godotenv.Load("./././.env")
}
