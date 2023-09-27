package initialize

import "igapp/pkg/app/database"

func pgSQL() {
	database.ConnectDBEcommerce()
}
