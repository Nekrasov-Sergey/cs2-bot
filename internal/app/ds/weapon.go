package ds

type Weapon struct {
	Name            string
	LongRange       bool // Указывает, подходит ли оружие для дальней стрельбы
	StealthAccuracy bool // Указывает, подходит ли оружие для скрытной и точной стрельбы
	ShootOnMove     bool // Указывает, подходит ли оружие для стрельбы на ходу
	Power           bool // Указывает, является ли оружие более мощным, чем скорострельным
	ShootInHead     bool // Указывает, подходит ли оружие для стрельбы в голову
	AccuracyDamage  bool // Указывает, является ли оружие точным и с высоким уроном на дальних дистанциях
	ShootBursts     bool // Указывает, подходит ли оружие для стрельбы очередями
	SmallCost       bool // Указывает, является ли оружие дешёвым
	Reward          bool // Указывает, имеет ли оружие большую награду за убийство
	CloseRange      bool // Указывает, подходит ли оружие для ближней стрельбы
	AmmoReserve     bool // Указывает, есть ли у оружия большой запас патронов
	Light           bool // Указывает, является ли оружие лёгким
	MachineGun      bool // Указывает, имеет ли оружие пулемётную стрельбу
	FastRecharge    bool // Указывает, имеет ли оружие быструю перезарядку
	LowRecoil       bool // Указывает, имеет ли оружие низкую отдачу
	Beautiful       bool // Указывает, является оружие красивым
	MediumRange     bool // Указывает, подходит ли оружие для стрельбы на средней дистанции
	RateOfFire      bool // Указывает, является ли оружие скорострельным
	Aggressive      bool // Указывает, подходит ли оружие для агрессивного стиля игры
	AccuracySlow    bool // Указывает, является ли оружие точным и с медленной стрельбой
}

var Weapons = []Weapon{
	{
		Name:            "AWP",
		LongRange:       true,
		StealthAccuracy: true,
		ShootOnMove:     false,
		Power:           true,
		ShootInHead:     true,
		AccuracyDamage:  true,
		ShootBursts:     false,
		SmallCost:       false,
		Reward:          false,
		CloseRange:      false,
		AmmoReserve:     false,
		Light:           false,
		MachineGun:      false,
		FastRecharge:    false,
		LowRecoil:       false,
		Beautiful:       true,
		MediumRange:     true,
		RateOfFire:      false,
		Aggressive:      false,
		AccuracySlow:    true,
	},
	{
		Name:            "Negev",
		LongRange:       false,
		StealthAccuracy: false,
		ShootOnMove:     true,
		Power:           false,
		ShootInHead:     false,
		AccuracyDamage:  false,
		ShootBursts:     false,
		SmallCost:       true,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     true,
		Light:           false,
		MachineGun:      true,
		FastRecharge:    false,
		LowRecoil:       false,
		Beautiful:       false,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      false,
		AccuracySlow:    false,
	},
	{
		Name:            "P90",
		LongRange:       false,
		StealthAccuracy: false,
		ShootOnMove:     true,
		Power:           false,
		ShootInHead:     false,
		AccuracyDamage:  false,
		ShootBursts:     false,
		SmallCost:       true,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     true,
		Light:           true,
		MachineGun:      false,
		FastRecharge:    true,
		LowRecoil:       true,
		Beautiful:       false,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      true,
		AccuracySlow:    false,
	},
	{
		Name:            "MAG-7",
		LongRange:       false,
		StealthAccuracy: false,
		ShootOnMove:     true,
		Power:           true,
		ShootInHead:     true,
		AccuracyDamage:  false,
		ShootBursts:     false,
		SmallCost:       true,
		Reward:          true,
		CloseRange:      true,
		AmmoReserve:     false,
		Light:           false,
		MachineGun:      false,
		FastRecharge:    false,
		LowRecoil:       false,
		Beautiful:       false,
		MediumRange:     false,
		RateOfFire:      false,
		Aggressive:      true,
		AccuracySlow:    false,
	},
	{
		Name:            "AK-47",
		LongRange:       true,
		StealthAccuracy: true,
		ShootOnMove:     false,
		Power:           true,
		ShootInHead:     true,
		AccuracyDamage:  true,
		ShootBursts:     false,
		SmallCost:       false,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     true,
		Light:           false,
		MachineGun:      false,
		FastRecharge:    false,
		LowRecoil:       false,
		Beautiful:       true,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      true,
		AccuracySlow:    false,
	},
	{
		Name:            "Desert Eagle",
		LongRange:       true,
		StealthAccuracy: true,
		ShootOnMove:     false,
		Power:           true,
		ShootInHead:     true,
		AccuracyDamage:  true,
		ShootBursts:     false,
		SmallCost:       true,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     false,
		Light:           true,
		MachineGun:      false,
		FastRecharge:    true,
		LowRecoil:       false,
		Beautiful:       true,
		MediumRange:     true,
		RateOfFire:      false,
		Aggressive:      false,
		AccuracySlow:    true,
	},
	{
		Name:            "M4A1-S",
		LongRange:       true,
		StealthAccuracy: true,
		ShootOnMove:     false,
		Power:           false,
		ShootInHead:     true,
		AccuracyDamage:  true,
		ShootBursts:     false,
		SmallCost:       false,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     false,
		Light:           false,
		MachineGun:      false,
		FastRecharge:    false,
		LowRecoil:       true,
		Beautiful:       true,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      false,
		AccuracySlow:    false,
	},
	{
		Name:            "M4A4",
		LongRange:       true,
		StealthAccuracy: true,
		ShootOnMove:     false,
		Power:           false,
		ShootInHead:     true,
		AccuracyDamage:  true,
		ShootBursts:     false,
		SmallCost:       false,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     true,
		Light:           false,
		MachineGun:      false,
		FastRecharge:    false,
		LowRecoil:       false,
		Beautiful:       false,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      false,
		AccuracySlow:    false,
	},
	{
		Name:            "MP7",
		LongRange:       false,
		StealthAccuracy: false,
		ShootOnMove:     true,
		Power:           false,
		ShootInHead:     false,
		AccuracyDamage:  false,
		ShootBursts:     false,
		SmallCost:       true,
		Reward:          true,
		CloseRange:      true,
		AmmoReserve:     true,
		Light:           true,
		MachineGun:      false,
		FastRecharge:    true,
		LowRecoil:       true,
		Beautiful:       false,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      true,
		AccuracySlow:    false,
	},
	{
		Name:            "Famas",
		LongRange:       true,
		StealthAccuracy: true,
		ShootOnMove:     false,
		Power:           false,
		ShootInHead:     true,
		AccuracyDamage:  false,
		ShootBursts:     true,
		SmallCost:       true,
		Reward:          false,
		CloseRange:      true,
		AmmoReserve:     true,
		Light:           false,
		MachineGun:      false,
		FastRecharge:    false,
		LowRecoil:       false,
		Beautiful:       false,
		MediumRange:     true,
		RateOfFire:      true,
		Aggressive:      false,
		AccuracySlow:    false,
	},
}
