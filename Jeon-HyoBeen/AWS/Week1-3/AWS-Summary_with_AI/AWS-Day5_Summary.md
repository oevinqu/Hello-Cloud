# 🚀 EC2 심화 (Spot & Dedicated)

이 노트는 EC2 구매 옵션 중 가장 복잡하고 시험에 자주 출제되는 **Spot Instance**와 **Dedicated Hosting**에 대해 다룹니다.

## 1. EC2 Spot Instance 상세

### Spot Instance Request (스팟 요청)

- **작동 원리:** 내가 지불할 의사가 있는 최대 가격(Max Price)을 설정하고, 현재 스팟 가격이 그보다 낮을 때 인스턴스를 확보합니다.
    
- **회수 알림:** AWS가 인스턴스를 회수할 때 **2분 전**에 알림(Grace Period)을 줍니다.
    
- 종료 방법 (매우 중요): 1. 반드시 **Spot Request를 먼저 취소(Cancel)**해야 합니다.
    
    2. 그 후에 실행 중인 **인스턴스를 종료(Terminate)**합니다.
    
    - _주의:_ 인스턴스만 먼저 종료하면, Request가 남아있어 자동으로 새로운 인스턴스를 다시 생성합니다.
        

## 2. Spot Fleet (스팟 플릿)

Spot Fleet은 여러 인스턴스 풀(Instance Types, AZ)을 조합하여 목표 용량(Target Capacity)을 자동으로 맞춰주는 집합체입니다.

### 할당 전략 (Allocation Strategy)

- **lowestPrice:** 가장 저렴한 풀에서만 할당. 비용 절감은 극대화되나 회수 위험이 높음.
    
- **diversified:** 선택한 모든 풀에 골고루 분산. 한 AZ가 마비되어도 안정성 유지.
    
- **capacityOptimized:** 현재 남는 용량이 가장 많은 풀 선택. 안정성이 높음.
    
- **priceCapacityOptimized (추천 ⭐):** * **원리:** 먼저 용량이 넉넉해서 회수될 확률이 낮은 풀들을 추려낸 뒤, 그중 가장 저렴한 풀을 선택.
    
    - **이유:** 비용 절감(Spot의 목적)과 안정성(작업 중단 방지) 사이의 가장 완벽한 균형을 제공하기 때문.
        

## 3. Dedicated Host vs Dedicated Instance

물리적 서버를 남과 공유하지 않는 두 옵션의 결정적 차이를 이해해야 합니다.

|   |   |   |
|---|---|---|
|**구분**|**Dedicated Instances (전용 인스턴스)**|**Dedicated Hosts (전용 호스트)**|
|**핵심 키워드**|하드웨어 비공유|**피지컬 코어, 소켓, BYOL**|
|**물리 제어권**|낮음 (어느 기계인지 모름)|**높음 (물리 서버 가시성 제공)**|
|**라이선스**|일반적인 인스턴스 과금|**기존 라이선스(소켓/코어 단위) 지입 가능**|
|**특징**|보안 규정 준수 목적|물리적 서버를 통째로 빌려 직접 관리|

> **💡 시험 팁:** 문제에 **"Socket", "Physical Core", "License (BYOL)"**라는 단어가 나오면 무조건 **Dedicated Host**가 정답입니다.

## 4. 구매 옵션 총정리 (Hotel 비유)

- **On-Demand:** 언제든 정가 내고 묵는 일반 객실 (유연성 최고, 비쌈)
    
- **Reserved (RI):** 1~3년 미리 예약하고 할인받는 회원권 (꾸준한 사용)
    
- **Savings Plans:** "난 얼마만큼 무조건 쓸게"라고 금액을 약정하는 유연한 할인권
    
- **Spot:** 빈방을 경매로 싸게 얻지만, 주인이 오면 나가야 함 (비용 90% 절감)
    
- **Dedicated Host:** 호텔 건물 한 채를 통째로 빌려 방 구조까지 직접 관리
    
- **Capacity Reservation:** "나 이때 방 쓸 거니까 비워놔"라고 예약 (할인은 없지만 가용성 보장)
    

## 💡 실습 시 주의사항

- **IAM Role 활용:** EC2 내부에서 AWS CLI를 쓸 때 Access Key를 입력하지 말고, 반드시 **IAM Role**을 할당하여 권한을 관리할 것 (보안의 핵심!).
    
- **Spot Fleet 혼합:** 순수 Spot으로만 구성하면 전체가 날아갈 수 있으므로, 안정성을 위해 **On-Demand 인스턴스를 일부 섞어서(Optional)** 운영할 수 있음.