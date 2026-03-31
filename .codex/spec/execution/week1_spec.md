# Go 8주 집중 학습 로드맵 - Week 1 실행 스펙형 Spec

## 1. 문서 목적

이 문서는 1주차를 실제 프로젝트 운영 문서처럼 정의한다.
목적은 "무엇을 공부할까"가 아니라 "이번 주가 끝나면 어떤 결과물이 남아 있어야 하는가"를 명확히 하는 데 있다.

이 스펙은 아래 질문에 답해야 한다.

- 이번 주의 미션은 무엇인가
- 매일 어떤 실행 항목을 완료해야 하는가
- 주차 종료 시 어떤 코드와 파일 구조가 존재해야 하는가
- 이후 주차 로드맵에서 1주차 결과물이 어떤 위치를 차지하는가

## 2. 대상 학습자와 전제

이 스펙은 아래 학습자를 기준으로 한다.

- 기본 프로그래밍 개념은 알고 있지만 Go는 처음이다.
- Java, Python, JavaScript 중 하나 이상에 익숙하다.
- 작은 CLI 프로그램과 파일 입출력은 이해하고 있다.
- 1주차 결과물을 이후 주차 프로젝트의 기반으로 활용할 계획이다.

## 3. 주차 미션

이번 주의 미션은 아래 한 문장으로 정의한다.

`JSON 기반 작업 명세를 읽고 검증하여 출력하는 Go CLI 프로그램 jobspec-cli를 직접 구현한다.`

이 결과물은 단순 문법 예제가 아니라 이후 주차 확장을 위한 최소 실행 가능한 베이스라인이다.

## 4. 주차 종료 시 최종 산출물

### 4.1 실행 가능한 CLI

- 프로그램명: `jobspec-cli`
- 실행 예시: `go run . ./examples/valid-job.json`
- 동작:
  - JSON 파일 경로를 입력받는다.
  - 파일을 읽는다.
  - `JobSpec` 구조체로 파싱한다.
  - 필수값을 검증한다.
  - 정상 입력이면 사람이 읽기 쉬운 형식으로 출력한다.
  - 실패 입력이면 원인에 맞는 에러 메시지를 출력하고 종료한다.

### 4.2 샘플 입력 파일

- 정상 입력 예시 최소 1개
- 실패 입력 예시 최소 2개

### 4.3 역할 분리된 소스코드

- main 진입점과 도메인 로직이 분리되어 있어야 한다.
- 읽기, 파싱, 검증, 출력 책임이 분리되어 있어야 한다.

### 4.4 학습 검증 기록

- Day 1~Day 7 진행 여부를 체크한 주간 실행표
- 최소 한 번 이상 작성한 짧은 회고 메모

## 5. 기능 요구사항

### 5.1 입력

- CLI 인자로 JSON 파일 경로를 입력받는다.
- 인자가 없으면 사용법 또는 입력 누락 에러를 출력한다.
- 파일이 존재하지 않으면 파일 읽기 실패 에러를 반환한다.

### 5.2 입력 JSON 스키마

1주차 기준 입력 스키마는 아래와 같다.

```json
{
  "name": "build-app",
  "command": "go build",
  "workdir": "./sample",
  "retry": 2,
  "enabled": true
}
```

### 5.3 구조체 요구사항

`JobSpec`는 최소 아래 필드를 가진다.

- `name string`
- `command string`
- `workdir string`
- `retry int`
- `enabled bool`

### 5.4 검증 규칙

- `name`은 빈 문자열이면 안 된다.
- `command`는 빈 문자열이면 안 된다.
- `retry`는 음수면 안 된다.
- `workdir`는 1주차에서는 선택값으로 둔다.
- `enabled`는 bool로 정상 파싱되어야 한다.

### 5.5 출력 규칙

정상 입력일 때는 사람이 읽기 쉬운 형식으로 출력한다.

예시:

```text
Job Name   : build-app
Command    : go build
Workdir    : ./sample
Retry      : 2
Enabled    : true
```

### 5.6 에러 처리 규칙

다음 경우 각각 다른 맥락의 에러 메시지를 출력해야 한다.

- 인자가 없는 경우
- 파일을 읽을 수 없는 경우
- JSON 형식이 잘못된 경우
- 필수값 검증에 실패한 경우

## 6. 비기능 요구사항

- 코드는 한 파일에 몰아넣지 않는다.
- 하나의 함수가 입력 처리, 검증, 출력까지 모두 담당하지 않게 분리한다.
- 정상 흐름의 에러 처리에 `panic`을 사용하지 않는다.
- 표준 라이브러리 중심으로 구현한다.
- 코드 길이가 짧아도 역할 분리가 드러나야 한다.

## 7. 일별 실행 계획

### Day 1 - 실행 구조와 문법 확인

**실행 항목**

- Go 설치 및 실행 확인
- `go version` 확인
- `hello.go`, `variable.go`, `zero.go` 작성
- zero value와 unused variable 에러 확인

**필수 산출물**

- 실행 가능한 `hello.go`
- 변수 선언 예제 파일
- zero value 확인 예제

