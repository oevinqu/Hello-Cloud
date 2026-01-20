# ☁️ AWS SAA-C03 복습 노트: 기초 & IAM


## 1. 클라우드 컴퓨팅의 핵심 이점 (Cloud Benefits)

기존의 온프레미스(On-premise) 레거시 인프라와 비교했을 때 클라우드는 다음과 같은 강점을 가집니다.

- **비용 효율성 (Cost-Effective):**
    
    - **CAPEX(자본 지출) → OPEX(운영 지출):** 데이터센터 구축을 위한 막대한 초기 투자 비용 없이, 사용한 만큼만 지불(Pay-as-you-go).
        
    - **규모의 경제 (Economies of Scale):** AWS의 대규모 인프라 덕분에 개별 리소스 단가가 낮아짐.
        
- **탄력성 및 민첩성 (Elasticity & Agility):**
    
    - 수요에 따라 리소스를 즉각적으로 확장(Scale-out)하거나 축소(Scale-in) 가능.
        
    - 하드웨어 조달 기간 없이 몇 분 만에 글로벌 배포 가능.
        
- **신뢰성 (Reliability):**
    
    - 데이터센터의 가용성 관리가 용이하며, 장애 발생 시에도 서비스 연속성 유지 가능.
        

## 2. AWS 글로벌 인프라 (Global Infrastructure)

### 리전 (Region) 선택 기준

리전은 전 세계에 퍼져 있는 데이터 센터의 집합체입니다. 선택 시 다음 4가지를 고려해야 합니다.

1. **Compliance (데이터 거버넌스 및 법적 요구사항):** 데이터가 특정 국가 내에 머물러야 하는가? (**가장 중요한 핵심 기준**)
    
2. **Proximity (근접성):** 고객과 얼마나 가까운가? (지연 시간/Latency 최소화)
    
3. **Available Services:** 해당 리전에서 내가 사용하려는 서비스(예: Bedrock, 특정 인스턴스 타입)를 지원하는가?
    
4. **Pricing (가격):** 리전마다 비용이 다르므로 예산에 맞춰 선택.
    

### 가용 영역 (Availability Zone, AZ)

- 하나의 리전은 보통 **3개 이상의 AZ**로 구성됩니다 (최소 3개).
    
- 각 AZ는 하나 이상의 개별 데이터 센터로 이루어져 있습니다.
    
- **재난 대비:** 각 AZ는 물리적으로 격리되어 있어(전력망, 수로 등), 한 곳에 문제가 생겨도 다른 AZ가 영향을 받지 않습니다.
    
- **연결:** AZ 간에는 초저지연(Ultra-low latency) 네트워킹으로 연결되어 실시간 데이터 동기화가 가능합니다.
    

## 3. IAM (Identity and Access Management)

IAM은 **글로벌 서비스**입니다. (리전별로 설정할 필요 없음)

### 핵심 구성 요소

- **Root Account:** 계정 생성 시 만들어지는 이메일 계정. 모든 권한을 가지므로 **절대 평소에 사용하지 말 것.**
    
- **Users:** 실제 사람(또는 서비스). 그룹에 속할 수 있고, 여러 그룹에 속할 수도 있음.
    
- **Groups:** 사용자들의 집합. 권한 관리를 효율적으로 하기 위해 사용 (그룹 내에 그룹을 넣을 수 없음).
    
- **Roles:** 사람이 아닌 **AWS 서비스**(예: EC2)가 다른 서비스(예: S3)에 접근할 때 부여하는 임시 권한.
    

### IAM Policy (JSON 구조)

정책은 JSON 문서 형식으로 권한을 정의합니다.

- **Version:** 정책 언어 버전 (항상 `"2012-10-17"`).
    
- **Id:** 정책 식별자 (선택 사항).
    
- **Statement (필수):** 실제 권한을 정의하는 부분.
    
    - **Sid:** Statement ID (선택 사항).
        
    - **Effect:** 접근 허용(`Allow`) 또는 거부(`Deny`).
        
    - **Principal:** 이 정책이 적용될 대상 (계정/사용자/역할).
        
    - **Action:** 허용/거부할 행위의 목록 (예: `s3:ListBucket`).
        
    - **Resource:** 정책이 적용될 리소스의 목록 (ARN 형식).
        
    - **Condition:** 정책이 효력을 발휘할 조건 (예: 특정 IP에서만 접근 가능).
        

> **⚠️ 보안 골든 룰:**
> 
> 1. **최소 권한 원칙 (Least Privilege):** 사용자에게 꼭 필요한 권한만 부여할 것.
>     
> 2. **Deny 우선순위:** 동일한 리소스에 대해 Allow와 Deny가 충돌하면 **무조건 Deny가 우선**함.
>     

## 4. 계정 보안 및 액세스 (MFA & Access)

### MFA (Multi-Factor Authentication)

- 비밀번호 외에 추가적인 인증 수단을 통해 보안 강화.
    
- **종류:** Virtual MFA(Google Authenticator, Authy), Physical MFA(Yubikey), Hardware Key Fob.
    

### AWS 접속 방법 3가지

1. **Management Console:** 웹 브라우저를 통한 그래픽 사용자 인터페이스 (비밀번호 + MFA 필요).
    
2. **CLI (Command Line Interface):** 터미널/커맨드 창에서 명령어로 관리. 스크립트 자동화에 유용.
    
3. **SDK (Software Development Kit):** 프로그래밍 언어(Java, Python 등) 내에서 코드로 AWS 리소스 관리.
    

> **Access Keys (보안 주의):** CLI와 SDK를 사용하기 위해서는 **Access Key ID**와 **Secret Access Key**가 필요합니다. 이는 비밀번호와 같으므로 절대 공유하거나 코드에 직접 노출해서는 안 됩니다.

## 💡 Exam Alert (시험 적중 포인트)

- **IAM은 글로벌 서비스**인가? (Yes)
    
- **그룹 안에 그룹**을 넣을 수 있는가? (No)
    
- **최소 권한 원칙**이 무엇인가? (필요한 것만 주는 것)
    
- **리전 선택** 시 법적 규제(Compliance)가 왜 중요한가? (가장 우선되는 제약 사항이므로)