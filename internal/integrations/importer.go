package integrations

// Ensure that integrations are registered via init()
import (
	_ "github.com/kiyanmair/shift-sync/internal/integrations/foochat"
	_ "github.com/kiyanmair/shift-sync/internal/integrations/foocode"
)
