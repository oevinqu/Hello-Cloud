
**가상환경 (venv)**
	- 프로젝트마다 필요한 패키지 버전이 다를 때 발생하는 충돌을 방지
	- 명령어: (난 python3으로 쉘에 입력해야함..)
		- python -m venv .venv
		- source .venv/bin/activate //꼭 activate해줘야함
		- pip list
		- pip freeze > requirements.txt
		- pip install -r requirements.txt

**Django?**
- Framework: 웹 개발에 필요한 반복적인 작업을 미리 구현해 둔 도구
- MVT 패턴
	- Model - 데이터베이스 구조 정의
	- View - 로직 처리 및 데이터 가공
	- Template -HTML(화면)
- Django ORM(Object-Relational Mapping)
	- SQL 쿼리를 직접 쓰지않고 Python 코드로 데이터베이스를 조작하는 기술
	- DB 종류가 바뀌어도 코드를 수정할 일이 적음
	- Migration:
		- makemigrations: Model의 변경 사항을기록
		- migrate: 기록된 변경 사항을 실제 DB에 반영
- -> 실습 코드보고 더 정리할 필요 있음

**DRF** 
- 현대적 웹: JSON 기반의 순수 데이터 반환
- API : 프로그램 간 데이터를 주고 받기 위한 약속

**Serializer** (직렬화)
- DB 객체를 JSON으로 변환하거나, 클라이언트가 보낸 JSON을 DB객체로 변환
- 데이터 타입의 유효성을 검증하고 안전하게 데이터를 변환하는 핵심 브릿지
**View Router**
- APIView : 특정 URL에 대해 GET,POST 등의 요청을 직접 처리하는 클래스 기반 뷰
- ViewSet & Router : 반복되는 CRUD 패턴을 자동화하여 URL 경로를 자동으로 생성해주는 도구

**API 테스트와 보안**
- Postman, Swagger, curl 등을 사용하여 API가 의도한대로 동작하는지 확인
- APITestCase를 활용하여 비즈니스 로직의 안정성 검증

---
실습 코드 추가 예정,.,,.