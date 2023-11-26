package example

import (
	exampeEntity "github.com/game-core/gocrafter/domain/entity/master/example"
)

func ToExamples(ers *exampeEntity.Examples) *Examples {
	var examples Examples
	for _, er := range *ers {
		example := &Example{
			ID:     er.ID,
			Name:   er.Name,
			Detail: er.Detail,
			Count:  er.Count,
		}
		examples = append(examples, *example)
	}

	return &examples
}
