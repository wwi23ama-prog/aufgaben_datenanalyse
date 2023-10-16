package testseries

// Average erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Durchschnittswert.
// Ist die Liste leer, wird 0.0 zurückgegeben.
func Average(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}
	// TODO
	return 0.0
}

// Median erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Median.
// Ist die Liste leer, wird 0 zurückgegeben.
func Median(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]int, len(values))
	copy(sorted, values)

	// TODO
	return sorted[len(sorted)/2]
}

// Mode erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Wert, der am häufigsten in der Liste vorkommt.
// Falls mehrere Werte am häufigsten vorkommen, wird der kleinste dieser Werte
// zurückgegeben. Ist die Liste leer, wird 0 zurückgegeben.
func Mode(values []int) int {
	if len(values) == 0 {
		return 0
	}
	maxpos := 0
	// TODO
	return maxpos
}

// GeometricMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das geometrische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das geometrische Mittel ist definiert als
// die n-te Wurzel aus dem Produkt der n Werte.
func GeometricMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}
	// TODO
	return 0.0
}

// HarmonicMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das harmonische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das harmonische Mittel ist definiert als
// die Kehrwert des Durchschnitts der Kehrwerte der n Werte.
func HarmonicMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}
	// TODO
	return 0.0
}
