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
│   └── publicpkg/
│       └── publicpkg.go
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
mkdir -p pkg/publicpkg
mkdir -p pkg/exinit
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

import (
    "text/template"
	htemplate "html/template"
)

func main() {
	template.New("foo").Parse(`{{define "T"}}Hello`)
	htemplate.New("foo").Parse(`{{define "T"}}Hello`)
}
```

이제 `make run` 명령을 사용하면 각 패키지에 따른 함수가 호출됩니다.

```bash
make run
```

### 5. 패키지명과 패키지 외부 공개

```go
// pkg/publicpkg/publicpkg.go
package publicpkg

import "fmt"

const (
	PI = 3.1415	// 공개되는 상수
	pi = 3.141516	// 공개되지 않는 상수
)

var ScreenSize int = 1080	// 공개되는 변수
var screenHeight int	// 공개되지 않는 변수

func PublicFunc() {	// 공개되는 함수
	const MyConst = 100	// 공개되지 않습니다.
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() { // 공개되지 않는 함수
	fmt.Println("This is a ㅔ걒ㅁㅅㄷ function")
}

type MyInt int	// 공개되는 별칭 타입
type myString string	// 공개되지 않는 별치 타입

type MyStruct struct {	// 공개되는 구조체
	Age	MyInt	// 공개되는 구조체 필드
	name string	// 공개되지 않는 구조체 필드
}

func (m MyStruct) PublicMethod() {	// 공개되는 메서드
	fmt.Println("This is a public method")
}

func (m MyStruct) privateMethod() {	// 공개되지 않는 메서드
	fmt.Println("This is a private method")
}

type myPrivateStruct struct {	// 공개되지 않는 구조체
	Age	MyInt	// 공개되지 않는 구조체 필드
	name string	// 공개되지 않는 구조체 필드
}

func (m myPrivateStruct) PrivateMethod() {	// 공개되지 않는 메서드
	fmt.Println("This is a private method")
}
```

```go
// cmd/golang-package/main.go
package main

import (
    "fmt"
	"pkg/publicpkg"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)	// 공개되는 상수 접근
	publicpkg.PublicFunc()		// 공개되는 함수 호출

	var myint publicpkg.MyInt = 10	// 공개되는 별칭 타입 사용
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}	// 구조체 사용
	fmt.Println("mystruct:", mystruct)
}
```

이제 `make run` 명령을 사용하면 패키지 공개 변수, 함수를 접근해서 출력됩니다.

```bash
make run
```


### 6. 패키지 초기화

```go
// pkg/exinit/exinit.go
package exinit

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init function", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}
```

```go
// cmd/golang-package/main.go
package main

import (
    "fmt"
	"golang-package/pkg/exinit"
)

func main() {
	fmt.Println("main function")
	exinit.PrintD()
}
```

이제 `make run` 명령을 사용하면 패키지의 초기화 흐름이 출력됩니다.

```bash
make run
```

패키지를 임포트하면 패키지 초기화가 시작되는데,
이때 먼저 패키지의 모든 전역 변수들이 초기화되고, 그다음에 init() 함수가 호출됩니다.