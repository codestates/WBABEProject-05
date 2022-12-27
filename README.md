# WBABEProject-05

# 띵동주문이요, 온라인 주문 시스템(Online Ordering System)

|         | 내용                                                    |
|---------|-------------------------------------------------------|
| 프로젝트 제목 | 띵동주문이요, 온라인 주문 시스템(Online Ordering System)            |
| 소요시간    | 3 ~ 5일                                                |
| 난이도     | 보통                                                    |
| 기술스택    | Golang, Gin Framework, RESTful, MVC패턴, MongoDB, Query |

---

## 요구사항
- gin-gonic framework를 사용하여 온라인 주문 시스템 API Server를 개발합니다.
- 온라인 주문 시스템에 관련한 데이터베이스 document를 직접 설계합니다.
- 주문자, 피주문자 CRUD API를 개발합니다.
- mongodb와 연동하여 실제 데이터베이스에 데이터가 저장될 수 있게 합니다.
- 디버깅을 위한 로그 출력이 가능하게 구성합니다.
- 스웨거를 이용해 API 문서화를 합니다.
- toml을 이용해 설정 파일을 구성합니다.


### 메뉴 신규 등록  - 피주문자

- **API |** 신규 메뉴 등록
    - 사업장에서 신규 메뉴 관련 정보를 등록하는 과정(ex. 메뉴 이름, 주문가능여부, 한정수량,  원산지, 가격, 맵기정도, etc)
    - 성공 여부를 리턴

### 메뉴 수정 / 삭제 - 피주문자

- **API |** 기존 메뉴 수정/삭제
    - 사업장에서 기존의 메뉴 정보 변경기능(ex. 가격변경, 원산지 변경, soldout)
    - 메뉴 삭제시, 실제 데이터 백업이나 뷰플래그를 이용한 안보임 처리
    - 금일 추천 메뉴 설정 변경, 리스트 출력
    - 성공 여부를 리턴

### 메뉴 리스트 출력 조회 - 주문자

- **API |** 메뉴 리스트 조회 및 정렬(추천/평점/주문수/최신)
    - 각 카테고리별  sort 리스트 출력(ex. order by 추천, 평점, 재주문수, 최신)
    - 결과 5~10여개 임의 생성 출력, sorting 여부 확인

### 메뉴별 평점 및 리뷰 조회 - 주문자

- **API |** 개별 메뉴별 평점 및 리뷰 보기
    - UI에서 메뉴 리스트에서 상기 리스트 출력에 따라 개별 메뉴를 선택했다고 가정
    - 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴

### 메뉴별 평점 작성 - 주문자

- **API |** 과거 주문 내역 중, 평점 및 리뷰 작성
    - 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장
    - 성공 여부 리턴

### 주문 - 주문자

- **API |** UI에서 메뉴 리스트에서 해당 메뉴 선택, 주문 요청 및 초기상태 저장
    - 주문정보를 입력받아 주문 저장(ex. 선택 메뉴 정보, 전화번호, 주소등 정보를 입력받아 DB 저장)
    - 주문 내역 초기상태 저장
    - 금일 주문 받은 일련번호-주문번호 리턴

### 주문 변경 - 주문자

- **API |** 메뉴 변경 및 추가
    - 메뉴 추가시 상태조회 후 `배달중`일 경우 실패 알림
        - 성공 실패 알림, 실패시 신규주문으로 전환
    - 메뉴 변경시 상태가 `조리중`, `배달중`일 경우 확인
        - 성공 실패 알림

### 주문 내역 조회 - 주문자

- **API |** 주문내역 조회
    - 현재 주문내역 리스트 및 상태 조회 - 하기 **주문 상태 조회**에서도 사용
        - ex. 접수중/조리중/배달중 etc
        - 없으면 null 리턴
    - 과거 주문내역 리스트 최신순으로 출력
        - 없으면 null 리턴

### 주문 상태 조회 - 피주문자

- **API |** 현재 주문내역 리스트 조회
- **API |** 각 메뉴별 상태 변경
    - ex. 상태 : 접수중/접수취소/추가접수/접수-조리중/배달중/배달완료 등을 이용 상태 저장
    - 각 단계별 사업장에서 상태 업데이트
        - **접수중 → 접수** or **접수취소 → 조리중** or **추가주문 → 배달중**
        - 성공여부 리턴


