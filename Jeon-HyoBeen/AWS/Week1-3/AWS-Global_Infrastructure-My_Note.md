Region , AZ , Edges

Region을 선택하는 조건
	Compliance: 법적 조건,, (가장 중요) 데이터가 특정 국가에 머물러야 하는지 판단.
	Proximity: 고객과의 거리(지연 시간 최소화) 아시아 대상 서비스인데 미국에 넣으면 ㄴㄴ(Latency)
	Available Service: 내가 이용하려는 서비스가 Region에 있는가?
	Pricing: 가격

AZ(Available Zone)
- 하나의 Region은 여러개의 AZ로 구성
- AZ는 하나 이상의 개별 데이터센터
- 재난 대비 / 고가용성 목적
- AZ간의 연결은 초저지연 네트워크/ 고대역폭 프라이빗 연결/ 물리적 분리

Edge Location
캐싱과 지연 시간 최소화를 위한 지점
콘텐츠 전달의 목적