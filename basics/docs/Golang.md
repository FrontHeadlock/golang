# Go 문법

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

### 0. Go 언어 특징

- 클래스 X , 상속 X 
- 메서드 O , 인터페이스 O , 익명 함수 O , 가비지 컬렉터 O, 포인터 O
- 코드 실행 과정
  - 폴더 생성 -> .go 파일 생성 -> go 모듈 생성 -> 빌드 -> 실행
  - 폴더생성 : 모든 코드가 패키지 단위로 작성 (같은 폴더면 같은 패키지에 포함)
  - Go 모듈 생성 : go mod init 으로 생성
  - 빌드 : go build 통해 실행 파일 생성

### 1. `package main`

```go
package main
```

모든 Go 프로그램은 패키지 선언(package declaration)으로 시작

패키지 선언은 이 코드가 어떤 패키지에 속하는 지 알려줌
패키지는 코드 묶음이고, 여러 기능을 제공 (Go의 모든 코드는 패키지 선언으로 시작)

패키지는 Go 언어에서 코드를 조직화하고 재사용하는 수단으로,
Go 프로그램에는 크게 두 가지 유형이 존재

- 실행 프로그램
- 라이브러리

실행 프로그램은 터미널에서 바로 실행할 수 있는 프로그램 (윈도우에서는 실행 파일 이름이 `.exe`로 끝난다.)
라이브러리는 다른 프로그램에서 사용할 수 있도록 묶어 놓은 코드의 모음


### 2. `import "fmt"`

```go
import "fmt"
```

fmt는 `format`의 약자로, 입력과 출력에 관련된 기능을 제공하는 표준 패키지

`fmt`가 큰따옴표로 둘러싸여 있다는 점도 중요하다.
이런 형태는 `문자열 리터럴(string literal)`이라고 하며, 수식에 해당한다.
Go에서 문자열은 문자, 숫자, 기호 등을 순서대로 나열한 값을 뜻한다.

지금 단계에서 기억해야 할 핵심은 다음과 같다.

- 문자열은 여는 큰따옴표 `"`로 시작하면 반드시 닫는 큰따옴표 `"`로 끝나야 한다.
- 두 큰따옴표 사이의 내용이 문자열의 값이 된다.
- 큰따옴표 자체는 문자열의 일부가 아니다.

지금은 `"fmt"` 역시 Go 문법 안에서 하나의 문자열이라는 점만 이해하면 충분하다.

### 3. 주석(comment)

주석은 `//`로 시작한다.
주석이 붙어있다면 Go 컴파일러가 무시해 프로그램 실행에 영향을 주지 않는다.

Go에서는 다음과 같은 주석을 지원한다
- `//`
- `/* */`

## 4. 함수 선언과 `main`

주석 다음으로 보게 되는 중요한 문법은 함수 선언이다.

```go
func main() {
	fmt.Println("Hello World")
}
```

`func` : 함수를 선언할 때 사용하는 키워드  
`main` : 함수 이름

Go에서 실행 프로그램은 `main` 패키지 안의 `main` 함수를 시작점으로 실행 (즉, 프로그램이 시작될 때 가장 먼저 실행되는 함수)

```go
fmt.Println("Hello World")
```

1. `fmt` 패키지에 접근
2. `Println` 함수 사용
3. `"Hello World"`라는 문자열을 인자로 전달

즉, `fmt.Println`은 `fmt` 패키지에 들어 있는 `Println` 함수를 뜻한다.

이 프로그램에서는 `Println` 함수가 실제 작업을 수행한다.  
그래서 프로그램을 실행하면 터미널에 `Hello World`가 표시된다.

### 5. `godoc fmt Println`

터미널에서 아래 명령을 입력하면 `Println` 함수에 대한 설명을 확인할 수 있다.

```go
godoc fmt Println
```

Go는 문서화가 매우 잘 되어 있는 언어로, doc 통해 빠르게 학습할 수 있다.

- `Println`은 전달받은 값을 기본 형식으로 출력한다.
- 값과 값 사이에는 공백이 들어간다.
- 마지막에는 줄바꿈이 추가된다.

- 출력한 바이트 수와 오류 정보를 반환할 수 있다.

`godoc` 명령을 사용시, Go 표준 라이브러리 통해 함수나 패키지를 빠르게 참고하면 유용


# Type

Go는 정적 타입 프로그래밍 언어로, 변수가 항상 특정 타입을 지니고 있고 해당 타입은 변경될 수 없다.

타입은 프로그램이 하려는 일을 추론하고, 실수를 잡아내는 데 도움된다.

### 1. 정수

- 소수부가 없는 숫자
- Go의 정수 타입으로는 uint8, uint16, uint32, uint64, int8, int16, int32, int64가 존재
  - 각 숫자는 타입이 사용하는 bit 수를 의미
  - uint : unsigned integer
  - int  : signed   integer 

- 부동 소수점 수
  - 소수부가 포함된 숫자
  - 정수와 달리 부동 소수점 수는 일정한 크기(32비트나 64비트)가 존재
  - Go에는 float32와 float64라는 두 가지 부동 소수점 타입과 complex64와 complex128이라고 하는 복소수 타입이 존재

```go
// 명시적 선언
var a int = 10
// 타입 추론 선언 (20을 보고 b의 타입을 int로 추론)
// := 는 함수 내부에서만 사용되며, 컴파일 시점에 타입이 결정됨
b := 20
```


### 2. 문자열

문자열은 텍스트를 표현하는 데 사용되는 문자의 나열이다.
Go의 문자열(`string`)은 바이트 시퀀스로 다뤄진다.
문자열 값은 생성 후 변경할 수 없다(immutable).