**완료 판정**

- 빈 파일에서 `main` 함수 포함 프로그램을 직접 만들 수 있다.
- `var`, `:=`, `const`를 각각 한 번 이상 사용했다.

### Day 2 - 함수와 제어 흐름 작성

**실행 항목**

- 함수와 제어문 예제 작성
- 다중 반환값 함수 작성
- 에러 반환 함수 작성
- 간단한 계산기 CLI 작성

**필수 산출물**

- 다중 반환값 함수 예제
- 에러 반환 함수 예제
- 분기와 반복을 포함한 계산기 예제

**완료 판정**

- `if err != nil`을 포함한 함수를 최소 1개 직접 작성했다.
- `switch`와 `for`를 모두 사용한 예제가 있다.

### Day 3 - 자료구조 감각 확보

**실행 항목**

- array, slice, map 예제 작성
- `append`, `range`, `len`, `cap` 실험
- 문자열과 rune 처리 실험

**필수 산출물**

- slice 성장 실험 코드
- `map[string]int` 카운팅 예제
- 문자열 길이 비교 예제

**완료 판정**

- `map[string]int` 카운팅 예제를 직접 만들었다.
- 문자열 길이와 rune 순회 차이를 출력으로 확인했다.

### Day 4 - 데이터 모델과 메서드 구성

**실행 항목**

- `Task` struct 정의
- method 작성
- pointer receiver 실험
- 상태 변경 메서드 작성

**필수 산출물**

- `Task` 구조체 예제
- value receiver / pointer receiver 비교 예제
- 상태 변경 메서드 예제

**완료 판정**

- 상태 변경 메서드와 조회 메서드를 구분해서 작성했다.
- value receiver와 pointer receiver 차이를 설명할 수 있다.

### Day 5 - 추상화와 에러 처리 패턴 적용

**실행 항목**

- `Runner` interface 정의
- 구현체 작성
- `errors.New`, `fmt.Errorf`, `defer` 예제 작성

**필수 산출물**

- interface와 구현체 예제
- 에러 생성/감싸기 예제
- `defer` 적용 예제

**완료 판정**

- interface를 만족하는 타입을 별도 선언 없이 구현했다.
- 에러 생성과 감싸기를 각각 한 번 이상 사용했다.

### Day 6 - 파일 입력과 JSON 파싱 조립

**실행 항목**

- `go mod init` 수행
- `JobSpec` 구조체 분리
- 파일 읽기 및 JSON 파싱 구현
- 입력 오류 케이스 처리

**필수 산출물**

- `go.mod`
- `JobSpec` 구조체 정의 파일
- JSON 파일 읽기/파싱 코드
- 잘못된 입력 처리 예제

**완료 판정**

- JSON 파일을 struct로 변환하는 코드가 동작한다.
- main 외 파일 또는 패키지로 역할이 분리되어 있다.

### Day 7 - 미니 프로젝트 통합 및 마감

**실행 항목**

- `jobspec-cli` 통합 완성
- 정상/실패 입력 예시 준비
- 출력 형식 정리
- 코드 구조 최종 정리

**필수 산출물**

- 실행 가능한 `jobspec-cli`
- 정상 입력 예시 JSON
- 실패 입력 예시 JSON 2개 이상
- 읽기/검증/출력 분리된 소스코드

**완료 판정**

- 정상 입력은 사람이 읽기 쉬운 결과를 출력한다.
- 실패 입력은 원인을 구분 가능한 메시지로 보여준다.
- 주차 최종 산출물 구조를 모두 충족한다.

## 8. Definition of Done

1주차 완료는 아래 조건을 모두 만족할 때만 인정한다.

### 구현 완료

- `jobspec-cli`가 실제로 실행된다.
- JSON 입력을 읽고 파싱한다.
- 검증 로직이 포함되어 있다.
- 에러 처리가 명시적으로 구현되어 있다.

### 구조 완료

- 코드가 단일 파일에 몰려 있지 않다.
- main, 도메인 구조체, 파일 읽기, 검증 로직, 출력 로직이 역할 기준으로 분리되어 있다.

### 학습 완료

- 아래 항목을 직접 설명할 수 있다.
  - `var`와 `:=` 차이
  - array와 slice 차이
  - value receiver와 pointer receiver 차이
  - Go interface의 암묵적 구현 방식
  - Go에서 error를 반환값으로 다루는 이유

### 기록 완료

- 주간 체크리스트가 체크되어 있다.
- 최소 하루 1회 이상 짧은 학습 메모가 남아 있다.

## 9. Week 1 종료 시점의 최종 파일 아키텍처

1주차 종료 시점에는 아래 정도 구조가 권장된다.

```text
jobspec-cli/
  go.mod
  main.go
  examples/
    valid-job.json
    invalid-job-empty-name.json
    invalid-job-negative-retry.json
  jobspec/
    spec.go
    reader.go
    validate.go
    printer.go
```

### 파일 역할

- `main.go`
  - CLI 인자 처리
  - 프로그램 진입점
  - 에러 출력 및 종료 처리
