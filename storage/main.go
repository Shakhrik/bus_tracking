package storage

import (
	"github.com/Shakhrik/bus_tracking/api/storage/postgres"
	"github.com/Shakhrik/bus_tracking/api/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	CompanyRepo() repo.CompanyRepoI
	CompanyBranchRepo() repo.CompanyBranchRepoI
	RegionRepo() repo.RegionRepoI
	FamilyRelationRepo() repo.FamilyRelationRepoI
	EducationRepo() repo.EducationRepoI
	PositionRepo() repo.PositionRepoI
	VacancyTemplateRepo() repo.VacancyTemplateRepoI
	GenderRepo() repo.GenderRepoI
	NotificationTemplateRepo() repo.EmployeeNotificationTemplateRepoI
	SystemNotificationTemplateRepo() repo.SystemNotificationTemplateRepoI
	EmployeeStatus() repo.EmployeeStatusRepoI
	ScheduleRepo() repo.ScheduleRepoI
	DirectorRepo() repo.DirectorStorageI
	DepartmentRepo() repo.DepartmentStorageI
	ArchievePositionRepo() repo.ArchievePositionsRepoI
	ShiftRepo() repo.ShiftRepoI
}

type storagePG struct {
	db                             *sqlx.DB
	companyRepo                    repo.CompanyRepoI
	companyBranchRepo              repo.CompanyBranchRepoI
	regionRepo                     repo.RegionRepoI
	familyRelationRepo             repo.FamilyRelationRepoI
	educationRepo                  repo.EducationRepoI
	positionRepo                   repo.PositionRepoI
	vacancyTemplateRepo            repo.VacancyTemplateRepoI
	genderRepo                     repo.GenderRepoI
	notificationTemplateRepo       repo.EmployeeNotificationTemplateRepoI
	systemNotificationTemplateRepo repo.SystemNotificationTemplateRepoI
	employeeStatus                 repo.EmployeeStatusRepoI
	scheduleRepo                   repo.ScheduleRepoI
	directorRepo                   repo.DirectorStorageI
	departmentRepo                 repo.DepartmentStorageI
	archievePositionRepo           repo.ArchievePositionsRepoI
	shiftRepo                      repo.ShiftRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:                             db,
		companyRepo:                    postgres.NewCompanyRepo(db),
		companyBranchRepo:              postgres.NewCompanyBranchRepo(db),
		regionRepo:                     postgres.NewRegionRepo(db),
		familyRelationRepo:             postgres.NewFamilyRelation(db),
		educationRepo:                  postgres.NewEducationRepo(db),
		positionRepo:                   postgres.NewPositionRepo(db),
		vacancyTemplateRepo:            postgres.NewVacancyTemplateRepo(db),
		genderRepo:                     postgres.NewGenderRepo(db),
		notificationTemplateRepo:       postgres.NewNotificationTemplateRepo(db),
		systemNotificationTemplateRepo: postgres.NewSystemNotificationTemplateRepo(db),
		employeeStatus:                 postgres.NewEmployeeStatusRepo(db),
		scheduleRepo:                   postgres.NewScheduleRepo(db),
		directorRepo:                   postgres.NewDirectorRepo(db),
		departmentRepo:                 postgres.NewDepartmentRepo(db),
		archievePositionRepo:           postgres.NewArchievePositionRepo(db),
		shiftRepo:                      postgres.NewShiftRepo(db),
	}
}

func (s storagePG) CompanyRepo() repo.CompanyRepoI {
	return s.companyRepo
}

func (s storagePG) CompanyBranchRepo() repo.CompanyBranchRepoI {
	return s.companyBranchRepo
}

func (s storagePG) RegionRepo() repo.RegionRepoI {
	return s.regionRepo
}

func (s storagePG) FamilyRelationRepo() repo.FamilyRelationRepoI {
	return s.familyRelationRepo
}

func (s storagePG) EducationRepo() repo.EducationRepoI {
	return s.educationRepo
}

func (s storagePG) PositionRepo() repo.PositionRepoI {
	return s.positionRepo
}

func (s storagePG) VacancyTemplateRepo() repo.VacancyTemplateRepoI {
	return s.vacancyTemplateRepo
}
func (s storagePG) GenderRepo() repo.GenderRepoI {
	return s.genderRepo
}

func (s storagePG) NotificationTemplateRepo() repo.EmployeeNotificationTemplateRepoI {
	return s.notificationTemplateRepo
}
func (s storagePG) SystemNotificationTemplateRepo() repo.SystemNotificationTemplateRepoI {
	return s.systemNotificationTemplateRepo
}

func (s storagePG) EmployeeStatus() repo.EmployeeStatusRepoI {
	return s.employeeStatus
}

func (s storagePG) ScheduleRepo() repo.ScheduleRepoI {
	return s.scheduleRepo
}

func (s storagePG) DirectorRepo() repo.DirectorStorageI {
	return s.directorRepo
}

func (s storagePG) DepartmentRepo() repo.DepartmentStorageI {
	return s.departmentRepo
}

func (s storagePG) ArchievePositionRepo() repo.ArchievePositionsRepoI {
	return s.archievePositionRepo
}
func (s storagePG) ShiftRepo() repo.ShiftRepoI {
	return s.shiftRepo
}
