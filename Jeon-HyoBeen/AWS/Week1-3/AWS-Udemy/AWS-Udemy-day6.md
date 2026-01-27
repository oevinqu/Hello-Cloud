Private vs Public IP (IPv4)
- Networking has two sorts of IPs. IPv4 and IPv6:
	- IPv4: 숫자 네개를 점 네개로 나눔
	- IPv6: 문자랑 숫자로
- 코스에서는 IPv4 만사용
- IPv4가 자주 사용
- IPv6는 보통 IoT나 이런거에 사용되나 우리가 지금 당장 쓸 이유가 없음

(Elastic)
공용 아이피 주소가 계속 바뀌는데 ㅌ탄성IP를 이용하면 우리가 인스턴스를 중단했다가 다시 시작할때마다 똑같은 IP주소(탄성주소)를 이용할 수 있다... associate와 disassociate를 하고 삭제하고 하면 된당 -> 돈 문제


Placement Groups (EC2 Network& Security 에 배치그룹)
- Sometimes you want control over the EC2 Instance placement strategy
- That strategy can be defined using placement groups
- When you create a placement group, you specify one of the following strategies for the group:
	- Cluster -- clusters instances into a low-latency group in a single AZ
	- Spread -- spreads instances across underlying hardware (max 7 instances per group per AZ) - critical applications
	- Partition -- spreads instances across many different partitions (which rely on different sets of racks) within an AZ. Scales to 100s of EC2 instances per group (Hadoop,Cassandra,Kafka)
-> AWS의 하드웨어와 직접 연결되지는 않지만 EC2 인스턴스를 어떻게 배치할지 AWS에게 알려주는것

Cluster -
<!-- ![[Pasted image 20260120155455.png]] -->
- Pros: Great network(10 Gbps bandwidth between instances with Enhanced Networking enabled - recommended)
- Cons: If the AZ fails, all instances fails at the same time
- Use case:
	- Big Data job that needs to complete fast
	- Application that needs extremely low latency and high network throughput

Spread - 실패 위험을 줄인것
<!-- ![[Pasted image 20260120160423.png]] -->
- Pros:
	- Can span across AZ
	- Reduced risk is simultaneous failure
	- EC2 Instances are on different physical hardware
- Cons:
	- Limited to 7 instances per AZ per placement group
- Use case:
	- Application that needs to maximize high availability
	- Critical Applications where each instance must be isolated from failure from each other

Partition - rack 단위로 파티션을 쪼개 같은 AZ라도 다른 하드웨어 사용
<!-- ![[Pasted image 20260120160730.png]] -->
- Up to 7 partitions per AZ
- Can span across multiple AZs in the same region
- Up to 100s of EC2 instances
- The instances in a partition do not share racks with the instances in the other partitions
- A partition failure can affect many EC2 but won't affect other partitions
- EC2 instances get access to the partition information as metadata
- Use cases: HDFS, HBase, Cassandra, Kafka

---
Elastic Network Interfaces (ENI)
- Logical component in a VPC that represents a virtual network card
- The ENI can have the following attributes:
	- Primary private IPv4, one or more secondary IPv4
	- One Elastic IP(IPv4) per private IPv4
	- One Public IPv4
	- One or more security groups
	- A MAC address
	- and etc...
-  You can create ENI independently and attach them on the fly (move them) on EC2 instances for failover
- Bound to a specific AZ
Example
EC2 Instance
 ├─ ENI #1 (Primary)
 │   ├─ Private IP 10.0.1.10
 │   ├─ (optional) Secondary Private IPs
 │
 ├─ ENI #2 (Secondary)
 │   ├─ Private IP 10.0.2.15
 │
 └─ ENI #3 (Secondary)
     ├─ Private IP 10.0.3.20

