package dummy

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gosuri/uiprogress"
	"github.com/linuxisnotunix/Vulnerobot/modules/models"
)

const id = "DUMMY"

//ModuleDUMMY dummy module do nothing
type ModuleDUMMY struct {
}

//New constructor for Module
func New(options map[string]interface{}) models.Collector {
	log.WithFields(log.Fields{
		"id":      id,
		"options": options,
	}).Debug("Creating new Module")
	return &ModuleDUMMY{}
}

//ID Return the id of the module
func (m *ModuleDUMMY) ID() string {
	return id
}

//IsAvailable Return the availability of the module
func (m *ModuleDUMMY) IsAvailable() bool {
	return true
}

//Collect collect and parse data to put in database
func (m *ModuleDUMMY) Collect(bar *uiprogress.Bar) error {
	//return nil
	return fmt.Errorf("Module '%s' does not implement Collect().", id)
}

//List display known CVE stored by this module in DB
func (m *ModuleDUMMY) List() (*arraylist.List, error) {
	return nil, fmt.Errorf("Module '%s' does not implement List().", id)
}
