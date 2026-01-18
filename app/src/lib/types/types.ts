export type Plant = {
	id: string;
	species: string;
	name: string;
	sunLight: SunlightRequirement;
	preferedTemperature: number;
	wateringIntervalDays: number;
	lastWatered: string | null; // ISO date string
	fertilizingIntervalDays: number;
	lastFertilized: string | null; // ISO date string
	preferedHumidity: number;
	sprayIntervalDays?: number;
	notes: string[] | null;
	flags: PlantFlag[] | null;
	photoIds: string[] | null;
};

export enum PlantFlag {
	NO_DRAUGHT = 'No Draught',
	REMOVE_BROWN_LEAVES = 'Remove Brown Leaves'
}

export enum SunlightRequirement {
	FULL_SUN = 'Full Sun',
	INDIRECT_SUN = 'Indirect Sun',
	PARTIAL_SHADE = 'Partial Shade',
	PARTIAL_TO_FULL_SHADE = 'Partial to Full Shade',
	FULL_SHADE = 'Full Shade'
}
