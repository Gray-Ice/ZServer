package db

const createAuthTableSQL = `CREATE TABLE auth (
    mac_address char(48),
    hostname varchar(255),
    )
`
