package integrations

// Ensure that integrations are registered via init()
import (
	_ "github.com/kiyanmair/shift-sync/internal/integrations/example_destination"
	_ "github.com/kiyanmair/shift-sync/internal/integrations/example_source"
)
