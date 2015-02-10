package pt

type Attenuation struct {
	Constant  float64
	Linear    float64
	Quadratic float64
}

var NoAttenuation = Attenuation{1, 0, 0}

func LinearAttenuation(x float64) Attenuation {
	return Attenuation{1, x, 0}
}

func QuadraticAttenuation(x float64) Attenuation {
	return Attenuation{1, 0, x}
}

func (a *Attenuation) Compute(d float64) float64 {
	return 1 / (a.Constant + a.Linear*d + a.Quadratic*d*d)
}
