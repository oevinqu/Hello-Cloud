# 💰 AWS 비용 관리 및 💻 EC2 기초 완벽 정리

이 노트는 **비용 모니터링**과 **EC2(Elastic Compute Cloud)**의 핵심 개념 및 인스턴스 유형에 대해 다룹니다.

## 1. AWS 비용 및 예산 관리 (Billing & Budgets)

AWS 환경을 안전하고 저렴하게 유지하기 위한 필수 도구들입니다.

- **AWS Budgets (예산 설정):** 템플릿을 사용하여 특정 금액 이상 사용 시 이메일 알림을 받거나, 서비스 중지 등의 액션을 설정할 수 있습니다.
    
- **AWS Cost Explorer (비용 분석):** 과거 사용량 분석 및 향후 비용 예측(Forecast)을 시각적으로 보여줍니다.
    
- **Free Tier Monitor:** 현재 프리티어 사용량을 확인하여 과금을 방지할 수 있습니다.
    
- **IAM Access to Billing:** 루트 계정이 아닌 IAM 사용자가 결제 정보를 보려면 반드시 **'Activate IAM Access'** 설정을 켜야 합니다.
    

## 2. EC2 (Elastic Compute Cloud) 개요

EC2는 AWS의 대표적인 **IaaS(Infrastructure as a Service)** 서비스로, 가상 서버를 대여하는 서비스입니다.

### 핵심 구성 요소

1. **EC2 (Virtual Machine):** 가상 서버 본체.
    
2. **EBS (Elastic Block Store):** 가상 드라이브 (데이터 저장소).
    
3. **ELB (Elastic Load Balancer):** 트래픽 분산 장치.
    
4. **ASG (Auto Scaling Group):** 수요에 따른 서버 자동 증설/감소.
    

### 부트스트래핑 (Bootstrapping) - EC2 User Data

- **개념:** 인스턴스가 시작될 때 자동으로 명령어를 실행하는 기능입니다.
    
- **특징:**
    
    - 인스턴스 **최초 실행 시 딱 한 번**만 실행됩니다.
        
    - **Root 사용자** 권한으로 실행됩니다. (모든 명령에 `sudo` 수준의 권한 부여)
        
    - 소프트웨어 설치, 업데이트, 파일 다운로드 등 초기 설정을 자동화할 때 사용합니다.
        

## 3. EC2 인스턴스 유형 (Instance Types)

용도에 따라 적절한 인스턴스 패밀리를 선택하는 것이 SAA 시험의 핵심입니다.

### 명명 규칙 (예: `m5.2xlarge`)

- **m:** 인스턴스 클래스 (패밀리)
    
- **5:** 세대 (숫자가 높을수록 최신/고성능)
    
- **2xlarge:** 인스턴스 크기 (CPU, RAM 용량)
    

### 주요 인스턴스 패밀리

|패밀리|유형|주요 사용 사례|
|---|---|---|
|**M**|**General Purpose**|웹 서버, 코드 저장소 등 균형 잡힌 작업 (t3.micro 등)|
|**C**|**Compute Optimized**|비디오 인코딩, 고성능 웹 서버, 머신러닝 (CPU 집중)|
|**R**|**Memory Optimized**|고성능 DB, 분산 캐시(Redis), 실시간 빅데이터 분석 (RAM 집중)|
|**I / D / H**|**Storage Optimized**|OLTP 시스템, NoSQL DB, 데이터 웨어하우징 (저장소 I/O 집중)|
|**T**|**Burstable Performance**|평소엔 저렴하게 쓰다가 갑자기 트래픽이 튈 때 성능 발휘|

## 4. EC2 생명 주기 및 네트워크 핵심

- **인스턴스 중지(Stop) 후 시작(Start):**
    
    - 물리적 호스트가 바뀔 수 있으며, **Public IP 주소가 변경**됩니다.
        
    - **Private IP 주소는 유지**됩니다.
        
    - EBS에 저장된 데이터는 유지됩니다.
        
- **인스턴스 재부팅(Reboot):**
    
    - 동일한 호스트에서 재시작되며, **Public IP 주소가 유지**됩니다.
        
- **Security Group (보안 그룹):**
    
    - **Stateful** 함: 인바운드 허용 시 아웃바운드는 자동 허용.
        
    - 기본적으로 모든 인바운드는 차단(Deny) 상태입니다.
        
    - HTTPS(443) 접속이 안 된다면 보안 그룹의 인바운드 규칙을 먼저 확인해야 합니다.
        

## 🛠️ 실습 및 세부 설정 팁

- **Key Pair:** SSH 접속을 위해 필요하며, OS 환경에 따라 형식이 다릅니다.
    
    - `.pem`: Mac, Linux, Windows 10 이상에서 사용.
        
    - `.ppk`: Windows 7/8 환경에서 PuTTY 사용 시 필요.
        
- **EBS Delete on Termination:** 인스턴스 종료(Terminate) 시 연결된 EBS 볼륨을 함께 삭제할지 여부를 결정할 수 있습니다. (기본은 Root 볼륨 삭제 활성)
    
- **Storage 옵션:** 네트워크 연결형(EBS, EFS)과 하드웨어 직접 연결형(Instance Store)으로 나뉩니다.
    

## 💡 Exam Alert (시험 적중 포인트)

- 특정 작업(예: 고성능 계산)에 최적화된 **인스턴스 유형을 매칭**하는 문제.
    
- **User Data**가 언제 실행되는지 (최초 부팅 시 1회).
    
- **비용 모니터링**을 위해 어떤 도구(Budgets, Cost Explorer)를 써야 하는지.
    
- 인스턴스 중지 시 **Public IP가 변한다는 사실**과 고정을 위한 **Elastic IP**의 존재.