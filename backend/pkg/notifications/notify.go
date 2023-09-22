package notifications

import (
	"backend/pkg/models"
	"fmt"
)

func NotifyOnShipmentStatusChange(business *models.Business, shipment *models.Shipment) {
	fmt.Printf("Notification for Business: %s, Shipment: %s, Status: %s\n", business.Name, shipment.ID, shipment.Status)
}
