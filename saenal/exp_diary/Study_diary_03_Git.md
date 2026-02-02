# Git

1. [Git의 의미, 목적, 기능](#1-git의-의미-목적-기능)
2. [용어 정리 및 기본 사용법](#2-용어-정리-및-기본-사용법)
3. [GitHub](#3-github)
4. [명령어 정리](#4-명령어-정리)
5. [추가 학습 사항](#5-추가-학습-사항)

*****

### 1. Git의 의미, 목적, 기능
  (1) 필요성 : 버전 관리(Version Control System), 버그/장애 검증 및 복구, 협업 보조<br>
  (2) 기능 : 변경 이력 기록, 추적, 복원<br>
  (3) 방법
  - 특정 시점의 파일 시스템을 스냅샷으로 기록
  - tree 상태(프로젝트 구조 및 파일 수정 사항), 메타 정보(작성자, 시간, 메시지), 히스토리
  - 기존 버전 관리는 중앙집중식(CVCS) Git은 분산식(DVCS)
  - Git : commit 관리, 인터넷 없이 local에서만도 기능
  - Github : 저장소 공유 플랫폼, PR, 리뷰, 이슈, CI/CD 가능  

[top](#git)

### 2. 용어 정리 및 기본 사용법
  - checksum : git의 기본 데이터 단위. SHA-1 해시값. 무결성 검사(데이터 변조 혹은 유실 체크)
  - repository
  - 영역
    - git directory (.git)
    - staging area (.git\index)
    - working directory
  - 파일 상태
  <table>
    <tr>
      <td rowspan="4"><b>untracked</b></td>
      <td><b>tracked</b></td>
    </tr>
    <tr>
      <td>committed</td>
    </tr>
    <tr>
      <td>staged</td>
    </tr>
    <tr>
      <td>modified</td>
    </tr>
  </table>

  - 기본 흐름 : (checkout >) modify > stage > commit (>merge)
  - .gitignore
    - 이 파일에서 설정한 규칙에 해당하는 파일들은 변경이 있어도 트래킹 X
    - 작성 규칙 : glob 패턴 ([참고](https://jw910911.tistory.com/136#google_vignette))
  - branch : 어느 commit 지점을 가리키는 포인터(reference)
    - ex. main or master, feature
    - 사용 예. feature branch 생성 > 수정, commit > main branch에 merge
    - 효과 : merge 전 테스트 가능, 협업 시 서로 다른 이슈를 동시에 여러 사람이 작업하기에 용이
    - HEAD : 현재 내가 보고 있는 포인터. 현재 checkout 한 브랜치. 새 commit이 생기면 이동
    - \*tag : realease 할 때 사용. 필요 시 push할 때 명시해야 함.
  - merge : 한 브랜치에서 commit한 내용을 다른 브랜치와 통합
    - fast-forward merge(직선형) vs. 3-way merge(공통 조상인 commit 이용)
  - conflict : 서로 다른 브랜치에서 같은 부분을 수정하여 merge 실패<br>
   \> 충돌한 부분을 직접 수정, commit > merge continue
  - stash : commit 하지 않고 임시 저장한 채로 checkout 가능
  - rebase : commit 이력을 선형으로 정리 (공개된 repo에 pussh 된 commit은 rebase 주의)
  - reset : commit 후 push 전 상태에서 HEAD/Index(staged)/Working directory 중 어느 단계로 되돌리기. soft/mixed/hard 중 선택. reset 히스토리는 남지 않음
  - revert : 특정 commit으로 되돌리되, revert commit이 히스토리에 남음. 실수가 이미 공유되었을 때 사용.
    - merge commit은 어떤 부모(mainline)로 되돌릴지 선택해야 함
    - revert commit을 다시 revert 가능

[top](#git)

### 3. GitHub
  - github에 repo 생성(원격, remotes/origin/...), 로컬과 연결하여 추적, 원격>로컬 복제, 로컬>원격 push
  - 브랜치 관리 전략 : git flow vs. github flow<br>
  \> 프로젝트 특성, 팀 성격에 따라 (릴리즈 주기, 버전 관리, 안정성 등)
  - PR : Pull Request. 수정 내용 설명, 검증, 리뷰, 합의 후 merge 가능
  - fork : 오픈소스를 내 github repo(origin)에 fork > 로컬에 clone > 수정, 내 repo에 push > upstream(원 소스)에 PR >> push 불가, PR 가능, upstream 추적 가능
  - collaborator
  - 토큰
  - organization
  - 문제 상황 ex.
    - feature branch보다 main branch가 앞설 때
    - hotfix 필요
    - commit 실수 (; 파일 누락)
    - 실수가 이미 공유됨(pushed)
  - Github Actions
    - CI/CD : Continuous Integration + Continuous Deployment 지속적 병합, 지속적 배포. ex. Jenkins, ArgoCD, Gitlab 등
    - 설정해둔 조건(PR, push 등)에 따라 자동 검증(test) 및 배포(AWS, Azure, GCP, Kubernetes, App Service, Container Service, On-prem server, Google Cloud etc.)
    - 구성 요소
      - workflow : 전체 자동화 시나리오. 한 repo 안에 여러 workflow 파일을 두고 역할별로 사용. 하나 이상의 job들로 구성된 YAML(.yml) 파일
      - event : 실행 조건, 트리거. push, PR, 스케줄 등으로 설정 또는 수동으로 실행. 브랜치를 한정하여 설정 가능 (on:)
      - job : 실행 단위. 한 대의 runner에서 실행
      - runner : job을 실행시키는 머신, github 제공 또는 사용자 직접 운영(내부망 대상 배포 시) (runs-on:)
      - step : 명령어 단위(한 줄/스크립트)
      - action : 기성품..패키지 같은 것.. (uses:)
      - script/command : (run:)
    - [YAML syntax](https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax)
    - CD 패턴 : build/test(CI) > artifact 생성 or container image 생성 > registry 업로드 > deploy 대상에 반영 > 상태 기록 (+알림)
    - CD 인증 : 배포 대상에 접근 위해 키나 토큰 등 필요

[top](#git)

### 4. 명령어 정리
  - (설정)
    - git config : 기본 설정
      - \--global user.name "user name"
      - \--global user.email "email\@address"
    - git init
  - (브랜치)
    - git branch : 현재 존재하는 브랜치 확인
      - \* main 또는 master 브랜치 하나만 존재할 때는 출력 내용 없음
      - 기본 또는 -l : 로컬 브랜치만
      - \-r : 원격 브랜치만
      - \-a : 모든 브랜치
      - \-v : 마지막 commit checksum과 commit message 함께 표시
      - \-vv : 어떤 로컬 브랜치가 어떤 원격 브랜치를 추적하는지 + diff 확인
      - git branch \[branch name] : 새 브랜치 생성
      - \-d : 삭제 (아직 merge 하지 않은 commit 있으면 삭제 불가 >> 강제 삭제 옵션 -D)
      - \-m \[원래 이름] \[새 이름] : 이름 변경 (덮어쓰기 옵션 -M)
      - \* 브랜치 생성/변경/삭제 후 push 해야만 원격에 반영 (git push \[remote branch] \[local branch])
    - git switch \[branch name]
      - \-c : 새 브랜치 만들고 바로 이동
    - git checkout \[branch name]
      - \-b : 새 브랜치 만들고 바로 이동
      - git checkout \[local branch(A)] \[remote branch(B)] : A가 B를 추적하게 함
      - git checkout -t \[remote branch]
  - (commit 관리)
    - git status
      - \-s  / \--short
    - git log
      - -p / --patch : diff 결과
      - \--stat
      - \--oneline
      - \--graph
      - \--short
      - \--full
      - \--fuller
      - \--format
      - \-S \[text] : "text"가 변경 내용에 포함되어 있는지 탐색
    - git diff : unstaged 파일의 변경 사항 확인
    - git add \[file name] : 변경 사항을 staging area에 올림
    - git commit
      - \-m \[commit message]
      - \-a : staging area 생략 (add + commit)
      - \--amend : 마지막 commit 수정 (파일 추가, 커밋 메시지 수정)
    - git rebase \[main/master] : feature branch에서 실행!!
    - git stash (save) "message" : 해당 브랜치의 변경 사항 임시 저장
      - git stash push -m "message"
      - \-u : untracked 파일 포함
      - list : stash 목록 확인
      - show : 변경된 파일 확인
      - show -p : 변경 내용 상세 확인
      - apply stash@{n} : 스택에 남긴 채로 불러오기 (stash index 미지정 시 가장 최근 것)
      - pop stash@{n} : 스택에서 지우고 불러오기 (stash index 미지정 시 가장 최근 것)
      - drop stash@{n} : 해당 stash 삭제
      - clear : stash stack 전체 삭제
    - git reset
      - \--soft HEAD~n : 최신부터 n개 커밋 취소하고 staged 상태로
      - \--mixed HEAD~n : 최신부터 n개 커밋 취소하고 modified 상태로 (working directory는 그대로)
      - \--hard HEAD~n : 최신부터 n개 커밋 취소하고 변경 내역도 모두 취소
      - HEAD~n 또는 HEAD^ (^의 개수만큼 이전으로)
      - 파일 특정하여 reset 가능
    - git restore \[filename]
      - \--staged : modified 상태로 되돌리기(unstaged)
    - git revert \[checksum]
  - (원격)
    - git remote : 등록된 원격 저장소 확인
      - git remote add \[remote branch] \[URL] : 연결 후 push
      - \-v : 이름 및 URL 확인
      - git remote rename \[이전 이름] \[새 이름]
      - git remote remove \[삭제할 원격 브랜치]
    - git clone \[URL]
    - git fetch \[remote branch] : 원격에서 로컬로 복사(X merge)
    - git pull \[remote branch] : fetch + merge
    - git push \[remote branch] \[local branch] 
      - \-u \[remote branch] \[local branch] 또는 \--set-upstream : 두 브랜치를 트래킹 설정하여 이후 "git push" 명령어만으로 자동 push
    - git tag
      - 기본 : 조회
      - git tag \[tag] : checksum 저장. 임시로 메타정보 없이 저장
      - \-a \[tag] : 메타정보(작성자, 날짜, 태그 메시지) 함께 저장
      - git push \[remote branch] \[tag] : push는 기본적으로 태그는 전송하지 않으므로 원할 때 명시해야 함
      - git checkout \[tag] : 가져오기

[top](#git)

### 5. 추가 학습 사항
- Remote Refs vs. Remote-tracking branch
- stash를 이용하면 한 branch 내에서 여러 버전 만들고 최종적으로 무엇을 commit할 지 결정해서 가는 것도 가능??
- 오픈소스 (ex.github actions의 action) 사용할 때 원 소스가 변경되면...?
