package helpers

// FindResource finds all the resources and returns them in the tabulable data structure
func FindResource(m map[string]interface{}) [][]string {
	data := [][]string{}

	for k, v := range m {
		// fmt.Printf("Terraform object type is is %v\n", k)
		switch k {
		case "resource":
			v1 := v.([]interface{})
			for _, v2 := range v1 {
				v3 := v2.(map[string]interface{})
				for k4, v4 := range v3 {
					v5 := v4.([]interface{})
					for _, v6 := range v5 {
						v7 := v6.(map[string]interface{})
						for k5 := range v7 {
							data = append(data, []string{k, k4, k5})
						}
					}
				}
			}
		default:
			v1 := v.([]interface{})
			for _, v2 := range v1 {
				v3 := v2.(map[string]interface{})
				for k4 := range v3 {
					data = append(data, []string{k, "N/A", k4})
				}
			}
		}
	}

	return data
}
