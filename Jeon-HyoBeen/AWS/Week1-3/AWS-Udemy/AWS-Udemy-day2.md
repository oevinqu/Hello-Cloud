Cloud Shell은 일부지역에서만 가능
그러나 터미널이 잘 되면 상관 없음

클라우드 셸이 파일을 업로드하고 다운로드하는거 쌉가능
pwd로 경로보고 파일 맞춰서 다운 혹은 업로드 가능 ;;

IAM Roles for Services

Some AWS service will need to perform actions on our behalf

그러기 위해선 권한을 줘야한다 AWS Service에게..
그게 IAM Role이다. 
IAM Role  -> Permission -> AWS Service

EC2 Instance ; virtual server

IAM Role을 만들고 하나의 엔터티를 만들고, aws에 접근할때, IAM 롤에 할당된 권한을 토대로 호출에 접근한다.

Common roles:
EC2, Lambda,Cloud,, -> 나중에 또 함

IAM 들어가서 역할 누르고 엔터티 유형 선택하고 적용할 서비스 찾아서 적용

핵심은 사람한테 하듯이 적용가능하네.

IAM Security Tools
- IAM Credentials Report (account-level) (자격증명보고서)
	- a report that list all your account's users and the status of their various credentials
- IAM Access Advisor (user-level) (마지막 액세스)
	- Access advisor shows the service permissions granted to a user and when those services were last accessed.
	- You can use this information to revise your policies
	  
자격증명보고서를 통해 알 수 있는 것은?
-> 계정 내 모든 사용자의 보안상태를 한눈에 파악할때 사용
마지막 액세스를 통해 알 수 있는 것은?
-> 사용자가 어떤 서비스를 마지막으로 언제 썼는지 확인하여, 안쓰는 권한 뺏는데 사용(최소 권한 원칙)

IAM Guidelines & Best practices
- Don't use the root account except for AWS account setup
  -> 설정할때빼고는 걍 개인계정 만들어서 써라..
- One physical user = One AWS user 
  -> 한사람당 하나의 계정만 만들어줘라 -> 관리를 위해 (책임 추적성)
- Assign users to groups and assign permissions to groups
  -> 그룹화를 통해 권한을 부여해라.(관리의 용이성을 위해)
- Create a strong password policy
  -> 비밀번호 정책을 강하게 만들어 보안을 지켜라
- Use and enforce the use of Multi Factor Authentication(MFA)
  -> MFA를 이용하여 보안을 강화해라
- Create and use Roles for giving permissions to AWS services
  -> 모든 서비스에 역할을 이용해 권한을 부여해라
- Use Access Keys for Programmatic Access (CLI / SDK)
  -> CLI/SDK 사용하여 AWS를 조작해야할때만 액세스 키를 생성해서 사용해라
- Audit permissions of your account using IAM Credentials Report & IAM Access Advisor
  -> 주기적으로 보안 도구들을 돌려서 안 쓰는 계정이나 과도한 권한이 있는 지 확인해야한다.
- Never share IAM users & Access Keys
  -> 절대 하지마!!
  
  ---
  IAM Section - Summary
- Users : mapped to a physical user, has a password for AWS Console
- Groups : contain users only
- Polices: JSON document that outlines permissions for user or groups
- Roles: for EC2 instances or AWS services
- Security: MFA + Password Policy
- AWS CLI : manage your AWS services using the command-line
- AWS SDK: manage your AWS services using a programming language
- Access Keys: access AWS using the CLI or SDK
- Audit: IAM Credential Reports & IAM Access Advisor(Last Access)

---
IAM 정리
Users
Groups
Polices
Roles 