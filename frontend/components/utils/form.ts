import {FieldValues} from "react-hook-form";

export function extractDirtyFields(data: FieldValues, dirtyFields: any): FieldValues {
	return Object.keys(dirtyFields).reduce<FieldValues>((acc, key) => {
		if (dirtyFields[key]) {
			acc[key] = data[key]
		}
		return acc
	}, {})
}
