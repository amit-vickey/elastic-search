package elasticsearch

var StudentMapping = `{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 1
	},
	"mappings": {
		"properties": {
			"roll_number": {
				"type": "integer"
			},
			"name": {
				"type": "text"
			},
			"age": {
				"type": "integer"
			},
			"gpa": {
				"type": "float"
			},
			"joined_on": {
				"type": "date",
				"format": "yyyy-MM-dd HH:mm:ss||epoch_millis"
			},
			"is_active": {
				"type": "boolean"
			}
		}
	}
}`