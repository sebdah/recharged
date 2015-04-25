package types

const (
	StopReasonLocalStop          string = "LocalStop"
	StopReasonRemoteStop         string = "RemoteStop"
	StopReasonEmergencyStop      string = "EmergencyStop"
	StopReasonEnergyLimitReached string = "EnergyLimitReached"
	StopReasonEVDisconnected     string = "EVDisconnected"
	StopReasonDeAuthorize        string = "DeAuthorize"
	StopReasonImmediateReset     string = "ImmediateReset"
	StopReasonLocalOutOfCredit   string = "LocalOutOfCredit"
	StopReasonLocalReboot        string = "LocalReboot"
	StopReasonPowerLoss          string = "PowerLoss"
	StopReasonPowerQuality       string = "PowerQuality"
	StopReasonSOCLimitReached    string = "SOCLimitReached"
	StopReasonTimeLimitReached   string = "TimeLimitReached"
	StopReasonUnlockCommand      string = "UnlockCommand"
	StopReasonOther              string = "Other"
)
