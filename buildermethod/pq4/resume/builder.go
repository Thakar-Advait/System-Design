package resume

type resume struct {
	name       string
	education  []string
	experience []string
}

type resumeBuilder struct {
	name       string
	education  []string
	experience []string
}

type Builder interface {
	SetName(name string) Builder
	AddEducation(education string) Builder
	AddExperience(experience string) Builder
	Build() resume
}

func NewResumeBuilder() *resumeBuilder {
	return &resumeBuilder{}
}

func (r *resumeBuilder) SetName(name string) Builder {
	r.name = name
	return r
	// panic("to be implemented")
}

func (r *resumeBuilder) AddEducation(education string) Builder {
	r.education = append(r.education, education)
	return r
	// panic("to be implemeted")
}

func (r *resumeBuilder) AddExperience(experience string) Builder {
	r.experience = append(r.experience, experience)
	return r
	// panic("to be implemented")
}

func (r *resumeBuilder) Build() resume {
	return resume{
		name:       r.name,
		experience: r.experience,
		education:  r.education,
	}
	// panic("to be implemented")
}
