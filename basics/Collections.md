# Collections

배열, 슬라이스, 맵은 여러 값을 묶어 다룰 때 사용하는 Go의 대표적인 내장 타입이다.

## 1. 배열(Array)

배열은 길이가 고정된 같은 타입 원소의 나열이다.

```go
var numbers [5]int
numbers[4] = 100
fmt.Println(numbers)
```

- `[5]int`
  길이가 5인 `int` 배열이라는 뜻이다.
- 배열의 인덱스는 0부터 시작한다.
- 값을 따로 넣지 않은 칸은 타입의 제로값을 가진다.
  - `int` 배열이면 `0`
  - `string` 배열이면 `""`
- 이때, 길이가 타입에 포함되므로 `[5]int`와 `[4]int`는 서로 다른 타입이다.

배열 전체 길이는 `len`으로 구할 수 있다.

```go
scores := [5]float64{98, 93, 77, 82, 83}

var total float64
for _, score := range scores {
	total += score
}

fmt.Println(total / float64(len(scores)))
```

- `range` : 배열을 순회할 때 사용한다.
- `_` : 현재 위치(index)가 필요 없을 때 버리는 값이다.
- `float64(len(scores))` : `len`은 `int`를 반환하므로 `float64`와 계산하려면 형변환이 필요하다.

현재 `Collections.go`에서는 배열 예제로 두 가지를 보여준다.

- `numbers`
  길이만 정하고 일부 값만 넣은 배열
- `scores`
  초기값을 한 번에 넣고 평균을 계산하는 배열

## 2. 슬라이스(Slice)

슬라이스는 배열 위에서 동작하는 더 유연한 컬렉션이다.
배열과 마찬가지로 인덱스로 접근할 수 있고 `len`을 가질 수 있지만,
배열과 달리 길이를 다루기 훨씬 편하다.

슬라이스는 다음과 같이 만들 수 있다.

```go
slice1 := []int{1, 2, 3}
slice2 := append(slice1, 4, 5)
```

- `[]int`
  길이를 쓰지 않고 타입만 적으면 슬라이스다.
- `append`
  기존 슬라이스 뒤에 값을 덧붙인 결과를 반환한다.

슬라이스를 미리 만들고 싶을 때는 `make`를 사용한다.

```go
slice4 := make([]int, 2)
```

- `make([]int, 2)`
  길이 2짜리 `int` 슬라이스를 만든다.
- `int`의 제로값은 `0`이므로 처음 상태는 `[0 0]`이다.

Go에는 슬라이스를 다룰 때 자주 쓰는 내장 함수가 있다.

- `append`
  기존 슬라이스 뒤에 값을 붙인다.
- `copy`
  한 슬라이스의 내용을 다른 슬라이스로 복사한다.

`copy`는 순서를 반드시 기억해야 한다.

```go
copy(dst, src)
```

- `dst` : 복사받는 대상
- `src` : 복사할 원본

현재 `Collections.go`의 코드는 아래와 같다.

```go
slice3 := []int{1, 2, 3}
slice4 := make([]int, 2)
copy(slice3, slice4)
fmt.Println("copy slice =", slice3, slice4)
```

이 코드에서 `slice4`는 `[0 0]`으로 시작하고, `copy(slice3, slice4)`는 `slice4`의 값을 `slice3`에 복사한다.
그래서 `slice3`의 앞 두 칸이 `0`으로 바뀌어 결과가 `[0 0 3]`처럼 보이게 된다.

즉, 이 예제는 "슬라이스 복사" 자체보다 `copy(dst, src)` 순서를 이해하는 데 초점을 두고 보면 된다.
만약 `slice3`의 내용을 `slice4`로 복사하고 싶다면 방향을 반대로 써야 한다.

```go
slice3 := []int{1, 2, 3}
slice4 := make([]int, len(slice3))
copy(slice4, slice3)
```

### 3. 맵(map)

맵은 순서가 없는 key-value의 집합이다.
맵은 연관 배열 또는 해시 테이블, 딕셔너리로,
연관 키를 통해 값을 찾는 데 사용한다.

맵 타입은 map 키워드에 이어 키 타입을 대괄호 안에,
값 타입을 지정해 표시한다.

```go
x := make(map[string]int)
x["key"] = 10
fmt.Println(x)
```

- `make(map[string]int)`
  key는 `string`, value는 `int`인 맵을 생성한다.
- `x["key"] = 10`
  `"key"`라는 key에 `10`을 저장한다.

맵의 특징
- 맵은 사용하기 전에 초기화가 필요하다.
- 맵은 길이가 가변적이라 새 항목을 추가할 때마다 바뀔 수 있다.
- 맵은 순서가 보장되지 않는다.
- 내장 `delete` 함수를 제공해 필요하면 항목 삭제가 가능하다.

맵에서 key 조회 시 자주 쓰는 패턴은 다음과 같다.

```go
if el, ok := elements["Li"]; ok {
	fmt.Println(el["name"], el["state"])
}
```

- `el`
  조회한 값
- `ok`
  해당 key가 실제로 존재하는지 나타내는 `bool`

이 패턴을 쓰면 key가 없을 때도 안전하게 처리할 수 있다.

예제의 `elements`는 `map[string]map[string]string` 형태의 중첩 맵이다.

- 바깥쪽 맵의 key
  원소 기호 (`"H"`, `"Li"` 등)
- 안쪽 맵의 key
  속성 이름 (`"name"`, `"state"`)

즉, `elements["Li"]["name"]`은 Lithium을 뜻하고, `elements["Li"]["state"]`는 solid를 뜻한다.

## 정리

- 배열은 길이가 고정된 컬렉션이다.
- 슬라이스는 배열보다 유연하며 `append`, `copy`, `make`와 함께 자주 사용된다.
- `copy`는 반드시 `copy(dst, src)` 순서를 기억해야 한다.
- 맵은 key로 value를 찾는 자료구조이며, key 존재 여부는 `value, ok := map[key]` 형태로 자주 확인한다.