---

### 프로젝트 구성

![프로젝트 구성](./readme_images/project-init.png)

- toml 파일 및 프로젝트 config 관련 사항은 config 폴더를 이용할 것이다.
- mvc 패턴을 활용하기에 controller 와 model 패키지를 만들었다.
- protocol 은 request 또는 response 로 내려줄 data struct 들을 넣을 것이다.
- util 패키지는 flag 및 범용적으로 사용하게될 사항들을 넣어주자.
- logger 관련사항은 logger 패키지에 넣어주자.
- router 와 관련된 사항은 router 패키지에서 다룰 것이다.

### git-branch 전략

![브랜치 전략](./readme_images/branch-strategy.png)

- 최종 main 브랜치에 merge 후 dev 브랜치는 main 브랜치와 싱크를 맞추기 위해 main 브랜치를 merge 한다.
- 위의 싸이클을 반복하며 개발한다.
- release 는 지금처럼 배포전 1차 개발단계에서는 생략하기도한다. 그러나 이 프로젝트에서는 생략하지 않기로 한다.


### naming 전략

- 인터페이스는 `~er`을 붙이고 구현체에는 붙이지 않는다.
- `.go` 파일은 두 단어 이상일 때 케밥케이스를 사용한다.
  - `ex) person_router.go`
- 코드는 카멜케이스를 원칙으로 한다.
  - `ex) personRouter`
- go의 `private`, `public` 네이밍을 기본으로 한다.
- 상수가 아닌이상 첫 글자는 소문자로 한다.
- 전역변수는 의미가 드러나게 작성하고 약어를 사용하지 않는다.
- 지역변수는 가능한한 약어로 하고, 리시버는 단어의 앞글자만을 사용한다.
- 만약 지역변수의 생존기간이 길어 코드의 가독성을 해치거나, 의미가 드러나야 한다면 약어를 사용하지 않는다.
- 메서드 이름은 최대한 의도가 드러나게 작성하자
  - ex)
  - ```go
    func (r *router) NewRouter(ctl Controller){
        n := ctl.Name
        ...
    }
    func (pr *PersonRouter) validatePerson(){
        pr.validate()
        ...
    }
    ```
    
# Docker 를 이용한 MongoDB 설치

- docker-compose-singlersdb.yml 파일을 이용해 mongoDB를 설치하자. 
- 여기서 우선 mongoDB의 replication 옵션을 이용할 것이다.
- replication 옵션을 이용하는 이유는 트랜젝션을 이용하게될지도 모르기 떄문이다.
- 트랜젝션은 mongoDB 를 standalone 상태로 구동하면 실행되지 않는다.
- 우선 트랜젝션을 mongoDB 에서는 권장하지 않기에 최대한 트랜젝션을 사용하지 않는 구조로 설계하겠지만 개발은 항상 유연해야하기에 가능성을 열어두는 측면에서 replication 옵션을 사용한다.
다음의 명령어들을 실행시키자.
```bash
$ docker-compose -f docker-compose-singlersdb.yml up -d
$ docker exec -it mongo-singlers /bin/bash
$ mongosh 127.0.0.1:27017/
test> rs.initiate()
```

    
### API 구현 중..
- 현재 구현된 API


    GET    /
	GET    /info           
    POST   /users/join     
    POST   /stores         
    POST   /stores/menu    
    DELETE /stores/menu    
    PUT    /stores/menu


- 몽고DB로 원하는 데이터를 뽑고 가공하는게 쉽지않다... 생각보다 더 쉽지 않아서 큰일이다.
- entity 설계만 끝나면 전체적인 도메인 구성과 쿼리는 어렵지 않을것이라 생각했는데.. 
- 정규화와 트랜젝션을 지양한다는 점에서 설계에 고민을 많이하기도 했고, 오히려 고민을 하다보니 더 복잡해진것 같기도하고, 무엇보다 mongo 쿼리가 처음 써보는거다보니 상당히 쉽지 않다.





