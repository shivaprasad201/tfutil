package helpers

import (
	"fmt"
)

// FindResource finds all the resources and returns them in the tabulable data structure
func FindResource(m map[string]interface{}) map[string][]string {
	c := make(map[string][]string)

	for k, v := range m {
		switch k {
		case "resource", "data":
			for _, v1 := range v.([]map[string]interface{}) {
				for k2, v2 := range v1 {
					for _, v3 := range v2.([]map[string]interface{}) {
						for k4 := range v3 {
							c[k] = append(c[k], fmt.Sprintf("%v:%v", k2, k4))
						}
					}
				}
			}

		default:
			for _, v1 := range v.([]map[string]interface{}) {
				for k2 := range v1 {
					c[k] = append(c[k], k2)
				}
			}
		}

	}
	return c
}
