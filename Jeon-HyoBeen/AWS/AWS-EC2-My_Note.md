
EC2 -> 사실상 서버 여는 거

EC2 인스턴스를 생성할때
OS,CPU&RAM,Storage,Network Setting, Security Group, Bootstrap
여기서 NetworkSetting 이란건 SubNet 얘기하는건가

Storage는 EBS와 EFS , Hardware가 있다.


인스턴스 패밀리 - CPU & RAM 어떻게 설정할건지 인스턴스 유형을 잡는거

일반적인 경우(제너럴 퍼포즈) -> t or m 
	t의 경우는 트래픽이 일정하지 않을때, m은 트래픽이 일정할때
고성능 CPU -> c
- CPU의 성능이 필요할때, 
고성능 RAM -> r 
- 많은 데이터를 처리해야할때, 초저지연, 데이터 분석, 캐시 작업 등에 적합
고성능 I/O -> I / D
- 고성능 저장소가 필요할때, I/O 성능 최적화 , I/O가 중요한 데이터베이스 적합

인스턴스의 가격 종류
On-Demand -> 제일 비쌈 사용하는거에 따른 돈 지불, 약정 없음, 단기/ 예측 불가 워크로드에 사용
RI (Reserved Instance) -> 1 or 3 years 로 구독권 느낌, 
Standard RI면 사고 팔수도 있음 가격 정해진 기한만큼 서비스 운용하기 좋음 - 인스턴스 타입과 리전 고정, 
Convertible RI는 인스턴스를 바꿀 수 있으나 할인이 적음, 사고 팔 수 없음
Saving Plans -> 얘도 1년 3년 약정, 그러나 사용량 기준으로 할인이 들어감, 인스턴스 타입과 리전 변경에 있어 자유로움 (RI와의 차이점)
Spot Instances -> 필요할때 필요한만큼 쓰는것 가장 싼 가격, 그러나 손실의 위험이 항상 존재, 뒤에 좀 더 정확하게 후술할거임(Spot Fleet이나 이뤄지는 과정이나..)
Dedicated Host -> 물리적 서버를 가짐, 통제가능, 그러나 비싼 가격//  라이선스 규정과 BYOL, 컴플라이언스 맞추기 위함,, 
Dedicated Instance -> 인스턴스에서 하드웨어는 분리됨, 그러나 따로 통제를 한다거나 물리적 서버를 갖지 않음

Spot Instances의 경우
어떻게 이뤄지나?

Spot Request set을 만듦 -> 규칙에 의해 생성 -> 인스턴트들의 집합 -> 예산설정 등에 의해 삭제 -> 다시 규칙에 맞추어 요청 생성 -> 반복
Spot Instances를 삭제하려면(안쓰려면) Request를 먼저 지우고 Spot Instances를 지워야 재생성 안됨
Spot Instance는 보통 Spot capacity에서 옴,  그때그때 최적의 가격을 찾아서 사용한다. 내가 설정한 금액을 벗어나면 인스턴스 삭제하게끔 설정 가능,
요즘은 maxPrice를 설정을 따로 안하고 On-Demand로 상향선을 걸고 사용한다.
그러면 스팟 인스턴스가 중단되는 이유는? -> AWS가 On-Demand 인스턴스를 위해 Spot capacity를 회수했기 때문
아하! AWS가 On-Demand로 쓰이지 않는 Spot Capacity를 주는데, On-Demand로 누가 쓴다고하면 회수해간다~

Spot Fleet은 뭘ㄹ까 - 여러 인스턴스를 조합해서 원하는 정도를 유지하게끔 해주는 기능
안정적이고 대규모적인 Spot 사용가능
원하는 용량(capacity)를 맞출 때 까지 여러 스팟 인스턴스를 자동으로 조합해서 유지하는 것
Spot Fleet이 없을 경우, Spot Instance 단일 요청에 대한 단점이 많음(용량 부족, 좁은 스코프(단일 AZ), 끊기면 다 날라감) -> 단점 보완하기 위한 Spot Fleet

