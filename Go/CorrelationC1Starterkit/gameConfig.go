package main

type Debug struct {
	PrintMapString           bool `json:"printMapString"`
	PrintTStrings            bool `json:"printTStrings"`
	PrintActStrings          bool `json:"printActStrings"`
	PrintHitStrings          bool `json:"printHitStrings"`
	PrintPlayerInputStrings  bool `json:"printPlayerInputStrings"`
	PrintBotErrors           bool `json:"printBotErrors"`
	PrintPlayerGetHitStrings bool `json:"printPlayerGetHitStrings"`
}

//config struct
//there is two string: a config one/turn one
type UnitInformation struct {
	Damage       float64 `json:"damage"`
	Cost         float64 `json:"cost"`
	GetHitRadius float64 `json:"getHitRadius"`
	ShieldAmount float64 `json:"shieldAmount"`
	Name         string  `json:"display"`
	Portee       float64 `json:"range"`
	Shorthand    string  `json:"shorthand"`
	Stability    float64 `json:"stability"`
	Speed        float64 `json:"speed"`
}
type TimingAndReplay struct {
	WaitTimeBotMax      float64 `json:"waitTimeBotMax"`
	PlayWaitTimeBotMax  float64 `json:"playWaitTimeBotMax"`
	WaitTimeManual      float64 `json:"waitTimeManual"`
	WaitForever         bool    `json:"waitForever"`
	WaitTimeBotSoft     float64 `json:"waitTimeBotSoft"`
	PlayWaitTimeBotSoft float64 `json:"playWaitTimeBotSoft"`
	ReplaySave          float64
	PlayReplaySave      float64
	StoreBotTimes       bool
	WaitTimeStartGame   float64
	WaitTimeEndGame     float64
}

type Resources struct {
	TurnIntervalForBitCapSchedule float64
	TurnIntervalForBitSchedule    float64
	BitRampBitCapGrowthRate       float64
	RoundStartBitRamp             float64
	BitGrowthRate                 float64
	StartingHP                    float64
	MaxBits                       float64
	BitsPerRound                  float64
	CoresPerRound                 float64
	CoresForPlayerDamage          float64
	StartingBits                  float64
	BitDecayPerRound              float64
	StartingCores                 float64
}
type Mechanics struct {
	BasePlayerHealthDamage    float64
	DamageGrowthBasedOnY      float64
	BitsCanStackOnDeployment  bool
	DestroyOwnUnitRefund      float64
	DestroyOwnUnitsEnabled    bool
	StepsRequiredSelfDestruct float64
	SelfDestructRadius        float64
	ShieldDecayPerFrame       float64
	MeleeMultiplier           float64
	DestroyOwnUnitDelay       float64
	RerouteMidRound           bool
	FirewallBuildTime         float64
}

type GameConfig struct {
	Debug            `json:"debug"`
	UnitsInformation []UnitInformation `json:"unitInformation"`
	TimingAndReplay  `json:"timingAndReplay"`
	Resources        `json:"resources"`
	Mechanics        `json:"mechanics"`
}
