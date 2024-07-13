package usecase

import (
	"cleanarchitecture/internal/entity"
	"cleanarchitecture/pkg/events"
)

type OrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() (OrdersOutputDTO, error) {
	orders, err := l.OrderRepository.FetchAll()
	if err != nil {
		return OrdersOutputDTO{}, nil
	}

	var ordersDTO OrdersOutputDTO
	for _, o := range orders {
		dto := OrderOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		}

		ordersDTO.Orders = append(ordersDTO.Orders, dto)
	}

	return ordersDTO, nil
}
