# Collections

배열, 슬라이스, 맵은 여러 값을 묶어 다룰 때 사용하는 Go의 대표적인 내장 타입이다.

## 1. 배열(Array)

배열은 길이가 고정된 같은 타입 원소의 나열

```go
var numbers [5]int
numbers[4] = 100
fmt.Println(numbers)
```

- `[5]int`
  길이가 5인 `int` 배열이라는 뜻이다.
- 배열의 인덱스는 0부터 시작한다.
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
- `float64(len(scores))` : `len`은 `int`를 반환하므로 `float64`와 계산하려면 형변환이 필요