Spot Fleet을 위한 핵심 개념 
Target Capacity, -> 인스턴스 개수가 아닌 capacity 단위
Allocation Strategy -> capacity-optimized , diversified, lowest-price, priceCapacityOptimized(recommended) 
On-Demand 포함(fallback 개념) <- Ensure baseline capacity

Allocation Strategy에 맞춰서 lowest cost의 인스턴스를 자동으로 요청해준다.
- lowest-price = 제일 싼 풀로 맞춰주는것 (short workload에 적합, cost optimization)
  -> 인터럽션 리스크가 크다.
- diversified = distributed across all pools -> 좋은 availability , long workload에 적합
- capacityOptimized = 말그대로 용량 최적화를 위한 용량 -> lowest interruption risk
- priceCapacityOptimized = 용량 사용에 최적화 + 가장 낮은 가격 (웬만하면 이게 최고임)
interruption risk = 개별 인스턴스의 위험
High Availability = 서비스 전체의 위험

Placement Group 
- Cluster
- Spread
- Partition

Cluster = 하나의 AZ안에 물리적으로 가까이 붙여놓는 것 -> 초저지연, 빠른 응답,  대량의 데이터 처리 가능 , 그런 AZ의 Fail이 일어날시 모든 인스턴스가 같이 다운됨
Spread = 하나의 하드웨어에 하나의 인스턴스, 실패 위험이 가장 적음  그러나 비용이 많이들고, 7개의 인스턴스가 최대임 (한 AZ 당)
Partition - rack 단위로 쪼갠것, 말그대로 파티션을 달아서 생각한다고 생각하면 된다. 하나의 AZ에 파티션을 둬서 관리할수도, 다른 AZ에 파티션 그룹을 통해 인스턴스들을 놓을 수도 있다.
안정성을 높이고, 스프레드보다 더 인스턴스간 소통이 용이함. 대신 같은 지역안에서만 사용가능


EC2 Hibernate 인스턴스 생성시 선택해서 옵션으로 킬 수 있음
instance를 terminate하면, 루트 볼륨이 삭제가 됨 (디폴트)
Stop이후 다시 키면, 기존의 RAM에 있던 내용이 날라감 (재부팅해야함) -> 시간 오래걸림,,,-> 절전모드 이용 -> 절전 모드 이용시 부트에 필요한 내용(RAM내용)을 root Volume에 담아놓음 -> 시간 절약!

그러나 Hibernate는 물리적 하드웨어에 부착되는 내용이 아니기에, 물리적 하드웨어를 쓰는 인스턴스는 불가능함 -> 인스턴스 패밀리에서 메탈 인스턴스는 적용이 안됨-> 물리적 하드웨어를 갖고있는 Dedicate Host는 사용이 불가,,
절전후 60일 지나면 안되고,  150GB까지 가능 -> EBS 볼륨에 대해서만 됨

Spot Instance로 Hibernate기능을 사용할 경우, 절전 모드 이후 다시 시작할때, 동일한 인스턴스로 실행 안될 수도 있다!, + Capacity가 부족한 경우 pending만  계속 뜰 수도 있음...불안정함..

---
## ENI
Elastic IP ? -> 인스턴스 공용 IP는 계속 바뀜, 중지 후 실행할때마다 바뀜
그럼 공용 IP로 접속해야할텐데 이거를 어떻게 관리해? -> 엘라스틱 IP
공용 IP에 탄력 IP를 (고정 외부 IP)부여함으로써 인스턴스에 계속 접근할 수 있도록 해주는것.
물론 비용이 든다. 또한 탄력 IP인 이유는 내가 원하는 인스턴스에 뗏다 붙였다 할 수 있어서.

