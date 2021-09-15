package a

func _()                   {}        // OK
func _(_ int)              {}        // OK
func _(_, _ int)           {}        // OK
func _(_, _ int, _ string) {}        // OK
func _(n int)              { _ = n } // OK
func _(_, n int)           { _ = n } // OK
func _(n int)              {}        // want "n is unused parameter"
func _(n, m int)           {}        // want "n is unused parameter" "m is unused parameter"