-> EC2 인스턴스를 만들때 고유한 ENI 하나 생성
ENI는 생성 시 고유한 Private IP를 가진다.
하나의 ENI에 여러개의 IP를 부여할 수 있다.
EC2 인스턴스 내에 여러개의 ENI를 연결 할 수 있으며 (인스턴스 타입 한도 내)
인스턴스 생성 시 생긴 고유한 ENI를 제외한 ENI는 다른 인스턴스에 뗏다 붙였다 할 수 있다.
-> 장애 대응 및 IP 유지를 위해.. IP는 ENI에 귀속 된다.
-> 1개의 리소스이기 때문에 여러개의 인스턴스에 동시 적용은 불가하다.
-> 같은 AZ안에서만 해야한다!!! ENI는 특정 Subnet에 속하기 때문에 !!! (bound to a specific AZ) 
물론 AZ에 하나의 서브넷만 존재하는 것은 아니지만,, 같은 AZ내의 인스턴스에 붙일 수 있다 ^^

---
#### 논리 계층: Subnet (가상의 방)

- **Subnet (서브넷):** VPC라는 네트워크를 용도(보안, 통제)에 맞게 쪼개놓은 **가상의 주소 구역**.
    
- **귀속 규칙:** 서브넷은 반드시 **하나의 AZ에만 귀속**됨. (AZ를 넘나들 수 없음)
    
- **확장성:** 하지만 **하나의 AZ 안에는 여러 개의 서브넷**을 만들 수 있음. (예: 2a AZ 안에 '퍼블릭 서브넷'과 '프라이빗 서브넷'을 동시에 생성)
    

####  두 계층의 연결 고리 (핵심!)

- **1:1 매칭 아님:** "서브넷 A = 데이터 센터 1동"이 아님.
    
- **추상화:** 서브넷 A에 인스턴스를 만들면, AWS는 해당 AZ에 속한 **여러 데이터 센터들 중** 가장 적절한 곳에 물리적으로 배치함.
    
- **ENI 이동의 비밀:** ENI가 "다른 서브넷인 인스턴스"로 옮겨갈 수 있는 이유는, 두 인스턴스가 **같은 AZ(물리적 건물 단지) 내에만 있다면** 네트워크 선을 끌어다 줄 수 있기 때문임.
---

Cf) Resource groups & Tag Editor를 모든 리소스에 대해 확인 가능하다.


EC2 Hibernate(절전모드) -> 인스턴스 생성시 Stop-Hibernate behavior(최대 절전 중지 방식)
- We know we can stop, terminate instances
	- Stop - the data on disk(EBS) is kept intact in the next start
	- Terminate - any EBS volumes (root) also set-up to be destroyed is lost
-  On start, the following happens:
	- First start: the OS boots & EC2 User Data script is run
	- Following starts: the OS boots up
	- Then your application starts, caches get warmed up, and that can take time!
- Introducing EC2 Hibernate:
	- The in-memory (RAM) state is preserved
	- The instance boot is much faster
	  (the OS is not stopped / restarted)
	- Under the hood: the RAM state is written to a file in the root EBS volume
	- The root EBS volume must be encrypted
- Use cases:
	- Long-running processing
	- Saving the RAM state
	- Services that take time to initialize
- EC2 Hibernate - Good to know -> 물리 자원을 직접 통제하는 경우 안됨// 일반 EC2 Hibernate는 깨어난 뒤, 아무 하드웨어에 띄우기 때문
	- Supported Instance Families - 거의 다 됨
	- Instance RAM Size - must be less than 150GB
	- Instance Size - not supported for bare metal instances. 베어 메탈 인스턴스는 적용 안됨 -> 하드웨어를 직접 할당받음
	- AMI - 거의 다 됨
	- Root Volume - must be EBS, encrypted, not instance store, and large
	- Available for On-Demand, Reserved and Spot Instance
	  -> Saving Plan은 Instance Size를 바꿀 수 있기에 안되고, Dedicated Host는 물리자원을 직접 통제하기에 불가 , Dedicated Instance는 될 수도 있지만...
	- An instance can NOT be hibernated more than 60 days.
	
STOP 과 비교하기
	- stop의 경우 작업 중이던 모든 문서, 브라우저 탭이 다 날아감
	  다시 킬때 OS 부팅부터,
	  인스턴스 값은 안나가지만 , EBS 비용은 계속 청구
	- Hibernate의 경우 RAM에 있던 모든 내용을 EBS에 파일 형태로 저장하고 전원을 끔
	  다시 킬때, 부팅 과정 없이 내가 멈췄던 화면 그대로 뜬다.
	  인스턴스 값은 안나가지만, EBS 비용이 조금 더 청구 될 수 있다.