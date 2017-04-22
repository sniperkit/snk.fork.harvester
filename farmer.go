package harvester

import (
	log "github.com/inconshreveable/log15"
)

type Farmer struct {
	pitchforks []Pitchfork
	maxLoops   int
}

func NewFarmer(pitchforks []Pitchfork) *Farmer {
	return &Farmer{
		pitchforks: pitchforks,
		// TODO parametrize this
		maxLoops: 100,
	}
}

// TODO change Data to Seeds
func (f *Farmer) Farm(s Seeds) (Data, error) {
	d := s.ToData()
	loops := 0
	for {
		c := d.Count()

		if f.maxLoops == loops {
			break
		}

		for _, p := range f.pitchforks {
			log.Info("harvest", "pitchfork", p.Name())
			if err := p.Harvest(d); err != nil {
				log.Warn("error executing pitchfork", "name", p.Name(), "error", err)
			}
		}

		finalCount := d.Count()
		if c == finalCount {
			log.Info("process finalized", "finalCount", finalCount)
			break
		}

		loops++
	}

	return d, nil
}
