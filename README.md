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

![프로젝트 구성 최종](./readme_images/최종프로젝트패키지들.png)

- 최종 프로젝트 구성이다.
- common 패키지는 flag 및 범용적으로 사용하게될 util, enum 등의 사항들을 넣어주었다.
- toml 파일 및 프로젝트 config 관련 사항은 config 폴더를 이용하였다.
- mvc 패턴을 활용하기에 controller 와 model 패키지를 만들었다.
- 조금 더 유연한 프로그램을 위해 controller 계층과 model 계층사이에 service 계층을 추가해 주었다.
- db 폴더는 메뉴삭제시 백업테이터로 삭제된 데이터를 파일로 저장하는 위치로 사용된다.(깃 소스에는 생략)
- docs 는 스웨거 관련 폴더이다.
- logger 관련사항은 logger 패키지에 넣어주자.
- logs 는 앱에서 사용되는 로그들이 기록되는 위치이다.(깃 소스에는 생략)
- protocol 은 request 또는 response 로 내려줄 data struct 들을 넣었다.
- router 와 관련된 사항은 router 패키지에서 다룰 것이다.
- test 는 간단하게 .http 를 활용해 테스트를 해보았기에 .http 파일들을 모아놓은 폴더이다.

---

- 추가적으로 아래와 같이 commend 또는 query 관련 파일이 있는데,   
  이는 한 파일이 가지고 있는 코드가 너무 많아 데이터의 조작을 하는 함수는 commend 로,    
  조회하는 함수는 query 로 분리시켜 주었다.

![commend-query 설명](./readme_images/commend-query.png)


---

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
- Golang 에서 두문자어는 대문자로 작성한다. ex) Database => DB
    - ex)
  ```go
    func (r *router) NewRouter(ctl Controller){
        n := ctl.Name
        ...
    }
    func (pr *PersonRouter) validatePerson(){
        pr.validate()
        ...
    }
    ```

## 실행환경의 편의를 위해 replica 를 사용한 MongoDB 가 아닌 standalone 으로 간단한 MongoDB로 전환했다.
- MongoDB 에 지향점에 맞춰 트랜잭션을 사용하지 않았다.
- 다음과 같이 docker-compose 파일을 실행하자
```bash
$ docker-compose up -d 
```
- 몽고DB 접속 URI :
  `mongodb://127.0.0.1:27017`




### API 

GET    /home/info   

POST   /app/v1/users/user
GET    /app/v1/users/user        
PUT    /app/v1/users/user        
DELETE /app/v1/users/user        

GET    /app/v1/stores            
GET    /app/v1/stores/store      
POST   /app/v1/stores/store       
PUT    /app/v1/stores/store         
GET    /app/v1/stores/store/recommends    
GET    /app/v1/stores/store/menus    
POST   /app/v1/stores/store/menus/menu    
PUT    /app/v1/stores/store/menus/menu    
DELETE /app/v1/stores/store/menus/menu    
GET    /app/v1/stores/store/menus/menu   

GET    /app/v1/orders/pages/store   
GET    /app/v1/orders/pages/customer   
POST   /app/v1/orders/order      
GET    /app/v1/orders/order      
PUT    /app/v1/orders/order/customer    
PUT    /app/v1/orders/order/store    
GET    /app/v1/orders/order/price   

POST   /app/v1/reviews/review      
GET    /app/v1/reviews/menu        
GET    /app/v1/reviews/customer     


## 느낀점

#### 아쉬운점