### 3. Boolean

Boolean(`bool`)은 참(`true`)과 거짓(`false`)을 나타내는 독립 타입이다.
Go에서 `bool`은 정수 타입이 아니며, 논리 연산 전용으로 사용한다.

# Variables

변수는 값을 저장하는 공간이며, 이름과 타입을 가진다.

```go
// 변수를 이용한 Hello World 개선
func main() {
    var message string = "Hello World"
    fmt.Println(message)
}
```

Go의 변수는 `var` 키워드를 이용해 선언할 수 있다.
`var message string = "Hello World"`는 다음 순서로 읽으면 된다.

1. `var` : 변수 선언
2. `message` : 변수 이름
3. `string` : 변수 타입
4. `"Hello World"` : 저장할 값

변수는 프로그램 실행 중 값을 다시 할당할 수 있다.

```go
func main() {
    var text string
    text = "first"
    fmt.Println(text)

    text = "second"
    fmt.Println(text)
}
```

또한 기존 값을 이용해 새로운 값을 만들 수도 있다.

```go
func main() {
    var text string
    text = "first "
    fmt.Println(text)

    text = text + "second"
    fmt.Println(text)
}
```

`=`는 오른쪽 식을 먼저 계산한 뒤, 그 결과를 왼쪽 변수에 저장한다.

Go에서는 `:=`를 활용하여 문장을 더 짧게 작성할 수 있다.
`:=`는 선언과 초기화, 그리고 타입 추론을 한 번에 처리한다.

1. 새 변수 선언 
2. 초기값 할당 
3. 타입 추론

위 3가지를 한 번에 처리한다.


Go 컴파일러가 변수에 할당하는 리터럴 값을 토대로 타입을 추론할 수 있기 때문에 타입은 필요하지 않다. 
(문자열 리터럴을 할당하고 있으므로 `message`에는 `string` 타입이 지정된다) 
컴파일러는 var 문장에서도 타입을 추론할 수 있다.

```go
message := "Hello World"
var otherMessage = "Hello World"

count := 5
fmt.Println(count)
```

`:=`는 함수 내부에서만 사용할 수 있으며, 새 변수를 만들 때 사용한다.
반면 `var`는 함수 내부와 패키지 레벨 모두에서 사용할 수 있다.

### 피드백 및 구체화

최근 작성한 `Variables.go`를 기준으로 보면, 입력을 먼저 받아야 하는 변수는 `var`, 계산 결과를 저장하는 변수는 `:=`로 두는 흐름이 자연스럽다.

```go
fmt.Print("화씨를 입력하시오: ")
var fahrenheit float64
fmt.Scanf("%f", &fahrenheit)

celsius := (fahrenheit - 32) * 5 / 9
fmt.Printf("섭씨: %gC\n", celsius)
```

- `var fahrenheit float64`
  아직 값은 없지만, 사용자 입력을 저장할 공간이 먼저 필요하기 때문에 `var`가 적합하다.
- `fmt.Scanf("%f", &fahrenheit)`
  `%f` 형식으로 입력을 읽고, `&fahrenheit`를 통해 해당 변수의 주소에 값을 저장한다.
- `celsius := ...`
  계산 결과를 새 변수에 바로 저장하므로 `:=`가 적합하다.
- `fmt.Printf("미터: %gm\n", meters)`
  숫자 뒤에 단위를 붙여 출력하고 싶을 때 사용하는 방식이다. 파이썬의 f-string과 목적은 비슷하지만, Go에서는 서식 지정자를 사용한다.

추가로 최근 예제에서는 `x`, `y`, `m`보다 `fahrenheit`, `celsius`, `feet`, `meters`처럼 의미가 드러나는 이름이 더 적절하다.
학습용 코드일수록 "짧은 이름"보다 "의도가 보이는 이름"이 낫다.

정리하면 다음처럼 구분하면 된다.

- `var`
  값을 나중에 넣을 예정이거나, 제로값 상태로 먼저 선언해야 할 때 사용
- `:=`
  함수 내부에서 새 변수를 선언하면서 바로 초기화할 때 사용
- `const`
  절대 바뀌지 않는 값을 선언할 때 사용

### 1. 변수의 이름을 짓는 법

변수 네이밍은 해당 변수를 보고, 개발자의 의도를 파악할 수 있어야 한다.

```go
// x가 무슨 역할인지 파악이 어려움
x := "Max"
fmt.Println("My dog's name is", x)

// 변수의 의미를 바로 파악하기 쉬움
name := "Max"
fmt.Println("My dog's name is", name)
```

Go 언어에서는 함수뿐 아니라 변수, 타입, 상수 같은 식별자도 첫 글자에 따라 가시성이 달라진다.

- 파스칼 케이스 (PascalCase): 
  이름이 대문자로 시작하는 경우 (예: UserName, PrintValue)
  다른 패키지에서도 접근 가능한 exported 식별자
- 카멜 케이스 (camelCase): 
  이름이 소문자로 시작하는 경우 (예: userName, printValue)
  같은 패키지 내부에서만 사용하는 unexported 식별자

### 2. 상수(constant)

기본적으로 상수는 불변이다. 
상수는 변수와 비슷하게 선언하지만, `var` 대신 `const` 키워드를 사용한다.
한 번 선언한 뒤에는 값을 변경할 수 없다.

```go
func main() {
    const greeting string = "Hello World"
    fmt.Println(greeting)

    // greeting = "Some other string" // 컴파일 오류
}
```

사용자 입력처럼 실행 중에 바뀌는 값은 `const`로 둘 수 없고, `var`로 선언해야 한다.

# Control Structure

### 종류
- For
- IF
- Switch
