package models

import "gorm.io/gorm"

type Responsibility struct {
	gorm.Model
	SkillID         uint
	Skill           Skill
	Priority        int
	Name            string
	Comments        string
	ExperienceLevel int
}

func (r *Responsibility) TableName() string {
	return "responsibility"
}

type ResponsibilitySynonim struct {
	gorm.Model
	ResponsibilityID uint
	Responsibility   Responsibility `gorm:"foreignKey:ResponsibilityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name             string
}

func (rs *ResponsibilitySynonim) TableName() string {
	return "responsibility_synonim"
}

type ResponsibilityConflict struct {
	gorm.Model
	Responsibility1ID uint
	Responsibility1   Responsibility `gorm:"foreignKey:Responsibility1ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Responsibility2ID uint
	Responsibility2   Responsibility `gorm:"foreignKey:Responsibility2ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Priority          int
}

func (rc *ResponsibilityConflict) TableName() string {
	return "responsibility_conflict"
}