- 처음에 Document 지향으로 최대한 해보려고 하였으나, MongoDB의 Document 데이터를 가공해서 가져오기가 쉽지않았다. 떄문에 RDBMS 방식처럼 _id 를 활용하였다. 
- 참고 : [리모델링](https://github.com/codestates/WBABEProject-05/issues/23)


#### 좋았던점
- 그래도 이번 프로젝트로 그동안 커리큘럼을 진행하면서 `Golang 에서는 어떻게 객체지향적으로 코드를 작성할 수 있을까?` 를 수 없이 고민했었는데 이번에 나름 만족할만한 구조로 개발을 했다.
- 각 계층은 의존성을 최소화하기 위해 interface 를 통해 소통을 하도록 DIP 지킬 수 있도록 작성했다. 이는 나중에 쉽게 구현로직을 변경할 수 있게 해준다. 변경된 Dependency 를 inject 해줘야 하기에 완벽한 OCP 는 아니겠지만
  초기에 Dependency 를 inject 받은 interface 를 Public 하게 열고 사용하여 최대한 OCP 가 이루어지도록 하였다. 이로써 사용하는 계층은 확장에는 열려있고 변경에는 닫혀있게 된다.
- 또한 각 구현체는 싱글톤으로 사용되도록 구현했으며, 좀 더 SRP 에 가깝도록 구현로직과 생산로직을 분리, 각 계층에 Dependency 를 Inject 해주는 manager 를 두었다.
- 그동안 main 메서드에 많은 로직이 들어가는 것 같아 불편했는데 이 부분도 init() 메서드와 app struct 를 활용해 깔끔하게 만들었고, main.go 파일 안에서 전체 프로젝트 코드들의 구성이 한눈에 드러날 수 있도록 했다.
- golang 은 상속이 안돼, 가끔 재상용의 불편을 느꼈는데, composite 패턴과 앞서 말한 Dependency 를 inject 받은 interface Public 하게 열어두어 재사용이 용이하게 하였다.
- 좀 더 개선사항들이 보여 이 부분도 시간을 더 갖고 싶지만.. 이점은 조금 아쉽지만,
- 그래도 어떤식으로 `Golang 에서 객체지향을 녹이면 좋을지 이번 프로젝트를 통해 어느정도 느끼고 꺠달은것 같아 재밌는 프로젝트였다.`


### Swagger 전체
- http://localhost:8080/swagger/index.html#/

- feat.추가적인 리팩토링을 하면서 아래 이미지와 실제는 아주 살짝 차이가 있을 수 있습니다. 자세한건 위 링크를 통해 스웨거로 확인해주세요.

![전체1](./readme_images/swagger/usecase/01주문기록전체.png)

![전체1](./readme_images/swagger/usecase/02메뉴기록전체.png)

![전체1](./readme_images/swagger/usecase/03가게전체.png)

![전체1](./readme_images/swagger/usecase/04사용자정보전체.png)


## UseCase 로 보는 Swagger 

- 추가적으로 아래 이미지들을 보면 페이지관련 query string 은 편의상 빈값으로 요청을 보냈는데, 서버에서 빈 값으로 요청이 오면 첫페이지 5개씩 최신순으로 기본값 처리한다.

![전체1](./readme_images/swagger/usecase/05사용자등록_유저.png)

![전체1](./readme_images/swagger/usecase/06사용자등록_유저응답.png)

![전체1](./readme_images/swagger/usecase/07사용자등록_가게.png)

![전체1](./readme_images/swagger/usecase/08사용자등록_가게_응답.png)

![전체1](./readme_images/swagger/usecase/09가게등록.png)

![전체1](./readme_images/swagger/usecase/10가게등록_응답.png)

![전체1](./readme_images/swagger/usecase/11메뉴등록.png)

![전체1](./readme_images/swagger/usecase/12메뉴등록_응답.png)

![전체1](./readme_images/swagger/usecase/13가게수정.png)

![전체1](./readme_images/swagger/usecase/14가게수정_응답.png)

![전체1](./readme_images/swagger/usecase/15가게보기.png)

![전체1](./readme_images/swagger/usecase/16가게보기_응답.png)

![전체1](./readme_images/swagger/usecase/17메뉴검색.png)

![전체1](./readme_images/swagger/usecase/17-1메뉴검색_응답.png)

![전체1](./readme_images/swagger/usecase/18-1주문전가격보기.png)

![전체1](./readme_images/swagger/usecase/18-2주문전가격보기_응답.png)

![전체1](./readme_images/swagger/usecase/19주문.png)

![전체1](./readme_images/swagger/usecase/20주문_응답.png)

![전체1](./readme_images/swagger/usecase/21최근주문저장확인리스트.png)

![전체1](./readme_images/swagger/usecase/22최근주문저장확인_응답.png)

![전체1](./readme_images/swagger/usecase/23주문조회유저.png)

![전체1](./readme_images/swagger/usecase/24주문조회유저_응답.png)

![전체1](./readme_images/swagger/usecase/25주문수정유저.png)

![전체1](./readme_images/swagger/usecase/26주문수정유저_응답.png)

![전체1](./readme_images/swagger/usecase/27주문수정가게.png)

![전체1](./readme_images/swagger/usecase/28주문수정가게_응답.png)

![전체1](./readme_images/swagger/usecase/29메뉴리뷰작성.png)

![전체1](./readme_images/swagger/usecase/30메뉴리뷰작성_응답.png)

![전체1](./readme_images/swagger/usecase/31메뉴리뷰확인.png)

![전체1](./readme_images/swagger/usecase/32메뉴리뷰확인_응답.png)

![전체1](./readme_images/swagger/usecase/33메뉴또한업데이트.png)

![전체1](./readme_images/swagger/usecase/34메뉴또한업데이트_확인.png)


