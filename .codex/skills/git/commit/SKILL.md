---
name: git-commit
description: Use when preparing a commit in a learning project and needing to choose an English commit prefix such as feat, fix, delete, refactor, or test based on the actual change
---

# 학습용 Git Commit

## 개요

- 학습용 프로젝트용 `git commit` 최소 규칙
- 작은 단위 기록, 명확한 메시지, 검증 결과 정리 목적

## 적용 조건

- 커밋 전 변경 범위 정리 시점
- 커밋 타입 및 메시지 결정 시점

## 규칙

- 하나의 커밋 = 하나의 논리적 작업
- 커밋 타입 = 정확히 하나만 선택
- 사용 가능 타입 = 다음 다섯 개만 허용
  - `feat`: 기능 추가
  - `fix`: 버그 수정
  - `delete`: 코드, 파일, 기능 제거
  - `refactor`: 동작 변화 없는 구조 정리
  - `test`: 테스트 추가 또는 테스트 코드 수정
- 한 커밋 메시지에 타입 혼합 금지
- 커밋 메시지 = 항상 영어
- 검증 = 현재 프로젝트에 맞는 명령 직접 실행
- 확인 없는 `PASS` 표기 금지

## 확인 순서

1. `git diff --staged` 확인
2. `feat/fix/delete/refactor/test` 중 타입 하나 선택
3. 하나의 타입으로 설명 어려우면 커밋 분리
4. 관련 검증 명령 실행
5. 짧은 영어 커밋 메시지 작성

## 출력 템플릿

```md
## Commit type
- feat / fix / delete / refactor / test

## Suggested message
- feat: feat: short summary
- fix: fix: short summary
- delete: delete: short summary
- refactor: refactor: short summary
- test: test: short summary

## Validation
- Gate: PASS/FAIL
- Mixed changes: YES/NO
```

## 메시지 규칙

- `feat`: `feat: short summary`
- `fix`: `fix: short summary`
- `delete`: `delete: short summary`
- `refactor`: `refactor: short summary`
- `test`: `test: short summary`

## 예시 메시지

- `feat: add basic variable example`
- `feat: implement slice iteration practice`
- `feat: add struct usage example`
- `fix: correct pointer example`
- `fix: resolve nil map initialization issue`
- `fix: correct loop condition in practice code`
- `delete: remove unused practice file`
- `delete: remove duplicate example code`
- `delete: remove obsolete notes file`
- `refactor: rename variables for clarity`
- `refactor: split examples by topic`
- `refactor: simplify helper function structure`
- `test: add function example test`
- `test: add map usage test case`
- `test: update loop practice test`
