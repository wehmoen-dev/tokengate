package tokengate

import "github.com/ethereum/go-ethereum/common"

type TokenGateClient interface {
	HasAxieFromCollection(wallet common.Address, collection AxieCollection) (bool, error)
	HasLandFromCollection(wallet common.Address, collection LandCollection) (bool, error)

	GetAxieSummary(wallet common.Address) (AxieSummary, error)
	GetLandSummary(wallet common.Address) (LandSummary, error)
}

type AxieCollection string

const (
	MysticAxie   AxieCollection = "mystic"
	XmasAxie     AxieCollection = "xmas"
	JapaneseAxie AxieCollection = "japanese"
	ShinyAxie    AxieCollection = "shiny"
	SummerAxie   AxieCollection = "summer"
	OriginAxie   AxieCollection = "origin"
)

type AxieSummary map[AxieCollection]int

type LandCollection string

const (
	SavannahLand LandCollection = "savannah"
	ForestLand   LandCollection = "forest"
	ArcticLand   LandCollection = "arctic"
	MysticLand   LandCollection = "mystic"
	GenesisLand  LandCollection = "genesis"
)

type LandSummary map[LandCollection]int

type StakingType string

const (
	RONStaking  StakingType = "ron"
	AXSStaking  StakingType = "axs"
	LANDStaking StakingType = "land"
)
