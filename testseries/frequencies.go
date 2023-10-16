package testseries

import "github.com/wwi23ama-prog/aufgaben_datenanalyse/intlists"

// AbsoluteFrequencies erwartet eine Liste mit den Werten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den absoluten Häufigkeiten für jeden Wert
// zwischen dem Minimum und dem Maximum der Messreihe.
func AbsoluteFrequencies(values []int) []int {
	// TODO
	return []int{}
}

// RelativeFrequencies erwartet eine Liste mit absoluten Häufigkeiten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den relativen Häufigkeiten.
func RelativeFrequencies(values []int) []float64 {
	freq := make([]float64, len(values))
	sum := float64(intlists.Sum(values))
	for i, r := range values {
		freq[i] = float64(r) / sum
	}
	return freq
}
