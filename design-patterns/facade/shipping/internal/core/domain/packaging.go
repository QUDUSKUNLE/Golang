package domain

type PackagingDTO struct {
	Height      float32 `json:"height" binding:"required" validate:"required"`
	Length      float32 `json:"length" binding:"required" validate:"required"`
	Name        string `json:"name" binding:"required" validate:"required"`
	Size_Unit   string `json:"size_unit" binding:"required" validate:"required"`
	Type        PACKAGE_TYPE `json:"type" binding:"required" validate:"required"`
	Width       float32 `json:"width" binding:"required" validate:"required"`
	Weight      float32 `json:"weight" binding:"required" validate:"required"`
	Weight_Unit string `json:"weight_unit" binding:"required" validate:"required"`
}

type TerminalPackagingDTO struct {
	height      float32 
	length      float32 
	name        string 
	size_unit   string 
	Type        PACKAGE_TYPE 
	width       float32 
	weight      float32 
	weight_unit string
}

func (packaging *PackagingDTO) BuildNewPackaging(pack PackagingDTO) *TerminalPackagingDTO {
	return &TerminalPackagingDTO{
		height: pack.Height,
		length: pack.Length,
		name: pack.Name,
		size_unit: pack.Size_Unit,
		Type: pack.Type,
		width: pack.Width,
		weight: pack.Weight,
		weight_unit: pack.Weight_Unit,
	}
}
