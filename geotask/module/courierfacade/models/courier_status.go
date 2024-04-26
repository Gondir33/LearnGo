package models

import (
	cm "gitlab.com/ptflp/geotask/module/courier/models"
	om "gitlab.com/ptflp/geotask/module/order/models"
)

type CourierStatus struct {
	Courier cm.Courier `json:"courier"`
	Orders  []om.Order `json:"orders"`
}
