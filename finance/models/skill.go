package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string
	Description string
}

func (s *Skill) TableName() string {
	return "skill"
}

type SkillDependency struct {
	ParentSkillID uint
	ParentSkill   Skill `gorm:"foreignKey:ParentSkillID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChildSkillID  uint
	ChildSkill    Skill `gorm:"foreignKey:ChildSkillID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (s *SkillDependency) TableName() string {
	return "skill_dependency"
}

type SkillConflict struct {
	gorm.Model
	Skill1ID uint
	Skill1   Skill `gorm:"foreignKey:Skill1ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Skill2ID uint
	Skill2   Skill `gorm:"foreignKey:Skill2ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comment  string
	Priority int
}

func (s *SkillConflict) TableName() string {
	return "skill_conflict"
}

type SkillDomain struct {
	SkillID  uint
	Skill    Skill
	DomainID uint
	Domain   Domain
	Priority int
}

func (sd *SkillDomain) TableName() string {
	return "skill_domain"
}
