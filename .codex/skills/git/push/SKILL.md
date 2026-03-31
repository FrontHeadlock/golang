---
name: git-push
description: Use when pushing a learning-project branch and needing to confirm the branch is ready, local verification is complete, and the pushed history stays clean and reviewable
---

# 학습용 Git Push

## 개요

- 학습용 프로젝트용 `git push` 최소 규칙
- 브랜치 상태, 최근 커밋, 검증 결과 확인 목적

## 적용 조건

- 최근 커밋 원격 반영 전 점검 시점
- 브랜치 상태 확인 시점

## 규칙

- 푸시 대상 = 설명 가능한 최근 커밋
- 현재 브랜치, 워킹트리 선확인
- 최근 커밋 메시지와 실제 작업 일치 필요
- 최근 커밋 접두어 = `feat`, `fix`, `delete`, `refactor`, `test` 중 하나
- 검증 미완료 상태 즉시 푸시 금지
- 설명 불가 변경 존재 시 선정리

## 확인 순서

1. `git branch --show-current` 확인
2. `git status --short` 확인
3. 최근 커밋 메시지 확인
4. 실행한 검증 결과 확인
5. 이상 없으면 푸시

## 출력 템플릿

```md
## Push readiness
- Branch: <name>
- Working tree: CLEAN/DIRTY
- Verification: PASS/FAIL
- Ready to push: YES/NO

## Notes
- command run: <verification command>
- latest commit: <hash> <message>
```

## 멈춰야 하는 경우

- 브랜치 불일치
- 검증 실패
- 최근 커밋 메시지 모호
- 워킹트리 내 설명 불가 변경 존재
