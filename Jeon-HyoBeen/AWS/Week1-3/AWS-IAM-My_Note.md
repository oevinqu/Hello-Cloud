
1. IAM의 기본
	-> 사용자(users),그룹(groups),정책(policies),역할(Role)에 대한 내용
보통 그룹으로 나눠서 권한을 부여하는 방식으로 관리
사용자는 여러개의 그룹에 속할 수 있다.
**그룹 내의 그룹 불가능 !!! (Nested Group X)** -> IAM 그룹은 Flat 구조
Policy -> 권한 (JSON 파일로 정리)
	Version
	Id(optional)
	Statements{
		SID
		Allow/Deny
		어디에?
		누가?
		Condition}
권한 JSON
Version, id, statement
version-> 언제
Statement -> 어디에 무엇을 누가 어떻게
Effects, Resources, User/Role/Services/ Actions
Role -> 권한 + 누가 쓸 수 있는지 (신뢰 관계) User나 AWS Service가 임시로 권한을 쓰는 방식

다시 정리 !
권한 Policy
Version, Statement
Statement -> SID, Effect, Action, Resource
	{Version
	Id(option)
	Statement:{
		SID(option)
		Effect
		Principal{ }
		Action
		Resource
		Condition(option)
		}
	}
ID,SID,Principal,Condition -> 있을 수 있다. 없어도 됨

User -> 장기 자격 증명 (Access Key)를 받음 -> 따라서 최소한의 키만 줘야함 !!!!!!
Role -> 임시 접근 키를 줌.
Service는 Role을 통해서만 권한 사용

Principal의 경우는 Role의 신뢰 정책에 반드시 필요(Trust Policy)
보통 service, 다른 사람에게 Role을 빌려줄때, 다른 AWS계정에 권한을 줄때 사용

**1인 1계정 원칙, 최소한의 권한 부여 잊지말기**
**자격 증명 보고서와 Last Access를 통해 수시 체크**
**MFA와 빡빡한 비밀번호 규정을 통해 보안강화**

권한을 주는 방식은 내 맘대로 커스텀 가능
Effect에서 Allow와 Deny중에선 Deny가 우선,, 명시적 deny > 명시적 allow > 비명시적 Deny

IAM은 글로벌 서비스이다.

**Root User는 셋업할때말곤 사용하지 말자!!!!!!!!!**

Access key란? AWS SDK나 CLI로 쓰기 위한 자격 증명
Access Key ID & Secret Access Key
User만 가질 수 있음.

보통 IAM Roles로 준다.
Access Key = 접근 방식
Security Credential로 보통 줌

인라인 정책
1:1 관계
재사용 불가
JSON 직접 작성

-> 재사용 가능한 권한 혹은 이미 있는 권한은 인라인이 아닌 그룹으로 권한을 준다.
-> 관리의 어려움 존재


Credentials Report 와 Access Analyzer ,, 액세스어드바이저도 있었나?
아무튼 Credentials Report는 IAM 유저의 Security를 확인할 수 있다.
Access Advisor(last Access)