package schemas

import "Finance/models"

//
// Skill Schema
//

type SkillSchema struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`

	Description string `json:"description" binding:"required"`
}

func (s *SkillSchema) ToModel(model *models.Skill) {
	model.ID = s.ID
	model.Name = s.Name

	model.Description = s.Description
}

func (s *SkillSchema) FromModel(model *models.Skill) {
	s.ID = model.ID
	s.Name = model.Name

	s.Description = model.Description
}

//
// Skill Dependency Schema
//

type SkillDependencySchema struct {
	ParentSkillID uint `json:"parent_skill" binding:"required"`
	ChildSkillID  uint `json:"child_skill" binding:"required"`
}

func (s *SkillDependencySchema) ToModel(model *models.SkillDependency) {
	model.ParentSkillID = s.ParentSkillID
	model.ChildSkillID = s.ChildSkillID
}

func (s *SkillDependencySchema) FromModel(model *models.SkillDependency) {
	s.ParentSkillID = model.ParentSkillID
	s.ChildSkillID = model.ChildSkillID
}

//
//	SkillConflict
//

type SkillConflictSchema struct {
	Skill1ID uint   `json:"skill_1" binding:"required"`
	Skill2ID uint   `json:"skill_2" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
	Priority int    `json:"priority" binding:"required"`
}

func (s *SkillConflictSchema) ToModel(model *models.SkillConflict) {
	model.Skill1ID = s.Skill1ID
	model.Skill2ID = s.Skill2ID
	model.Comment = s.Comment
	model.Priority = s.Priority
}

func (s *SkillConflictSchema) FromModel(model *models.SkillConflict) {
	s.Skill1ID = model.Skill1ID
	s.Skill2ID = model.Skill2ID
	s.Comment = model.Comment
	s.Priority = model.Priority
}

//
//	SkillDomain
//

type SkillDomainSchema struct {
	SkillID  uint `json:"skill_id" binding:"required"`
	DomainID uint `json:"domain_id" binding:"required"`
	Priority int  `json:"priority" binding:"required"`
}

func (s *SkillDomainSchema) ToModel(model *models.SkillDomain) {
	model.SkillID = s.SkillID
	model.DomainID = s.DomainID

	model.Priority = s.Priority
}

func (s *SkillDomainSchema) FromModel(model *models.SkillDomain) {
	s.SkillID = model.SkillID
	s.DomainID = model.DomainID

	s.Priority = model.Priority
}
