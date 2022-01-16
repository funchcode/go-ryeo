package scheme

const SchemeReceipt = `
	CREATE TABLE IF NOT EXISTS RECEIPT (
		id VARCHAR(255),
		name VARCHAR(255) NOT NULL,
		date DATE NOT NULL,
		kind VARCHAR(255) NOT NULL,
		category VARCHAR(255) NOT NULL,
		contents VARCHAR(255),
		amount INT NOT NULL,
		assets VARCHAR(255) NOT NULL,
		disuse tinyint(1)
		created_at DATE NOT NULL,
		updated_at DATE,
		PRIMARY KEY(id)
	)
`
