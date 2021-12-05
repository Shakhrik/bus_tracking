package repo

import "github.com/Shakhrik/bus_tracking/api/models"

type VacancyTemplateRepoI interface {
	Create(models.VacancyTemplate) error
	Get(guid string) (models.VacancyTemplate, error)
	GetByPosition(positionId string, company_id string) (models.GetVacancyTemplate, error)
	Update(models.VacancyTemplate) error
	Delete(models.VacancyTemplate) error
}
