package manage

import (
	"context"
	"fmt"
	parkinglot "parking-lot/pkg/parking-lot"
)

func (m *Manager) CreateParkingLot(ctx context.Context, slot uint8) error {
	newParkingLot := make([]parkinglot.CarSlot, slot)
	m.ParkingLot = &newParkingLot
	return nil
}

func (m *Manager) AllocateParkingLot(ctx context.Context, regisNum, color string) (int, error) {
	if _, err := m.GetParkingLot(ctx); err != nil {
		return 0, err
	}
	for no, slot := range *m.ParkingLot {
		if slot.ID == 0 {
			(*m.ParkingLot)[no] = parkinglot.CarSlot{
				ID:          no + 1,
				PlateNumber: regisNum,
				Color:       color,
			}
			return no + 1, nil
		}
	}
	return 0, fmt.Errorf("Sorry, parking lot is full")
}

func (m *Manager) GetParkingLot(ctx context.Context) (*[]parkinglot.CarSlot, error) {
	if m.ParkingLot == nil {
		return nil, fmt.Errorf("Empty parking lot")
	}
	return m.ParkingLot, nil
}