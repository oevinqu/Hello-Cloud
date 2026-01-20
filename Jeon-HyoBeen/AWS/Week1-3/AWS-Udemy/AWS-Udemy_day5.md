EC2 Spot Instance Request
WTP 를 결정해서 정한다.
spot instnace 가격에 따라 옮겨다니며 하면 된다.
2분안에 끝내야함

EC2 Spot Instance Requests
- Can get a discount of up to 90% compared to On-demand
- Define max spot price and get the instance while current spot price < max
	- The hourly spot price varies based on offer and capacity
	- If the current spot price > your max price you can choose to stop or terminate your instance with a 2 minutes grace period.
- Other strategy: Spot Block (사라진 서비스,,시험엔 나올수도?)
	- "block" spot instance during a specified time frame ( 1 to 6 hours) without interruptions
	- In rare situations, the instance may be reclaimed
- Used for batch jobs, data analysis, or workloads that are resilient to failures.
- Not great for critical jobs or databases

AZ따라 가격이 달라진다. 정의한 맥스 가격에 따라서 스팟 가격이 비싸지면, 사용하지 않게 된다. 맥스 가격보다 낮아지면 인스턴스를 확보한다. 작업 비용을 최적화 할 수 있다.  가격이 어떻게 될지는 아무도 모르지만, 전반적으로 on demand 보다 싸게 쓸 수 있다.

How to terminate Spot Instance
<!-- ![[IMG_4745.jpeg]] -->


이런 스팟 인스턴스를 하는것은, 우리의 책임이고, 
Spot Request를 먼저 없애고, 인스턴스를 없애야한다. 인스턴스 먼저 없애면 리퀘스트가 들어가서 다시 생기니까...

Spot Fleets
- Spot Fleets = set of Spot Instance + (optional) On- Demand Instances
- The Spot Fleet will try to meet the target capacity with price constraints
	- Define possible launch pools: instance type(m5.large), OS, Availability Zone
	- Can have multiple launch pools, so that the fleet can choose
	- Spot Fleet stops launching instances when reaching capacity or max cost
- Strategies to allocate Spot Instance:
	- lowestPrice : from the pool with the lowest price (cost optimization, short workload)
	- diversified : distributed across all pools (great for availability, long workloads)
	- capacityOptimized: pool with the optimal capacity for the number of instances
	- priceCapacityOptimized (recommended): pools with highest capacity available, then select the pool with the lowest price(best choice for most workloads)
- Spot Fleets allow us to automatically request Spot Instances with the lowest price

스팟 인스턴스에서 하는것 = Spot Fleets.내가 원하는 설정을 한뒤, 스팟 요청이 생성되고 조건에 맞는 인스턴스들이 생성
인스턴스 생성에서 max price설정해서 하는거 -> 테스트, 인스턴스 중심 (개별 요청)
예약 인스턴스로 가서 카트에 추가 가능. 하고 예약 인스턴스를 주문 가능
근데 절감형 플랜 (Saving plans)가 더 이득이라서,,, 사라질수도/?

전용 호스트 = Dedicated Host


---

You would like to deploy a database technology on an EC2 instance and the vendor license bills you based on the physical cores and underlying network socket visibility. Which EC2 Purchasing Option allows you to get visibility into them?

걍 피지컬 코어, 소켓 이런거 나오면 Dedicated Host ㄱㄱ, 
Dedicated Instance는 다른 사람과 하드웨어를 공유하지 않는거뿐임..
물리적 서버 어쩌고 저쩌고, 실제 소켓을 볼 수 있다 ㅇㅈㄹ하면 무조건 Dedicated Host ㄱㄱ

Spot Fleet is a set of Spot Instances and optionally On-Demand Instances
-> 다 Spot Instance로 설정하면 내 서버가 통째로 날라갈수있으니, On-Demand instance를 넣어서 안정적으로 관리할 수도 있다(옵션 ~!!)

---
