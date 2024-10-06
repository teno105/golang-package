아래는 실습 순서에 맞춰 다시 작성한 `README.md`입니다.

---

# golang-package

`golang-package`는 Golang으로 작성된 간단한 애플리케이션으로, 포인터의 사용을 익히기 위한 실습입니다.


## 프로젝트 구조

```plaintext
golang-package/
│
├── cmd/
│   └── golang-package/
│       └── main.go
│
├── pkg/
│   └── custompkg/
│       └── custompkg.go
│
├── go.mod
├── Makefile
└── README.md
```

## 실습 순서

### 1. 패키지 구조를 위한 디렉토리 생성

먼저 프로젝트 디렉터리를 설정하고 필요한 디렉터리들을 생성합니다.

```bash
mkdir golang-package
cd golang-package
go mod init golang-package

mkdir -p cmd/golang-package
mkdir -p pkg/custompkg
```

### 2. 패키지 사용하기

`cmd/golang-package/` 디렉터리 아래에 `main.go` 파일을 생성하고,
struct 를 선언 및 활용하는 코드를 작성합니다.

```go
// cmd/golang-package/main.go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
	fmt.Println(rand.Int()) // 랜덤한 숫자값을 출력합니다.
}
```

이 코드를 통해 프로그램을 실행하면 rand.Int()의 값이 출력됩니다.

### 3. `Makefile` 작성

이제 프로젝트의 빌드 및 실행을 자동화하기 위한 `Makefile`을 프로젝트 루트에 작성합니다.

```makefile
# Go 관련 변수 설정
APP_NAME := golang-package
CMD_DIR := ./cmd/golang-package
BUILD_DIR := ./build

.PHONY: all clean build run test fmt vet install

all: build

# 빌드 명령어
build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

# 실행 명령어
run: build
	@echo "==> Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# 코드 포맷팅
fmt:
	@echo "==> Formatting code..."
	go fmt ./...

# 코드 분석
vet:
	@echo "==> Running go vet..."
	go vet ./...

# 의존성 설치
install:
	@echo "==> Installing dependencies..."
	go mod tidy

# 테스트 실행
test:
	@echo "==> Running tests..."
	go test -v ./...

# 빌드 정리
clean:
	@echo "==> Cleaning build directory..."
	rm -rf $(BUILD_DIR)
```

`Makefile`을 이용하여 코드를 빌드하고 실행할 수 있습니다.

```bash
make run
```

이 명령어를 통해 `main.go`에서 작성한 struct 정보를 확인할 수 있습니다.

### 4. 겹치는 패키지 문제 별칭으로 풀기

패키지명이 겹칠때는 별칭(aliasing)을 줘서 구별해줍니다.

```go
// cmd/golang-package/main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

이제 `make run` 명령을 사용하면 각 패키지에 따른 함수가 호출됩니다.

```bash
make run
```

### 5. 패키지명과 패키지 외부 공개


### 6. 패키지 초기화

