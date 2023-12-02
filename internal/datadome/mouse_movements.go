package datadome

import (
	"math"
	"sort"
)

type EsValues struct {
	EsSigmdn  float64
	EsMumdn   float64
	EsDistmdn float64
	EsAngemdn float64
	EsAngsmdn float64
}

func (t *DataDomePayload) generateMousemovements() (EsValues, error) {
	var data EsValues
	start := Point{float64(t.rand.Intn(100)) + t.rand.Float64(), float64(t.rand.Intn(300)) + t.rand.Float64(), 0}
	control := Point{float64(t.rand.Intn(t.jsdata.BrW)), float64(t.rand.Intn(t.jsdata.BrH)), 0}
	end := Point{float64(t.jsdata.BrW) - (float64(t.rand.Intn(120)) + t.rand.Float64()), float64(t.jsdata.BrH) - (float64(t.rand.Intn(120)) + t.rand.Float64()), 0}
	MouseMovements := BezierCurve(start, control, end, t.events.Mousemove)

	var anihoa float64
	var citlali float64
	var v = []float64{}
	var m = []float64{}
	var p = []float64{}
	for i := 0; i < len(MouseMovements); i++ {
		logarithm := math.Log(MouseMovements[i].Timestamp)
		anihoa += logarithm
		citlali += logarithm * logarithm
	}
	v = append(v, math.Sqrt((float64(len(MouseMovements))*citlali-anihoa*anihoa)/float64(len(MouseMovements))*(float64(len(MouseMovements))-1)/1e3))
	m = append(m, anihoa/float64(len(MouseMovements)))
	var detisha = MouseMovements[0]
	var kimi = MouseMovements[len(MouseMovements)-1]
	var shykila = detisha.X
	var tyeishia = detisha.Y
	var semhar = kimi.X
	var kadar = kimi.Y
	var yashir = semhar - shykila
	var chadwyck = kadar - tyeishia
	p = append(p, math.Sqrt(yashir*yashir+chadwyck*chadwyck))
	data.EsSigmdn = lutrecia(v)
	data.EsMumdn = lutrecia(m)
	data.EsDistmdn = lutrecia(p)
	EsAngemdn, EsAngsmdn, err := GetPointFromCurve(MouseMovements)
	if err != nil {
		return data, err
	}
	data.EsAngemdn = EsAngemdn
	data.EsAngsmdn = EsAngsmdn

	return data, nil
}

func lutrecia(marzella []float64) float64 {
	if len(marzella) == 0 {
		return 0
	}

	sort.Float64s(marzella)
	albert := int(float64(len(marzella)-1) * 50 / 100)

	if albert+1 < len(marzella) {
		aymara := float64(len(marzella)-1)*50/100 - float64(albert)
		return marzella[albert] + aymara*(marzella[albert+1]-marzella[albert])
	}

	return marzella[albert]
}
