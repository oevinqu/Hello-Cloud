EFS - Elastic File System
- managed NFS (network file system) that can be mounted on many EC2
- EFS works with EC2 instances in multi-AZ
- Highly available, scalable, expensive (3x gp2), pay per use

Security Group 안에 있는 EFS FileSystem은 AZ를 초월하여 EC2에 mount할 수 있다.

Use cases: content management, web serving ,data sharing, Wordpress
Uses NFSv4.1 protocol
Use security group to control access to EFS
Compatible with Linux based AMI (not Windows)
Encryption at rest using KMS
POSIX file system (~Linux) that has a standard file API
File system scales automatically, pay-per-use, no capacity planning!

EFS - Performance & Storage Classes

EBS Multi-Attach는 io1/ io2 말하는건가?

---
EFS Hands On

EFS 로 들어가고 생성하고
VPC 있다..
Customize로 설정 가능
File System type으로
Regional 할지 One Zone  할지
특정한 AZ로 할 수 있다.

자동 백업 기능 활성화

라이프사이클 관리
액세스했을때
액세스 안했을때 -> 저장소이동
첫 표준 액세스 관리


퍼포먼스 세팅
쓰루풋 모드
엘라스틱, 프로비전 ,
버스팅? -> 처리량 기반에대한 내용
향상 모드 -> 탄성(파일크기에 상관없이 조정)

프로비전이나 버스팅은 필요한 만큼 우리가 조정하는것
-대가를 미리 줌

추가설정
EI에 대해선 General Purpose하고,
Max I/O - 빅데이터 설정에서 필요

일단 인핸스의 엘라스틱과 프로비전 그리고 버스팅 까지 이렇게 3개는 알아야한다.

네트워크 액세스 설정
Mounts targets로 AZ 놓고, 서브넷 ID 놓고, 특정 보안 그룹 걸고
-> EFS를 위한 시큐리티 그룹을 만드는게 낫다.

파일 시스템 정책 - 고급 기술임

Review 가능

EC2 인스턴스 만들고,
Configure Storage? 에서 하면 되네
edit 으로 EFS하고  add share filesystem하면 되고, 서브넷을 먼저 설정해줘야 가능
네트워크 설정에서 서브넷 설정 가능

ㅇㅎ EFS로 