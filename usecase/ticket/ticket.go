package ticket

import (
	"log"

	interopaD "github.com/mohfahrur/interop-core/domain/interopa"
	interopbD "github.com/mohfahrur/interop-core/domain/interopb"
	interopcD "github.com/mohfahrur/interop-core/domain/interopc"
	"github.com/mohfahrur/interop-core/entity"
)

type TicketAgent interface {
	UpdateData(req entity.User) (err error)
	SendNotifikasi(req entity.User) (err error)
}

type TicketUsecase struct {
	InteropaDomain interopaD.InteropaDomain
	InteropbDomain interopbD.InteropbDomain
	InteropcDomain interopcD.InteropcDomain
}

func NewTicketUsecase(
	interopaDomain interopaD.InteropaDomain,
	interopbDomain interopbD.InteropbDomain,
	interopcDomain interopcD.InteropcDomain) *TicketUsecase {

	return &TicketUsecase{
		InteropaDomain: interopaDomain,
		InteropbDomain: interopbDomain,
		InteropcDomain: interopcDomain}
}

func (uc *TicketUsecase) UpdateData(req entity.User) (err error) {

	err = uc.InteropcDomain.UpdateData(req)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (uc *TicketUsecase) SendNotifikasi(req entity.User) (err error) {

	err = uc.InteropaDomain.SendEmail(req)
	if err != nil {
		log.Println(err)
		return
	}

	err = uc.InteropbDomain.SendTelegram(req)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
