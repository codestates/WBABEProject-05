package receipt

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var ReceiptModel ReceiptModeler

type ReceiptModeler interface {
	// InsertReceipt : 메뉴 리스트에서 해당 메뉴 선택, 주문 요청 및 초기상태 저장, 주문 내역 초기상태 저장, / - 금일 주문 받은 일련번호-주문번호 리턴
	InsertReceipt(receipt *entity.Receipt) (string, error)
	// UpdateReceipt :  메뉴 변경 및 추가, 메뉴 추가시 상태조회 후 배달중일 경우 실패 알림 / 성공 실패 알림, 실패시 신규주문으로 전환, ; 메뉴 변경시 상태가 조리중, 배달중일 경우 확인, / 성공 실패 알림
	UpdateReceiptStatus(receipt *entity.Receipt) (int, error)

	UpdateCancelReceipt(receipt *entity.Receipt) (int, error)

	// SelectReceiptByID : 주문내역 조회, 현재 주문내역 리스트 및 상태 조회 - 하기 주문 상태 조회에서도 사용, / ex. 접수중/조리중/배달중 etc 없으면 null 리턴 ; 과거 주문내역 리스트 최신순으로 출력, / 없으면 null 리턴
	SelectReceiptByID(receiptID string) (*entity.Receipt, error)
	// SelectReceipts : 현재 주문내역 리스트 조회 , 각 메뉴별 상태 변경 , / ex. 상태 : 접수중/접수취소/추가접수/접수-조리중/배달중/배달완료 등을 이용 상태 저장, 각 단계별 사업장에서 상태 업데이트 - 접수중 → 접수 or 접수취소 → 조리중 or 추가주문 → 배달중, /  성공여부 리턴

	SelectToDayTotalCount() (int, error)

	SelectSortLimitedReceipt(userID string, sort *page.Sort, skip, limit int) ([]*entity.Receipt, error)
	SelectTotalCount(userID string) (int, error)
}
