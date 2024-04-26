package main

type OrganicWorld interface {
	Live()
	NonCellular
	Cell
}
type Cell interface {
	Prokaryote
	Eukaryote
	Grow()
	Divide() Cell
}
type Prokaryote interface {
	ProduceToxins()
	Bacteria
	Archaea
}
type Eukaryote interface {
	CloneGenome()
	Animal
	Fungus
	Plant
}
type Animal interface {
	Move()
	Eat()
}
type Fungus interface{}

type Plant interface{}

type Bacteria interface{}

type Archaea interface{}

type NonCellular interface {
	Replicate() NonCellular
	Virus
}
type Virus interface {
	Infect()
}
type InfluenzaVirus struct {
	Virus
}

func (v *InfluenzaVirus) Infect() {
}
func (v *InfluenzaVirus) Replicate() NonCellular {
	return &InfluenzaVirus{}
}

type AnimalCat struct {
	AnimalEukaryote
}

func (tmp *AnimalCat) Move() {
}
func (tmp *AnimalCat) Eat() {
}

type AnimalProkaryote struct {
	AnimalCell
}

func (tmp *AnimalProkaryote) ProduceToxins() {
}

type AnimalEukaryote struct {
	*AnimalCell
}

func (tmp *AnimalEukaryote) CloneGenome() {
}

type AnimalCell struct{ AnimalCellOrganicWorld }

func (tmp *AnimalCell) Grow() {
}
func (tmp *AnimalCell) Divide() Cell {
	return &AnimalCell{}
}

type AnimalCellOrganicWorld struct {
	OrganicWorld
}

func (tmp *AnimalCellOrganicWorld) Live() {
}
