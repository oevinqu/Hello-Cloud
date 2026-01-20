
AWS - Budget setup
billing and Cost Management - 과금 정보 및 비용 관리

밑에서 IAM user role access to Billing information 에서 Activate IAM Access
레거시 홈의 빌링 홈에 있음;;

active service 볼 수 있고, 어디서 가격이 쓰이는지, 이번달 예상가격, 지난달가격 이런거 볼 수 있다. -> Cost Analysis 지금은 AWS Cost Explore -> 분석

무료 계층 사용을 받을 수 있나도 볼 수 있음 -> Free tire
Current usage와 Forecast Usage도 볼 수 있다. 

Budgets 로 들어가면 template도 짤 수 있다. -> 얼마 쓸때마다 이메일 보냄 혹은 얼마이상쓰면 돈을 보냄, (예산), Instance 정지도 설정 가능(예산 초과시 특정 액션 취하게 할 수 있음)

---
EC2로 웹사이트 만들거라고??
EC2란?

EC2 = Elastic Compute Cloud = Infrastructure as a Service
It mainly consists in the capability of:
- Renting virtual machines(EC2)
- Storing data on virtual drives(EBS)
- Distributing load across machines (ELB)
- Scaling the services using an auto-scaling group (ASG)

Knowing EC2 is fundamental to understand how the Cloud works

EC2 sizing & configuration options
- Operating System (OS) : Linux, Windows or Mac OS
- How much compute power & cores (CPU)
- How much random-access memory (RAM)
- How much storage space:
  - **Network-attached(**EBS & EFS)
  - hardware (EC2 Instance Store)
- Network card: speed of the card, Public IP address
- Firewall rules: **security group**
- Bootstrap script (configure at first launch):EC2 User Data

EC2 User Data -> **재부팅때는 실행 안된다!!!** -> 중지 후 시작 모두 이미 한번 부팅되었던 인스턴스로 간주.. User Data는 다시 실행되지 않는다. 따라서 다시하려면 아예 삭제하고 새로만들거나, 복잡한 설정을 해야함
It is possible to bootstrap our instances using EC2 User data scripts.
bootstrap means lauching commands when a machine starts
That script is only run once at the instance first start
EC2 user data is used to automate boot tasks such as:
- Installing updates
- Installing software
- Downloading common files from the internet
- Anything you can think of -> 부팅 시 할일이 많아지는 것만 기억하자
The EC2 User Data Script runs with the root user -> 어떤 명령이든 sudo 권한을 갖는다.
부팅 자동화 -> bootstrap

Hands-On: -> 루트계정으로 하는거 아님! 루트 계정으로 권한만 넘기고 IAM User로 진행
Launching an EC2 Instance running Linux
- We'll be launching our first virtual server using the AWS Console
- We'll get a first high-level approach to the various parameters
- We'll see that our web server is launched using EC2 user data
- We'll learn how to start/ stop / terminate our instance.
EC2 Console -> Instances
Key pair 하는게 좋다  SSH 유틸리티를 사용하기에
키를 한쌍으로 mac,Linux,window 10환경에선 private key file format 을 .pem으로
PuTTY.... window 7,8은 이거 쓴다

HTTPs 를 쓰는데 왜 얘는 안하는거지 이걸 차단한다는건강? 방화벽 보안그룹에 ㅐㄷ해 좀 더 봐야겠네 -> 체크 안하면 못들어오게 하는거맞는듯,, https로 접근하면 무한 로딩 띄우네,, http 접근허용,,
보안 그룹의 경우는 Stateful 해서 ,inbound 설정을 했으면 outbound는 자동으로 설정된다.

Configure storage를 통해 볼륨 설정이 가능한데, 고급을 통해서 인스턴스를 끌때, 볼륨을 지울지말지 정할 수도 있음

고급 세부 정보에서 bootstrap 설정 가능 - User data 에 코드 작성하면 됨

공영 IPv4 address -> 인스턴스에 접근하는 주소
Private 은 네트워크 내부에서 비공개로 접근하는 주소

안쓸땐 중지하고, 클라우드에선 인스턴스를 시작하고 없애는 게 일반적이다 클라우드니까 !
인스턴스를 정지하고 시작하면 공용 ip주소가 바뀐다

EC2 Instance Types - Overview
- You can use different types of EC2 instances that are optimised for different use cases (https://aws.amazon.com/ec2/instance-types/)
- AWS has the following naming convention:
	- m5.2xlarge
	- m: instance class
	- 5:generation (AWS improves them over time)
	- 2xlarge: size within the instance class -> 크면 CPU 많이 잡아먹음

시험의 관점
EC2 Instance types - General Purpose
- Great for a diversity of workloads such as web servers or code repositories
- Balance between:
	- Compute
	- Memory
	- Networking
- In the course, we will be using the t2.micro(지원 끊겨서 t3.micro 쓰는중) which is a General Purpose EC2 instance

EC2 Instance Types - Compute Optimized
- Great for compute-intensive tasks that require high performance processors:
	- Batch processing workloads
	- Media transcoding
	- High performance web servers
	- High performance computing(HPC)
	- Scientific modeling & machine learning
	- Dedicated gaming servers
-> C5, C6 등 C라는 이름을 사용합니다?

EC2 Instance Types - Memory Optimized -> 진짜 메모리, 데이터를 최대한 RAM에 가두고 싶을때, 대규모 관계형 DB등 캐칭 쿼리 속도를 높일때,
- Fast performance for workloads that process large data sets in memory
- Use cases:
	- High performance, relational/non-relational databases
	- Distributed web scale cache stores
	- In-memory databases optimized for BI (business intelligence)
	- Applications performing real-time processing of big unstructured data

EC2 Instance Types - Storage Optimized -> 읽기,쓰기 등 저장소(디스크)에서 데이터 가져오는 속도를 높이는 게 목표일때
- Great for storage-intensive tasks that require high, sequential read and write access to large data sets on local storage
- Use cases:
	- High frequency online transaction processing (OLTP) systems
	- Relational & NoSQL databases
	- Cache for in-memory databases (for example, Redis)
	- Data warehousing applications
	- Distributed file systems

Note: other type of instances may exist but only the most important ones have been discussed so you get the general idea

EC2 Instance Types:example
t2.micro
t2.xlarge
c5d.4xlarge
r5.16xlarge
m5.8xlarge

Great website: https://www.ec2instances.info/ -> 레거시라 지워짐 요 ㅠㅠ
now redirects to https://instaces.vantage.sh -> 여깃 ㅓ내용 비교 가능 !

Elastic IP?