ENI = 네트워크 카드, IPs = ENI가 갖고 있는 속성
하나의 ENI는 고정 IP를 가짐, 추가적으로 서브 IP도 가질 수 있음.
또한 하나의 EC2 인스턴스에 ENI를 여러개 붙일 수도 있음. 
EC2를 생성하면 하나의 ENI가 생김, -> Primary ENI는 못 넘김
Secondary ENI만 넘길 수 있음 
하나의 AZ안에서만 가능! 왜? ENI는 서브넷에 귀속되기 때문

---
## EBS Volume
음... 하나의 EC2 인스턴스에 여러개의 EBS볼륨 붙일 수 있었고,,,
첫 생성시 생기는 볼륨은 기본적으로 인스턴스 삭제시 삭제되는 옵션(해제가능)
EBS하나에 여러개의 인스턴스 붙이는건 io2/ io3만 가능하고,,
다른 AZ로 못옮기는데 옮기려면 스냅샷써야하고,,
암호화를 설정 안한 EBS볼륨에 대해 암호화하려면 스냅샷이용하면 되고,,

스냅샷 -> 복제기능 (지역 종속이지만, 카피해서 다른 지역으로 옮길 수 있음)
EBS 볼륨은 기본적으로 AZ에 종속되는데, 다른 AZ로 옮길거면 snapshot기능을 이용해서 복사해서 쓰면 된다.
EBS Snapshot은 지우면 Recycle Bin 에 들어간다  - > 복구 가능 (옵션임)
Fast Snapshot Restore? -> 뭐 스냅샷 복사할때 시간걸리는게 덜걸리는거겠지? 가격은 늘 trade-off로 오는거 기억하자..
스냅샷은 증분방식이다. 처음에만 모든걸 가져오고, 이후엔 변경되는 애들만 변경

AMI -> 라이브러리 만드는거임
pre-packaged -> 다른 Region에 EC2 인스턴스를 가져올 수 있음(패키지된 모습으로,, OS,own software, configuration,,, etc) 
라이브러리랑 똑같이 내가 직접 만들수도, 남이 만든걸 쓸 수도(살 수도 있음, 마켓존재) 
Region 종속인데, 카피가능 다른 지역으로 돌리게



EC2 Instance Store 
- 진짜 개빠르고 높은 퍼포먼스 disk
- I/O performance 최상, 쓰로풋도 최상
- 근데 일시적인 저장소임, EC2 스탑하면 사라짐, 따라서 백업안하면 사용자 책임
EBS 타입
6개 타입
gp2/ gp3 - 가장 일반적인 목적의 SSD, gp2의 경우는 비례해서 늘어남 IOPS와 Disk 용량?
gp3는 개별적으로 내가 설정 가능
-> 일단 쌈, 싸고 좋음
io1/ io2 Block Express -> SSD, 고성능 SSD임, 말그대로 고성능임, 또한 multi-attach 가능
st 1 - HDD, 가격이 싼 HDD 
sc 1 - 제일 싼 HDD

일단 루트볼륨으로 할 수 있는 애는 gp2/gp3, io1,io2 Block Express임

IOPS는 SSD가 높음, HDD는 낮음

|**타입**|**IOPS**|**Throughput**|**용도**|
|---|---|---|---|
|gp3|중~높|중|대부분|
|io2|매우 높음|높음|DB|
|st1|낮음|높음|로그/빅데이터|
|sc1|매우 낮음|낮음|아카이브|
IOPS 256,000 이상이면 Instance Storage인거 기억하자
나머지는 문제 풀면서 이해하기 

EBS multi-Attach -> 한번에 16개 인스턴스 연결이 최대임, 클러스터화 된 파일 시스템만 가능

---
## EFS
진짜 말그대로 파일 시스템, 
EBS볼륨보다 사용성 높고, 사용 범위 넓은데, 가격이 나감.
Security Group 설정 가능, 

여러개의 EC2에 접근가능, 또한 지역범위임
보안 그룹 사용 가능
리눅스에서만 사용가능
확장성 높고 사용성도 높음
자동 용량 확장 -> 내가 용량을 미리 정하는게 아님, 
Auto Scaling 그룹에서 여러 EC2가 같은 파일을 써야할때 사용하면 됨