- `jobspec/spec.go`
  - `JobSpec` 구조체 정의
- `jobspec/reader.go`
  - 파일 읽기 및 JSON 파싱
- `jobspec/validate.go`
  - 입력값 검증 로직
- `jobspec/printer.go`
  - 사람이 읽기 쉬운 출력 형식 생성
- `examples/*.json`
  - 정상/실패 입력 예시

## 10. 전체 로드맵 관점에서 본 1주차 결과물

1주차 결과물은 단순 예제가 아니라 이후 8주 로드맵의 출발점이다.

```text
study/
  .codex/
    spec/
      week1_learning_guide_spec.md
      week1_execution_spec.md
      week2_spec.md
      ...
  jobspec-cli/                # Week 1 결과물
    go.mod
    main.go
    examples/
    jobspec/
  task-runner/                # Week 2~3 예상 확장
  process-executor/           # Week 4~5 예상 확장
  cicd-simulator/             # Week 6~8 예상 확장
```

1주차는 아래 기반을 만든다.

- 작은 CLI 진입점 설계 경험
- package 분리 경험
- 입력 파싱과 검증 흐름 경험
- 명시적 error 처리 경험

이 기반 위에서 이후 주차는 아래처럼 확장된다.

- `Week 2`: package 구조 정리, 테스트 추가, 리팩터링
- `Week 3`: 파일 입력 확장, 설정 분리, 다중 명령 처리
- `Week 4~5`: 프로세스 실행, 표준 출력/표준 에러 처리
- `Week 6~8`: CI/CD 도구 형태의 작은 오케스트레이션 프로그램으로 확장

즉, 1주차의 파일 아키텍처는 완성형이 아니라 "확장 가능한 최소 구조"여야 한다.

## 11. 리스크 및 주의사항

- Java/Python 스타일로 모든 것을 객체화하려고 하지 않는다.
- 처음부터 복잡한 디렉토리 구조를 만들지 않는다.
- interface를 남용하지 않는다. 필요한 지점에서만 도입한다.
- panic/recover를 일반적인 에러 처리 대체 수단으로 사용하지 않는다.
- CLI가 작더라도 `입력 -> 파싱 -> 검증 -> 출력` 흐름이 드러나도록 작성한다.

## 12. Notion용 주간 실행표

아래 항목은 Notion에 그대로 붙여 넣어 체크박스 형태로 사용할 수 있다.

### Week 1 전체 체크리스트

- [ ] Day 1: `package main`, `func main()`, `var`, `:=`, `const`, zero value를 직접 예제로 작성했다.
- [ ] Day 1: unused variable 에러를 직접 확인하고 이유를 설명할 수 있다.
- [ ] Day 2: 다중 반환값 함수와 `if err != nil` 패턴을 직접 작성했다.
- [ ] Day 2: `for`, `switch`를 포함한 작은 CLI 예제를 만들었다.
- [ ] Day 3: slice의 `len`, `cap`, `append` 동작을 출력으로 확인했다.
- [ ] Day 3: `map[string]int` 기반 카운팅 예제를 작성했다.
- [ ] Day 3: 문자열과 rune 처리 차이를 확인했다.
- [ ] Day 4: `struct`와 method를 사용해 데이터를 모델링했다.
- [ ] Day 4: value receiver와 pointer receiver 차이를 실험했다.
- [ ] Day 5: interface와 구현체를 작성했다.
- [ ] Day 5: `errors.New`, `fmt.Errorf`, `defer`를 각각 사용했다.
- [ ] Day 6: `go mod init`으로 모듈을 만들었다.
- [ ] Day 6: JSON 파일을 읽어 `JobSpec` struct로 파싱했다.
- [ ] Day 6: 검증 로직을 별도 함수 또는 파일로 분리했다.
- [ ] Day 7: `jobspec-cli`를 통합해 정상 입력을 처리했다.
- [ ] Day 7: 실패 입력에 대해 구분 가능한 에러 메시지를 출력했다.
- [ ] Day 7: 코드가 최소 역할 단위로 분리되어 있다.
- [ ] 회고: 이번 주에 Go에서 가장 낯설었던 개념 2가지를 기록했다.
- [ ] 회고: 다음 주에 보강할 주제를 1개 이상 적었다.

### 일일 종료 기록 템플릿

- [ ] 오늘 학습 완료
- [ ] 오늘 새로 이해한 개념 1개 이상 기록
- [ ] 오늘 헷갈린 개념 1개 이상 기록
- [ ] 오늘 작성한 코드 파일명 기록
- [ ] 내일 다시 볼 포인트 1개 기록

## 13. 실행 스펙형 기준 요약

1주차의 완료 상태는 아래로 정의한다.

- Go 기본 문법을 직접 쓸 수 있다.
- Go다운 구조와 명시적 에러 처리 흐름을 이해한다.
- `jobspec-cli`라는 실행 가능한 작은 결과물이 존재한다.
- 이후 주차 확장을 감당할 최소 파일 아키텍처가 준비되어 있다.
